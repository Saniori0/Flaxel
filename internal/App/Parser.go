package App

import "strings"

func (app App) ParseCommand(command string) Command {
	commandComponents := strings.Split(command, " ")

	commandId := commandComponents[0]
	commandVariables := strings.Split(commandComponents[1], ",")

	switch commandId {
	case "Cache":
		return &Cache{
			vars: commandVariables,
		}
	case "Read":
		return &Read{
			vars: commandVariables,
		}
	}

	return nil
}
