package mysql_redeem

import (
	"database/sql"
	"github.com/khalil-farashiani/fusion-wallet-hub/discount/internal/entity"
	"github.com/khalil-farashiani/fusion-wallet-hub/pkg/errmsg"
	"github.com/khalil-farashiani/fusion-wallet-hub/pkg/richerror"
	"time"
)

func (db *DB) GetRedeem(title string) (entity.Redeem, error) {
	const op = "mysql.GetRedeem"

	var result entity.Redeem
	if err := db.conn.Conn().QueryRow("SELECT amount, created_at FROM redeem WHERE title = ?", title).
		Scan(&result.Amount, &result.CreatedAt); err != nil {
		if err == sql.ErrNoRows {
			return entity.Redeem{}, richerror.New(op).WithErr(err).
				WithMessage(errmsg.ErrorMsgNotFound).WithKind(richerror.NotFound)
		}
		return entity.Redeem{}, richerror.New(op).WithErr(err).
			WithMessage(errmsg.ErrorMsgSomethingWentWrong).WithKind(richerror.Unexpected)
	}
	return result, nil
}

func (db *DB) GetReports(status string) ([]entity.RedeemReport, error) {
	const op = "mysql.GetReports"

	var results []entity.RedeemReport
	rows, err := db.conn.Conn().Query("SELECT id, user_id, amount, created_at FROM redeem_report WHERE status = ?", status)
	if err != nil {
		return nil, richerror.New(op).WithErr(err).WithMessage(errmsg.ErrorMsgSomethingWentWrong).WithKind(richerror.Unexpected)
	}
	defer rows.Close()

	for rows.Next() {
		var report entity.RedeemReport
		if err := rows.Scan(&report.ID, &report.UserId, &report.Amount, &report.CreatedAt); err != nil {
			return nil, richerror.New(op).WithErr(err).WithMessage(errmsg.ErrorMsgSomethingWentWrong).WithKind(richerror.Unexpected)
		}
		results = append(results, report)
	}

	if err := rows.Err(); err != nil {
		return nil, richerror.New(op).WithErr(err).WithMessage(errmsg.ErrorMsgSomethingWentWrong).WithKind(richerror.Unexpected)
	}

	return results, nil
}
func (db *DB) CreateUserRedeemRecord(redeem entity.RedeemReport) error {
	const op = "mysql.CreateUserRedeemRecord"

	stmt, err := db.conn.Conn().Prepare("INSERT INTO redeem_report (title,user_id, amount, status, created_at) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return richerror.New(op).WithErr(err).
			WithMessage(errmsg.ErrorMsgSomethingWentWrong).WithKind(richerror.Unexpected)
	}
	if _, err = stmt.Exec(redeem.Title, redeem.UserId, redeem.Amount, redeem.Status, time.Now()); err != nil {
		return richerror.New(op).WithErr(err).
			WithMessage(errmsg.ErrorMsgSomethingWentWrong).WithKind(richerror.Unexpected)
	}
	return nil
}

func (db *DB) IsUserUseRedeemBefore(redeem entity.RedeemReport) (bool, error) {
	const op = "mysql.SetUserRedeemRecord"
	var isUserUseRedeemBefore bool
	err := db.conn.Conn().QueryRow("SELECT IF(COUNT(*),'true','false') FROM redeem_report WHERE user_id = ? AND status = ? AND title = ?",
		redeem.UserId,
		redeem.Status,
		redeem.Title).Scan(&isUserUseRedeemBefore)
	if err != nil {
		return false, richerror.New(op).WithErr(err).
			WithMessage(errmsg.ErrorMsgSomethingWentWrong).WithKind(richerror.Unexpected)
	}
	return isUserUseRedeemBefore, nil
}

func (db *DB) CreateRedeem(redeem entity.Redeem) error {
	const op = "mysql.CreateRedeem"

	stmt, err := db.conn.Conn().Prepare("INSERT INTO redeem (title, amount, quantity, created_at) VALUES (?, ?, ?, ?)")
	if err != nil {
		return richerror.New(op).WithErr(err).
			WithMessage(errmsg.ErrorMsgSomethingWentWrong).WithKind(richerror.Unexpected)
	}
	if _, err = stmt.Exec(redeem.Title, redeem.Amount, redeem.Quantity, redeem.CreatedAt); err != nil {
		return richerror.New(op).WithErr(err).
			WithMessage(errmsg.ErrorMsgSomethingWentWrong).WithKind(richerror.Unexpected)
	}
	return nil
}
