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
	asRoleJobFieldNames          = builder.RawFieldNames(&AsRoleJob{})
	asRoleJobRows                = strings.Join(asRoleJobFieldNames, ",")
	asRoleJobRowsExpectAutoSet   = strings.Join(stringx.Remove(asRoleJobFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	asRoleJobRowsWithPlaceHolder = strings.Join(stringx.Remove(asRoleJobFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheAssetAsRoleJobIdPrefix = "cache:asset:asRoleJob:id:"
)

type (
	asRoleJobModel interface {
		Insert(ctx context.Context, data *AsRoleJob) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*AsRoleJob, error)
		Update(ctx context.Context, data *AsRoleJob) error
		Delete(ctx context.Context, id int64) error
	}

	defaultAsRoleJobModel struct {
		sqlc.CachedConn
		table string
	}

	AsRoleJob struct {
		RoleId int64 `db:"role_id"` // 角色id
		JobId  int64 `db:"job_id"`  // 岗位id
		Id     int64 `db:"id"`
	}
)

func newAsRoleJobModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultAsRoleJobModel {
	return &defaultAsRoleJobModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`as_role_job`",
	}
}

func (m *defaultAsRoleJobModel) Insert(ctx context.Context, data *AsRoleJob) (sql.Result, error) {
	assetAsRoleJobIdKey := fmt.Sprintf("%s%v", cacheAssetAsRoleJobIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?)", m.table, asRoleJobRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.RoleId, data.JobId)
	}, assetAsRoleJobIdKey)
	return ret, err
}

func (m *defaultAsRoleJobModel) FindOne(ctx context.Context, id int64) (*AsRoleJob, error) {
	assetAsRoleJobIdKey := fmt.Sprintf("%s%v", cacheAssetAsRoleJobIdPrefix, id)
	var resp AsRoleJob
	err := m.QueryRowCtx(ctx, &resp, assetAsRoleJobIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", asRoleJobRows, m.table)
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

func (m *defaultAsRoleJobModel) Update(ctx context.Context, data *AsRoleJob) error {
	assetAsRoleJobIdKey := fmt.Sprintf("%s%v", cacheAssetAsRoleJobIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, asRoleJobRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.RoleId, data.JobId, data.Id)
	}, assetAsRoleJobIdKey)
	return err
}

func (m *defaultAsRoleJobModel) Delete(ctx context.Context, id int64) error {
	assetAsRoleJobIdKey := fmt.Sprintf("%s%v", cacheAssetAsRoleJobIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, assetAsRoleJobIdKey)
	return err
}

func (m *defaultAsRoleJobModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheAssetAsRoleJobIdPrefix, primary)
}

func (m *defaultAsRoleJobModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", asRoleJobRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultAsRoleJobModel) tableName() string {
	return m.table
}
