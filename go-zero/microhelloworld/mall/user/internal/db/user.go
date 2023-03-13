package db

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"user/internal/model"
	"fmt"
	"context"
)

type DB struct {
	Conn sqlx.SqlConn
}

func ConnectMysql(datasource string) *DB {
	sqlconn := sqlx.NewMysql(datasource)
	return &DB{
		Conn : sqlconn,
	}
}

func(db *DB) Insert(ctx context.Context, data *model.User) error {
	sql := fmt.Sprintf("insert to %s(number, name, password, gender) values(?,?,?,?)", data.TableName())
	_, err := db.Conn.ExecCtx(ctx, sql, data.Id, data.Name, data.Password, data.Gender)
	if err != nil {
		fmt.Println("db.Conn insert error:", err)
		return err
	}
	return nil
}

func(db *DB) Find(ctx context.Context, data *model.User)