// Code generated by entc, DO NOT EDIT.

package asmarketappusertemplate

import (
	"orginone/common/schema/predicate"
	"orginone/common/tools/date"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.AsMarketAppUserTemplate {
	return predicate.AsMarketAppUserTemplate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.AsMarketAppUserTemplate {
	return predicate.AsMarketAppUserTemplate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.AsMarketAppUserTemplate {
	return predicate.AsMarketAppUserTemplate(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.AsMarketAppUserTemplate {
	return predicate.AsMarketAppUserTemplate(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.AsMarketAppUserTemplate {
	return predicate.AsMarketAppUserTemplate(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.AsMarketAppUserTemplate {
	return predicate.AsMarketAppUserTemplate(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.AsMarketAppUserTemplate {
	return predicate.AsMarketAppUserTemplate(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.AsMarketAppUserTemplate {
	return predicate.AsMarketAppUserTemplate(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.AsMarketAppUserTemplate {
	return predicate.AsMarketAppUserTemplate(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v int64) predicate.AsMarketAppUserTemplate {
	return predicate.AsMarketAppUserTemplate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserID), v))
	})
}

// TemplateID applies equality check predicate on the "template_id" field. It's identical to TemplateIDEQ.
func TemplateID(v int64) predicate.AsMarketAppUserTemplate {
	return predicate.AsMarketAppUserTemplate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTemplateID), v))
	})
}

// UseStatus applies equality check predicate on the "use_status" field. It's identical to UseStatusEQ.
func UseStatus(v int64) predicate.AsMarketAppUserTemplate {
	return predicate.AsMarketAppUserTemplate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUseStatus), v))
	})
}

// IsDeleted applies equality check predicate on the "is_deleted" field. It's identical to IsDeletedEQ.
func IsDeleted(v int64) predicate.AsMarketAppUserTemplate {
	return predicate.AsMarketAppUserTemplate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIsDeleted), v))
	})
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v int64) predicate.AsMarketAppUserTemplate {
	return predicate.AsMarketAppUserTemplate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStatus), v))
	})
}

// CreateUser applies equality check predicate on the "create_user" field. It's identical to CreateUserEQ.
func CreateUser(v int64) predicate.AsMarketAppUserTemplate {
	return predicate.AsMarketAppUserTemplate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateUser), v))
	})
}

// UpdateUser applies equality check predicate on the "update_user" field. It's identical to UpdateUserEQ.
func UpdateUser(v int64) predicate.AsMarketAppUserTemplate {
	return predicate.AsMarketAppUserTemplate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateUser), v))
	})
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v date.DateTime) predicate.AsMarketAppUserTemplate {
	vc := time.Time(v)
	return predicate.AsMarketAppUserTemplate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateTime), vc))
	})
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v date.DateTime) predicate.AsMarketAppUserTemplate {
	vc := time.Time(v)
	return predicate.AsMarketAppUserTemplate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateTime), vc))
	})
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v int64) predicate.AsMarketAppUserTemplate {
	return predicate.AsMarketAppUserTemplate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserID), v))
	})
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v int64) predicate.AsMarketAppUserTemplate {
	return predicate.AsMarketAppUserTemplate(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUserID), v))
	})
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...int64) predicate.AsMarketAppUserTemplate {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AsMarketAppUserTemplate(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUserID), v...))
	})
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...int64) predicate.AsMarketAppUserTemplate {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AsMarketAppUserTemplate(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUserID), v...))
	})
}

// UserIDIsNil applies the IsNil predicate on the "user_id" field.
func UserIDIsNil() predicate.AsMarketAppUserTemplate {
	return predicate.AsMarketAppUserTemplate(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldUserID)))
	})
}

// UserIDNotNil applies the NotNil predicate on the "user_id" field.
func UserIDNotNil() predicate.AsMarketAppUserTemplate {
	return predicate.AsMarketAppUserTemplate(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldUserID)))
	})
}

// TemplateIDEQ applies the EQ predicate on the "template_id" field.
func TemplateIDEQ(v int64) predicate.AsMarketAppUserTemplate {
	return predicate.AsMarketAppUserTemplate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTemplateID), v))
	})
}

// TemplateIDNEQ applies the NEQ predicate on the "template_id" field.
func TemplateIDNEQ(v int64) predicate.AsMarketAppUserTemplate {
	return predicate.AsMarketAppUserTemplate(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTemplateID), v))
	})
}

