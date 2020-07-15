package storage

import (
	"context"
	"fmt"

	"github.com/ekozlova94/parking/internal/ctxutils"
	"github.com/ekozlova94/parking/internal/model"
)

func Save(autoNumber int, ctx context.Context) error {
	db := ctxutils.DbFromContext(ctx)
	if _, err := db.Exec("insert into parking(auto_number) values ($1)", autoNumber); err != nil {
		return fmt.Errorf("got wrong while saving auto number: %w", err)
	}
	return nil
}

func Get(autoNumber int, ctx context.Context) (*model.Subscription, error) {
	db := ctxutils.DbFromContext(ctx)
	result, err := db.Query("select id, auto_number from parking where auto_number = $1 limit 1", autoNumber)
	if err != nil {
		return nil, fmt.Errorf("got error while saving: %w", err)
	}
	//noinspection GoUnhandledErrorResult
	defer result.Close()
	for result.Next() {
		var m model.Subscription
		err = result.Scan(&m.ID, &m.AutoNumber)
		if err != nil {
			return nil, fmt.Errorf("got wrong while getting auto number: %w", err)
		}
		return &m, nil
	}
	return nil, nil
}
