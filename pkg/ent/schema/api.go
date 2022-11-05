package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"

	"github.com/suyuan32/simple-admin-core/pkg/ent/schema/mixins"
)

type Api struct {
	ent.Schema
}

func (Api) Fields() []ent.Field {
	return []ent.Field{
		field.String("path").Comment("API path | API 路径"),
		field.String("description").Comment("API description | API 描述"),
		field.String("api_group").Comment("API group | API 分组"),
		field.String("method").Default("POST").Comment("HTTP method | HTTP 请求类型"),
	}
}

func (Api) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
	}
}

func (Api) Edges() []ent.Edge {
	return nil
}
