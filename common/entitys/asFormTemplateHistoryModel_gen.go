// Code generated by goctl. DO NOT EDIT!

package entitys

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	asFormTemplateHistoryFieldNames          = builder.RawFieldNames(&AsFormTemplateHistory{})
	asFormTemplateHistoryRows                = strings.Join(asFormTemplateHistoryFieldNames, ",")
	asFormTemplateHistoryRowsExpectAutoSet   = strings.Join(stringx.Remove(asFormTemplateHistoryFieldNames, "`create_time`", "`update_time`"), ",")
	asFormTemplateHistoryRowsWithPlaceHolder = strings.Join(stringx.Remove(asFormTemplateHistoryFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheAssetAsFormTemplateHistoryIdPrefix = "cache:asset:asFormTemplateHistory:id:"
)

type (
	asFormTemplateHistoryModel interface {
		Insert(ctx context.Context, data *AsFormTemplateHistory) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*AsFormTemplateHistory, error)
		Update(ctx context.Context, data *AsFormTemplateHistory) error
		Delete(ctx context.Context, id int64) error
	}

	defaultAsFormTemplateHistoryModel struct {
		sqlc.CachedConn
		table string
	}

	AsFormTemplateHistory struct {
		Id             int64          `db:"id"`               // 主键
		Name           string         `db:"name"`             // 模板名称
		IconCls        sql.NullString `db:"icon_cls"`         // 对应单据图标编码
		ModelSheetJson string         `db:"model_sheet_json"` // 单据的json源数据
		Description    sql.NullString `db:"description"`      // 描述信息
		ModelType      sql.NullInt64  `db:"model_type"`       // 单据模板分类
		Version        int64          `db:"version"`          // 模板版本
		TenantCode     sql.NullString `db:"tenant_code"`      // 模板所属租户的id
		CreateTime     time.Time      `db:"create_time"`      // 创建时间
		CreateUser     sql.NullString `db:"create_user"`      // 创建者id
		UpdateTime     time.Time      `db:"update_time"`      // 更新时间
		UpdateUser     sql.NullString `db:"update_user"`      // 更新用户id
		Status         int64          `db:"status"`           // 数据状态：0：不可用；1：可用；
		IsDeleted      int64          `db:"is_deleted"`       // 是否被删除：0：未删除；1：已删除
	}
)

func newAsFormTemplateHistoryModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultAsFormTemplateHistoryModel {
	return &defaultAsFormTemplateHistoryModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`as_form_template_history`",
	}
}

func (m *defaultAsFormTemplateHistoryModel) Insert(ctx context.Context, data *AsFormTemplateHistory) (sql.Result, error) {
	assetAsFormTemplateHistoryIdKey := fmt.Sprintf("%s%v", cacheAssetAsFormTemplateHistoryIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, asFormTemplateHistoryRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Id, data.Name, data.IconCls, data.ModelSheetJson, data.Description, data.ModelType, data.Version, data.TenantCode, data.CreateUser, data.UpdateUser, data.Status, data.IsDeleted)
	}, assetAsFormTemplateHistoryIdKey)
	return ret, err
}

func (m *defaultAsFormTemplateHistoryModel) FindOne(ctx context.Context, id int64) (*AsFormTemplateHistory, error) {
	assetAsFormTemplateHistoryIdKey := fmt.Sprintf("%s%v", cacheAssetAsFormTemplateHistoryIdPrefix, id)
	var resp AsFormTemplateHistory
	err := m.QueryRowCtx(ctx, &resp, assetAsFormTemplateHistoryIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", asFormTemplateHistoryRows, m.table)
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

func (m *defaultAsFormTemplateHistoryModel) Update(ctx context.Context, data *AsFormTemplateHistory) error {
	assetAsFormTemplateHistoryIdKey := fmt.Sprintf("%s%v", cacheAssetAsFormTemplateHistoryIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, asFormTemplateHistoryRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.Name, data.IconCls, data.ModelSheetJson, data.Description, data.ModelType, data.Version, data.TenantCode, data.CreateUser, data.UpdateUser, data.Status, data.IsDeleted, data.Id)
	}, assetAsFormTemplateHistoryIdKey)
	return err
}

func (m *defaultAsFormTemplateHistoryModel) Delete(ctx context.Context, id int64) error {
	assetAsFormTemplateHistoryIdKey := fmt.Sprintf("%s%v", cacheAssetAsFormTemplateHistoryIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, assetAsFormTemplateHistoryIdKey)
	return err
}

func (m *defaultAsFormTemplateHistoryModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheAssetAsFormTemplateHistoryIdPrefix, primary)
}

func (m *defaultAsFormTemplateHistoryModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", asFormTemplateHistoryRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultAsFormTemplateHistoryModel) tableName() string {
	return m.table
}
