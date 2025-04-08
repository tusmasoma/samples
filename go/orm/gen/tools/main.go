package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

const (
	DB_USER = "root"
	DB_PASS = "root"
	DB_NAME = "go_orm_gen_setup_db"
)

func main() {
	dsn := fmt.Sprintf("%s:%s@tcp(localhost:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", DB_USER, DB_PASS, DB_NAME)
	db, _ := gorm.Open(mysql.Open(dsn))
	g := gen.NewGenerator(gen.Config{
		OutPath:           "../database/mysql/query",
		Mode:              gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
		FieldWithIndexTag: true,
		FieldWithTypeTag:  true,
		FieldNullable:     true,
	})
	g.UseDB(db)
	all := g.GenerateAllTable()
	g.ApplyBasic(all...)
	g.Execute()
}
