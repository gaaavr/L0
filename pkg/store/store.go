package store

import (
	"L0/pkg/cache"
	"L0/pkg/model"
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
)

// Store структура для базы данных
type Store struct {
	dbpool *pgxpool.Pool
}

// NewStore возвращает структуру с инициализированной базой данных
func NewStore(url string) (*Store, error) {
	var err error
	dbpool, err := pgxpool.Connect(context.Background(), url)
	if err != nil {
		return &Store{}, err
	}

	return &Store{dbpool: dbpool}, nil
}

// AddOrder добавляет заказ в базу данных
func (s *Store) AddOrder(order model.Order) error {
	_, err := s.dbpool.Exec(context.Background(), "insert into orders (\"order\") values ($1)", order)
	if err != nil {
		return err
	}
	return nil
}


// RestoreCache заполняет кэш заказами из базы данных
func (s *Store) RestoreCache(cache *cache.Cache) error {
	rows, err := s.dbpool.Query(context.Background(), "select \"order\" from orders")
	if err != nil {
		return err
	}

	for rows.Next() {
		var order model.Order
		err = rows.Scan(&order)
		if err != nil {
			return err
		}

		cache.Put(order)
	}

	return rows.Err()
}
