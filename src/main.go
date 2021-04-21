package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
)
type Person struct {
	First_Name     string `json:"first_name"`
	Last_name string `json:"last_name"`
	Age int `json:"age"`
	Phone int64 `json:"phone"`
	Address string `json:"address"`


}


func GETHandler(w http.ResponseWriter, r *http.Request) {
	db := OpenConnection()

	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		log.Fatal(err)
	}

	var people []Person

	for rows.Next() {
		var person Person
		rows.Scan(&person.First_Name, &person.Last_name,&person.Age,&person.Phone,&person.Address)
		people = append(people, person)
	}

	peopleBytes, _ := json.MarshalIndent(people, "", "\t")

	w.Header().Set("Content-Type", "application/json")
	w.Write(peopleBytes)

	defer rows.Close()
	defer db.Close()
}

	func main() {
		http.HandleFunc("/", GETHandler)

		log.Fatal(http.ListenAndServe(":8080", nil))
	}

	// load .env file from given path
	// we keep it empty it will load .env from current directory
func OpenConnection() *sql.DB {
	err := godotenv.Load("ENV.env")
	if err != nil {
		log.Fatal("Error loading env file \n", err)
	}
	var db *sql.DB

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))

	log.Print("Connecting to PostgreSQL DB...")
	db, err = sql.Open("postgres",dsn)


	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2)

	}
	log.Println("connected")
	return db;

}