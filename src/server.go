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
	"strings"
	"os"
	"io"
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
	staticDirectory string
	fileDirectory string
}

func (srv *Server) Initialize(port int, databasePath string, databasePort int, username string, password string, database string, staticDirectory string, fileDirectory string) error {
	//We are our own handler
	srv.staticDirectory = staticDirectory
	srv.fileDirectory = fileDirectory
	srv.srv = &http.Server{Addr: fmt.Sprintf(":%d", port), Handler: srv}
	srv.db = &Database{}
	return srv.db.Connect(databasePath, databasePort, username, password, database)
}

func (srv *Server) Run() error {
	return srv.srv.ListenAndServe()
}

func (srv *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var s []string
	var p string

	method := r.Method
	path := r.URL.Path

	fmt.Printf("'%s' <%s>\n", method, path)

	fmt.Printf("sf %v\n", staticRegex.FindStringSubmatch(path))

	if path == "/" {
		fmt.Println("SERVING HOMEPAGE")
		srv.ServeHomePage(w, r)
	} else if s=blogPostRegex.FindStringSubmatch(path); method=="GET" && len(s)>1 {
		fmt.Printf("SERVING BLOG POST <%v>\n", s)
		if s[3] != "" {
			fmt.Println("YEAR-MONTH-TITLE")
			y, _ := strconv.Atoi(s[1])
			m, _ := strconv.Atoi(s[2])
			st := s[3]
			srv.ServeBlogPost(w, r, y, m, st)
		} else if s[2] != "" {
			fmt.Println("YEAR-MONTH")
			y, _ := strconv.Atoi(s[1])
			m, _ := strconv.Atoi(s[2])
			srv.ServeBlogMonth(w, r, y, m)
		} else {
			fmt.Println("YEAR")
			y, _ := strconv.Atoi(s[1])
			srv.ServeBlogYear(w, r, y)
		}
	//} else if s=fileRegex.FindStringSubmatch(path); method=="GET" && len(s)>0 {
	} else if p=strings.TrimPrefix(path, "/file/"); p!=path {
		fmt.Printf("SERVING FILE <%s>\n", p)
		//srv.ServeFile(w, r, s[1])
	//} else if s=staticRegex.FindStringSubmatch(path); method=="GET" && len(s)>0 {
		http.ServeFile(w, r, fmt.Sprintf("%s/%s", srv.fileDirectory, s[1]))
	} else if p=strings.TrimPrefix(path, "/static/"); p!=path {
		fmt.Printf("SERVING STATIC FILE <%v>\n", p)
		fmt.Printf("%s/%s", srv.staticDirectory, p)
		//srv.ServeStatic(w, r, s[1])
		//http.ServeFile(w, r, fmt.Sprintf("/home/nicolas/www/static/%s", s[1]))
		http.ServeFile(w, r, fmt.Sprintf("%s/%s", srv.staticDirectory, p))
	} else if method=="GET" && path=="/newblogpost" {
		srv.ServeNewBlogPost(w, r)
	} else if method=="POST" && path=="/newblogpost" {
		srv.CreateBlogPost(w, r)
	} else if method=="GET" && path=="/uploadfile" {
		srv.ServeUploadFile(w, r)
	} else if method=="POST" && path=="/uploadfile" {
		srv.UploadFile(w, r)
	} else if path == "/login" {
		srv.ServeLogin(w, r)
	} else if path == "/disclaimer" {

	} else if path == "/contact" {

	} else if path == "/about" {

	} else if method=="POST" && path=="/newblogpostcomment" {
		srv.ServeError(w, NotImplementedError())
	} else if method=="POST" && path=="/likeblogpost" {
		srv.ServeError(w, NotImplementedError())
	} else if (method=="GET"||method=="POST") && path=="/editblogpost" {
		srv.ServeError(w, NotImplementedError())
	} else if method=="POST" && path=="/likeblogpost" {
		srv.ServeError(w, NotImplementedError())
	} else if method=="POST" && path=="/search" {
		srv.ServeError(w, NotImplementedError())
		//srv.ServeSearch(w, r)
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
	tx, _ := srv.db.Begin()
	//Returns
	postids, err := tx.GetBlogPostsForMonth(year, month)
	if err != nil {
		srv.ServeError(w, err)
		return
	}
	fmt.Printf("BLOG POST PER MONTH %04d/%02d\n", year, month)

	WriteGeneralHeader(w, fmt.Sprintf("Blog Posts for %d/%02d", year, month), "")

	for _, id := range postids {
		title, shorttitle, authorusername, author, created, modified, body, _ := tx.GetBlogPostBody(id)
		WriteBlogPostBodyOverview(w, id, shorttitle, strings.ToUpper(title), authorusername, author, created, modified,  body, nil, nil)
	}

	pyear, pmonth, err := tx.GetPreviousBlogYearMonth(year, month)
	nyear, nmonth, err := tx.GetNextBlogYearMonth(year, month)

	fmt.Printf("%d %d %d %d", pyear, pmonth, nyear, nmonth)

	WriteBlogPostPreviousNextMonth(w, pyear, pmonth, nyear, nmonth)

	WriteGeneralTrailer(w)

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

	title, _, authorusername, author, created, modified,  body, err := tx.GetBlogPostBody(id)
	if err != nil {
		srv.ServeError(w, err)
	}

	WriteGeneralHeader(w, title, "")
	WriteBlogPostBody(w, id, strings.ToUpper(title), authorusername, author, created, modified,  body, nil, nil)
	WriteGeneralTrailer(w)
}

//
//
//

func (srv *Server) ServeLogin(w http.ResponseWriter, r *http.Request) {
	WriteLoginScreen(w, "", nil, nil)
}

func (srv *Server) ServeUploadFile(w http.ResponseWriter, r *http.Request) {
	WriteGeneralHeader(w, "Upload File", "")
	WriteUploadFile(w, "", nil, nil)
	WriteGeneralTrailer(w)
}

func (srv *Server) UploadFile(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10000000)
	fmt.Println(r.Form["filename"])
	filename := r.Form["filename"]
	file, _, err := r.FormFile("file")
	if err != nil {
		WriteError(w, err)
		return
	}
	defer file.Close()
	target, err := os.OpenFile(fmt.Sprintf("%s/%s", srv.fileDirectory, filename), os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		WriteError(w, err)
		return
	}
	defer target.Close()
	io.Copy(target, file)
	WriteGeneralHeader(w, "Upload File", "")
	WriteUploadFile(w, "", nil, nil)
	WriteGeneralTrailer(w)
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
	s.Initialize(4000, "localhost", 5432, "postgres", "postgres", "www", "/home/nicolas/www/static", "/home/nicolas/www/file")
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
