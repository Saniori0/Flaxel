package App

func (app App) ExecCommand(unparsedCommand string) string {
	command := app.ParseCommand(unparsedCommand)

	return command.executeInApp(app)
}
