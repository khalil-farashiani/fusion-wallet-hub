package mysql_balance

import (
	"database/sql"
	"github.com/khalil-farashiani/fusion-wallet-hub/pkg/errmsg"
	"github.com/khalil-farashiani/fusion-wallet-hub/pkg/richerror"
	"github.com/khalil-farashiani/fusion-wallet-hub/wallet/internal/balance/entity"
)

func (d *DB) GetUserBalance(userID string) (entity.Balance, error) {
	const op = "mysql.GetUserBalance"

	var balance entity.Balance
	if err := d.conn.Conn().QueryRow("SELECT balance FROM users WHERE user_id = ?", userID).Scan(&balance); err != nil {
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
