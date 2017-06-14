package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"database/sql"
	_ "github.com/lib/pq"

	"github.com/edmore/my-band/models"
	"github.com/julienschmidt/httprouter"
)

var db *sql.DB

// Member variables
var name string
var surname string
var speciality string
var id int

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

// Member Controllers
func MembersIndex(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	rows, err := db.Query("SELECT * FROM members")
	if err == sql.ErrNoRows {
		http.Error(rw, err.Error(), http.StatusNotFound)
		return
	}

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Response
	members := make([]*models.Member, 0)
	for rows.Next() {
		member := new(models.Member)
		err := rows.Scan(&member.Id, &member.Name, &member.Surname, &member.Speciality)
		if err != nil {
			log.Fatal(err)
		}
		members = append(members, member)
	}

	js, err := json.Marshal(members)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(js)
}

func MembersCreate(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	decoder := json.NewDecoder(r.Body)
	member := models.Member{}
	err := decoder.Decode(&member)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	_, err = db.Exec("INSERT INTO members(name, surname, speciality) VALUES($1, $2, $3)",
		member.Name, member.Surname, member.Speciality)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
}

func MemberShow(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Process
	err := db.QueryRow("SELECT * FROM members WHERE id=$1", p.ByName("id")).Scan(&id, &name, &surname, &speciality)
	if err == sql.ErrNoRows {
		http.Error(rw, "Member not found", http.StatusNotFound)
		return
	}
	if err != nil {
		log.Fatal(err)
	}

	// Response
	member := models.Member{id, name, surname, speciality}
	js, err := json.Marshal(member)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(js)
}

func MemberUpdate(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	decoder := json.NewDecoder(r.Body)
	member := models.Member{}
	err := decoder.Decode(&member)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()
	var result sql.Result

	// Some refactoring required here ...
	if member.Name != "" {
		result, err = db.Exec("UPDATE members SET name=$1 where id=$2",
			member.Name, p.ByName("id"))
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	if member.Surname != "" {
		result, err = db.Exec("UPDATE members SET surname=$1 where id=$2",
			member.Surname, p.ByName("id"))
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	if member.Speciality != "" {
		result, err = db.Exec("UPDATE members SET speciality=$1 where id=$2",
			member.Speciality, p.ByName("id"))
		if err != nil {
			log.Fatal(err)
		}
	}

	rowsAffected, err := result.RowsAffected()
	if rowsAffected == 0 {
		http.Error(rw, "Member not found", http.StatusNotFound)
		return
	}
}

func MemberDelete(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	result, err := db.Exec("DELETE FROM members where id=$1", p.ByName("id"))
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	rowsAffected, err := result.RowsAffected()
	if rowsAffected == 0 {
		http.Error(rw, "Member not found", http.StatusNotFound)
		return
	}
}
