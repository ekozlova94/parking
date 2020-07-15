package parking

import (
	"context"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ekozlova94/parking/internal/ctxutils"
	"github.com/ekozlova94/parking/pkg/forms"
	"github.com/stretchr/testify/assert"
)

func TestSubscription_Conflict(t *testing.T) {
	// arrange
	t.Parallel()
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	//noinspection GoUnhandledErrorResult
	defer db.Close()

	ctx := ctxutils.NewDbContext(context.Background(), db)
	req := &forms.Request{AutoNumber: 1}

	mock.ExpectQuery("^select (.+) from parking").
		WithArgs(req.AutoNumber).
		WillReturnRows(sqlmock.NewRows([]string{"id", "auto_number"}).AddRow(1, 1))

	// act
	errSub := Subscription(ctx, req)

	// assert
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
	assert.True(t, errors.Is(errSub, ErrConflict))
}

func TestSubscription_Success(t *testing.T) {
	// arrange
	t.Parallel()
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	//noinspection GoUnhandledErrorResult
	defer db.Close()

	ctx := ctxutils.NewDbContext(context.Background(), db)
	req := &forms.Request{AutoNumber: 1}

	mock.ExpectQuery("^select (.+) from parking").
		WithArgs(req.AutoNumber).
		WillReturnRows(sqlmock.NewRows([]string{"id", "auto_number"}))
	mock.ExpectExec("insert into parking").
		WithArgs(req.AutoNumber).
		WillReturnResult(sqlmock.NewResult(1, 1))

	// act
	errSub := Subscription(ctx, req)

	// assert
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
	assert.NoError(t, errSub)
}

func TestCheck_NoPass(t *testing.T) {
	// arrange
	t.Parallel()
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	//noinspection GoUnhandledErrorResult
	defer db.Close()

	ctx := ctxutils.NewDbContext(context.Background(), db)
	req := &forms.Request{AutoNumber: 1}

	mock.ExpectQuery("^select (.+) from parking").
		WithArgs(req.AutoNumber).
		WillReturnRows(sqlmock.NewRows([]string{"id", "auto_number"}))

	// act
	errSub := Check(ctx, req)

	// assert
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
	assert.True(t, errors.Is(errSub, ErrNotFound))
}

func TestCheck_Pass(t *testing.T) {
	// arrange
	t.Parallel()
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	//noinspection GoUnhandledErrorResult
	defer db.Close()

	ctx := ctxutils.NewDbContext(context.Background(), db)
	req := &forms.Request{AutoNumber: 1}

	mock.ExpectQuery("^select (.+) from parking").
		WithArgs(req.AutoNumber).
		WillReturnRows(sqlmock.NewRows([]string{"id", "auto_number"}).AddRow(1, 1))

	// act
	errSub := Check(ctx, req)

	// assert
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
	assert.NoError(t, errSub)
}
