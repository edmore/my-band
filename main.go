/* My Band App for Angie */

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"database/sql"
	_ "github.com/lib/pq"

	"github.com/julienschmidt/httprouter"
)

type App struct {
	Name    string `json:"app_name"`
	Version string `json:"version"`
}

type Member struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Speciality string `json:"speciality"`
}

type Members []Member

var db *sql.DB
var name string
var surname string
var speciality string
var id int

func main() {
	r := httprouter.New()
	r.GET("/api/v1", Root)

	// Members
	r.GET("/api/v1/members", MembersIndex)
	r.POST("/api/v1/members", MembersCreate)

	// Member singular
	r.GET("/api/v1/member/:id", MemberShow)
	r.PUT("/api/v1/member/:id", MemberUpdate)
	r.DELETE("/api/v1/member/:id", MemberDelete)
	r.GET("/api/v1/member/:id/edit", MemberEdit)

	fmt.Println("Starting server on :8080")
	http.ListenAndServe(":8080", r)
}

func init() {
	var err error

	db, err = sql.Open("postgres", "user=edmoremoyo dbname=band sslmode=disable")
	if err != nil {
		log.Fatal("Error: The data source arguments are not valid")
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
}

func Root(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	app := App{"My Band", "1.0"}
	js, err := json.Marshal(app)

	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.Write(js)
}

func MembersIndex(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintln(rw, "Members index")
}

func MembersCreate(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintln(rw, "Member create")
}

func MemberShow(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Process
	err := db.QueryRow("SELECT * FROM members WHERE id=$1", p.ByName("id")).Scan(&id, &name, &surname, &speciality)
	if err == sql.ErrNoRows {
		http.Error(rw, err.Error(), http.StatusNotFound)
		return
	}
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	// Response
	member := Member{id, name, surname, speciality}
	js, err := json.Marshal(member)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(js)
}

func MemberUpdate(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintln(rw, "Member update")
}

func MemberDelete(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintln(rw, "Member delete")
}

func MemberEdit(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintln(rw, "Member edit")
}
