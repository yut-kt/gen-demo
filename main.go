package main

import (
	"gen-demo/gen/query"
	"log"
	"os"

	slogGorm "github.com/orandin/slog-gorm"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log/slog"
)

func main() {
	l := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))
	slog.SetDefault(l)

	handler := slog.With("layer", "database")
	gormLogger := slogGorm.New(slogGorm.WithHandler(handler.Handler()))

	db, err := gorm.Open(mysql.Open("root:passwd@tcp(localhost:3306)/db?charset=utf8mb4&parseTime=True&loc=Local"),
		&gorm.Config{Logger: gormLogger},
	)
	if err != nil {
		panic(err)
	}

	q := query.Use(db)

	p := q.Parent
	records, err := p.Joins(p.Child).Where(q.Child.Name.Eq("test")).Find()
	if err != nil {
		panic(err)
	}

	for _, record := range records {
		log.Printf("record: %+v", record)
	}
}
