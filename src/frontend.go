package main

import (
	"bytes"
	"fmt"
	"time"
	"strings"
	"net/http"
)

func format_time(t time.Time) string {
	//return t.Format("Mon, Jan 2, 2006  03:04:05 PM")
	return t.Format("Mon, Jan 2, 2006")
}

func format_now() string {
	return format_time(time.Now())
}

func WriteGeneralHeader(w http.ResponseWriter, title string, username string) {
	str := `<!DOCTYPE HTML>
	<HTML lang="en">
	<HEAD>
		<TITLE>%title%</TITLE>
		<LINK REL="stylesheet" HREF="/static/style.css">
		<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/4.7.0/css/font-awesome.min.css">
		<link href="https://fonts.googleapis.com/css?family=IBM+Plex+Serif:400,400i,700&amp;subset=latin-ext" rel="stylesheet">
	</HEAD>
	<BODY>
		<DIV CLASS="mainframe">
			<DIV CLASS="statusbar">
				<UL CLASS="statusbar">
					<LI>%datetime%</LI>
					<LI STYLE="float:right;">%userorlogin%</LI>
				</UL>
			</DIV>
			<SCRIPT>
				//taken from https://www.w3schools.com/howto/howto_js_topnav_responsive.asp
				function resizeHeaderNavBar() {
					var x = document.getElementById("headernavbar")
						if (x.className == "headernavbar") {
							x.className += "responsive"
						} else {
							x.className = "headernavbar"
						}
				}
			</SCRIPT>
			<DIV CLASS="header">
				<H1 CLASS="headertitle">NICOLAS GRENIER</H1>
				<DIV CLASS="headernavbar" ID="headernavbar">
					<A HREF="/" CLASS="active">HOME</A>
					<A HREF="/blog">BLOG</A>
					<A HREF="/notes">NOTES</A>
					<A HREF="/projects">PROJECTS</A>
					<A HREF="/texts">TEXTS</A>
					<FORM ACTION="/search" METHOD="post">
						<INPUT TYPE="text" CLASS="headersearch" PLACEHOLDER="Search" name="search" ID="search">
						<BUTTON TYPE="submit"><I CLASS="fa fa-search"></I></BUTTON>
					</FORM>
					<A HREF="javascript:void(0);" CLASS="icon" ONCLICK="resizeHeaderNavBar()">
						<I CLASS="fa fa-bars"></I>
					</A>
				</DIV>
			</DIV>
			<DIV CLASS="content">
`
	str = strings.Replace(str, "%datetime%", format_now(), -1)
	str = strings.Replace(str, "%title%", title, -1)
	if username == "" {
		str = strings.Replace(str, "%userorlogin%", "<A HREF=\"/login\">NOT LOGGED IN</A>", -1)
	} else {
		str = strings.Replace(str, "%userorlogin%", fmt.Sprintf("<A HREF=\"/account\">%s</A>", username), -1)
	}
	w.Write([]byte(str))
}

func WriteGeneralTrailer(w http.ResponseWriter) {
	w.Write([]byte(`</DIV>
	<DIV CLASS="bottombar">
		<UL CLASS="bottombar">
			<LI>&#9400; Nicolas Grenier, 2019</LI>
			<LI STYLE="float:right;"><A HREF="/about">ABOUT</A></LI>
			<LI STYLE="float:right;"><A HREF="/contact">CONTACT</A></LI>
			<LI STYLE="float:right;"><A HREF="/disclaimer">DISCLAIMER</A></LI>
		</UL>
	</DIV>
	</DIV>
	</HTML>`))
}

