package main

import (
	"c0fee-api/db"
	"c0fee-api/db/seeds"
	"log"
)

func main() {
	// データベース接続
	dbConn := db.NewDB()

	// シードデータの挿入
	if err := seeds.CreateUserSeeds(dbConn); err != nil {
		log.Fatalf("failed to seed users: %v", err)
	}
	if err := seeds.CreateRoasterSeeds(dbConn); err != nil {
		log.Fatalf("failed to seed roasters: %v", err)
	}
	// if err := seeds.CreateVarietySeeds(dbConn); err != nil {
	// 	log.Fatalf("failed to seed varieties: %v", err)
	// }
	// if err := seeds.CreateCountriesSeeds(dbConn); err != nil {
	// 	log.Fatalf("failed to seed countries: %v", err)
	// }
	if err := seeds.CreateProcessMethodSeeds(dbConn); err != nil {
		log.Fatalf("failed to seed process methods: %v", err)
	}
	if err := seeds.CreateBeanSeeds(dbConn); err != nil {
		log.Fatalf("failed to seed beans: %v", err)
	}

	log.Println("All seeds have been successfully inserted.")
}
