// Code generated by entc, DO NOT EDIT.

package user

import (
	"short-link/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
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
func IDGT(id int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "createdAt" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updatedAt" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAt applies equality check predicate on the "deletedAt" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// Username applies equality check predicate on the "username" field. It's identical to UsernameEQ.
func Username(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUsername), v))
	})
}

// Password applies equality check predicate on the "password" field. It's identical to PasswordEQ.
func Password(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPassword), v))
	})
}

// Avatar applies equality check predicate on the "avatar" field. It's identical to AvatarEQ.
func Avatar(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAvatar), v))
	})
}

// Phone applies equality check predicate on the "phone" field. It's identical to PhoneEQ.
func Phone(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPhone), v))
	})
}

// Email applies equality check predicate on the "email" field. It's identical to EmailEQ.
func Email(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEmail), v))
	})
}

// Token applies equality check predicate on the "token" field. It's identical to TokenEQ.
func Token(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldToken), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "createdAt" field.
func CreatedAtEQ(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "createdAt" field.
func CreatedAtNEQ(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "createdAt" field.
func CreatedAtIn(vs ...time.Time) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "createdAt" field.
func CreatedAtNotIn(vs ...time.Time) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "createdAt" field.
func CreatedAtGT(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "createdAt" field.
func CreatedAtGTE(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "createdAt" field.
func CreatedAtLT(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "createdAt" field.
func CreatedAtLTE(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updatedAt" field.
func UpdatedAtEQ(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updatedAt" field.
func UpdatedAtNEQ(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updatedAt" field.
func UpdatedAtIn(vs ...time.Time) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updatedAt" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updatedAt" field.
func UpdatedAtGT(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updatedAt" field.
func UpdatedAtGTE(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updatedAt" field.
func UpdatedAtLT(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updatedAt" field.
func UpdatedAtLTE(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAtEQ applies the EQ predicate on the "deletedAt" field.
func DeletedAtEQ(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtNEQ applies the NEQ predicate on the "deletedAt" field.
func DeletedAtNEQ(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIn applies the In predicate on the "deletedAt" field.
func DeletedAtIn(vs ...time.Time) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtNotIn applies the NotIn predicate on the "deletedAt" field.
func DeletedAtNotIn(vs ...time.Time) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtGT applies the GT predicate on the "deletedAt" field.
func DeletedAtGT(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtGTE applies the GTE predicate on the "deletedAt" field.
func DeletedAtGTE(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLT applies the LT predicate on the "deletedAt" field.
func DeletedAtLT(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLTE applies the LTE predicate on the "deletedAt" field.
func DeletedAtLTE(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIsNil applies the IsNil predicate on the "deletedAt" field.
func DeletedAtIsNil() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldDeletedAt)))
	})
}

// DeletedAtNotNil applies the NotNil predicate on the "deletedAt" field.
func DeletedAtNotNil() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldDeletedAt)))
	})
}

// UsernameEQ applies the EQ predicate on the "username" field.
func UsernameEQ(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUsername), v))
	})
}

// UsernameNEQ applies the NEQ predicate on the "username" field.
func UsernameNEQ(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUsername), v))
	})
}

// UsernameIn applies the In predicate on the "username" field.
func UsernameIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUsername), v...))
	})
}

// UsernameNotIn applies the NotIn predicate on the "username" field.
func UsernameNotIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUsername), v...))
	})
}

// UsernameGT applies the GT predicate on the "username" field.
func UsernameGT(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUsername), v))
	})
}

// UsernameGTE applies the GTE predicate on the "username" field.
func UsernameGTE(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUsername), v))
	})
}

// UsernameLT applies the LT predicate on the "username" field.
func UsernameLT(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUsername), v))
	})
}

// UsernameLTE applies the LTE predicate on the "username" field.
func UsernameLTE(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUsername), v))
	})
}

// UsernameContains applies the Contains predicate on the "username" field.
func UsernameContains(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldUsername), v))
	})
}

// UsernameHasPrefix applies the HasPrefix predicate on the "username" field.
func UsernameHasPrefix(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldUsername), v))
	})
}

// UsernameHasSuffix applies the HasSuffix predicate on the "username" field.
func UsernameHasSuffix(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldUsername), v))
	})
}

// UsernameEqualFold applies the EqualFold predicate on the "username" field.
func UsernameEqualFold(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldUsername), v))
	})
}

// UsernameContainsFold applies the ContainsFold predicate on the "username" field.
func UsernameContainsFold(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldUsername), v))
	})
}

// RoleEQ applies the EQ predicate on the "role" field.
func RoleEQ(v Role) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRole), v))
	})
}

// RoleNEQ applies the NEQ predicate on the "role" field.
func RoleNEQ(v Role) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRole), v))
	})
}