// TemplateIDIn applies the In predicate on the "template_id" field.
func TemplateIDIn(vs ...int64) predicate.AsMarketAppUserTemplate {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AsMarketAppUserTemplate(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTemplateID), v...))
	})
}

// TemplateIDNotIn applies the NotIn predicate on the "template_id" field.
func TemplateIDNotIn(vs ...int64) predicate.AsMarketAppUserTemplate {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AsMarketAppUserTemplate(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTemplateID), v...))
	})
}

// TemplateIDIsNil applies the IsNil predicate on the "template_id" field.
func TemplateIDIsNil() predicate.AsMarketAppUserTemplate {
	return predicate.AsMarketAppUserTemplate(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldTemplateID)))
	})
}

// TemplateIDNotNil applies the NotNil predicate on the "template_id" field.
func TemplateIDNotNil() predicate.AsMarketAppUserTemplate {
	return predicate.AsMarketAppUserTemplate(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldTemplateID)))
	})
}

// UseStatusEQ applies the EQ predicate on the "use_status" field.
func UseStatusEQ(v int64) predicate.AsMarketAppUserTemplate {
	return predicate.AsMarketAppUserTemplate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUseStatus), v))
	})
}

// UseStatusNEQ applies the NEQ predicate on the "use_status" field.
func UseStatusNEQ(v int64) predicate.AsMarketAppUserTemplate {
	return predicate.AsMarketAppUserTemplate(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUseStatus), v))
	})
}

// UseStatusIn applies the In predicate on the "use_status" field.
func UseStatusIn(vs ...int64) predicate.AsMarketAppUserTemplate {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AsMarketAppUserTemplate(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUseStatus), v...))
	})
}

// UseStatusNotIn applies the NotIn predicate on the "use_status" field.
func UseStatusNotIn(vs ...int64) predicate.AsMarketAppUserTemplate {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AsMarketAppUserTemplate(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUseStatus), v...))
	})
}

// UseStatusGT applies the GT predicate on the "use_status" field.
func UseStatusGT(v int64) predicate.AsMarketAppUserTemplate {
	return predicate.AsMarketAppUserTemplate(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUseStatus), v))
	})
}

// UseStatusGTE applies the GTE predicate on the "use_status" field.
func UseStatusGTE(v int64) predicate.AsMarketAppUserTemplate {
	return predicate.AsMarketAppUserTemplate(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUseStatus), v))
	})
}

// UseStatusLT applies the LT predicate on the "use_status" field.
func UseStatusLT(v int64) predicate.AsMarketAppUserTemplate {
	return predicate.AsMarketAppUserTemplate(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUseStatus), v))
	})
}

// UseStatusLTE applies the LTE predicate on the "use_status" field.
func UseStatusLTE(v int64) predicate.AsMarketAppUserTemplate {
	return predicate.AsMarketAppUserTemplate(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUseStatus), v))
	})
}

// UseStatusIsNil applies the IsNil predicate on the "use_status" field.
func UseStatusIsNil() predicate.AsMarketAppUserTemplate {
	return predicate.AsMarketAppUserTemplate(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldUseStatus)))
	})
}

// UseStatusNotNil applies the NotNil predicate on the "use_status" field.
func UseStatusNotNil() predicate.AsMarketAppUserTemplate {
	return predicate.AsMarketAppUserTemplate(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldUseStatus)))
	})
}

// IsDeletedEQ applies the EQ predicate on the "is_deleted" field.
func IsDeletedEQ(v int64) predicate.AsMarketAppUserTemplate {
	return predicate.AsMarketAppUserTemplate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIsDeleted), v))
	})
}

// IsDeletedNEQ applies the NEQ predicate on the "is_deleted" field.
func IsDeletedNEQ(v int64) predicate.AsMarketAppUserTemplate {
	return predicate.AsMarketAppUserTemplate(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldIsDeleted), v))
	})
}

// IsDeletedIn applies the In predicate on the "is_deleted" field.
func IsDeletedIn(vs ...int64) predicate.AsMarketAppUserTemplate {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AsMarketAppUserTemplate(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldIsDeleted), v...))
	})
}

// IsDeletedNotIn applies the NotIn predicate on the "is_deleted" field.
func IsDeletedNotIn(vs ...int64) predicate.AsMarketAppUserTemplate {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AsMarketAppUserTemplate(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldIsDeleted), v...))
	})
}

