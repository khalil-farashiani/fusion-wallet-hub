package mysql_transaction

import (
	"database/sql"
	"fmt"
	"github.com/khalil-farashiani/fusion-wallet-hub/pkg/errmsg"
	"github.com/khalil-farashiani/fusion-wallet-hub/pkg/richerror"
	"github.com/khalil-farashiani/fusion-wallet-hub/wallet/internal/transaction/entity"
	"strings"
)

func (db *DB) CreateTransaction(tx entity.Transaction) error {
	const op = "mysql.CreateTransaction"

	stmt, err := db.conn.Conn().Prepare("INSERT INTO transaction(type, account_id, amount, created_at) VALUES (?, ?, ?, ?)")
	if err != nil {
		return richerror.New(op).WithErr(err).
			WithMessage(errmsg.ErrorMsgSomethingWentWrong).WithKind(richerror.Unexpected)
	}
	if _, err := stmt.Exec(tx.Type, tx.AccountID, tx.Amount, tx.CreatedAt); err != nil {
		return richerror.New(op).WithErr(err).
			WithMessage(errmsg.ErrorMsgSomethingWentWrong).WithKind(richerror.Unexpected)
	}

	return nil
}

func (db *DB) GetUserTransactions(userID string, pg entity.Paginate) ([]entity.Transaction, error) {
	const op = "mysql.GetUserTransactions"

	var result []entity.Transaction
	if err := db.conn.Conn().QueryRow(buildGetUserTransactionQuery(userID, pg)).Scan(&result); err != nil {
		if err == sql.ErrNoRows {
			return nil, richerror.New(op).WithErr(err).
				WithMessage(errmsg.ErrorMsgNotFound).WithKind(richerror.NotFound)
		}

		// TODO - log unexpected error for better observability
		return nil, richerror.New(op).WithErr(err).
			WithMessage(errmsg.ErrorMsgCantScanQueryResult).WithKind(richerror.Unexpected)
	}
	return result, nil
}

func (db *DB) GetTransactions(pg entity.Paginate) ([]entity.Transaction, error) {
	const op = "mysql.GetTransactions"

	var result []entity.Transaction
	if err := db.conn.Conn().QueryRow(buildGetAllTransactionQuery(pg)).Scan(&result); err != nil {
		if err == sql.ErrNoRows {
			return nil, richerror.New(op).WithErr(err).
				WithMessage(errmsg.ErrorMsgNotFound).WithKind(richerror.NotFound)
		}

		// TODO - log unexpected error for better observability
		return nil, richerror.New(op).WithErr(err).
			WithMessage(errmsg.ErrorMsgCantScanQueryResult).WithKind(richerror.Unexpected)
	}
	return result, nil
}

func buildGetAllTransactionQuery(pg entity.Paginate) (string, []interface{}) {
	baseQuery := "SELECT id, type, amount, created_at FROM transactions"
	var conditions []string
	var args []interface{}

	// Cursor pagination
	if pg.Cursor > 0 {
		conditions = append(conditions, "id > ?")
		args = append(args, pg.Cursor)
	}

	// Date filters
	if !pg.AfterDateTime.IsZero() {
		conditions = append(conditions, "created_at > ?")
		args = append(args, pg.AfterDateTime)
	}
	if !pg.BeforeDateTime.IsZero() {
		conditions = append(conditions, "created_at < ?")
		args = append(args, pg.BeforeDateTime)
	}

	// Assembling the query
	if len(conditions) > 0 {
		baseQuery += " WHERE " + strings.Join(conditions, " AND ")
	}
	baseQuery += " ORDER BY id"
	baseQuery += fmt.Sprintf(" LIMIT %d", pg.Limit)

	return baseQuery, args
}

func buildGetUserTransactionQuery(userID string, pg entity.Paginate) (string, []interface{}) {
	baseQuery := "SELECT id, type, amount, created_at FROM transactions"
	var conditions []string
	var args []interface{}

	// AccountID is mandatory in your base query example
	conditions = append(conditions, "account_id = ?")
	args = append(args, userID)

	// Cursor pagination
	if pg.Cursor > 0 {
		conditions = append(conditions, "id > ?")
		args = append(args, pg.Cursor)
	}

	// Date filters
	if !pg.AfterDateTime.IsZero() {
		conditions = append(conditions, "created_at > ?")
		args = append(args, pg.AfterDateTime)
	}
	if !pg.BeforeDateTime.IsZero() {
		conditions = append(conditions, "created_at < ?")
		args = append(args, pg.BeforeDateTime)
	}

	// Assembling the query
	if len(conditions) > 0 {
		baseQuery += " WHERE " + strings.Join(conditions, " AND ")
	}
	baseQuery += " ORDER BY id"
	baseQuery += fmt.Sprintf(" LIMIT %d", pg.Limit)

	return baseQuery, args
}
