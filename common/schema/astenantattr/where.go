// Code generated by entc, DO NOT EDIT.

package astenantattr

import (
	"orginone/common/schema/predicate"
	"orginone/common/tools/date"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.AsTenantAttr {
	return predicate.AsTenantAttr(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.AsTenantAttr {
	return predicate.AsTenantAttr(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.AsTenantAttr {
	return predicate.AsTenantAttr(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.AsTenantAttr {
	return predicate.AsTenantAttr(func(s *sql.Selector) {
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
func IDNotIn(ids ...int64) predicate.AsTenantAttr {
	return predicate.AsTenantAttr(func(s *sql.Selector) {
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
func IDGT(id int64) predicate.AsTenantAttr {
	return predicate.AsTenantAttr(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.AsTenantAttr {
	return predicate.AsTenantAttr(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.AsTenantAttr {
	return predicate.AsTenantAttr(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.AsTenantAttr {
	return predicate.AsTenantAttr(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// ParentID applies equality check predicate on the "parent_id" field. It's identical to ParentIDEQ.
func ParentID(v int64) predicate.AsTenantAttr {
	return predicate.AsTenantAttr(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldParentID), v))
	})
}

// AttrName applies equality check predicate on the "attr_name" field. It's identical to AttrNameEQ.
func AttrName(v string) predicate.AsTenantAttr {
	return predicate.AsTenantAttr(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAttrName), v))
	})
}

// AttrRemark applies equality check predicate on the "attr_remark" field. It's identical to AttrRemarkEQ.
func AttrRemark(v string) predicate.AsTenantAttr {
	return predicate.AsTenantAttr(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAttrRemark), v))
	})
}

// IsDeleted applies equality check predicate on the "is_deleted" field. It's identical to IsDeletedEQ.
func IsDeleted(v int64) predicate.AsTenantAttr {
	return predicate.AsTenantAttr(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIsDeleted), v))
	})
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v int64) predicate.AsTenantAttr {
	return predicate.AsTenantAttr(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStatus), v))
	})
}

// CreateUser applies equality check predicate on the "create_user" field. It's identical to CreateUserEQ.
func CreateUser(v int64) predicate.AsTenantAttr {
	return predicate.AsTenantAttr(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateUser), v))
	})
}

// UpdateUser applies equality check predicate on the "update_user" field. It's identical to UpdateUserEQ.
func UpdateUser(v int64) predicate.AsTenantAttr {
	return predicate.AsTenantAttr(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateUser), v))
	})
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v date.DateTime) predicate.AsTenantAttr {
	vc := time.Time(v)
	return predicate.AsTenantAttr(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateTime), vc))
	})
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v date.DateTime) predicate.AsTenantAttr {
	vc := time.Time(v)
	return predicate.AsTenantAttr(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateTime), vc))
	})
}

// ParentIDEQ applies the EQ predicate on the "parent_id" field.
func ParentIDEQ(v int64) predicate.AsTenantAttr {
	return predicate.AsTenantAttr(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldParentID), v))
	})
}

// ParentIDNEQ applies the NEQ predicate on the "parent_id" field.
func ParentIDNEQ(v int64) predicate.AsTenantAttr {
	return predicate.AsTenantAttr(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldParentID), v))
	})
}

// ParentIDIn applies the In predicate on the "parent_id" field.
func ParentIDIn(vs ...int64) predicate.AsTenantAttr {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AsTenantAttr(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldParentID), v...))
	})
}

// ParentIDNotIn applies the NotIn predicate on the "parent_id" field.
func ParentIDNotIn(vs ...int64) predicate.AsTenantAttr {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AsTenantAttr(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldParentID), v...))
	})
}

// ParentIDIsNil applies the IsNil predicate on the "parent_id" field.
func ParentIDIsNil() predicate.AsTenantAttr {
	return predicate.AsTenantAttr(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldParentID)))
	})
}