// IsDeletedGT applies the GT predicate on the "is_deleted" field.
func IsDeletedGT(v int64) predicate.AsMarketAppUserTemplate {
	return predicate.AsMarketAppUserTemplate(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldIsDeleted), v))
	})
}

// IsDeletedGTE applies the GTE predicate on the "is_deleted" field.
func IsDeletedGTE(v int64) predicate.AsMarketAppUserTemplate {
	return predicate.AsMarketAppUserTemplate(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldIsDeleted), v))
	})
}

// IsDeletedLT applies the LT predicate on the "is_deleted" field.
func IsDeletedLT(v int64) predicate.AsMarketAppUserTemplate {
	return predicate.AsMarketAppUserTemplate(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldIsDeleted), v))
	})
}

// IsDeletedLTE applies the LTE predicate on the "is_deleted" field.
func IsDeletedLTE(v int64) predicate.AsMarketAppUserTemplate {
	return predicate.AsMarketAppUserTemplate(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldIsDeleted), v))
	})
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v int64) predicate.AsMarketAppUserTemplate {
	return predicate.AsMarketAppUserTemplate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStatus), v))
	})
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v int64) predicate.AsMarketAppUserTemplate {
	return predicate.AsMarketAppUserTemplate(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStatus), v))
	})
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...int64) predicate.AsMarketAppUserTemplate {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AsMarketAppUserTemplate(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldStatus), v...))
	})
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...int64) predicate.AsMarketAppUserTemplate {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AsMarketAppUserTemplate(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldStatus), v...))
	})
}

// StatusGT applies the GT predicate on the "status" field.
func StatusGT(v int64) predicate.AsMarketAppUserTemplate {
	return predicate.AsMarketAppUserTemplate(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldStatus), v))
	})
}

// StatusGTE applies the GTE predicate on the "status" field.
func StatusGTE(v int64) predicate.AsMarketAppUserTemplate {
	return predicate.AsMarketAppUserTemplate(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldStatus), v))
	})
}

// StatusLT applies the LT predicate on the "status" field.
func StatusLT(v int64) predicate.AsMarketAppUserTemplate {
	return predicate.AsMarketAppUserTemplate(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldStatus), v))
	})
}

// StatusLTE applies the LTE predicate on the "status" field.
func StatusLTE(v int64) predicate.AsMarketAppUserTemplate {
	return predicate.AsMarketAppUserTemplate(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldStatus), v))
	})
}

// StatusIsNil applies the IsNil predicate on the "status" field.
func StatusIsNil() predicate.AsMarketAppUserTemplate {
	return predicate.AsMarketAppUserTemplate(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldStatus)))
	})
}

// StatusNotNil applies the NotNil predicate on the "status" field.
func StatusNotNil() predicate.AsMarketAppUserTemplate {
	return predicate.AsMarketAppUserTemplate(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldStatus)))
	})
}

// CreateUserEQ applies the EQ predicate on the "create_user" field.
func CreateUserEQ(v int64) predicate.AsMarketAppUserTemplate {
	return predicate.AsMarketAppUserTemplate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateUser), v))
	})
}

// CreateUserNEQ applies the NEQ predicate on the "create_user" field.
func CreateUserNEQ(v int64) predicate.AsMarketAppUserTemplate {
	return predicate.AsMarketAppUserTemplate(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreateUser), v))
	})
}

// CreateUserIn applies the In predicate on the "create_user" field.
func CreateUserIn(vs ...int64) predicate.AsMarketAppUserTemplate {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AsMarketAppUserTemplate(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreateUser), v...))
	})
}

// CreateUserNotIn applies the NotIn predicate on the "create_user" field.
func CreateUserNotIn(vs ...int64) predicate.AsMarketAppUserTemplate {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AsMarketAppUserTemplate(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreateUser), v...))
	})
}

// CreateUserGT applies the GT predicate on the "create_user" field.
func CreateUserGT(v int64) predicate.AsMarketAppUserTemplate {
	return predicate.AsMarketAppUserTemplate(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreateUser), v))
	})
}

// CreateUserGTE applies the GTE predicate on the "create_user" field.
func CreateUserGTE(v int64) predicate.AsMarketAppUserTemplate {
	return predicate.AsMarketAppUserTemplate(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreateUser), v))
	})
}

// CreateUserLT applies the LT predicate on the "create_user" field.
func CreateUserLT(v int64) predicate.AsMarketAppUserTemplate {
	return predicate.AsMarketAppUserTemplate(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreateUser), v))
	})
}

