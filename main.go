/* My Band App for Angie */

package main

import (
	"encoding/json"
	"fmt"
	"github.com/edmore/my-band/controllers"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type App struct {
	Name    string `json:"app_name"`
	Version string `json:"version"`
}

func main() {
	r := httprouter.New()
	r.GET("/api/v1", Root)

	// Members
	r.GET("/api/v1/members", controllers.MembersIndex)
	r.POST("/api/v1/members", controllers.MembersCreate)

	// Member singular
	r.GET("/api/v1/member/:id", controllers.MemberShow)
	r.PUT("/api/v1/member/:id", controllers.MemberUpdate)
	r.DELETE("/api/v1/member/:id", controllers.MemberDelete)

	fmt.Println("Starting server on :8080")
	http.ListenAndServe(":8080", r)
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