// ParentIDNotNil applies the NotNil predicate on the "parent_id" field.
func ParentIDNotNil() predicate.AsTenantAttr {
	return predicate.AsTenantAttr(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldParentID)))
	})
}

// AttrNameEQ applies the EQ predicate on the "attr_name" field.
func AttrNameEQ(v string) predicate.AsTenantAttr {
	return predicate.AsTenantAttr(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAttrName), v))
	})
}

// AttrNameNEQ applies the NEQ predicate on the "attr_name" field.
func AttrNameNEQ(v string) predicate.AsTenantAttr {
	return predicate.AsTenantAttr(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAttrName), v))
	})
}

// AttrNameIn applies the In predicate on the "attr_name" field.
func AttrNameIn(vs ...string) predicate.AsTenantAttr {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AsTenantAttr(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAttrName), v...))
	})
}

// AttrNameNotIn applies the NotIn predicate on the "attr_name" field.
func AttrNameNotIn(vs ...string) predicate.AsTenantAttr {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AsTenantAttr(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAttrName), v...))
	})
}

// AttrNameGT applies the GT predicate on the "attr_name" field.
func AttrNameGT(v string) predicate.AsTenantAttr {
	return predicate.AsTenantAttr(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAttrName), v))
	})
}

// AttrNameGTE applies the GTE predicate on the "attr_name" field.
func AttrNameGTE(v string) predicate.AsTenantAttr {
	return predicate.AsTenantAttr(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAttrName), v))
	})
}

// AttrNameLT applies the LT predicate on the "attr_name" field.
func AttrNameLT(v string) predicate.AsTenantAttr {
	return predicate.AsTenantAttr(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAttrName), v))
	})
}

// AttrNameLTE applies the LTE predicate on the "attr_name" field.
func AttrNameLTE(v string) predicate.AsTenantAttr {
	return predicate.AsTenantAttr(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAttrName), v))
	})
}

// AttrNameContains applies the Contains predicate on the "attr_name" field.
func AttrNameContains(v string) predicate.AsTenantAttr {
	return predicate.AsTenantAttr(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldAttrName), v))
	})
}

// AttrNameHasPrefix applies the HasPrefix predicate on the "attr_name" field.
func AttrNameHasPrefix(v string) predicate.AsTenantAttr {
	return predicate.AsTenantAttr(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldAttrName), v))
	})
}

// AttrNameHasSuffix applies the HasSuffix predicate on the "attr_name" field.
func AttrNameHasSuffix(v string) predicate.AsTenantAttr {
	return predicate.AsTenantAttr(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldAttrName), v))
	})
}

// AttrNameEqualFold applies the EqualFold predicate on the "attr_name" field.
func AttrNameEqualFold(v string) predicate.AsTenantAttr {
	return predicate.AsTenantAttr(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldAttrName), v))
	})
}

// AttrNameContainsFold applies the ContainsFold predicate on the "attr_name" field.
func AttrNameContainsFold(v string) predicate.AsTenantAttr {
	return predicate.AsTenantAttr(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldAttrName), v))
	})
}

// AttrRemarkEQ applies the EQ predicate on the "attr_remark" field.
func AttrRemarkEQ(v string) predicate.AsTenantAttr {
	return predicate.AsTenantAttr(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAttrRemark), v))
	})
}

// AttrRemarkNEQ applies the NEQ predicate on the "attr_remark" field.
func AttrRemarkNEQ(v string) predicate.AsTenantAttr {
	return predicate.AsTenantAttr(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAttrRemark), v))
	})
}

// AttrRemarkIn applies the In predicate on the "attr_remark" field.
func AttrRemarkIn(vs ...string) predicate.AsTenantAttr {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AsTenantAttr(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAttrRemark), v...))
	})
}

// AttrRemarkNotIn applies the NotIn predicate on the "attr_remark" field.
func AttrRemarkNotIn(vs ...string) predicate.AsTenantAttr {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AsTenantAttr(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAttrRemark), v...))
	})
}

