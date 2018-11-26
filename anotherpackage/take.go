package takedatabase

import (
	"context"
	"database/sql"
	"log"
)

var userDatabase []string

var ctx = context.Background()

var databaseString string = "root:1111@tcp(127.0.0.1:3306)/oop"

// Inserttodatabase
func InsertToDatbase(studentInfor [6]string) {

	for i := 0; i < 6; i++ {
		log.Println("database:", studentInfor[i])
	}

	//opendatabase
	db, err := sql.Open("mysql", databaseString)
	CheckErr(err)

	stmt, err := db.Prepare("insert student set idstudent = ?,name = ?,more_info= ?, lesson =? ")
	CheckErr(err)

	// stmt, err := db.Prepare("insert student set idstudent = ?,name = ?,phone= ?, gmail =? ")
	// CheckErr(err)

	res, err := stmt.Exec(studentInfor[3], studentInfor[0], studentInfor[2], studentInfor[1])
	CheckErr(err)

	lastID, err := res.LastInsertId()
	CheckErr(err)

	log.Println(lastID)

	db.SetMaxOpenConns(1)

	defer db.Close()
}

func CancelLesson(ID string) string {

	var id string
	var name string
	var lesson string
	db, err := sql.Open("mysql", databaseString)
	CheckErr(err)

	rows, err := db.Query("SELECT name, idstudent, lesson FROM student WHERE idstudent =?", ID)

	for rows.Next() {
		err := rows.Scan(&name, &id, &lesson)
		CheckErr(err)

		log.Printf("name  and id : %s , %s  ", name, id)
	}
	defer rows.Close()

	return lesson
}

// CheckErr
func CheckErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
