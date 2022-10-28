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
	tSysUserFieldNames          = builder.RawFieldNames(&TSysUser{})
	tSysUserRows                = strings.Join(tSysUserFieldNames, ",")
	tSysUserRowsExpectAutoSet   = strings.Join(stringx.Remove(tSysUserFieldNames, "`create_time`", "`update_time`"), ",")
	tSysUserRowsWithPlaceHolder = strings.Join(stringx.Remove(tSysUserFieldNames, "`USER_ID`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheAssetTSysUserUSERIDPrefix = "cache:asset:tSysUser:uSERID:"
)

type (
	tSysUserModel interface {
		Insert(ctx context.Context, data *TSysUser) (sql.Result, error)
		FindOne(ctx context.Context, uSERID string) (*TSysUser, error)
		Update(ctx context.Context, data *TSysUser) error
		Delete(ctx context.Context, uSERID string) error
	}

	defaultTSysUserModel struct {
		sqlc.CachedConn
		table string
	}

	TSysUser struct {
		USERID     string         `db:"USER_ID"`     // 用户ID
		USERNAME   string         `db:"USER_NAME"`   // 用户姓名
		PASSWORD   string         `db:"PASSWORD"`    // 密码
		SALT       sql.NullString `db:"SALT"`        // 密码盐
		SEX        sql.NullString `db:"SEX"`         // 性别
		ROLEID     sql.NullString `db:"ROLE_ID"`     // 所属角色ID
		ORGID      sql.NullString `db:"ORG_ID"`      // 所属机构ID
		MOBILE     sql.NullString `db:"MOBILE"`      // 手机号码
		IDCARDNO   sql.NullString `db:"ID_CARD_NO"`  // 身份证号码
		EMAIL      sql.NullString `db:"EMAIL"`       // 邮箱
		STATUS     sql.NullString `db:"STATUS"`      // 用户状态
		SORTNO     sql.NullInt64  `db:"SORT_NO"`     // 排序号
		REMARK     sql.NullString `db:"REMARK"`      // 备注
		CREATEBY   sql.NullString `db:"CREATE_BY"`   // 创建人
		CREATEDATE sql.NullTime   `db:"CREATE_DATE"` // 创建日期
		CREATETIME sql.NullTime   `db:"CREATE_TIME"` // 创建时间
		UPDATEBY   sql.NullString `db:"UPDATE_BY"`   // 修改人
		UPDATEDATE sql.NullTime   `db:"UPDATE_DATE"` // 修改日期
		UPDATETIME sql.NullTime   `db:"UPDATE_TIME"` // 修改时间
	}
)

func newTSysUserModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultTSysUserModel {
	return &defaultTSysUserModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`t_sys_user`",
	}
}

func (m *defaultTSysUserModel) Insert(ctx context.Context, data *TSysUser) (sql.Result, error) {
	assetTSysUserUSERIDKey := fmt.Sprintf("%s%v", cacheAssetTSysUserUSERIDPrefix, data.USERID)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, tSysUserRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.USERID, data.USERNAME, data.PASSWORD, data.SALT, data.SEX, data.ROLEID, data.ORGID, data.MOBILE, data.IDCARDNO, data.EMAIL, data.STATUS, data.SORTNO, data.REMARK, data.CREATEBY, data.CREATEDATE, data.CREATETIME, data.UPDATEBY, data.UPDATEDATE, data.UPDATETIME)
	}, assetTSysUserUSERIDKey)
	return ret, err
}

func (m *defaultTSysUserModel) FindOne(ctx context.Context, uSERID string) (*TSysUser, error) {
	assetTSysUserUSERIDKey := fmt.Sprintf("%s%v", cacheAssetTSysUserUSERIDPrefix, uSERID)
	var resp TSysUser
	err := m.QueryRowCtx(ctx, &resp, assetTSysUserUSERIDKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `USER_ID` = ? limit 1", tSysUserRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, uSERID)
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

func (m *defaultTSysUserModel) Update(ctx context.Context, data *TSysUser) error {
	assetTSysUserUSERIDKey := fmt.Sprintf("%s%v", cacheAssetTSysUserUSERIDPrefix, data.USERID)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `USER_ID` = ?", m.table, tSysUserRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.USERNAME, data.PASSWORD, data.SALT, data.SEX, data.ROLEID, data.ORGID, data.MOBILE, data.IDCARDNO, data.EMAIL, data.STATUS, data.SORTNO, data.REMARK, data.CREATEBY, data.CREATEDATE, data.CREATETIME, data.UPDATEBY, data.UPDATEDATE, data.UPDATETIME, data.USERID)
	}, assetTSysUserUSERIDKey)
	return err
}

func (m *defaultTSysUserModel) Delete(ctx context.Context, uSERID string) error {
	assetTSysUserUSERIDKey := fmt.Sprintf("%s%v", cacheAssetTSysUserUSERIDPrefix, uSERID)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `USER_ID` = ?", m.table)
		return conn.ExecCtx(ctx, query, uSERID)
	}, assetTSysUserUSERIDKey)
	return err
}

func (m *defaultTSysUserModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheAssetTSysUserUSERIDPrefix, primary)
}

func (m *defaultTSysUserModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `USER_ID` = ? limit 1", tSysUserRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultTSysUserModel) tableName() string {
	return m.table
}