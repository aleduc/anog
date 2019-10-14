package handlers

import (
	"github.com/mailru/easyjson"
	"sync"
)

const nullValue = "null"

// anagramer интерфейс для работы со словарем анаграм
type anagramer interface {
	get(string) string
	load(loadData)
}

// anagram содержит загруженный словарь и мюьтекс для совместного доступа
type anagram struct {
	dict map[string]string
	sync.RWMutex
}

// NewAnagram конструктор для anagramer
func NewAnagram() anagramer {
	return &anagram{}
}

func(a *anagram) get(word string) string{
	a.RLock()
	defer a.RUnlock() // в случае борьбы за наносекунды/микросекунды(в случае многократного вызова), есть смысл не использовать дефер
	if val, ok := a.dict[normalize(word)]; ok {
		return val
	} else {
		return nullValue
	}
}

func(a *anagram) load(data loadData) {
	dict := make(map[string]loadData) // нет смысла задавать длину, предполагается что делается редко, а какие слова - не ясно
	for _, word := range data {
		normWord := normalize(word)
		dict[normWord] = append(dict[normWord], word)
	}
	result := make(map[string]string, len(dict))
	for k, dictLine := range dict {
		line,_ := easyjson.Marshal(dictLine)
		result[k] = string(line)
	}
	a.Lock()
	a.dict = result
	a.Unlock()
}
