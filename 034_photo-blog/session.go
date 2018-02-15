package main

import (
	"fmt"
	"log"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

func createCookie(w http.ResponseWriter, req *http.Request) {
	sID, _ := uuid.NewV4()
	c := &http.Cookie{
		Name:  "session",
		Value: sID.String(),
	}
	http.SetCookie(w, c)
	stmt, err := db.Prepare("INSERT INTO session (user_id, session_id) VALUES(?, ?)")
	check(err)
	defer stmt.Close()

	r, err := stmt.Exec(u.ID, c.Value)
	check(err)
	n, err := r.RowsAffected()
	check(err)
	fmt.Println(n)
}

// func getUser(w http.ResponseWriter, req *http.Request) user {
// 	// get cookie
// 	c, err := req.Cookie("session")
// 	if err != nil {
// 		sID, _ := uuid.NewV4()
// 		c = &http.Cookie{
// 			Name:  "session",
// 			Value: sID.String(),
// 		}
//
// 	}
// 	http.SetCookie(w, c)
//
// 	// if the user exists already, get user
// 	var u user
// 	if un, ok := dbSessions[c.Value]; ok {
// 		u = dbUsers[un]
// 	}
// 	return u
// }

//
func alreadyLoggedIn(req *http.Request) bool {
	c, err := req.Cookie("session")
	if err != nil {
		return false
	}
	id := findUserId(c.Value)
	findUserById(id)
	return true
}

func findUserId(i string) int64 {
	var id int64
	err := db.QueryRow("SELECT user_id FROM session where session_id = ?", i).Scan(&id)
	if err != nil {
		log.Fatal(err)
	}
	return id
}

func findUserById(i int64) {
	var name string
	var p []byte
	var id int
	err := db.QueryRow("SELECT username, password, _id FROM user where _id = ?", i).Scan(&name, &p, &id)
	if err != nil {
		log.Fatal(err)
	}
	u = user{name, p, id}
}

func findByUsername(un string) {
	var name string
	var p []byte
	var id int
	err := db.QueryRow("SELECT username, password, _id FROM user where username = ?", un).Scan(&name, &p, &id)
	if err != nil {
		log.Fatal(err)
	}
	u = user{name, p, id}
}
