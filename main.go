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
	mc := controllers.NewMemberController()

	// Members
	r.GET("/api/v1/members", mc.MembersIndex)
	r.POST("/api/v1/members", mc.MembersCreate)

	// Member singular
	r.GET("/api/v1/member/:id", mc.MemberShow)
	r.PUT("/api/v1/member/:id", mc.MemberUpdate)
	r.DELETE("/api/v1/member/:id", mc.MemberDelete)

	fmt.Println("Starting server on :8080")
	http.ListenAndServe(":8080", &Server{r})
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

type Server struct {
  r *httprouter.Router
}

func (s *Server) ServeHTTP (w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Access-Control-Allow-Origin", "*")
  s.r.ServeHTTP(w, r)
}
