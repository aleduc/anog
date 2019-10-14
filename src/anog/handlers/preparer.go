package handlers

import (
	"github.com/mailru/easyjson"
)

//easyjson:json
type loadData []string

// preparer интерфейс для подготвки данных
type preparer interface {
	prepare([]byte) (loadData, error)
}

func NewLoadPrepare() preparer{
	return &load{}
}
// load имплементирует интерфейс preparer для загружаемых данных
type load struct {}

func(l *load) prepare(input []byte) (loadData, error){
	var result loadData
	// easyjson, т.к. она удобней всего тут зашла(после стандартной)
	err := easyjson.Unmarshal(input, &result)
	return result, err
}

