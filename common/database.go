package common

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/hewo233/cxsj_homework1/model"
	"log"
)

const (
	dbName = "cxsj1db"
	dbUser = "hewo"
	dbHost = "localhost:3306"
)

var UserDB *sql.DB

// InitDB init DB, return a pointer to sql.DB
func InitDB(dbPassword string) (*sql.DB, error) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s", dbUser, dbPassword, dbHost, dbName))
	if err != nil {
		log.Println("Error opening database connection:", err)
		return nil, err
	}
	if err = db.Ping(); err != nil {
		log.Println("Error connecting to database:", err)
		return nil, err
	}
	println("success connect to doc")
	return db, nil
}

// CloseDB close DB by pointer
func CloseDB(db *sql.DB) {
	if db != nil {
		err := db.Close()
		if err != nil {
			log.Println("Error closing database connection:", err)
			return
		}
	}
}

// AddRecord add record to DB
func AddRecord(info *model.CxsjUser, db *sql.DB) error {
	result, err := db.Exec("INSERT INTO users (email, name, password, gender) VALUES (?, ?, ?, ?)", info.Email, info.Name, info.Password, info.Gender)
	if err != nil {
		log.Println("Error adding record to database:", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Error getting rows affected:", err)
		return err
	}
	fmt.Printf("Add Rows affected: %d\n", rowsAffected)

	return nil

}

// QueryRecordByEmail query record by email
func QueryRecordByEmail(email string, db *sql.DB) (*model.CxsjUser, error) {
	rows, err := db.Query("SELECT * FROM users WHERE email = ?", email)
	if err != nil {
		log.Println("Error querying record by email: ", err)
		return nil, err
	}

	defer rows.Close()

	var user model.CxsjUser
	for rows.Next() {
		err := rows.Scan(&user.ID, &user.Email, &user.Name, &user.Password, &user.Gender)
		if err != nil {
			log.Println("Error scanning record: ", err)
			return nil, err
		}

		return &user, nil
	}
	log.Println("No record found by email: ", email)
	return nil, nil
}

// ModifyRecord modify record (email is key, should not be modified)
func ModifyRecord(info *model.CxsjUser, db *sql.DB) error {
	result, err := db.Exec("UPDATE users SET name = ?, password = ?, Gender = ? WHERE email = ?",
		info.Name, info.Password, info.Gender, info.Email)
	if err != nil {
		log.Println("Error modifying record: ", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Error getting rows affected: ", err)
		return err
	}

	fmt.Printf("Modify Rows affected: %d\n", rowsAffected)
	return nil

}

// DeleteRecord delete record
func DeleteRecord(email string, db *sql.DB) error {
	result, err := db.Exec("DELETE FROM users WHERE email = ?", email)
	if err != nil {
		log.Println("Error deleting record: ", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Error getting rows affected: ", err)
		return err
	}

	fmt.Printf("Delete Rows affected: %d\n", rowsAffected)

	return nil
}

func ListAllRecords(db *sql.DB) ([]model.CxsjUser, error) {

	var usersList []model.CxsjUser

	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		log.Println("Error querying all records: ", err)
		return nil, err
	}
	for rows.Next() {

		var user model.CxsjUser
		err := rows.Scan(&user.ID, &user.Email, &user.Name, &user.Password, &user.Gender)

		usersList = append(usersList, user)

		if err != nil {
			log.Println("Error scanning record: ", err)
			return nil, err
		}
	}

	return usersList, nil
}
