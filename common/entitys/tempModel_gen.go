// Code generated by goctl. DO NOT EDIT!

package entitys

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	tempFieldNames          = builder.RawFieldNames(&Temp{})
	tempRows                = strings.Join(tempFieldNames, ",")
	tempRowsExpectAutoSet   = strings.Join(stringx.Remove(tempFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	tempRowsWithPlaceHolder = strings.Join(stringx.Remove(tempFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheAssetTempIdPrefix = "cache:asset:temp:id:"
)

type (
	tempModel interface {
		Insert(ctx context.Context, data *Temp) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Temp, error)
		Update(ctx context.Context, data *Temp) error
		Delete(ctx context.Context, id int64) error
	}

	defaultTempModel struct {
		sqlc.CachedConn
		table string
	}

	Temp struct {
		Tablename sql.NullString `db:"tablename"`
		LieName   sql.NullString `db:"LieName"`
		Id        int64          `db:"id"`
	}
)

func newTempModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultTempModel {
	return &defaultTempModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`temp`",
	}
}

func (m *defaultTempModel) Insert(ctx context.Context, data *Temp) (sql.Result, error) {
	assetTempIdKey := fmt.Sprintf("%s%v", cacheAssetTempIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?)", m.table, tempRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Tablename, data.LieName)
	}, assetTempIdKey)
	return ret, err
}

func (m *defaultTempModel) FindOne(ctx context.Context, id int64) (*Temp, error) {
	assetTempIdKey := fmt.Sprintf("%s%v", cacheAssetTempIdPrefix, id)
	var resp Temp
	err := m.QueryRowCtx(ctx, &resp, assetTempIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", tempRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultTempModel) Update(ctx context.Context, data *Temp) error {
	assetTempIdKey := fmt.Sprintf("%s%v", cacheAssetTempIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, tempRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.Tablename, data.LieName, data.Id)
	}, assetTempIdKey)
	return err
}

func (m *defaultTempModel) Delete(ctx context.Context, id int64) error {
	assetTempIdKey := fmt.Sprintf("%s%v", cacheAssetTempIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, assetTempIdKey)
	return err
}

func (m *defaultTempModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheAssetTempIdPrefix, primary)
}

func (m *defaultTempModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", tempRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultTempModel) tableName() string {
	return m.table
}
