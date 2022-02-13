package main

import (
	cli "echoserver/echoserver/cli"

	"os"
)

func main() {
	args := os.Args[1:]
	cli.Run(args)
}