// AttrRemarkGT applies the GT predicate on the "attr_remark" field.
func AttrRemarkGT(v string) predicate.AsTenantAttr {
	return predicate.AsTenantAttr(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAttrRemark), v))
	})
}

// AttrRemarkGTE applies the GTE predicate on the "attr_remark" field.
func AttrRemarkGTE(v string) predicate.AsTenantAttr {
	return predicate.AsTenantAttr(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAttrRemark), v))
	})
}

// AttrRemarkLT applies the LT predicate on the "attr_remark" field.
func AttrRemarkLT(v string) predicate.AsTenantAttr {
	return predicate.AsTenantAttr(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAttrRemark), v))
	})
}

// AttrRemarkLTE applies the LTE predicate on the "attr_remark" field.
func AttrRemarkLTE(v string) predicate.AsTenantAttr {
	return predicate.AsTenantAttr(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAttrRemark), v))
	})
}

// AttrRemarkContains applies the Contains predicate on the "attr_remark" field.
func AttrRemarkContains(v string) predicate.AsTenantAttr {
	return predicate.AsTenantAttr(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldAttrRemark), v))
	})
}

// AttrRemarkHasPrefix applies the HasPrefix predicate on the "attr_remark" field.
func AttrRemarkHasPrefix(v string) predicate.AsTenantAttr {
	return predicate.AsTenantAttr(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldAttrRemark), v))
	})
}

// AttrRemarkHasSuffix applies the HasSuffix predicate on the "attr_remark" field.
func AttrRemarkHasSuffix(v string) predicate.AsTenantAttr {
	return predicate.AsTenantAttr(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldAttrRemark), v))
	})
}

// AttrRemarkIsNil applies the IsNil predicate on the "attr_remark" field.
func AttrRemarkIsNil() predicate.AsTenantAttr {
	return predicate.AsTenantAttr(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldAttrRemark)))
	})
}

// AttrRemarkNotNil applies the NotNil predicate on the "attr_remark" field.
func AttrRemarkNotNil() predicate.AsTenantAttr {
	return predicate.AsTenantAttr(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldAttrRemark)))
	})
}

// AttrRemarkEqualFold applies the EqualFold predicate on the "attr_remark" field.
func AttrRemarkEqualFold(v string) predicate.AsTenantAttr {
	return predicate.AsTenantAttr(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldAttrRemark), v))
	})
}

// AttrRemarkContainsFold applies the ContainsFold predicate on the "attr_remark" field.
func AttrRemarkContainsFold(v string) predicate.AsTenantAttr {
	return predicate.AsTenantAttr(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldAttrRemark), v))
	})
}

// IsDeletedEQ applies the EQ predicate on the "is_deleted" field.
func IsDeletedEQ(v int64) predicate.AsTenantAttr {
	return predicate.AsTenantAttr(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIsDeleted), v))
	})
}

// IsDeletedNEQ applies the NEQ predicate on the "is_deleted" field.
func IsDeletedNEQ(v int64) predicate.AsTenantAttr {
	return predicate.AsTenantAttr(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldIsDeleted), v))
	})
}

// IsDeletedIn applies the In predicate on the "is_deleted" field.
func IsDeletedIn(vs ...int64) predicate.AsTenantAttr {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AsTenantAttr(func(s *sql.Selector) {
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
func IsDeletedNotIn(vs ...int64) predicate.AsTenantAttr {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AsTenantAttr(func(s *sql.Selector) {
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
func IsDeletedGT(v int64) predicate.AsTenantAttr {
	return predicate.AsTenantAttr(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldIsDeleted), v))
	})
}

// IsDeletedGTE applies the GTE predicate on the "is_deleted" field.
func IsDeletedGTE(v int64) predicate.AsTenantAttr {
	return predicate.AsTenantAttr(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldIsDeleted), v))
	})
}

