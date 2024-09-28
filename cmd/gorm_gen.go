package main

import (
	"gorm.io/gen"
	"gorm.io/gen/field"

	"gen-demo/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	generator := gen.NewGenerator(gen.Config{
		OutPath:      "gen/query",
		ModelPkgPath: "gen/model",
		WithUnitTest: false,
		// generate model global configuration
		FieldNullable:     true,
		FieldCoverable:    true,
		FieldSignable:     true,
		FieldWithIndexTag: true,
		FieldWithTypeTag:  true,
		Mode:              gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
	})

	db, err := gorm.Open(mysql.Open("root:passwd@tcp(localhost:3306)/db?charset=utf8mb4&parseTime=True&loc=Local"))
	if err != nil {
		panic(err)
	}

	if err := db.AutoMigrate(model.Parent{}, model.Child{}); err != nil {
		panic(err)
	}

	generator.UseDB(db)

	generator.ApplyBasic(
		generator.GenerateModel(new(model.Child).TableName()),
		generator.GenerateModel(
			new(model.Parent).TableName(),
			gen.FieldRelateModel(field.HasOne, "Child", new(model.Child), &field.RelateConfig{
				RelatePointer: true,
				GORMTag:       field.GormTag{}.Append("foreignKey", "ID"),
			}),
		),
	)
	generator.Execute()
}
