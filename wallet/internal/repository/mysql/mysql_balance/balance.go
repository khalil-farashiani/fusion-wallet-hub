package mysql_balance

import (
	"database/sql"
	"github.com/khalil-farashiani/fusion-wallet-hub/pkg/errmsg"
	"github.com/khalil-farashiani/fusion-wallet-hub/pkg/richerror"
	"github.com/khalil-farashiani/fusion-wallet-hub/wallet/internal/balance/entity"
)

func (d *DB) IncreaseBalance(userID string, newAmount uint64) error {
	const op = "mysql.IncreaseBalance"

	stmt, err := d.conn.Conn().Prepare("UPDATE balance SET amount = ? WHERE account_id = ?")
	if err != nil {
		return richerror.New(op).WithErr(err).
			WithMessage(errmsg.ErrorMsgSomethingWentWrong).WithKind(richerror.Unexpected)
	}
	if _, err = stmt.Exec(newAmount, userID); err != nil {
		if err == sql.ErrNoRows {
			return richerror.New(op).WithErr(err).
				WithMessage(errmsg.ErrorMsgNotFound).WithKind(richerror.NotFound)
		}
		return richerror.New(op).WithErr(err).
			WithMessage(errmsg.ErrorMsgCantScanQueryResult).WithKind(richerror.Unexpected)
	}
	return nil
}

func (d *DB) GetUserBalance(userID string) (entity.Balance, error) {
	const op = "mysql.GetUserBalance"

	var balance entity.Balance
	if err := d.conn.Conn().QueryRow("SELECT amount FROM balance WHERE account_id = ?", userID).Scan(&balance.Amount); err != nil {
		if err == sql.ErrNoRows {
			return entity.Balance{}, richerror.New(op).WithErr(err).
				WithMessage(errmsg.ErrorMsgNotFound).WithKind(richerror.NotFound)
		}

		// TODO - log unexpected error for better observability
		return entity.Balance{}, richerror.New(op).WithErr(err).
			WithMessage(errmsg.ErrorMsgCantScanQueryResult).WithKind(richerror.Unexpected)
	}

	return balance, nil
}
