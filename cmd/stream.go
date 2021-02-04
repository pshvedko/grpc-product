package cmd

import (
	"context"
	"github.com/eiannone/keyboard"
	"github.com/pshvedko/grpc-product/product"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"log"
)

var streamCmd = &cobra.Command{
	Use:   "stream",
	Short: "A brief description of your command",
	Long:  "FIXME",
	RunE:  runStream,
}

func runStream(*cobra.Command, []string) error {
	queries := make(chan product.ListQuery)
	errors := make(chan error)
	defer func() { <-errors }()
	defer close(queries)

	go do(queries, errors)

	log.Println("Press any key...")

	keys, err := keyboard.GetKeys(10)
	if err != nil {
		return err
	}
	defer keyboard.Close()

	for {
		select {
		case err = <-errors:
			return err
		case key := <-keys:
			if key.Err != nil {
				return key.Err
			}
			switch key.Key {
			case keyboard.KeyEsc, keyboard.KeyCtrlC, keyboard.KeyEnter:
				return err
			case keyboard.KeySpace:
				log.Println("continue!")
			}
		}
	}
}

func do(queries <-chan product.ListQuery, errors chan<- error) {
	defer close(errors)
	dial, err := grpc.Dial(addrFlag.String(), grpc.WithInsecure())
	if err != nil {
		return
	}
	defer dial.Close()
	client := product.NewProductServiceClient(dial)

	var stream product.ProductService_ListStreamClient
	stream, err = client.ListStream(context.TODO())
	if err != nil {
		log.Println("1", err)
		return
	}

	for range queries {
		read(stream)
	}

	err = stream.CloseSend()
	if err != nil {
		log.Println("3", err)
		return
	}
	read(stream)

	log.Println("END")
}

func read(stream product.ProductService_ListStreamClient) {
	for {
		_, err := stream.Recv()
		if err != nil {
			log.Println("2", err)
			break
		}
	}
}

func init() {
	listCmd.AddCommand(streamCmd)
}
