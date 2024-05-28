package main

import (
	"fmt"
	"log"
	"strings"
	"unicode"

	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

// DBType database type
type DBType string

const (
	// dbMySQL Gorm Drivers mysql || postgres || sqlite || sqlserver
	dbMySQL      DBType = "mysql"
	dbPostgres   DBType = "postgres"
	dbSQLite     DBType = "sqlite"
	dbSQLServer  DBType = "sqlserver"
	dbClickHouse DBType = "clickhouse"
)

const (
	// defaultQueryPath 模型文件存放目录
	defaultQueryPath = "./gen/dao/query"
)

const Dsn = "root:passw0rd@tcp(127.0.0.1:3306)/db_goods_center?charset=utf8mb4&parseTime=True&loc=Local"

// const Dsn = "root:abcd1234@tcp(192.168.28.75:3306)/db_token?charset=utf8mb4&parseTime=True&loc=Local"

func main() {
	// mysql
	db, err := gorm.Open(mysql.Open(Dsn))

	// clickhouse
	// db, err := gorm.Open(clickhouse.Open(Dsn))
	// postgres
	// db, err := gorm.Open(postgres.Open(Dsn))

	if err != nil {
		log.Fatalln("connect db server fail:", err)
	}

	g := gen.NewGenerator(gen.Config{
		OutPath: defaultQueryPath,
		// 默认为：gen.go
		// OutFile: "genb.go",
		// ModelPkgPath 生成模型的包名
		ModelPkgPath: "model",
		// 生成单元测试。
		// WithUnitTest: true,
		// FieldNullable 字段可为空时使用指针生成
		FieldNullable: true,
		// 字段有默认值时使用指针生成
		// FieldCoverable: true,
		// FieldWithIndexTag 使用GROM索引标记生成字段
		FieldWithIndexTag: true,
		// FieldWithTypeTag 使用gorm列类型标记生成字段
		FieldWithTypeTag: true,
		// 基于数据表定义的数据类型，生成对应的数据类型
		FieldSignable: true,
	})

	g.UseDB(db)

	g.WithJSONTagNameStrategy(
		ToLowerCamelCase, // 表字段 json tag 驼峰, 默认是下划线分割
	)
	// 部分表
	tables := []string{
		"apple_config",
		"apple_product_price",
		"user_score",
		// "gp_product_price",
	}
	// tables := []string{} //全部表
	models, err := genModels(g, db, tables)
	if err != nil {
		log.Fatalln("get tables info fail:", err)
	}

	// false 指生成model,不生成query文件; true 全都生成
	onlyModel := true
	if onlyModel {
		g.ApplyBasic(models...)
	}

	g.Execute()
}

// genModels is gorm/gen generated models
func genModels(g *gen.Generator, db *gorm.DB, tables []string) (models []interface{}, err error) {
	if len(tables) == 0 {
		// Execute tasks for all tables in the database
		tables, err = db.Migrator().GetTables()
		if err != nil {
			return nil, fmt.Errorf("GORM migrator get all tables fail: %w", err)
		}
	}

	// Execute some data table tasks
	models = make([]interface{}, len(tables))
	for i, tableName := range tables {
		models[i] = g.GenerateModel(tableName)
	}
	return models, nil
}

// ToUpperCamelCase 下划线单词转为大写驼峰单词
func ToUpperCamelCase(s string) string {
	s = strings.Replace(s, "_", " ", -1)
	s = strings.Title(s)
	return strings.Replace(s, " ", "", -1)
}

// ToLowerCamelCase 下划线单词转为小写驼峰单词
func ToLowerCamelCase(s string) string {
	s = ToUpperCamelCase(s)
	return string(unicode.ToLower(rune(s[0]))) + s[1:]
}
