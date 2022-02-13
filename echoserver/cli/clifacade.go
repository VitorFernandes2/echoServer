package cli

import (
	api "echoserver/echoserver/api"
	"fmt"
)

func ExecuteCommand(option string) {
	switch option {
	case "server":
		api.Run(8082)
	case "help":
		fmt.Println("COMMAND \t\t DESCRIPTION")
		fmt.Println("main server \t\t run HTTP server API\nmain other_string \t create a POST request")
	default:
		CreateMessage(option)
	}
}
