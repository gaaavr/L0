package main

import (
	"github.com/nats-io/stan.go"
	"io/ioutil"
	"log"
)

// скрипт для публикации данных в канале
func main() {
	sc, err := stan.Connect("test-cluster", "stan", stan.NatsURL("nats://0.0.0.0:4222"))
	if err != nil {
		log.Println(err)
		return
	}

	fileNames := []string{
		"script/data1.json",
		"script/data2.json",
		"script/data3.json",
		"script/invalid.txt",
		"script/invalidData1.json",
		"script/invalidData2.json",
	}

	for _, name := range fileNames {
		data, err := ioutil.ReadFile(name)
		if err != nil {
			log.Println(err)
			return
		}
		err = sc.Publish("service", data)
		if err != nil {
			log.Println(err)
			return
		}
	}
	err = sc.Close()
	if err != nil {
		log.Println(err)
		return
	}
}