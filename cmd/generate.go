package main

import (
	"intikom-interview/model"

	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

type Querier interface {
	// SELECT page_count * page_size FROM pragma_page_count('events')
	GetSizeEvents() (gen.T, error)
}

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath:       "../dal",
		FieldNullable: true,
		/* Mode: gen.WithoutContext|gen.WithDefaultQuery*/
		//if you want the nullable field generation property to be pointer type, set FieldNullable true
		/* FieldNullable: true,*/
		//if you want to assign field which has default value in `Create` API, set FieldCoverable true, reference: https://gorm.io/docs/create.html#Default-Values
		/* FieldCoverable: true,*/
		//if you want to generate index tags from database, set FieldWithIndexTag true
		/* FieldWithIndexTag: true,*/
		//if you want to generate type tags from database, set FieldWithTypeTag true
		/* FieldWithTypeTag: true,*/
		//if you need unit tests for query code, set WithUnitTest true
		WithUnitTest: true,
		Mode:         gen.WithoutContext | gen.WithDefaultQuery,
	})

	db, _ := gorm.Open(mysql.Open("root:tT£91F]y3[LV@tcp(127.0.0.1:3306)/intikom_test?charset=utf8mb4&parseTime=True&loc=Local"))

	g.UseDB(db)

	db.AutoMigrate(&model.Task{}, &model.User{})

	//g.ApplyInterface(func(Querier) {}, model.Event{})

	// apply basic crud api on structs or table models which is specified by table name with function
	// GenerateModel/GenerateModelAs. And generator will generate table models' code when calling Excute.
	//g.ApplyBasic(model.User{}, g.GenerateModel("company"), g.GenerateModelAs("people", "Person", gen.FieldIgnore("address")))
	g.ApplyBasic(model.User{}, model.Task{})

	// execute the action of code generation
	g.Execute()
}
