/* My Band App for Angie */

package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Member struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Speciality string `json:"speciality"`
}

type Members []Member

func main() {
	r := httprouter.New()
	r.GET("/api/", Home)

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

func Home(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintln(rw, "My Band ... My band!!!")
}

func MembersIndex(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintln(rw, "Members index")
}

func MembersCreate(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintln(rw, "Member create")
}

func MemberShow(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	member := Member{1, "Edmore", "Moyo", "vocalist"}

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
