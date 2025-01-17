// Code generated by entc, DO NOT EDIT.

package schema

import (
	"fmt"
	"orginone/common/schema/asinputdata"
	"orginone/common/tools/date"
	"strings"

	"entgo.io/ent/dialect/sql"
)

// AsInputData is the model entity for the AsInputData schema.
type AsInputData struct {
	config `json:"-"`
	// ID of the ent.
	// 主键
	ID int64 `json:"id,string"`
	// FileID holds the value of the "file_id" field.
	// 导入文件ID
	FileID int64 `json:"fileId"`
	// FileName holds the value of the "file_name" field.
	// 导入的文件名称
	FileName string `json:"fileName"`
	// TableName holds the value of the "table_name" field.
	// 表名称
	TableName string `json:"tableName"`
	// Type holds the value of the "type" field.
	// 导入类型 0-文件 1-数据
	Type int64 `json:"type"`
	// TCount holds the value of the "t_count" field.
	// 导入成功条数
	TCount int64 `json:"tCount"`
	// FCount holds the value of the "f_count" field.
	// 导入失败条数
	FCount int64 `json:"fCount"`
	// Context holds the value of the "context" field.
	// 导入失败原因
	Context string `json:"context"`
	// EndTime holds the value of the "end_time" field.
	// 结束时间
	EndTime date.DateTime `json:"endTime"`
	// TotalTime holds the value of the "total_time" field.
	// 总共耗时
	TotalTime int64 `json:"totalTime"`
	// TenantCode holds the value of the "tenant_code" field.
	// 所属租户ID
	TenantCode string `json:"tenantCode"`
	// ImportType holds the value of the "import_type" field.
	ImportType int64 `json:"importType"`
	// IsDeleted holds the value of the "is_deleted" field.
	// 是否被删除
	IsDeleted int64 `json:"isDeleted"`
	// Status holds the value of the "status" field.
	// 状态
	Status int64 `json:"status"`
	// CreateUser holds the value of the "create_user" field.
	// 创建用户
	CreateUser int64 `json:"createUser"`
	// UpdateUser holds the value of the "update_user" field.
	// 更新用户
	UpdateUser int64 `json:"updateUser"`
	// CreateTime holds the value of the "create_time" field.
	// 创建时间
	CreateTime date.DateTime `json:"createTime"`
	// UpdateTime holds the value of the "update_time" field.
	// 更新时间
	UpdateTime date.DateTime `json:"updateTime"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*AsInputData) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case asinputdata.FieldID, asinputdata.FieldFileID, asinputdata.FieldType, asinputdata.FieldTCount, asinputdata.FieldFCount, asinputdata.FieldTotalTime, asinputdata.FieldImportType, asinputdata.FieldIsDeleted, asinputdata.FieldStatus, asinputdata.FieldCreateUser, asinputdata.FieldUpdateUser:
			values[i] = new(sql.NullInt64)
		case asinputdata.FieldFileName, asinputdata.FieldTableName, asinputdata.FieldContext, asinputdata.FieldTenantCode:
			values[i] = new(sql.NullString)
		case asinputdata.FieldEndTime, asinputdata.FieldCreateTime, asinputdata.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type AsInputData", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the AsInputData fields.
func (aid *AsInputData) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case asinputdata.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			aid.ID = int64(value.Int64)
		case asinputdata.FieldFileID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field file_id", values[i])
			} else if value.Valid {
				aid.FileID = value.Int64
			}
		case asinputdata.FieldFileName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field file_name", values[i])
			} else if value.Valid {
				aid.FileName = value.String
			}
		case asinputdata.FieldTableName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field table_name", values[i])
			} else if value.Valid {
				aid.TableName = value.String
			}
		case asinputdata.FieldType:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				aid.Type = value.Int64
			}
		case asinputdata.FieldTCount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field t_count", values[i])
			} else if value.Valid {
				aid.TCount = value.Int64
			}
		case asinputdata.FieldFCount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field f_count", values[i])
			} else if value.Valid {
				aid.FCount = value.Int64
			}
		case asinputdata.FieldContext:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field context", values[i])
			} else if value.Valid {
				aid.Context = value.String
			}
		case asinputdata.FieldEndTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field end_time", values[i])
			} else if value.Valid {
				aid.EndTime = date.DateTime(value.Time)
			}
		case asinputdata.FieldTotalTime:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field total_time", values[i])
			} else if value.Valid {
				aid.TotalTime = value.Int64
			}
		case asinputdata.FieldTenantCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field tenant_code", values[i])
			} else if value.Valid {
				aid.TenantCode = value.String
			}
		case asinputdata.FieldImportType:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field import_type", values[i])
			} else if value.Valid {
				aid.ImportType = value.Int64
			}
		case asinputdata.FieldIsDeleted:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field is_deleted", values[i])
			} else if value.Valid {
				aid.IsDeleted = value.Int64
			}
		case asinputdata.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				aid.Status = value.Int64
			}
		case asinputdata.FieldCreateUser:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field create_user", values[i])
			} else if value.Valid {
				aid.CreateUser = value.Int64
			}
		case asinputdata.FieldUpdateUser:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field update_user", values[i])
			} else if value.Valid {
				aid.UpdateUser = value.Int64
			}
		case asinputdata.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				aid.CreateTime = date.DateTime(value.Time)
			}
		case asinputdata.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				aid.UpdateTime = date.DateTime(value.Time)
			}
		}
	}
	return nil
}

// Update returns a builder for updating this AsInputData.
// Note that you need to call AsInputData.Unwrap() before calling this method if this AsInputData
// was returned from a transaction, and the transaction was committed or rolled back.
func (aid *AsInputData) Update() *AsInputDataUpdateOne {
	return (&AsInputDataClient{config: aid.config}).UpdateOne(aid)
}

// Unwrap unwraps the AsInputData entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (aid *AsInputData) Unwrap() *AsInputData {
	tx, ok := aid.config.driver.(*txDriver)
	if !ok {
		panic("schema: AsInputData is not a transactional entity")
	}
	aid.config.driver = tx.drv
	return aid
}

// String implements the fmt.Stringer.
func (aid *AsInputData) String() string {
	var builder strings.Builder
	builder.WriteString("AsInputData(")
	builder.WriteString(fmt.Sprintf("id=%v", aid.ID))
	builder.WriteString(", file_id=")
	builder.WriteString(fmt.Sprintf("%v", aid.FileID))
	builder.WriteString(", file_name=")
	builder.WriteString(aid.FileName)
	builder.WriteString(", table_name=")
	builder.WriteString(aid.TableName)
	builder.WriteString(", type=")
	builder.WriteString(fmt.Sprintf("%v", aid.Type))
	builder.WriteString(", t_count=")
	builder.WriteString(fmt.Sprintf("%v", aid.TCount))
	builder.WriteString(", f_count=")
	builder.WriteString(fmt.Sprintf("%v", aid.FCount))
	builder.WriteString(", context=")
	builder.WriteString(aid.Context)
	builder.WriteString(", end_time=")
	builder.WriteString(fmt.Sprintf("%v", aid.EndTime))
	builder.WriteString(", total_time=")
	builder.WriteString(fmt.Sprintf("%v", aid.TotalTime))
	builder.WriteString(", tenant_code=")
	builder.WriteString(aid.TenantCode)
	builder.WriteString(", import_type=")
	builder.WriteString(fmt.Sprintf("%v", aid.ImportType))
	builder.WriteString(", is_deleted=")
	builder.WriteString(fmt.Sprintf("%v", aid.IsDeleted))
	builder.WriteString(", status=")
	builder.WriteString(fmt.Sprintf("%v", aid.Status))
	builder.WriteString(", create_user=")
	builder.WriteString(fmt.Sprintf("%v", aid.CreateUser))
	builder.WriteString(", update_user=")
	builder.WriteString(fmt.Sprintf("%v", aid.UpdateUser))
	builder.WriteString(", create_time=")
	builder.WriteString(fmt.Sprintf("%v", aid.CreateTime))
	builder.WriteString(", update_time=")
	builder.WriteString(fmt.Sprintf("%v", aid.UpdateTime))
	builder.WriteByte(')')
	return builder.String()
}

// AsInputDataSlice is a parsable slice of AsInputData.
type AsInputDataSlice []*AsInputData

func (aid AsInputDataSlice) config(cfg config) {
	for _i := range aid {
		aid[_i].config = cfg
	}
}
