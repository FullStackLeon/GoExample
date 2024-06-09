package main

import (
	"github.com/rs/zerolog/log"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:root@tcp(127.0.0.1:3306)/kratos?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal().Err(err).Stack()
		return
	}

	g := gen.NewGenerator(gen.Config{
		OutPath: "GORM/dal",
	})

	g.UseDB(db)

	// 一对一关系(一个用户只在一个部门)
	//depModel := g.GenerateModel("deps")
	//userModel := g.GenerateModel("users",
	//	gen.FieldRelate(field.HasOne, "org_id", depModel,
	//		&field.RelateConfig{
	//			RelatePointer: true,
	//			GORMTag: map[string][]string{
	//				"foreignKey": {"org_id"}, 在 users 表中的关联Deps表主键的id字段
	//				"references": {"id"},
	//			},
	//		},
	//	))

	//一对多关系(一个用户可以在多个部门)
	//depModel := g.GenerateModel("deps")
	//userModel := g.GenerateModel("users",
	//	gen.FieldRelate(field.HasMany, "Deps", depModel,
	//		&field.RelateConfig{
	//			RelateSlice: true,
	//			GORMTag: map[string][]string{
	//				"references": {"user_id"}, // 在 users 表中的主键字段
	//				"foreignKey": {"user_id"}, // 在 deps 表中的外键字段
	//			},
	//		}))

	// 多对多关系(一个用户可以在多个部门,一个部门可以有多个用户)
	depModel := g.GenerateModel("deps")
	userModel := g.GenerateModel("users",
		gen.FieldRelate(field.Many2Many, "Deps", depModel,
			&field.RelateConfig{
				RelateSlice: true,
				GORMTag: map[string][]string{
					"many2many":      {"user_dep_relations"}, // 中间表的名称
					"foreignKey":     {"user_id"},            // foreignKey指定当前模型(User)在关联表(user_dep_relations)中的外键字段
					"joinForeignKey": {"user_id"},            // 指定中间表(user_dep_relations)中关联到另一个模型(Dep)的外键字段名称
					"references":     {"id"},                 // 指定当前模型(User)的主键字段id
					"joinReferences": {"id"},                 // 指定关联模型(Dep)的主键字段id
				},
			},
		),
	)
	g.ApplyBasic(
		userModel,
		depModel,
	)
	g.Execute()
}
