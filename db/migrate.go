package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// 建立数据库连接
func Migrate() {
	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/taobao?multiStatements=true")
	if err != nil {
		fmt.Println("1", err)
		return
	}
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		fmt.Println("2", err)
		return
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://./db/migrations",
		"mysql",
		driver,
	)
	if err != nil {
		fmt.Println("3", err)
		return
	}
	err = m.Up()
	if err != nil {
		fmt.Println("4", err)
	}
}