//Write error
func WriteError(w http.ResponseWriter, err error) {
	var b bytes.Buffer
	b.WriteString("<!DOCTYPE HTML>\n<HTML>\n<HEAD>\n<TITLE>ERROR</TITLE>\n")
	b.WriteString("<LINK REL=\"stylesheet\" HREF=\"/static/style.css\">")
	b.WriteString("</HEAD>\n<BODY>\n")
	b.WriteString("<DIV CLASS=\"errorcontainer\">\n")
	b.WriteString("<H1>* * *  ERROR  * * *</H1>")
	b.WriteString("<P>One or more errors occured while fetching the ressource.</P>\n")
	b.WriteString(fmt.Sprintf("<P CLASS=\"errorline\">%s</P>\n",err.Error()))
	b.WriteString("</DIV>\n</BODY>\n</HTML>")
	w.Write(b.Bytes())
}

//writes the blog post body
func WriteBlogPostBody(w http.ResponseWriter, id int, title string, username string, name string, created time.Time, modified time.Time, body string, errors []string, messages []string) {
	var b bytes.Buffer
	b.WriteString("<DIV CLASS=\"blogpostbody\">\n")
	b.WriteString(fmt.Sprintf("<H1>%s</H1>\n", title))
	b.WriteString(fmt.Sprintf("<P CLASS=\"byline\">By %s</P>\n",name))
	b.WriteString(body)
	b.WriteString("</DIV>\n")
	w.Write(b.Bytes())
}

//Writes out the comments
func WriteBlogPostComments() {

}

//Writes out a summary x likes n comments
func WriteBlogPostCommentSummary() {

}

//Write out the login screen.
func WriteLoginScreen(w http.ResponseWriter, username string, errors []string, messages []string) {

}

//Write out the blog post creation form.
func WriteNewBlogPost(w http.ResponseWriter, title string, shortTitle string, body string, formatter string, publish bool, errors []string, messages []string) {
	var sb bytes.Buffer
	sb.WriteString("<DIV CLASS=\"newblogpost\">\n")
	//Write out the errors
	if errors != nil || messages != nil {
		sb.WriteString("<DIV CLASS=\"errortable\">\n")
		for _, e := range errors {
			sb.WriteString("<P>")
			sb.WriteString(e)
			sb.WriteString("</P>\n")
		}
		for _, m := range messages {
			sb.WriteString("<P>")
			sb.WriteString(m)
			sb.WriteString("</P>\n")
		}
		sb.WriteString("</DIV>\n")
	}
	//Write out the title
	//sb.WriteString("<FORM CLASS=\"newblogpost\" ACTION=\"http://localhost:8080/newblogpost\" METHOD=\"post\">\n")
	sb.WriteString("<FORM CLASS=\"newblogpost\" ACTION=\"/newblogpost\" METHOD=\"post\">\n")
	sb.WriteString(fmt.Sprintf("<LABEL FOR=\"title\">Title</LABEL>\n"))
	sb.WriteString(fmt.Sprintf("<INPUT CLASS=\"newblogposttitle\" ID=\"title\" NAME=\"title\" VALUE=\"%s\"><BR>\n", title))
	sb.WriteString(fmt.Sprintf("<LABEL FOR=\"shorttitle\">Short Title</LABEL>\n"))
	sb.WriteString(fmt.Sprintf("<INPUT CLASS=\"newblogposttitle\" ID=\"shorttitle\" NAME=\"shorttitle\" VALUE=\"%s\">\n", shortTitle))
	//Write out the body
	sb.WriteString("<TEXTAREA ID=\"body\" ROWS=\"25\" NAME=\"body\">")
	sb.WriteString(body)
	sb.WriteString("</TEXTAREA>\n")
	//Write out the formatter selector
	sb.WriteString("<SELECT ID=\"formatter\">\n")
	sb.WriteString("<OPTION VALUE=\"raw\">RAW</OPTION>\n")
	sb.WriteString("<OPTION VALUE=\"markdown\">MARKDOWN</OPTION>\n")
	sb.WriteString("</SELECT>")
	//Publish button
	sb.WriteString("<INPUT TYPE=\"submit\" VALUE=\"submit\">")
	//End
	sb.WriteString("</FORM>\n")
	sb.WriteString("</DIV>\n")

	//We write
	w.Write(sb.Bytes())
}
