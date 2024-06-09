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
		Mode:    gen.WithDefaultQuery,
	})

	g.UseDB(db)

	depModel := g.GenerateModel("deps")
	userModel := g.GenerateModel("users", gen.FieldRelate(field.HasOne, "org_id", depModel,
		&field.RelateConfig{
			RelatePointer: true,
			GORMTag: map[string][]string{
				"foreignKey": {"org_id"},
				"references": {"id"},
			},
		},
	))

	g.ApplyBasic(
		userModel,
		depModel,
	)
	g.Execute()
}
