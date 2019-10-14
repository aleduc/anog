package handlers

import (
	"github.com/valyala/fasthttp"
	"fmt"
)

// константы неэкспортируемые и очевидные, поэтому без комментов
const (
	getLocation = "/get"
	loadLocation = "/load"
	wordParam = "word"
)

// router интерфейс для роутинга запросов
type router interface {
	Route(ctx *fasthttp.RequestCtx)
}

// route роутинг с анаграммами
type route struct {
	anagram anagramer
	load preparer
}

// NewRouter конструткор для роутера, прнимает anagramer. Нужно для мокирования в тестах
func NewRouter(anagram anagramer, load preparer) router {
	return &route{anagram: anagram, load:load}
}

// Route функция роутинга запросов, обрабатывает запросы и данные из них, потом уже работает со словарем анаграмм
func(r *route) Route(ctx *fasthttp.RequestCtx) {
	switch string(ctx.Path()) {
	case getLocation:
		word := string(ctx.QueryArgs().Peek(wordParam))
		if word != ""{
			ctx.SetContentType("application/json")
			_, _ = fmt.Fprint(ctx, r.anagram.get(word))
		}
	case loadLocation:
		words, err := r.load.prepare(ctx.PostBody())
		if err == nil {
			r.anagram.load(words)
		}
	}
}