// IsDeletedLT applies the LT predicate on the "is_deleted" field.
func IsDeletedLT(v int64) predicate.AsTenantAttr {
	return predicate.AsTenantAttr(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldIsDeleted), v))
	})
}

// IsDeletedLTE applies the LTE predicate on the "is_deleted" field.
func IsDeletedLTE(v int64) predicate.AsTenantAttr {
	return predicate.AsTenantAttr(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldIsDeleted), v))
	})
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v int64) predicate.AsTenantAttr {
	return predicate.AsTenantAttr(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStatus), v))
	})
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v int64) predicate.AsTenantAttr {
	return predicate.AsTenantAttr(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStatus), v))
	})
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...int64) predicate.AsTenantAttr {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AsTenantAttr(func(s *sql.Selector) {
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
func StatusNotIn(vs ...int64) predicate.AsTenantAttr {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AsTenantAttr(func(s *sql.Selector) {
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
func StatusGT(v int64) predicate.AsTenantAttr {
	return predicate.AsTenantAttr(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldStatus), v))
	})
}

// StatusGTE applies the GTE predicate on the "status" field.
func StatusGTE(v int64) predicate.AsTenantAttr {
	return predicate.AsTenantAttr(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldStatus), v))
	})
}

// StatusLT applies the LT predicate on the "status" field.
func StatusLT(v int64) predicate.AsTenantAttr {
	return predicate.AsTenantAttr(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldStatus), v))
	})
}

// StatusLTE applies the LTE predicate on the "status" field.
func StatusLTE(v int64) predicate.AsTenantAttr {
	return predicate.AsTenantAttr(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldStatus), v))
	})
}

// StatusIsNil applies the IsNil predicate on the "status" field.
func StatusIsNil() predicate.AsTenantAttr {
	return predicate.AsTenantAttr(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldStatus)))
	})
}

// StatusNotNil applies the NotNil predicate on the "status" field.
func StatusNotNil() predicate.AsTenantAttr {
	return predicate.AsTenantAttr(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldStatus)))
	})
}

// CreateUserEQ applies the EQ predicate on the "create_user" field.
func CreateUserEQ(v int64) predicate.AsTenantAttr {
	return predicate.AsTenantAttr(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateUser), v))
	})
}

// CreateUserNEQ applies the NEQ predicate on the "create_user" field.
func CreateUserNEQ(v int64) predicate.AsTenantAttr {
	return predicate.AsTenantAttr(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreateUser), v))
	})
}

// CreateUserIn applies the In predicate on the "create_user" field.
func CreateUserIn(vs ...int64) predicate.AsTenantAttr {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AsTenantAttr(func(s *sql.Selector) {
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
func CreateUserNotIn(vs ...int64) predicate.AsTenantAttr {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AsTenantAttr(func(s *sql.Selector) {
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
func CreateUserGT(v int64) predicate.AsTenantAttr {
	return predicate.AsTenantAttr(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreateUser), v))
	})
}

// CreateUserGTE applies the GTE predicate on the "create_user" field.
func CreateUserGTE(v int64) predicate.AsTenantAttr {
	return predicate.AsTenantAttr(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreateUser), v))
	})
}

// CreateUserLT applies the LT predicate on the "create_user" field.
func CreateUserLT(v int64) predicate.AsTenantAttr {
	return predicate.AsTenantAttr(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreateUser), v))
	})
}

// CreateUserLTE applies the LTE predicate on the "create_user" field.
func CreateUserLTE(v int64) predicate.AsTenantAttr {
	return predicate.AsTenantAttr(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreateUser), v))
	})
}

// CreateUserIsNil applies the IsNil predicate on the "create_user" field.
func CreateUserIsNil() predicate.AsTenantAttr {
	return predicate.AsTenantAttr(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldCreateUser)))
	})
}

