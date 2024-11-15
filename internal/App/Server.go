package App

import (
	"bufio"
	"net"
)

func (app App) runServer() {
	listener, err := net.Listen("tcp", "localhost:9090")

	if err != nil {
		return
	}

	defer listener.Close()

	for {
		connection, err := listener.Accept()

		if err != nil {
			return
		}

		go app.handleConnection(connection)
	}
}

func (app App) handleConnection(connection net.Conn) {
	defer connection.Close()

	for {
		input, err := bufio.NewReader(connection).ReadString('\n')

		if err != nil {
			return
		}

		output := app.ExecCommand(input)

		connection.Write([]byte(output))
	}
}
