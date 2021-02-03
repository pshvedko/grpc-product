package service

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"os"
	"testing"

	"github.com/golang/mock/gomock"
	mockService "github.com/pshvedko/grpc-product/service/mock"
	mockStorage "github.com/pshvedko/grpc-product/storage/mock"
)

type query struct {
	url string
}

func (q query) GetUrl() string {
	return q.url
}

type body struct {
	csv [][]string
	row int
	col int
	pos int
	end bool
}

func (b *body) Read(p []byte) (int, error) {
	if b.end {
		return 0, os.ErrClosed
	} else if len(p) == 0 {
		return 0, nil
	}
	for _, r := range b.csv[b.row:] {
		n := len(r)
		for _, c := range r[b.col:] {
			w := c[b.pos:]
			if len(w) == 0 {
				b.col++
				b.pos = -1
				if b.col < n {
					w = ";"
				} else {
					w = "\n"
				}
			}
			copy(p, w[:1])
			b.pos++
			return 1, nil
		}
		b.row++
		b.col = 0
	}
	return 0, io.EOF
}

func (b *body) Close() error {
	b.end = true
	return nil
}

func TestService_Fetch(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	type fields struct {
		Browser Browser
		Storage Storage
	}
	type args struct {
		ctx   context.Context
		query FetchQuery
	}
	ctx := context.TODO()
	tests := []struct {
		name        string
		fields      fields
		args        args
		wantLoaded  uint32
		wantChanged uint32
		wantAdded   uint32
		wantErr     bool
	}{
		// TODO: Add test cases.
		{
			name: "Ok",
			fields: fields{
				Browser: func() Browser {
					m := mockService.NewMockBrowser(ctrl)
					r := &http.Request{
						Method: http.MethodGet,
						URL: &url.URL{
							Scheme: "http",
							Host:   "localhost",
							Path:   "/",
						},
						Proto:      "HTTP/1.1",
						ProtoMajor: 1,
						ProtoMinor: 1,
						Header:     http.Header{},
						Host:       "localhost",
					}
					r = r.WithContext(ctx)
					m.EXPECT().Do(gomock.Eq(r)).Return(&http.Response{
						StatusCode: 200,
						Body: &body{
							csv: [][]string{
								{"PRODUCT NAME", "PRICE"},
								{"USD", "75.905"},
								{"EUR", "91.625"},
								{"CNY", "11.749"},
								{"RUR", "1"},
								{"RUR", "1"},
								{"RUR", "2"},
							},
						},
					}, nil)
					return m
				}(),
				Storage: func() Storage {
					m := mockService.NewMockStorage(ctrl)
					p := mockStorage.NewMockProducts(ctrl)
					gomock.InOrder(
						p.EXPECT().Push("USD", 75.905).Return(0, nil),
						p.EXPECT().Push("EUR", 91.625).Return(0, nil),
						p.EXPECT().Push("CNY", 11.749).Return(0, nil),
						p.EXPECT().Push("RUR", 1.).Return(0, nil),
						p.EXPECT().Push("RUR", 1.).Return(1, nil),
						p.EXPECT().Push("RUR", 2.).Return(2, nil),
					)
					m.EXPECT().Products().Return(p).Times(6)
					return m
				}(),
			},
			args: args{
				ctx:   ctx,
				query: query{url: "http://localhost/"},
			},
			wantChanged: 1,
			wantLoaded:  6,
			wantAdded:   4,
			wantErr:     false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Service{
				Browser: tt.fields.Browser,
				Storage: tt.fields.Storage,
			}
			gotLoaded, gotChanged, gotAdded, err := s.Fetch(tt.args.ctx, tt.args.query)
			if (err != nil) != tt.wantErr {
				t.Errorf("Fetch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotLoaded != tt.wantLoaded {
				t.Errorf("Fetch() gotLoaded = %v, want %v", gotLoaded, tt.wantLoaded)
			}
			if gotChanged != tt.wantChanged {
				t.Errorf("Fetch() gotChanged = %v, want %v", gotChanged, tt.wantChanged)
			}
			if gotAdded != tt.wantAdded {
				t.Errorf("Fetch() gotAdded = %v, want %v", gotAdded, tt.wantAdded)
			}
		})
	}
}