// CreateUserNotNil applies the NotNil predicate on the "create_user" field.
func CreateUserNotNil() predicate.AsTenantAttr {
	return predicate.AsTenantAttr(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldCreateUser)))
	})
}

// UpdateUserEQ applies the EQ predicate on the "update_user" field.
func UpdateUserEQ(v int64) predicate.AsTenantAttr {
	return predicate.AsTenantAttr(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateUser), v))
	})
}

// UpdateUserNEQ applies the NEQ predicate on the "update_user" field.
func UpdateUserNEQ(v int64) predicate.AsTenantAttr {
	return predicate.AsTenantAttr(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdateUser), v))
	})
}

// UpdateUserIn applies the In predicate on the "update_user" field.
func UpdateUserIn(vs ...int64) predicate.AsTenantAttr {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AsTenantAttr(func(s *sql.Selector) {
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
func UpdateUserNotIn(vs ...int64) predicate.AsTenantAttr {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AsTenantAttr(func(s *sql.Selector) {
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
func UpdateUserGT(v int64) predicate.AsTenantAttr {
	return predicate.AsTenantAttr(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdateUser), v))
	})
}

// UpdateUserGTE applies the GTE predicate on the "update_user" field.
func UpdateUserGTE(v int64) predicate.AsTenantAttr {
	return predicate.AsTenantAttr(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdateUser), v))
	})
}

// UpdateUserLT applies the LT predicate on the "update_user" field.
func UpdateUserLT(v int64) predicate.AsTenantAttr {
	return predicate.AsTenantAttr(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdateUser), v))
	})
}

// UpdateUserLTE applies the LTE predicate on the "update_user" field.
func UpdateUserLTE(v int64) predicate.AsTenantAttr {
	return predicate.AsTenantAttr(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdateUser), v))
	})
}

// UpdateUserIsNil applies the IsNil predicate on the "update_user" field.
func UpdateUserIsNil() predicate.AsTenantAttr {
	return predicate.AsTenantAttr(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldUpdateUser)))
	})
}

// UpdateUserNotNil applies the NotNil predicate on the "update_user" field.
func UpdateUserNotNil() predicate.AsTenantAttr {
	return predicate.AsTenantAttr(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldUpdateUser)))
	})
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v date.DateTime) predicate.AsTenantAttr {
	vc := time.Time(v)
	return predicate.AsTenantAttr(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateTime), vc))
	})
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v date.DateTime) predicate.AsTenantAttr {
	vc := time.Time(v)
	return predicate.AsTenantAttr(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreateTime), vc))
	})
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...date.DateTime) predicate.AsTenantAttr {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = time.Time(vs[i])
	}
	return predicate.AsTenantAttr(func(s *sql.Selector) {
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
func CreateTimeNotIn(vs ...date.DateTime) predicate.AsTenantAttr {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = time.Time(vs[i])
	}
	return predicate.AsTenantAttr(func(s *sql.Selector) {
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
func CreateTimeGT(v date.DateTime) predicate.AsTenantAttr {
	vc := time.Time(v)
	return predicate.AsTenantAttr(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreateTime), vc))
	})
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v date.DateTime) predicate.AsTenantAttr {
	vc := time.Time(v)
	return predicate.AsTenantAttr(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreateTime), vc))
	})
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v date.DateTime) predicate.AsTenantAttr {
	vc := time.Time(v)
	return predicate.AsTenantAttr(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreateTime), vc))
	})
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v date.DateTime) predicate.AsTenantAttr {
	vc := time.Time(v)
	return predicate.AsTenantAttr(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreateTime), vc))
	})
}

// CreateTimeIsNil applies the IsNil predicate on the "create_time" field.
func CreateTimeIsNil() predicate.AsTenantAttr {
	return predicate.AsTenantAttr(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldCreateTime)))
	})
}

