package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"time"
)

type DefaultMix struct {
	mixin.Schema
}

func (DefaultMix) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Comment("主键ID"),
		field.Time("createdAt").Comment("创建时间").Default(time.Now()),
		field.Time("updatedAt").Comment("更新时间").Default(time.Now).UpdateDefault(time.Now),
		field.Time("deletedAt").Comment("删除时间").Nillable().Optional(),
	}
}
