package model

import (
	"context"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"short-link/ent"
	"short-link/ent/migrate"
)

var client *ent.Client

// InitDatabase 初始化数据库连接
func InitDatabase() {
	var err error
	client, err = ent.Open("mysql", "root:123456@tcp(localhost:3306)/short_link?parseTime=true")
	if err != nil {
		log.Fatalf("failed connecting to mysql: %v", err)
	}
	ctx := context.Background()
	// Run migration.
	err = client.Schema.Create(
		ctx,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	)
	if err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
