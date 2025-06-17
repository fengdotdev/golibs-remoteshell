package shell

import (
	"context"
	"fmt"
)

func Terminal(ctx *context.Context) {

	myshell := New()

	newctx, cancel := context.WithCancel(*ctx)
	defer cancel()

	for {
		select {
		case <-*newctx.Done():
			fmt.Println("Exiting terminal...")
			return
		default:
			var command string

			print("Enter command (or 'exit' to quit): ")
			fmt.Scanln(&command)

			if command == "exit" {

			}

			output, err := myshell.RunCommand(command)
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}

			fmt.Println(output)
		}
	}

}
