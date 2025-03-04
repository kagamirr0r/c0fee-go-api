package main

import (
	"bufio"
	"c0fee-api/infrastructure/db"
	"c0fee-api/infrastructure/db/seeds"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	fmt.Println("テーブル名を入力してください (users/roasters/varieties/countries/areas/farms/farmers/process_methods/beans/bean_ratings/all):")

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("入力の読み取りに失敗しました: %v", err)
	}

	// 改行文字を取り除く
	targetTable := strings.TrimSpace(input)
	// データベース接続
	dbConn := db.NewDB()

	switch targetTable {
	case "users":
		if err := seeds.CreateUserSeeds(dbConn); err != nil {
			log.Fatalf("failed to seed users: %v", err)
		}
	case "roasters":
		if err := seeds.CreateRoasterSeeds(dbConn); err != nil {
			log.Fatalf("failed to seed roasters: %v", err)
		}
	case "varieties":
		if err := seeds.CreateVarietySeeds(dbConn); err != nil {
			log.Fatalf("failed to seed varieties: %v", err)
		}
	case "countries":
		if err := seeds.CreateCountriesSeeds(dbConn); err != nil {
			log.Fatalf("failed to seed countries: %v", err)
		}
	case "areas":
		if err := seeds.CreateAreasSeeds(dbConn); err != nil {
			log.Fatalf("failed to seed areas: %v", err)
		}
	case "farms":
		if err := seeds.CreateFarmsSeeds(dbConn); err != nil {
			log.Fatalf("failed to seed farms: %v", err)
		}
	case "farmers":
		if err := seeds.CreateFarmersSeeds(dbConn); err != nil {
			log.Fatalf("failed to seed farmers: %v", err)
		}
	case "process_methods":
		if err := seeds.CreateProcessMethodSeeds(dbConn); err != nil {
			log.Fatalf("failed to seed process methods: %v", err)
		}
	case "beans":
		if err := seeds.CreateBeanSeeds(dbConn); err != nil {
			log.Fatalf("failed to seed beans: %v", err)
		}
	case "bean_ratings":
		if err := seeds.CreateBeanRatingSeeds(dbConn); err != nil {
			log.Fatalf("failed to seed beans: %v", err)
		}
	case "all":
		if err := seeds.CreateUserSeeds(dbConn); err != nil {
			log.Fatalf("failed to seed users: %v", err)
		}
		if err := seeds.CreateRoasterSeeds(dbConn); err != nil {
			log.Fatalf("failed to seed roasters: %v", err)
		}
		if err := seeds.CreateVarietySeeds(dbConn); err != nil {
			log.Fatalf("failed to seed varieties: %v", err)
		}
		if err := seeds.CreateCountriesSeeds(dbConn); err != nil {
			log.Fatalf("failed to seed countries: %v", err)
		}
		if err := seeds.CreateAreasSeeds(dbConn); err != nil {
			log.Fatalf("failed to seed areas: %v", err)
		}
		if err := seeds.CreateFarmsSeeds(dbConn); err != nil {
			log.Fatalf("failed to seed farms: %v", err)
		}
		if err := seeds.CreateFarmersSeeds(dbConn); err != nil {
			log.Fatalf("failed to seed farmers: %v", err)
		}
		if err := seeds.CreateProcessMethodSeeds(dbConn); err != nil {
			log.Fatalf("failed to seed process methods: %v", err)
		}
		if err := seeds.CreateBeanSeeds(dbConn); err != nil {
			log.Fatalf("failed to seed beans: %v", err)
		}
		if err := seeds.CreateBeanRatingSeeds(dbConn); err != nil {
			log.Fatalf("failed to seed beans: %v", err)
		}
	default:
		log.Fatalf("invalid table name: %s", targetTable)
	}

	log.Println("Seeds have been successfully inserted.")
}
