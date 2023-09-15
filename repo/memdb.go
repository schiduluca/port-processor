package repo

import (
	"context"
	"encoding/json"
	"github.com/redis/go-redis/v9"
	"github.com/schiduluca/port-processor/models"
	"os"
)

type MemDB struct {
	client *redis.Client
}

func NewMemDB() *MemDB {
	return &MemDB{
		client: redis.NewClient(&redis.Options{
			Addr:     os.Getenv("REDIS_HOST"),
			Password: "", // no password set
			DB:       0,  // use default DB
		}),
	}
}

func (m *MemDB) Save(ctx context.Context, key string, val models.Port) error {
	b, err := json.Marshal(&val)
	if err != nil {
		return err
	}
	err = m.client.Set(ctx, key, string(b), 0).Err()
	if err != nil {
		return err
	}
	return nil
}
