package streaming

import (
	"L0/pkg/cache"
	"L0/pkg/model"
	"L0/pkg/store"
	"bytes"
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/nats-io/stan.go"
	"log"
)

// Streaming структура для стримингового сервиса
type Streaming struct {
	sc    stan.Conn
	cache *cache.Cache
	store *store.Store
}

// InitStreaming инициализирует стриминговый сервис с кэшом и бд
func (s *Streaming) InitStreaming(cache *cache.Cache, store *store.Store) {
	s.cache = cache
	s.store = store
}

// ConnectAndSubscribe коннектится к nats-streaming и подписывается на канал
func (s *Streaming) ConnectAndSubscribe(stanClusterID, clientID, URL string) error {
	var err error
	s.sc, err = stan.Connect(stanClusterID, clientID, stan.NatsURL(URL))
	if err != nil {
		return err
	}

	validate := validator.New()

	_, err = s.sc.Subscribe("service", func(m *stan.Msg) {
		var msg model.Order
		dec := json.NewDecoder(bytes.NewReader(m.Data))
		dec.DisallowUnknownFields()
		err := dec.Decode(&msg)
		if err != nil {
			log.Println(err)
			err = m.Ack()
			if err != nil {
				return
			}
			return
		}
		err = validate.Struct(msg)
		if err != nil {
			log.Println(err)
			err = m.Ack()
			if err != nil {
				return
			}
			return
		}

		if !s.cache.IsExist(msg.OrderUID) {
			err := s.store.AddOrder(msg)
			if err != nil {
				log.Println(err)
				err = m.Ack()
				if err != nil {
					return
				}
				return
			}
			s.cache.Put(msg)
		}

		err = m.Ack()
		if err != nil {
			return
		}
	}, stan.DurableName("durable-service"), stan.SetManualAckMode())
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
