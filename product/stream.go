package product

import (
	"io"
	"log"
)

func (a Application) ListStream(stream ProductService_ListStreamServer) error {
	for {
		query, err := stream.Recv()
		log.Printf("stream: %v %v", query, err)
		if err != nil {
			if err == io.EOF {
				err = nil
			}
			return err
		}
		//	_, _ = a.Service.List(stream.Context(), query)
	}
}
