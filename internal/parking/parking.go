package parking

import (
	"context"
	"fmt"

	"github.com/ekozlova94/parking/internal/storage"
	"github.com/ekozlova94/parking/pkg/forms"
)

var ErrConflict = fmt.Errorf("you can not subscribe a second time")
var ErrNotFound = fmt.Errorf("you are not subscribed")

func Subscription(ctx context.Context, req *forms.Request) error {
	result, err := storage.Get(req.AutoNumber, ctx)
	if err != nil {
		return err
	}
	if result != nil {
		return ErrConflict
	}
	if err := storage.Save(req.AutoNumber, ctx); err != nil {
		return err
	}
	return nil
}

func Check(ctx context.Context, req *forms.Request) error {
	result, err := storage.Get(req.AutoNumber, ctx)
	if err != nil {
		return err
	}
	if result == nil {
		return ErrNotFound
	}
	return nil
}
