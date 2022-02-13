package cli

func Run(args []string) {
	var option string = args[0]
	ExecuteCommand(option)
}
