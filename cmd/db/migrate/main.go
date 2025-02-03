package main

import (
	"fmt"
	"log"
	"os"

	_ "c0fee-api/db/migrations" // マイグレーション関数の登録

	"github.com/joho/godotenv"
	"github.com/pressly/goose/v3"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln(err)
	}
	// データベース接続文字列
	dbString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PW"), os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_DB"))

	// PostgreSQLへの接続
	db, err := goose.OpenDBWithDriver("postgres", dbString)
	if err != nil {
		log.Fatalf("failed to open DB: %v", err)
	}
	defer db.Close()

	// コマンドライン引数を解析してgooseコマンドを実行
	if len(os.Args) < 2 {
		log.Fatalf("usage: %s", "\n up\n down\n status\n version\n")
	}

	var args []string
	if len(os.Args) >= 2 {
		for _, arg := range os.Args[2:] {
			args = append(args, arg)
		}
	}

	if err := goose.Run(os.Args[1], db, "db/migrations", args...); err != nil {
		log.Fatalf("failed to run goose command: %v", err)
	}
}
