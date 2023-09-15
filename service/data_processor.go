package service

import (
	"context"
	"encoding/json"
	"github.com/schiduluca/port-processor/models"
	"io"
)

type Storage interface {
	Save(ctx context.Context, key string, val models.Port) error
}

type JSONProcessor struct {
	storage Storage
}

func NewJSONProcessor(repo Storage) JSONProcessor {
	return JSONProcessor{
		storage: repo,
	}
}

func (p JSONProcessor) Process(ctx context.Context, r io.Reader) error {
	decoder := json.NewDecoder(r)

	// Read opening delimiter. `[` or `{`
	if _, err := decoder.Token(); err != nil {
		return err
	}

	for decoder.More() {
		// Read opening delimiter. Key to object, like `AEAJM`
		token, err := decoder.Token()
		if err != nil {
			return err
		}

		var port models.Port
		if err := decoder.Decode(&port); err != nil {
			return err
		}

		err = p.storage.Save(ctx, token.(string), port)
		if err != nil {
			return err
		}

	}

	// Read closing delimiter. `]` or `}`
	if _, err := decoder.Token(); err != nil {
		return err
	}

	return nil
}
