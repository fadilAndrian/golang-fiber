package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strings"
)

var DB *sql.DB

func Connect() {
	dbUser := os.Getenv("DBUSER")
	dbName := os.Getenv("DBNAME")
	password := os.Getenv("PASSWORD")
	fmt.Println(dbUser, dbName, password, "test")
	connStr := fmt.Sprintf("user=%s dbname=%s sslmode=disable password=%s", dbUser, dbName, password)
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err.Error())
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Database connected")
	runMigration(db, "migrations")

	DB = db
}

func runMigration(db *sql.DB, migrationsDir string) {
	files, err := ioutil.ReadDir(migrationsDir)
	if err != nil {
		log.Fatal(err)
	}

	sort.Slice(files, func(i, j int) bool {
		return strings.Compare(files[i].Name(), files[j].Name()) < 0
	})

	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".sql") {
			migrationFilePath := fmt.Sprintf("%s/%s", migrationsDir, file.Name())

			migrationSQL, err := readMigrationFile(migrationFilePath)
			if err != nil {
				log.Fatalf("failed to read migration file %s: %v", file.Name(), err)
			}

			_, err = db.Exec(migrationSQL)
			if err != nil {
				log.Fatalf("failed to execute migration %s: %v", file.Name(), err)
			}

			fmt.Printf("Migration %s executed successfully\n", file.Name())
		}
	}
}

func readMigrationFile(filePath string) (string, error) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to read migration file: %v", err)
	}
	return string(content), nil
}