// RoleIn applies the In predicate on the "role" field.
func RoleIn(vs ...Role) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldRole), v...))
	})
}

// RoleNotIn applies the NotIn predicate on the "role" field.
func RoleNotIn(vs ...Role) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldRole), v...))
	})
}

// PasswordEQ applies the EQ predicate on the "password" field.
func PasswordEQ(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPassword), v))
	})
}

// PasswordNEQ applies the NEQ predicate on the "password" field.
func PasswordNEQ(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPassword), v))
	})
}

// PasswordIn applies the In predicate on the "password" field.
func PasswordIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPassword), v...))
	})
}

// PasswordNotIn applies the NotIn predicate on the "password" field.
func PasswordNotIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPassword), v...))
	})
}

// PasswordGT applies the GT predicate on the "password" field.
func PasswordGT(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPassword), v))
	})
}

// PasswordGTE applies the GTE predicate on the "password" field.
func PasswordGTE(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPassword), v))
	})
}

// PasswordLT applies the LT predicate on the "password" field.
func PasswordLT(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPassword), v))
	})
}

// PasswordLTE applies the LTE predicate on the "password" field.
func PasswordLTE(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPassword), v))
	})
}

// PasswordContains applies the Contains predicate on the "password" field.
func PasswordContains(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldPassword), v))
	})
}

// PasswordHasPrefix applies the HasPrefix predicate on the "password" field.
func PasswordHasPrefix(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldPassword), v))
	})
}

// PasswordHasSuffix applies the HasSuffix predicate on the "password" field.
func PasswordHasSuffix(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldPassword), v))
	})
}

// PasswordIsNil applies the IsNil predicate on the "password" field.
func PasswordIsNil() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldPassword)))
	})
}

// PasswordNotNil applies the NotNil predicate on the "password" field.
func PasswordNotNil() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldPassword)))
	})
}

// PasswordEqualFold applies the EqualFold predicate on the "password" field.
func PasswordEqualFold(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldPassword), v))
	})
}

// PasswordContainsFold applies the ContainsFold predicate on the "password" field.
func PasswordContainsFold(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldPassword), v))
	})
}

// AvatarEQ applies the EQ predicate on the "avatar" field.
func AvatarEQ(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAvatar), v))
	})
}

// AvatarNEQ applies the NEQ predicate on the "avatar" field.
func AvatarNEQ(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAvatar), v))
	})
}

// AvatarIn applies the In predicate on the "avatar" field.
func AvatarIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAvatar), v...))
	})
}

// AvatarNotIn applies the NotIn predicate on the "avatar" field.
func AvatarNotIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAvatar), v...))
	})
}

// AvatarGT applies the GT predicate on the "avatar" field.
func AvatarGT(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAvatar), v))
	})
}

// AvatarGTE applies the GTE predicate on the "avatar" field.
func AvatarGTE(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAvatar), v))
	})
}

// AvatarLT applies the LT predicate on the "avatar" field.
func AvatarLT(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAvatar), v))
	})
}

// AvatarLTE applies the LTE predicate on the "avatar" field.
func AvatarLTE(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAvatar), v))
	})
}

// AvatarContains applies the Contains predicate on the "avatar" field.
func AvatarContains(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldAvatar), v))
	})
}

// AvatarHasPrefix applies the HasPrefix predicate on the "avatar" field.
func AvatarHasPrefix(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldAvatar), v))
	})
}

// AvatarHasSuffix applies the HasSuffix predicate on the "avatar" field.
func AvatarHasSuffix(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldAvatar), v))
	})
}

// AvatarEqualFold applies the EqualFold predicate on the "avatar" field.
func AvatarEqualFold(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldAvatar), v))
	})
}

// AvatarContainsFold applies the ContainsFold predicate on the "avatar" field.
func AvatarContainsFold(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldAvatar), v))
	})
}

// PhoneEQ applies the EQ predicate on the "phone" field.
func PhoneEQ(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPhone), v))
	})
}

// PhoneNEQ applies the NEQ predicate on the "phone" field.
func PhoneNEQ(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPhone), v))
	})
}

// PhoneIn applies the In predicate on the "phone" field.
func PhoneIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPhone), v...))
	})
}

// PhoneNotIn applies the NotIn predicate on the "phone" field.
func PhoneNotIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPhone), v...))
	})
}

// PhoneGT applies the GT predicate on the "phone" field.
func PhoneGT(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPhone), v))
	})
}

// PhoneGTE applies the GTE predicate on the "phone" field.
func PhoneGTE(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPhone), v))
	})
}

// PhoneLT applies the LT predicate on the "phone" field.
func PhoneLT(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPhone), v))
	})
}

// PhoneLTE applies the LTE predicate on the "phone" field.
func PhoneLTE(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPhone), v))
	})
}

