// Code generated by entc, DO NOT EDIT.

package schema

import (
	"fmt"
	"orginone/common/schema/asproperties"
	"orginone/common/tools/date"
	"strings"

	"entgo.io/ent/dialect/sql"
)

// AsProperties is the model entity for the AsProperties schema.
type AsProperties struct {
	config `json:"-"`
	// ID of the ent.
	// 主键
	ID int64 `json:"id,string"`
	// PropertiesName holds the value of the "properties_name" field.
	// 岗位名称
	PropertiesName string `json:"propertiesName"`
	// GroupID holds the value of the "group_id" field.
	// 租户id
	GroupID int64 `json:"groupId"`
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
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the AsPropertiesQuery when eager-loading is set.
	Edges AsPropertiesEdges `json:"edges"`
}

// AsPropertiesEdges holds the relations/edges for other nodes in the graph.
type AsPropertiesEdges struct {
	// AllTenants holds the value of the allTenants edge.
	AllTenants []*AsPropertiesDistribution `json:"alltenants"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// AllTenantsOrErr returns the AllTenants value or an error if the edge
// was not loaded in eager-loading.
func (e AsPropertiesEdges) AllTenantsOrErr() ([]*AsPropertiesDistribution, error) {
	if e.loadedTypes[0] {
		return e.AllTenants, nil
	}
	return nil, &NotLoadedError{edge: "allTenants"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*AsProperties) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case asproperties.FieldID, asproperties.FieldGroupID, asproperties.FieldIsDeleted, asproperties.FieldStatus, asproperties.FieldCreateUser, asproperties.FieldUpdateUser:
			values[i] = new(sql.NullInt64)
		case asproperties.FieldPropertiesName:
			values[i] = new(sql.NullString)
		case asproperties.FieldCreateTime, asproperties.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type AsProperties", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the AsProperties fields.
func (ap *AsProperties) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case asproperties.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ap.ID = int64(value.Int64)
		case asproperties.FieldPropertiesName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field properties_name", values[i])
			} else if value.Valid {
				ap.PropertiesName = value.String
			}
		case asproperties.FieldGroupID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field group_id", values[i])
			} else if value.Valid {
				ap.GroupID = value.Int64
			}
		case asproperties.FieldIsDeleted:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field is_deleted", values[i])
			} else if value.Valid {
				ap.IsDeleted = value.Int64
			}
		case asproperties.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				ap.Status = value.Int64
			}
		case asproperties.FieldCreateUser:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field create_user", values[i])
			} else if value.Valid {
				ap.CreateUser = value.Int64
			}
		case asproperties.FieldUpdateUser:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field update_user", values[i])
			} else if value.Valid {
				ap.UpdateUser = value.Int64
			}
		case asproperties.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				ap.CreateTime = date.DateTime(value.Time)
			}
		case asproperties.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				ap.UpdateTime = date.DateTime(value.Time)
			}
		}
	}
	return nil
}

// QueryAllTenants queries the "allTenants" edge of the AsProperties entity.
func (ap *AsProperties) QueryAllTenants() *AsPropertiesDistributionQuery {
	return (&AsPropertiesClient{config: ap.config}).QueryAllTenants(ap)
}

// Update returns a builder for updating this AsProperties.
// Note that you need to call AsProperties.Unwrap() before calling this method if this AsProperties
// was returned from a transaction, and the transaction was committed or rolled back.
func (ap *AsProperties) Update() *AsPropertiesUpdateOne {
	return (&AsPropertiesClient{config: ap.config}).UpdateOne(ap)
}

// Unwrap unwraps the AsProperties entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ap *AsProperties) Unwrap() *AsProperties {
	tx, ok := ap.config.driver.(*txDriver)
	if !ok {
		panic("schema: AsProperties is not a transactional entity")
	}
	ap.config.driver = tx.drv
	return ap
}

// String implements the fmt.Stringer.
func (ap *AsProperties) String() string {
	var builder strings.Builder
	builder.WriteString("AsProperties(")
	builder.WriteString(fmt.Sprintf("id=%v", ap.ID))
	builder.WriteString(", properties_name=")
	builder.WriteString(ap.PropertiesName)
	builder.WriteString(", group_id=")
	builder.WriteString(fmt.Sprintf("%v", ap.GroupID))
	builder.WriteString(", is_deleted=")
	builder.WriteString(fmt.Sprintf("%v", ap.IsDeleted))
	builder.WriteString(", status=")
	builder.WriteString(fmt.Sprintf("%v", ap.Status))
	builder.WriteString(", create_user=")
	builder.WriteString(fmt.Sprintf("%v", ap.CreateUser))
	builder.WriteString(", update_user=")
	builder.WriteString(fmt.Sprintf("%v", ap.UpdateUser))
	builder.WriteString(", create_time=")
	builder.WriteString(fmt.Sprintf("%v", ap.CreateTime))
	builder.WriteString(", update_time=")
	builder.WriteString(fmt.Sprintf("%v", ap.UpdateTime))
	builder.WriteByte(')')
	return builder.String()
}

// AsPropertiesSlice is a parsable slice of AsProperties.
type AsPropertiesSlice []*AsProperties

func (ap AsPropertiesSlice) config(cfg config) {
	for _i := range ap {
		ap[_i].config = cfg
	}
}
