package mysql_transaction

import (
	"database/sql"
	"fmt"
	"github.com/khalil-farashiani/fusion-wallet-hub/pkg/errmsg"
	"github.com/khalil-farashiani/fusion-wallet-hub/pkg/richerror"
	entity2 "github.com/khalil-farashiani/fusion-wallet-hub/wallet/internal/entity"
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

func (db *DB) GetUserTransactions(userID string, pg entity2.Paginate) ([]entity.Transaction, error) {
	const op = "mysql.GetUserTransactions"

	var result []entity.Transaction
	query, args := buildGetUserTransactionQuery(userID, pg)
	rows, err := db.conn.Conn().Query(query, args...)
	defer rows.Close()
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, richerror.New(op).WithErr(err).
				WithMessage(errmsg.ErrorMsgNotFound).WithKind(richerror.NotFound)
		}

		// TODO - log unexpected error for better observability
		return nil, richerror.New(op).WithErr(err).
			WithMessage(errmsg.ErrorMsgCantScanQueryResult).WithKind(richerror.Unexpected)
	}

	for rows.Next() {
		var tx entity.Transaction
		if err := rows.Scan(&tx.ID, &tx.Type, &tx.Amount, &tx.CreatedAt); err != nil {
			return nil, richerror.New(op).WithErr(err).WithMessage(errmsg.ErrorMsgSomethingWentWrong).WithKind(richerror.Unexpected)
		}
		result = append(result, tx)
	}
	return result, nil
	return result, nil
}

func (db *DB) GetTransactions(pg entity2.Paginate) ([]entity.Transaction, error) {
	const op = "mysql.GetTransactions"

	var result []entity.Transaction
	query, args := buildGetAllTransactionQuery(pg)
	rows, err := db.conn.Conn().Query(query, args...)
	defer rows.Close()

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, richerror.New(op).WithErr(err).
				WithMessage(errmsg.ErrorMsgNotFound).WithKind(richerror.NotFound)
		}

		// TODO - log unexpected error for better observability
		return nil, richerror.New(op).WithErr(err).
			WithMessage(errmsg.ErrorMsgCantScanQueryResult).WithKind(richerror.Unexpected)
	}

	for rows.Next() {
		var tx entity.Transaction
		if err := rows.Scan(&tx.ID, &tx.Type, &tx.Amount, &tx.AccountID, &tx.CreatedAt); err != nil {
			return nil, richerror.New(op).WithErr(err).WithMessage(errmsg.ErrorMsgSomethingWentWrong).WithKind(richerror.Unexpected)
		}
		result = append(result, tx)
	}
	return result, nil
}

func buildGetAllTransactionQuery(pg entity2.Paginate) (string, []interface{}) {
	baseQuery := "SELECT id, type, amount, account_id, created_at FROM transaction"
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

func buildGetUserTransactionQuery(userID string, pg entity2.Paginate) (string, []interface{}) {
	baseQuery := "SELECT id, type, amount, created_at FROM transaction"
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