// PhoneContains applies the Contains predicate on the "phone" field.
func PhoneContains(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldPhone), v))
	})
}

// PhoneHasPrefix applies the HasPrefix predicate on the "phone" field.
func PhoneHasPrefix(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldPhone), v))
	})
}

// PhoneHasSuffix applies the HasSuffix predicate on the "phone" field.
func PhoneHasSuffix(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldPhone), v))
	})
}

// PhoneIsNil applies the IsNil predicate on the "phone" field.
func PhoneIsNil() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldPhone)))
	})
}

// PhoneNotNil applies the NotNil predicate on the "phone" field.
func PhoneNotNil() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldPhone)))
	})
}

// PhoneEqualFold applies the EqualFold predicate on the "phone" field.
func PhoneEqualFold(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldPhone), v))
	})
}

// PhoneContainsFold applies the ContainsFold predicate on the "phone" field.
func PhoneContainsFold(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldPhone), v))
	})
}

// EmailEQ applies the EQ predicate on the "email" field.
func EmailEQ(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEmail), v))
	})
}

// EmailNEQ applies the NEQ predicate on the "email" field.
func EmailNEQ(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEmail), v))
	})
}

// EmailIn applies the In predicate on the "email" field.
func EmailIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldEmail), v...))
	})
}

// EmailNotIn applies the NotIn predicate on the "email" field.
func EmailNotIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldEmail), v...))
	})
}

// EmailGT applies the GT predicate on the "email" field.
func EmailGT(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEmail), v))
	})
}

// EmailGTE applies the GTE predicate on the "email" field.
func EmailGTE(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEmail), v))
	})
}

// EmailLT applies the LT predicate on the "email" field.
func EmailLT(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEmail), v))
	})
}

// EmailLTE applies the LTE predicate on the "email" field.
func EmailLTE(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEmail), v))
	})
}

// EmailContains applies the Contains predicate on the "email" field.
func EmailContains(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldEmail), v))
	})
}

// EmailHasPrefix applies the HasPrefix predicate on the "email" field.
func EmailHasPrefix(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldEmail), v))
	})
}

// EmailHasSuffix applies the HasSuffix predicate on the "email" field.
func EmailHasSuffix(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldEmail), v))
	})
}

// EmailIsNil applies the IsNil predicate on the "email" field.
func EmailIsNil() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldEmail)))
	})
}

// EmailNotNil applies the NotNil predicate on the "email" field.
func EmailNotNil() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldEmail)))
	})
}

// EmailEqualFold applies the EqualFold predicate on the "email" field.
func EmailEqualFold(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldEmail), v))
	})
}

// EmailContainsFold applies the ContainsFold predicate on the "email" field.
func EmailContainsFold(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldEmail), v))
	})
}

// TokenEQ applies the EQ predicate on the "token" field.
func TokenEQ(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldToken), v))
	})
}

// TokenNEQ applies the NEQ predicate on the "token" field.
func TokenNEQ(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldToken), v))
	})
}

// TokenIn applies the In predicate on the "token" field.
func TokenIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldToken), v...))
	})
}

// TokenNotIn applies the NotIn predicate on the "token" field.
func TokenNotIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldToken), v...))
	})
}

// TokenGT applies the GT predicate on the "token" field.
func TokenGT(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldToken), v))
	})
}

// TokenGTE applies the GTE predicate on the "token" field.
func TokenGTE(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldToken), v))
	})
}

// TokenLT applies the LT predicate on the "token" field.
func TokenLT(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldToken), v))
	})
}

// TokenLTE applies the LTE predicate on the "token" field.
func TokenLTE(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldToken), v))
	})
}

// TokenContains applies the Contains predicate on the "token" field.
func TokenContains(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldToken), v))
	})
}

// TokenHasPrefix applies the HasPrefix predicate on the "token" field.
func TokenHasPrefix(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldToken), v))
	})
}

// TokenHasSuffix applies the HasSuffix predicate on the "token" field.
func TokenHasSuffix(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldToken), v))
	})
}

// TokenIsNil applies the IsNil predicate on the "token" field.
func TokenIsNil() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldToken)))
	})
}

// TokenNotNil applies the NotNil predicate on the "token" field.
func TokenNotNil() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldToken)))
	})
}

// TokenEqualFold applies the EqualFold predicate on the "token" field.
func TokenEqualFold(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldToken), v))
	})
}

// TokenContainsFold applies the ContainsFold predicate on the "token" field.
func TokenContainsFold(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldToken), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.User) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.User) predicate.User {
	return predicate.User(func(s *sql.Selector) {
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
func Not(p predicate.User) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		p(s.Not())
	})
}