// CreateUserLTE applies the LTE predicate on the "create_user" field.
func CreateUserLTE(v int64) predicate.AsMarketAppUserTemplate {
	return predicate.AsMarketAppUserTemplate(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreateUser), v))
	})
}

// CreateUserIsNil applies the IsNil predicate on the "create_user" field.
func CreateUserIsNil() predicate.AsMarketAppUserTemplate {
	return predicate.AsMarketAppUserTemplate(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldCreateUser)))
	})
}

// CreateUserNotNil applies the NotNil predicate on the "create_user" field.
func CreateUserNotNil() predicate.AsMarketAppUserTemplate {
	return predicate.AsMarketAppUserTemplate(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldCreateUser)))
	})
}

// UpdateUserEQ applies the EQ predicate on the "update_user" field.
func UpdateUserEQ(v int64) predicate.AsMarketAppUserTemplate {
	return predicate.AsMarketAppUserTemplate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateUser), v))
	})
}

// UpdateUserNEQ applies the NEQ predicate on the "update_user" field.
func UpdateUserNEQ(v int64) predicate.AsMarketAppUserTemplate {
	return predicate.AsMarketAppUserTemplate(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdateUser), v))
	})
}

// UpdateUserIn applies the In predicate on the "update_user" field.
func UpdateUserIn(vs ...int64) predicate.AsMarketAppUserTemplate {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AsMarketAppUserTemplate(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUpdateUser), v...))
	})
}

// UpdateUserNotIn applies the NotIn predicate on the "update_user" field.
func UpdateUserNotIn(vs ...int64) predicate.AsMarketAppUserTemplate {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AsMarketAppUserTemplate(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUpdateUser), v...))
	})
}

// UpdateUserGT applies the GT predicate on the "update_user" field.
func UpdateUserGT(v int64) predicate.AsMarketAppUserTemplate {
	return predicate.AsMarketAppUserTemplate(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdateUser), v))
	})
}

// UpdateUserGTE applies the GTE predicate on the "update_user" field.
func UpdateUserGTE(v int64) predicate.AsMarketAppUserTemplate {
	return predicate.AsMarketAppUserTemplate(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdateUser), v))
	})
}

// UpdateUserLT applies the LT predicate on the "update_user" field.
func UpdateUserLT(v int64) predicate.AsMarketAppUserTemplate {
	return predicate.AsMarketAppUserTemplate(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdateUser), v))
	})
}

// UpdateUserLTE applies the LTE predicate on the "update_user" field.
func UpdateUserLTE(v int64) predicate.AsMarketAppUserTemplate {
	return predicate.AsMarketAppUserTemplate(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdateUser), v))
	})
}

// UpdateUserIsNil applies the IsNil predicate on the "update_user" field.
func UpdateUserIsNil() predicate.AsMarketAppUserTemplate {
	return predicate.AsMarketAppUserTemplate(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldUpdateUser)))
	})
}

// UpdateUserNotNil applies the NotNil predicate on the "update_user" field.
func UpdateUserNotNil() predicate.AsMarketAppUserTemplate {
	return predicate.AsMarketAppUserTemplate(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldUpdateUser)))
	})
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v date.DateTime) predicate.AsMarketAppUserTemplate {
	vc := time.Time(v)
	return predicate.AsMarketAppUserTemplate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateTime), vc))
	})
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v date.DateTime) predicate.AsMarketAppUserTemplate {
	vc := time.Time(v)
	return predicate.AsMarketAppUserTemplate(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreateTime), vc))
	})
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...date.DateTime) predicate.AsMarketAppUserTemplate {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = time.Time(vs[i])
	}
	return predicate.AsMarketAppUserTemplate(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreateTime), v...))
	})
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...date.DateTime) predicate.AsMarketAppUserTemplate {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = time.Time(vs[i])
	}
	return predicate.AsMarketAppUserTemplate(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreateTime), v...))
	})
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v date.DateTime) predicate.AsMarketAppUserTemplate {
	vc := time.Time(v)
	return predicate.AsMarketAppUserTemplate(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreateTime), vc))
	})
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v date.DateTime) predicate.AsMarketAppUserTemplate {
	vc := time.Time(v)
	return predicate.AsMarketAppUserTemplate(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreateTime), vc))
	})
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v date.DateTime) predicate.AsMarketAppUserTemplate {
	vc := time.Time(v)
	return predicate.AsMarketAppUserTemplate(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreateTime), vc))
	})
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v date.DateTime) predicate.AsMarketAppUserTemplate {
	vc := time.Time(v)
	return predicate.AsMarketAppUserTemplate(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreateTime), vc))
	})
}

