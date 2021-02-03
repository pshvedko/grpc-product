package cmd

import (
	"github.com/eiannone/keyboard"
	"github.com/spf13/cobra"
	"log"
)

var streamCmd = &cobra.Command{
	Use:   "stream",
	Short: "A brief description of your command",
	Long:  "FIXME",
	RunE:  runStream,
}

func runStream(*cobra.Command, []string) (err error) {
	log.Println("Press any key...")
	var key keyboard.Key
	for {
		_, key, err = keyboard.GetSingleKey()
		if err != nil {
			return
		}
		switch key {
		case keyboard.KeyEsc:
			return
		case keyboard.KeyEnter:
			log.Println("ENTER key pressed!")
		case keyboard.KeySpace:
			log.Println("SPACE key pressed!")
		}
	}
}

func init() {
	listCmd.AddCommand(streamCmd)
}
