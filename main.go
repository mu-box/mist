package main

import (
	"fmt"

	"github.com/mu-box/mist/commands"
)

func main() {
	if err := commands.MistCmd.Execute(); err != nil && err.Error() != "" {
		fmt.Println(err.Error())
	}
}