// CreateTimeIsNil applies the IsNil predicate on the "create_time" field.
func CreateTimeIsNil() predicate.AsMarketAppUserTemplate {
	return predicate.AsMarketAppUserTemplate(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldCreateTime)))
	})
}

// CreateTimeNotNil applies the NotNil predicate on the "create_time" field.
func CreateTimeNotNil() predicate.AsMarketAppUserTemplate {
	return predicate.AsMarketAppUserTemplate(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldCreateTime)))
	})
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v date.DateTime) predicate.AsMarketAppUserTemplate {
	vc := time.Time(v)
	return predicate.AsMarketAppUserTemplate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateTime), vc))
	})
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v date.DateTime) predicate.AsMarketAppUserTemplate {
	vc := time.Time(v)
	return predicate.AsMarketAppUserTemplate(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdateTime), vc))
	})
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...date.DateTime) predicate.AsMarketAppUserTemplate {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = time.Time(vs[i])
	}
	return predicate.AsMarketAppUserTemplate(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUpdateTime), v...))
	})
}

// UpdateTimeNotIn applies the NotIn predicate on the "update_time" field.
func UpdateTimeNotIn(vs ...date.DateTime) predicate.AsMarketAppUserTemplate {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = time.Time(vs[i])
	}
	return predicate.AsMarketAppUserTemplate(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUpdateTime), v...))
	})
}

// UpdateTimeGT applies the GT predicate on the "update_time" field.
func UpdateTimeGT(v date.DateTime) predicate.AsMarketAppUserTemplate {
	vc := time.Time(v)
	return predicate.AsMarketAppUserTemplate(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdateTime), vc))
	})
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v date.DateTime) predicate.AsMarketAppUserTemplate {
	vc := time.Time(v)
	return predicate.AsMarketAppUserTemplate(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdateTime), vc))
	})
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v date.DateTime) predicate.AsMarketAppUserTemplate {
	vc := time.Time(v)
	return predicate.AsMarketAppUserTemplate(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdateTime), vc))
	})
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v date.DateTime) predicate.AsMarketAppUserTemplate {
	vc := time.Time(v)
	return predicate.AsMarketAppUserTemplate(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdateTime), vc))
	})
}

// UpdateTimeIsNil applies the IsNil predicate on the "update_time" field.
func UpdateTimeIsNil() predicate.AsMarketAppUserTemplate {
	return predicate.AsMarketAppUserTemplate(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldUpdateTime)))
	})
}

// UpdateTimeNotNil applies the NotNil predicate on the "update_time" field.
func UpdateTimeNotNil() predicate.AsMarketAppUserTemplate {
	return predicate.AsMarketAppUserTemplate(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldUpdateTime)))
	})
}

// HasUserx applies the HasEdge predicate on the "userx" edge.
func HasUserx() predicate.AsMarketAppUserTemplate {
	return predicate.AsMarketAppUserTemplate(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserxTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserxTable, UserxColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserxWith applies the HasEdge predicate on the "userx" edge with a given conditions (other predicates).
func HasUserxWith(preds ...predicate.AsUser) predicate.AsMarketAppUserTemplate {
	return predicate.AsMarketAppUserTemplate(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserxInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserxTable, UserxColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasTemplatex applies the HasEdge predicate on the "templatex" edge.
func HasTemplatex() predicate.AsMarketAppUserTemplate {
	return predicate.AsMarketAppUserTemplate(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TemplatexTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, TemplatexTable, TemplatexColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTemplatexWith applies the HasEdge predicate on the "templatex" edge with a given conditions (other predicates).
func HasTemplatexWith(preds ...predicate.AsMarketAppComponentTemplate) predicate.AsMarketAppUserTemplate {
	return predicate.AsMarketAppUserTemplate(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TemplatexInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, TemplatexTable, TemplatexColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.AsMarketAppUserTemplate) predicate.AsMarketAppUserTemplate {
	return predicate.AsMarketAppUserTemplate(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.AsMarketAppUserTemplate) predicate.AsMarketAppUserTemplate {
	return predicate.AsMarketAppUserTemplate(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.AsMarketAppUserTemplate) predicate.AsMarketAppUserTemplate {
	return predicate.AsMarketAppUserTemplate(func(s *sql.Selector) {
		p(s.Not())
	})
}
