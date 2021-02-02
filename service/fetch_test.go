package service

//go:generate mockgen --destination=mock/browser.go . Browser
//go:generate mockgen --destination=mock/storage.go . Storage

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"testing"

	"github.com/golang/mock/gomock"
	service "github.com/pshvedko/grpc-product/service/mock"
)

type query string

func (q query) GetUrl() string {
	return string(q)
}

type body [][]string

func (b body) Read(p []byte) (n int, err error) {
	return 0, io.EOF
}

func (b body) Close() error {
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
					m := service.NewMockBrowser(ctrl)
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
						Body:       body{},
					}, nil)
					return m
				}(),
				Storage: func() Storage {
					m := service.NewMockStorage(ctrl)
					m.EXPECT()
					return m
				}(),
			},
			args: args{
				ctx:   ctx,
				query: query("http://localhost/"),
			},
			wantLoaded:  0,
			wantChanged: 0,
			wantAdded:   0,
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
