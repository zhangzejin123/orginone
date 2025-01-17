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
	actCmmnRuPlanItemInstFieldNames          = builder.RawFieldNames(&ActCmmnRuPlanItemInst{})
	actCmmnRuPlanItemInstRows                = strings.Join(actCmmnRuPlanItemInstFieldNames, ",")
	actCmmnRuPlanItemInstRowsExpectAutoSet   = strings.Join(stringx.Remove(actCmmnRuPlanItemInstFieldNames, "`create_time`", "`update_time`"), ",")
	actCmmnRuPlanItemInstRowsWithPlaceHolder = strings.Join(stringx.Remove(actCmmnRuPlanItemInstFieldNames, "`ID_`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheAssetActCmmnRuPlanItemInstIDPrefix = "cache:asset:actCmmnRuPlanItemInst:iD:"
)

type (
	actCmmnRuPlanItemInstModel interface {
		Insert(ctx context.Context, data *ActCmmnRuPlanItemInst) (sql.Result, error)
		FindOne(ctx context.Context, iD string) (*ActCmmnRuPlanItemInst, error)
		Update(ctx context.Context, data *ActCmmnRuPlanItemInst) error
		Delete(ctx context.Context, iD string) error
	}

	defaultActCmmnRuPlanItemInstModel struct {
		sqlc.CachedConn
		table string
	}

	ActCmmnRuPlanItemInst struct {
		ID                  string         `db:"ID_"`
		REV                 int64          `db:"REV_"`
		CASEDEFID           sql.NullString `db:"CASE_DEF_ID_"`
		CASEINSTID          sql.NullString `db:"CASE_INST_ID_"`
		STAGEINSTID         sql.NullString `db:"STAGE_INST_ID_"`
		ISSTAGE             byte           `db:"IS_STAGE_"`
		ELEMENTID           sql.NullString `db:"ELEMENT_ID_"`
		NAME                sql.NullString `db:"NAME_"`
		STATE               sql.NullString `db:"STATE_"`
		STARTTIME           sql.NullTime   `db:"START_TIME_"`
		STARTUSERID         sql.NullString `db:"START_USER_ID_"`
		REFERENCEID         sql.NullString `db:"REFERENCE_ID_"`
		REFERENCETYPE       sql.NullString `db:"REFERENCE_TYPE_"`
		TENANTID            string         `db:"TENANT_ID_"`
		ITEMDEFINITIONID    sql.NullString `db:"ITEM_DEFINITION_ID_"`
		ITEMDEFINITIONTYPE  sql.NullString `db:"ITEM_DEFINITION_TYPE_"`
		ISCOMPLETEABLE      byte           `db:"IS_COMPLETEABLE_"`
		ISCOUNTENABLED      byte           `db:"IS_COUNT_ENABLED_"`
		VARCOUNT            sql.NullInt64  `db:"VAR_COUNT_"`
		SENTRYPARTINSTCOUNT sql.NullInt64  `db:"SENTRY_PART_INST_COUNT_"`
	}
)

func newActCmmnRuPlanItemInstModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultActCmmnRuPlanItemInstModel {
	return &defaultActCmmnRuPlanItemInstModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`act_cmmn_ru_plan_item_inst`",
	}
}

func (m *defaultActCmmnRuPlanItemInstModel) Insert(ctx context.Context, data *ActCmmnRuPlanItemInst) (sql.Result, error) {
	assetActCmmnRuPlanItemInstIDKey := fmt.Sprintf("%s%v", cacheAssetActCmmnRuPlanItemInstIDPrefix, data.ID)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, actCmmnRuPlanItemInstRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.ID, data.REV, data.CASEDEFID, data.CASEINSTID, data.STAGEINSTID, data.ISSTAGE, data.ELEMENTID, data.NAME, data.STATE, data.STARTTIME, data.STARTUSERID, data.REFERENCEID, data.REFERENCETYPE, data.TENANTID, data.ITEMDEFINITIONID, data.ITEMDEFINITIONTYPE, data.ISCOMPLETEABLE, data.ISCOUNTENABLED, data.VARCOUNT, data.SENTRYPARTINSTCOUNT)
	}, assetActCmmnRuPlanItemInstIDKey)
	return ret, err
}

func (m *defaultActCmmnRuPlanItemInstModel) FindOne(ctx context.Context, iD string) (*ActCmmnRuPlanItemInst, error) {
	assetActCmmnRuPlanItemInstIDKey := fmt.Sprintf("%s%v", cacheAssetActCmmnRuPlanItemInstIDPrefix, iD)
	var resp ActCmmnRuPlanItemInst
	err := m.QueryRowCtx(ctx, &resp, assetActCmmnRuPlanItemInstIDKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `ID_` = ? limit 1", actCmmnRuPlanItemInstRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, iD)
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

func (m *defaultActCmmnRuPlanItemInstModel) Update(ctx context.Context, data *ActCmmnRuPlanItemInst) error {
	assetActCmmnRuPlanItemInstIDKey := fmt.Sprintf("%s%v", cacheAssetActCmmnRuPlanItemInstIDPrefix, data.ID)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `ID_` = ?", m.table, actCmmnRuPlanItemInstRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.REV, data.CASEDEFID, data.CASEINSTID, data.STAGEINSTID, data.ISSTAGE, data.ELEMENTID, data.NAME, data.STATE, data.STARTTIME, data.STARTUSERID, data.REFERENCEID, data.REFERENCETYPE, data.TENANTID, data.ITEMDEFINITIONID, data.ITEMDEFINITIONTYPE, data.ISCOMPLETEABLE, data.ISCOUNTENABLED, data.VARCOUNT, data.SENTRYPARTINSTCOUNT, data.ID)
	}, assetActCmmnRuPlanItemInstIDKey)
	return err
}

func (m *defaultActCmmnRuPlanItemInstModel) Delete(ctx context.Context, iD string) error {
	assetActCmmnRuPlanItemInstIDKey := fmt.Sprintf("%s%v", cacheAssetActCmmnRuPlanItemInstIDPrefix, iD)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `ID_` = ?", m.table)
		return conn.ExecCtx(ctx, query, iD)
	}, assetActCmmnRuPlanItemInstIDKey)
	return err
}

func (m *defaultActCmmnRuPlanItemInstModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheAssetActCmmnRuPlanItemInstIDPrefix, primary)
}

func (m *defaultActCmmnRuPlanItemInstModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `ID_` = ? limit 1", actCmmnRuPlanItemInstRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultActCmmnRuPlanItemInstModel) tableName() string {
	return m.table
}
