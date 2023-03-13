package db

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"user/rpc/model"
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
	sql := fmt.Sprintf("insert into %s(number, name, password, gender) values(?,?,?,?)", data.TableName())
	_, err := db.Conn.ExecCtx(ctx, sql, data.Id, data.Name, data.Password, data.Gender)
	if err != nil {
		fmt.Println("db.Conn insert error:", err)
		return err
	}
	return nil
}

func(db *DB) FindByName(ctx context.Context, data *model.User) error {
	sql := fmt.Sprintf("select * from %s where name = ?", data.TableName())
	fmt.Println("查询前的data:", data)
	err := db.Conn.QueryRowCtx(ctx, data, sql, data.Name)
	if err != nil {
		fmt.Println("db.Conn findbyname error:", err)
		return err
	}
	fmt.Println("查询后的data:", data)
	return nil
}