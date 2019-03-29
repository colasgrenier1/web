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
	"strconv"
	"time"
)

//Static regexes for dispatching
var staticRegex = regexp.MustCompile(`/static/(.+)$`)
var fileRegex = regexp.MustCompile(`/file/(.+)$`)
var blogPostRegex = regexp.MustCompile(`(?:([0-9]{4})(?:/([0-9]{1,2})(?:/(?:(\w*)))?)?)$`)
var textRegex = regexp.MustCompile(``)

//Server structure
type Server struct {
	//db Database
	srv *http.Server
}

func (srv *Server) Initialize(port int) {
	//We are our own handler
	srv.srv = &http.Server{Addr: fmt.Sprintf(":%d", port), Handler: srv}
}

func (srv *Server) Run() {
	srv.srv.ListenAndServe()
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
	} else if method=="POST" && path=="/newpost" {
		srv.CreateBlogPost(w, r)
	} else if method=="POST" && path=="/newblogpostcomment" {
	} else if method=="POST" && path=="/likeblogpost" {
	} else if (method=="GET"||method=="POST") && path=="/editblogpost" {
	} else if method=="POST" && path=="/likeblogpost" {
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

func (srv *Server) ServeBlogPost(w http.ResponseWriter, r *http.Request, year int, month int, title string) {
	WriteGeneralHeader(w, "a", "")
	WriteBlogPostBody(w, 123, "Some Blog Post", "coalsgrenier", "Nicolas Grenier", time.Now(), time.Now(), `
<p>Lorem ipsum dolor sit amet, consectetur adipiscing elit. Morbi ullamcorper, magna non rutrum mollis, purus metus rhoncus est, ac tempus nisl est et dui. Quisque feugiat semper leo, nec sollicitudin ex porttitor a. Pellentesque et turpis quis sapien elementum pharetra. Nulla tellus enim, ullamcorper sed lectus eu, luctus tempor elit. Quisque lacinia tellus ac dui faucibus maximus. Ut et lacus semper, mollis mauris sit amet, auctor nunc. Nunc nec porttitor leo. Phasellus consequat nulla dui, tristique tincidunt est mollis nec. Class aptent taciti sociosqu ad litora torquent per conubia nostra, per inceptos himenaeos. Cras posuere non odio a volutpat. Mauris sed rhoncus enim. Morbi vitae elit a sapien tempor vestibulum. Curabitur ullamcorper elit ex, quis pellentesque nulla pellentesque ut.

<p>Nulla facilisi. Praesent ullamcorper eget eros quis commodo. Etiam non lobortis nulla. Nullam ut sollicitudin felis. Morbi lacus ante, rutrum a facilisis et, hendrerit ut augue. Curabitur leo risus, pellentesque et rutrum sit amet, sagittis sit amet tortor. Suspendisse sed nulla vel est aliquam mollis ac ac urna. Praesent at nulla mauris. Cras lacinia nulla eget purus vestibulum varius eget ac velit. Sed vel est quis ante hendrerit ultrices vel vitae justo. Donec at orci purus. Suspendisse euismod finibus metus in tempor. Aenean accumsan tincidunt metus, nec pharetra elit vulputate quis.

<p>Vestibulum vitae justo viverra, mollis magna et, scelerisque metus. Nunc vel malesuada neque. Sed lacinia laoreet erat quis bibendum. Sed metus massa, bibendum et velit eget, pellentesque porta nisi. Suspendisse tincidunt magna pulvinar nulla dignissim consequat. Interdum et malesuada fames ac ante ipsum primis in faucibus. Aenean maximus tellus et justo viverra fermentum. Cras id justo ultrices, hendrerit sem et, condimentum leo. Etiam id ornare metus. Vivamus blandit at neque ullamcorper sagittis. Vestibulum ultrices bibendum porta. Donec non dictum dolor. Quisque fermentum ligula ut aliquet imperdiet. Curabitur sit amet nisi hendrerit, venenatis mauris sit amet, elementum neque.

<p>Phasellus eleifend nisi mauris, vitae dapibus dui egestas finibus. Phasellus ac lacus ipsum. Cras tristique sapien id turpis aliquet luctus. Praesent et leo volutpat, aliquet lorem non, pulvinar purus. Interdum et malesuada fames ac ante ipsum primis in faucibus. Nullam interdum facilisis enim, non malesuada libero commodo eu. Fusce pulvinar quam tincidunt facilisis sodales. Fusce blandit elit id orci tristique, in condimentum nunc commodo. Cras sed finibus eros. Suspendisse est elit, varius eget feugiat non, scelerisque vel magna. Aliquam in tristique neque, in lacinia mauris. Integer et imperdiet felis, ut semper ex. Duis venenatis porttitor vulputate. Vivamus eu porta nisi, eu convallis quam. Sed at lacus neque.

<p>Praesent aliquam ut leo sit amet volutpat. Sed a est rhoncus, tempor metus sit amet, volutpat neque. Maecenas molestie dictum metus, non porttitor enim pretium in. Maecenas arcu nisl, malesuada in ex vel, aliquet aliquet nisl. Curabitur vel est purus. Suspendisse porta diam lacus, eget scelerisque nisi aliquet a. Quisque ac maximus risus, quis volutpat arcu. Praesent ac turpis quis augue laoreet bibendum vel sed lectus. Nulla commodo urna ac dolor iaculis, sit amet venenatis ipsum congue. In hac habitasse platea dictumst. Vivamus mi urna, egestas malesuada eros ut, condimentum eleifend sapien. Maecenas sodales vulputate risus, at laoreet enim condimentum sit amet. Donec eget lectus suscipit, vulputate nunc vitae, pretium libero. Integer vitae vulputate odio. Interdum et malesuada fames ac ante ipsum primis in faucibus. `, nil, nil)
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







func main() {
	s := &Server{}
	s.Initialize(4000)
	s.Run()
}







//
// S T A T I C   A N D   F I L E    S E R V I N G
//

func (srv *Server) ServeStatic(w http.ResponseWriter, r *http.Request, file string) {

}

func (srv *Server) ServeFile(w http.ResponseWriter, r *http.Request, file string) {

}
