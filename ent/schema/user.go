package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "user"},
	}
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		DefaultMix{},
	}
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("username").Comment("用户名"),
		field.Enum("role").Comment("用户角色").Values("ADMIN", "USER"),
		field.String("password").Comment("密码").Optional().Nillable(),
		field.String("avatar").Comment("头像").Default("/avatar/default.png"),
		field.String("phone").Comment("手机号").Optional().Nillable(),
		field.String("email").Comment("邮箱").Optional().Nillable(),
		field.String("token").Comment("登录token").Optional().Nillable(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
