package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

type user struct {
	UserName string
	Password []byte
	ID       int
}

var tpl *template.Template
var db *sql.DB
var err error
var u user

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	db, err = sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/photo_blog")
	check(err)
	defer db.Close()

	err = db.Ping()
	check(err)

	http.HandleFunc("/", index)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	if !alreadyLoggedIn(r) {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(w, "index.gohtml", u.UserName)
}

func login(w http.ResponseWriter, r *http.Request) {
	if alreadyLoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	if r.Method == http.MethodPost {
		un := r.FormValue("username")
		p := r.FormValue("password")
		findByUsername(un)
		err := bcrypt.CompareHashAndPassword(u.Password, []byte(p))
		if err != nil {
			http.Error(w, "Username and/or password do not match", http.StatusForbidden)
			return
		}
		// create session
		createCookie(w, r)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	tpl.ExecuteTemplate(w, "login.gohtml", nil)
}

func signup(w http.ResponseWriter, r *http.Request) {
	if alreadyLoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	if r.Method == http.MethodPost {
		un := r.FormValue("username")
		p := r.FormValue("password")
		bs, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.MinCost)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		stmt, err := db.Prepare("INSERT INTO user (username, password) VALUES(?, ?)")
		check(err)
		defer stmt.Close()

		res, err := stmt.Exec(un, bs)
		check(err)
		id, _ := res.LastInsertId()
		findUserById(id)
		createCookie(w, r)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	tpl.ExecuteTemplate(w, "signup.gohtml", nil)
}

func logout(w http.ResponseWriter, r *http.Request) {
	c, _ := r.Cookie("session")
	stmt, err := db.Prepare("DELETE FROM session where session_id = ?")
	check(err)
	defer stmt.Close()

	res, err := stmt.Exec(c.Value)
	check(err)
	fmt.Println(res)

	c = &http.Cookie{
		Name:   "session",
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(w, c)
	u = user{}
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

// func getUserId(i int) int {
// 	var id int
// 	err := db.QueryRow("SELECT user_id FROM session where session_id = ?", i).Scan(&id)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	u = user{name, p, id}
// }
