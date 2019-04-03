//
// S E R V E R
//
// Basic handling, etc. Page formation is the domain of frontend.go
//

package main

import (
	"fmt"
	"regexp"
	"net/http"
	"net/http/httputil"
	"strconv"
	//"time"
	"errors"
)

//Static regexes for dispatching
var staticRegex = regexp.MustCompile(`/static/(.+)$`)
var fileRegex = regexp.MustCompile(`/file/(.+)$`)
var blogPostRegex = regexp.MustCompile(`(?:([0-9]{4})(?:/([0-9]{1,2})(?:/(?:(\w*)))?)?)$`)
var textRegex = regexp.MustCompile(``)

//Server structure
type Server struct {
	db *Database
	srv *http.Server
}

func (srv *Server) Initialize(port int, databasePath string, databasePort int, username string, password string, database string) error {
	//We are our own handler
	srv.srv = &http.Server{Addr: fmt.Sprintf(":%d", port), Handler: srv}
	srv.db = &Database{}
	return srv.db.Connect(databasePath, databasePort, username, password, database)
}

func (srv *Server) Run() error {
	return srv.srv.ListenAndServe()
}

func (srv *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var s []string

	method := r.Method
	path := r.URL.Path

	fmt.Printf("'%s' <%s>\n", method, path)

	fmt.Printf("sf %v\n", staticRegex.FindStringSubmatch(path))

	if path == "/" {
		fmt.Println("SERVING HOMEPAGE")
		srv.ServeHomePage(w, r)
	} else if s=blogPostRegex.FindStringSubmatch(path); method=="GET" && len(s)>1 {
		fmt.Printf("SERVING BLOG POST <%v>\n", s)
		switch len(s) {
			case 2:
				y, _ := strconv.Atoi(s[1])
				srv.ServeBlogYear(w, r, y)
			case 3:
				y, _ := strconv.Atoi(s[1])
				m, _ := strconv.Atoi(s[2])
				srv.ServeBlogMonth(w, r, y, m)
			case 4:
				fmt.Println("YEAR-MONTH-TITLE")
				y, _ := strconv.Atoi(s[1])
				m, _ := strconv.Atoi(s[2])
				st := s[3]
				srv.ServeBlogPost(w, r, y, m, st)
		}
	} else if s=fileRegex.FindStringSubmatch(path); method=="GET" && len(s)>0 {
		fmt.Printf("SERVING FILE <%s>\n", s[1])
		//srv.ServeFile(w, r, s[1])
	} else if s=staticRegex.FindStringSubmatch(path); method=="GET" && len(s)>0 {
		fmt.Printf("SERVING STATIC FILE <%v>\n", s)
		//srv.ServeStatic(w, r, s[1])
		http.ServeFile(w, r, fmt.Sprintf("/home/nicolas/www/static/%s", s[1]))
	} else if method=="GET" && path=="/newblogpost" {
		srv.ServeNewBlogPost(w, r)
	} else if method=="POST" && path=="/newblogpost" {
		srv.ServeReflection(w, r)
		//srv.CreateBlogPost(w, r)
	} else if method=="POST" && path=="/newblogpostcomment" {
	} else if method=="POST" && path=="/likeblogpost" {
	} else if (method=="GET"||method=="POST") && path=="/editblogpost" {
	} else if method=="POST" && path=="/likeblogpost" {
	} else if method=="POST" && path=="/search" {
		srv.ServeSearch(w, r)
	} else if path=="/error" {
		srv.ServeError(w, errors.New("E9999 ERROR PAGE REQUESTED"))
	} else {
		w.Write([]byte("NOT FOUND"))
	}
}

func (srv *Server) ServeHomePage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ALLO"))
}

func (srv *Server) ServeBlogYear(w http.ResponseWriter, r *http.Request, year int) {

}

func (srv *Server) ServeBlogMonth(w http.ResponseWriter, r *http.Request, year int, month int) {

}

func (srv *Server) ServeBlogPost(w http.ResponseWriter, r *http.Request, year int, month int, shortTitle string) {
	//We first get the id of the post
	tx, _ := srv.db.Begin()
	id, err := tx.GetBlogPostNumber(year, month, shortTitle)
	if err != nil {
		srv.ServeError(w, err)
		return
	}
	fmt.Printf("BLOG POST ID %d\n", id)

	title, authorusername, author, created, modified,  body, err := tx.GetBlogPostBody(id)
	if err != nil {
		srv.ServeError(w, err)
	}

	WriteGeneralHeader(w, title, "")
	WriteBlogPostBody(w, id, title, authorusername, author, created, modified,  body, nil, nil)
	WriteGeneralTrailer(w)
}

//
//
//

func (srv *Server) ServeLogin(w http.ResponseWriter, r *http.Request) {

}

func (srv *Server) ServeLoginPOST(w http.ResponseWriter, r *http.Request) {

}

//
//New blog post
//

//Serves the edit screen.
func (srv *Server) ServeNewBlogPost(w http.ResponseWriter, r *http.Request) {
	WriteGeneralHeader(w, "New Blog Post", "")
	WriteNewBlogPost(w, "", "", "", "MARKDOWN", false, []string{"E0158 SYNTAX ERROR"}, nil)
	WriteGeneralTrailer(w)
}

//This creates a blog post. If successful it redirects to its url or if not
//successful redirects to edit screen (with error msg).
func (srv *Server) CreateBlogPost(w http.ResponseWriter, r *http.Request) {

}

//Search
func (srv *Server) ServeSearch(w http.ResponseWriter, r *http.Request) {

}





func main() {
	s := &Server{}
	s.Initialize(4000, "localhost", 5432, "postgres", "postgres", "www")
	s.Run()
}







//
// S T A T I C   A N D   F I L E    S E R V I N G
//

func (srv *Server) ServeStatic(w http.ResponseWriter, file string) {

}

func (srv *Server) ServeFile(w http.ResponseWriter, file string) {

}


//
// E R R O R S
//
func (srv *Server) ServeError(w http.ResponseWriter, err error) {
	w.WriteHeader(500)
	WriteError(w, err)
}

//
// D E B U G
//
func (srv *Server) ServeReflection(w http.ResponseWriter, r *http.Request) {
	b, e := httputil.DumpRequest(r, true)
	if e != nil {
		panic(e)
	}
	w.Write(b)
}