// CreateTimeNotNil applies the NotNil predicate on the "create_time" field.
func CreateTimeNotNil() predicate.AsTenantAttr {
	return predicate.AsTenantAttr(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldCreateTime)))
	})
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v date.DateTime) predicate.AsTenantAttr {
	vc := time.Time(v)
	return predicate.AsTenantAttr(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateTime), vc))
	})
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v date.DateTime) predicate.AsTenantAttr {
	vc := time.Time(v)
	return predicate.AsTenantAttr(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdateTime), vc))
	})
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...date.DateTime) predicate.AsTenantAttr {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = time.Time(vs[i])
	}
	return predicate.AsTenantAttr(func(s *sql.Selector) {
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
func UpdateTimeNotIn(vs ...date.DateTime) predicate.AsTenantAttr {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = time.Time(vs[i])
	}
	return predicate.AsTenantAttr(func(s *sql.Selector) {
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
func UpdateTimeGT(v date.DateTime) predicate.AsTenantAttr {
	vc := time.Time(v)
	return predicate.AsTenantAttr(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdateTime), vc))
	})
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v date.DateTime) predicate.AsTenantAttr {
	vc := time.Time(v)
	return predicate.AsTenantAttr(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdateTime), vc))
	})
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v date.DateTime) predicate.AsTenantAttr {
	vc := time.Time(v)
	return predicate.AsTenantAttr(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdateTime), vc))
	})
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v date.DateTime) predicate.AsTenantAttr {
	vc := time.Time(v)
	return predicate.AsTenantAttr(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdateTime), vc))
	})
}

// UpdateTimeIsNil applies the IsNil predicate on the "update_time" field.
func UpdateTimeIsNil() predicate.AsTenantAttr {
	return predicate.AsTenantAttr(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldUpdateTime)))
	})
}

// UpdateTimeNotNil applies the NotNil predicate on the "update_time" field.
func UpdateTimeNotNil() predicate.AsTenantAttr {
	return predicate.AsTenantAttr(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldUpdateTime)))
	})
}

// HasParentx applies the HasEdge predicate on the "parentx" edge.
func HasParentx() predicate.AsTenantAttr {
	return predicate.AsTenantAttr(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ParentxTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ParentxTable, ParentxColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasParentxWith applies the HasEdge predicate on the "parentx" edge with a given conditions (other predicates).
func HasParentxWith(preds ...predicate.AsTenantAttr) predicate.AsTenantAttr {
	return predicate.AsTenantAttr(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ParentxTable, ParentxColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasChildrens applies the HasEdge predicate on the "childrens" edge.
func HasChildrens() predicate.AsTenantAttr {
	return predicate.AsTenantAttr(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ChildrensTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ChildrensTable, ChildrensColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasChildrensWith applies the HasEdge predicate on the "childrens" edge with a given conditions (other predicates).
func HasChildrensWith(preds ...predicate.AsTenantAttr) predicate.AsTenantAttr {
	return predicate.AsTenantAttr(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ChildrensTable, ChildrensColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasAttrRoles applies the HasEdge predicate on the "attrRoles" edge.
func HasAttrRoles() predicate.AsTenantAttr {
	return predicate.AsTenantAttr(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AttrRolesTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, AttrRolesTable, AttrRolesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAttrRolesWith applies the HasEdge predicate on the "attrRoles" edge with a given conditions (other predicates).
func HasAttrRolesWith(preds ...predicate.AsTenantAttrRole) predicate.AsTenantAttr {
	return predicate.AsTenantAttr(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AttrRolesInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, AttrRolesTable, AttrRolesColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.AsTenantAttr) predicate.AsTenantAttr {
	return predicate.AsTenantAttr(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.AsTenantAttr) predicate.AsTenantAttr {
	return predicate.AsTenantAttr(func(s *sql.Selector) {
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
func Not(p predicate.AsTenantAttr) predicate.AsTenantAttr {
	return predicate.AsTenantAttr(func(s *sql.Selector) {
		p(s.Not())
	})
}