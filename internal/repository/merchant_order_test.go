package repository_test

import (
	"context"
	"database/sql"
	"errors"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	"github.com/Pos-Tech-Challenge-48/delivery-merchant-api/internal/entities"
	"github.com/Pos-Tech-Challenge-48/delivery-merchant-api/internal/repository"
)

func TestMerchantRepository_Update(t *testing.T) {
	tests := []struct {
		name        string
		order       *entities.MerchantOrder
		setupMock   func(sqlmock.Sqlmock, *entities.MerchantOrder)
		expectedErr error
	}{
		{
			name: "successful update",
			order: &entities.MerchantOrder{
				MerchantID: "123",
				OrderID:    "456",
				Status:     "completed",
			},
			setupMock: func(mock sqlmock.Sqlmock, order *entities.MerchantOrder) {
				mock.ExpectExec(`UPDATE merchant_order`).
					WithArgs(order.Status, order.MerchantID, order.OrderID).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			expectedErr: nil,
		},
		{
			name: "update fails",
			order: &entities.MerchantOrder{
				MerchantID: "123",
				OrderID:    "456",
				Status:     "failed",
			},
			setupMock: func(mock sqlmock.Sqlmock, order *entities.MerchantOrder) {
				mock.ExpectExec(`UPDATE merchant_order`).
					WithArgs(order.Status, order.MerchantID, order.OrderID).
					WillReturnError(errors.New("update failed"))
			},
			expectedErr: errors.New("update failed"),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()

			r := repository.NewMerchantRepository(db)
			tc.setupMock(mock, tc.order)

			err = r.Update(context.Background(), tc.order)
			assert.Equal(t, tc.expectedErr, err)
			assert.Nil(t, mock.ExpectationsWereMet())
		})
	}
}

func TestMerchantRepository_Get(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	r := repository.NewMerchantRepository(db)

	query := `SELECT merchant_id, order_id,status, last_modified_date, created_date
	FROM merchant_order
	WHERE merchant_id = $1 and order_id = $2`

	ctx := context.Background()
	merchantID := "123"
	orderID := "456"

	tests := []struct {
		name           string
		mockSetup      func()
		expectedResult *entities.MerchantOrder
		expectedErr    error
	}{
		{
			name: "order found",
			mockSetup: func() {
				timestamp := time.Now().Round(time.Second)
				rows := sqlmock.NewRows([]string{"merchant_id", "order_id", "status", "last_modified_date", "created_date"}).
					AddRow(merchantID, orderID, "received", timestamp, timestamp)

				mock.ExpectQuery(query).
					WithArgs(merchantID, orderID).
					WillReturnRows(rows)
			},
			expectedResult: &entities.MerchantOrder{
				MerchantID:       merchantID,
				OrderID:          orderID,
				Status:           "received",
				LastModifiedDate: time.Now().Round(time.Second),
				CreatedDate:      time.Now().Round(time.Second),
			},
			expectedErr: nil,
		},
		{
			name: "order not found",
			mockSetup: func() {
				rows := sqlmock.NewRows([]string{"merchant_id", "order_id", "status", "last_modified_date", "created_date"})
				mock.ExpectQuery(query).
					WithArgs(merchantID, orderID).
					WillReturnRows(rows)
			},
			expectedResult: nil,
			expectedErr:    nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockSetup()
			result, err := r.Get(ctx, merchantID, orderID)
			if result != nil {
				result.LastModifiedDate = result.LastModifiedDate.Round(time.Second)
				result.CreatedDate = result.CreatedDate.Round(time.Second)
			}
			assert.Equal(t, tc.expectedResult, result)
			assert.Equal(t, tc.expectedErr, err)
			assert.Nil(t, mock.ExpectationsWereMet())
		})
	}
}

func TestMerchantRepository_Insert(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	r := repository.NewMerchantRepository(db)

	ctx := context.Background()
	merchantOrder := &entities.MerchantOrder{
		MerchantID: "123",
		OrderID:    "456",
		Status:     "received",
	}

	tests := []struct {
		name        string
		order       *entities.MerchantOrder
		mockSetup   func()
		expectedErr error
	}{
		{
			name:  "successful insert",
			order: merchantOrder,
			mockSetup: func() {
				mock.ExpectExec(`INSERT INTO merchant_order`).
					WithArgs(merchantOrder.MerchantID, merchantOrder.OrderID, merchantOrder.Status).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			expectedErr: nil,
		},
		{
			name:  "insert fails due to DB error",
			order: merchantOrder,
			mockSetup: func() {
				mock.ExpectExec(`INSERT INTO merchant_order`).
					WithArgs(merchantOrder.MerchantID, merchantOrder.OrderID, merchantOrder.Status).
					WillReturnError(sql.ErrConnDone)
			},
			expectedErr: sql.ErrConnDone,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockSetup()
			err := r.Insert(ctx, tc.order)
			assert.Equal(t, tc.expectedErr, err)
			assert.Nil(t, mock.ExpectationsWereMet())
		})
	}
}

func TestMerchantRepository_GetByStatus(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	r := repository.NewMerchantRepository(db)

	ctx := context.Background()
	merchantID := "123"
	status := "received"

	// Common setup for date-time fields
	timestamp := time.Now()

	tests := []struct {
		name           string
		mockSetup      func()
		expectedResult []entities.MerchantOrder
		expectedErr    error
	}{
		{
			name: "successful data retrieval",
			mockSetup: func() {
				rows := sqlmock.NewRows([]string{"merchant_id", "order_id", "status", "last_modified_date", "created_date"}).
					AddRow(merchantID, 1, status, timestamp, timestamp).
					AddRow(merchantID, 2, status, timestamp, timestamp)
				mock.ExpectQuery(`SELECT merchant_id, order_id, status, last_modified_date, created_date FROM merchant_order WHERE merchant_id = \$1 AND status = \$2`).
					WithArgs(merchantID, status).
					WillReturnRows(rows)
			},
			expectedResult: []entities.MerchantOrder{
				{MerchantID: merchantID, OrderID: "1", Status: status, LastModifiedDate: timestamp, CreatedDate: timestamp},
				{MerchantID: merchantID, OrderID: "2", Status: status, LastModifiedDate: timestamp, CreatedDate: timestamp},
			},
			expectedErr: nil,
		},
		{
			name: "no data found",
			mockSetup: func() {
				rows := sqlmock.NewRows([]string{"merchant_id", "order_id", "status", "last_modified_date", "created_date"})
				mock.ExpectQuery(`SELECT merchant_id, order_id, status, last_modified_date, created_date FROM merchant_order WHERE merchant_id = \$1 AND status = \$2`).
					WithArgs(merchantID, status).
					WillReturnRows(rows)
			},
			expectedResult: nil,
			expectedErr:    nil,
		},
		{
			name: "database error on query",
			mockSetup: func() {
				mock.ExpectQuery(`SELECT merchant_id, order_id, status, last_modified_date, created_date FROM merchant_order WHERE merchant_id = \$1 AND status = \$2`).
					WithArgs(merchantID, status).
					WillReturnError(sql.ErrConnDone)
			},
			expectedResult: nil,
			expectedErr:    sql.ErrConnDone,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockSetup()
			result, err := r.GetByStatus(ctx, merchantID, status)
			assert.Equal(t, tc.expectedResult, result)
			assert.Equal(t, tc.expectedErr, err)
			assert.Nil(t, mock.ExpectationsWereMet())
		})
	}
}
