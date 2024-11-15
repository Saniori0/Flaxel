package App

import (
	"context"
	"fmt"
	"github.com/xuri/excelize/v2"
	"strconv"
	time "time"
)

type Command interface {
	executeInApp(app App) string
	filePath() string
}

type Cache struct {
	Command
	vars []string
}

func (command *Cache) executeInApp(app App) string {
	file, err := command.file()

	fmt.Println("Start cache " + file.Path + "...")

	if err != nil {
		fmt.Println(err)
		return ""
	}

	list, _ := command.list()
	listName := file.GetSheetList()[list]

	var rows, err2 = file.GetRows(listName)
	rowsLen := len(rows)

	fmt.Println("Rows quantity: " + strconv.Itoa(rowsLen))

	if err2 != nil {
		return ""
	}

	defer file.Close()

	chunkIndex := 0
	chunk := NewChunk(chunkIndex, command.filePath(), listName, list)

	for _, row := range rows {
		if chunk.ReadyToStore() {
			chunkIndex += 1

			fmt.Println("Chunk " + chunk.GetIndex() + " saving...")

			chunk.StoreToApp(app)

			chunk = NewChunk(chunkIndex, command.filePath(), strconv.Itoa(list), list)
		}

		chunk.Add(row)
	}

	fmt.Println(command.filePath() + "-" + strconv.Itoa(list) + "-rowsQuantity")

	app.Redis.Set(context.Background(), command.filePath()+"-"+strconv.Itoa(list)+"-rowsQuantity", rowsLen, time.Hour)

	return "ok bro!"
}

func (command *Cache) file() (*excelize.File, error) {
	return excelize.OpenFile(command.filePath())
}

func (command *Cache) list() (int, error) {
	return strconv.Atoi(command.vars[1])
}

func (command *Cache) filePath() string {
	return command.vars[0]
}

type Read struct {
	Command
	vars []string
}

func (command *Read) executeInApp(app App) string {
	return "take ur data bro!"
}
func (command *Read) filePath() string {
	return command.vars[0]
}

type Count struct {
	Command
	vars []string
}

func (command *Count) executeInApp(app App) string {
	return "so many rows bro!"
}
