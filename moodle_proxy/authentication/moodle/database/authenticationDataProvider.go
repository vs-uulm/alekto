package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/ma-zero-trust-prototype/moodle_proxy/env"
	"log"
)

func getUserIdBySessionId(sessionId string) int64 {

	var userId sql.NullInt64

	db, err := sql.Open("mysql", getDatabaseInfo())
	defer closeConnection(db, err)

	query := "SELECT t.userid " +
		"FROM bitnami_moodle.mdl_sessions t " +
		"WHERE sid='" + sessionId + "';"

	rows, err := db.Query(query)

	if rows != nil {

		for rows.Next() {

			err = rows.Scan(&userId)
			checkErr(err)
		}
	}

	return userId.Int64
}

func getValidSessionBySessionIdAndUsername(sessionId string, username string) bool {

	var SessionForUserExists sql.NullBool

	db, err := sql.Open("mysql", getDatabaseInfo())
	defer closeConnection(db, err)

	query := fmt.Sprintf("SELECT 1 "+
		"FROM bitnami_moodle.mdl_sessions s "+
		"JOIN bitnami_moodle.mdl_user u on u.id = s.userid "+
		"WHERE sid='%v' AND u.username = '%v'", sessionId, username)

	rows, err := db.Query(query)

	if rows != nil {
		for rows.Next() {
			err = rows.Scan(&SessionForUserExists)
			checkErr(err)
		}
	}

	return SessionForUserExists.Bool
}

func getAllRunningSessions() {

	db, err := sql.Open("mysql", getDatabaseInfo())
	defer closeConnection(db, err)

	rows, err := db.Query(
		"SELECT t.* " +
			"FROM bitnami_moodle.mdl_sessions t " +
			"ORDER BY state ASC LIMIT 501")

	for rows.Next() {

		var id sql.NullInt64
		var state sql.NullInt64
		var sid sql.NullString
		var userid sql.NullInt64
		var sessdata sql.NullString
		var timecreated sql.NullInt64
		var timemodified sql.NullInt64
		var firstip sql.NullString
		var lastip sql.NullString

		err = rows.Scan(&id, &state, &sid, &userid, &sessdata, &timecreated, &timemodified,
			&firstip, &lastip)

		checkErr(err)

		fmt.Printf("------------ \n")
		fmt.Printf("id: %d \n", id.Int64)
		fmt.Printf("userid: %d \n", userid.Int64)
		fmt.Println("sessionid: " + sid.String)
	}
}

func CheckConnection() bool {

	db, err := sql.Open("mysql", getDatabaseInfo())
	defer closeConnection(db, err)

	if err := db.Ping(); err != nil {
		log.Fatalln(err)
		return false
	}

	fmt.Println("Successfully connected!")
	return true
}

func getDatabaseInfo() string {

	mariaDbInfo := env.GetMariaDbInfo()

	fmt.Println(mariaDbInfo)

	return mariaDbInfo
}

func closeConnection(db *sql.DB, err error) {
	if err != nil {
		log.Fatalln(err)
	}

	err = db.Close()

	if err != nil {
		log.Fatalln(err)
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
