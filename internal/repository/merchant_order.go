package repository

import (
	"context"
	"database/sql"

	"github.com/Pos-Tech-Challenge-48/delivery-merchant-api/internal/entities"
)

type MerchantRepository struct {
	db *sql.DB
}

func NewMerchantRepository(db *sql.DB) *MerchantRepository {
	return &MerchantRepository{
		db: db,
	}
}

func (r *MerchantRepository) Update(ctx context.Context, merchantOrder *entities.MerchantOrder) error {
	query := `
		UPDATE merchant_order
		SET status = $1,
			last_modified_date = NOW()
		WHERE merchant_id = $2 and order_id = $3
	`

	_, err := r.db.ExecContext(
		ctx,
		query,
		merchantOrder.Status,
		merchantOrder.MerchantID,
		merchantOrder.OrderID,
	)
	if err != nil {
		return err
	}
	return nil
}

func (r *MerchantRepository) Get(ctx context.Context, merchantID string, orderID string) (*entities.MerchantOrder, error) {
	query := `SELECT merchant_id, order_id,status, last_modified_date, created_date
	FROM merchant_order
	WHERE merchant_id = $1 and order_id = $2`

	row := r.db.QueryRowContext(ctx, query, merchantID, orderID)

	merchantOrder := &entities.MerchantOrder{}
	err := row.Scan(
		&merchantOrder.MerchantID,
		&merchantOrder.OrderID,
		&merchantOrder.Status,
		&merchantOrder.LastModifiedDate,
		&merchantOrder.CreatedDate,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return merchantOrder, nil
}

func (r *MerchantRepository) Insert(ctx context.Context, merchantOrder *entities.MerchantOrder) error {
	query := `
	INSERT INTO merchant_order (merchant_id, order_id, status, last_modified_date, created_date)
	VALUES ($1, $2, $3, NOW(), NOW())
`
	_, err := r.db.ExecContext(ctx, query,
		merchantOrder.MerchantID,
		merchantOrder.OrderID,
		merchantOrder.Status,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *MerchantRepository) GetByStatus(ctx context.Context, merchantID string, status string) ([]entities.MerchantOrder, error) {
	query := `
	SELECT merchant_id, order_id, status, last_modified_date, created_date
	FROM merchant_order
	WHERE merchant_id = $1 AND status = $2
`
	rows, err := r.db.QueryContext(ctx, query, merchantID, status)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []entities.MerchantOrder
	for rows.Next() {
		var merchantOrder entities.MerchantOrder
		if err := rows.Scan(
			&merchantOrder.MerchantID,
			&merchantOrder.OrderID,
			&merchantOrder.Status,
			&merchantOrder.LastModifiedDate,
			&merchantOrder.CreatedDate,
		); err != nil {
			return nil, err
		}
		orders = append(orders, merchantOrder)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return orders, nil
}
