package App

import (
	"container/list"
	"context"
	"encoding/json"
	"strconv"
	"time"
)

type Chunk struct {
	Index     int        `json:"index,omitempty"`
	Rows      *list.List `json:"rows,omitempty"`
	file      string
	listName  string
	listIndex int
}

func NewChunk(index int, file string, listName string, listIndex int) Chunk {
	return Chunk{
		Index:     index,
		Rows:      list.New(),
		file:      file,
		listName:  listName,
		listIndex: listIndex,
	}
}

func (chunk Chunk) Add(row []string) {
	chunk.Rows.PushBack(row)
}

func (chunk Chunk) StoreToApp(app App) {
	key := chunk.file + "-" + strconv.Itoa(chunk.listIndex) + "-" + chunk.GetIndex()

	for e := chunk.Rows.Front(); e != nil; e = e.Next() {
		jsoned, _ := json.Marshal(e)

		app.Redis.LPush(context.Background(), key, jsoned)
		app.Redis.Expire(context.Background(), key, time.Hour)
	}
}

func (chunk Chunk) ReadyToStore() bool {
	return chunk.Count() >= 100
}

func (chunk Chunk) GetIndex() string {
	return strconv.Itoa(chunk.Index)
}

func (chunk Chunk) ToJson() ([]byte, error) {
	return json.Marshal(chunk.Rows)
}

func (chunk Chunk) Count() int {
	return chunk.Rows.Len()
}
