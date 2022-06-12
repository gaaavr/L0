package pkg

import (
	"L0/pkg/api"
	"L0/pkg/cache"
	"L0/pkg/store"
	"L0/pkg/streaming"
	"log"
	"net/http"
)

// App главная структура сервиса
type App struct {
	stream streaming.Streaming
	api    api.Api
}

// Run запускает стриминговый сервис и сервер
func (a *App) Run(stanClusterID, clientID, URL string) {
	err := a.Init()
	if err != nil {
		log.Fatal(err)
	}

	err = a.stream.ConnectAndSubscribe(stanClusterID, clientID, URL)
	if err != nil {
		log.Fatal(err)
	}

	err = http.ListenAndServe(":8080", a.api.Router)
	if err != nil {
		log.Fatal(err)
	}
}

// Init инициализирует кэш, бд, апи и стриминговый сервис
func (a *App) Init() error {
	cache := cache.NewCache()


	newStore, err := store.NewStore("postgres://postgres:1488@localhost:5432/wb")
	if err != nil {
		return err
	}

	a.api.InitRouter(cache)

	a.stream.InitStreaming(cache, newStore)

	err = newStore.RestoreCache(cache)
	if err != nil {
		return err
	}

	return nil
}
