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
			<DIV CLASS="header">
				<H1 CLASS="headertitle">NICOLAS GRENIER</H1>
				<UL CLASS="headernavbar">
					<LI><A HREF="/blog">BLOG</A></LI>
					<LI><A HREF="/notes">NOTES</A></LI>
					<LI><A HREF="/projects">PROJECTS</A></LI>
					<LI><A HREF="/texts">TEXTS</A></LI>
					<LI><FORM ACTION="/search">
						<INPUT TYPE="text" CLASS="headersearch" PLACEHOLDER="Search">
						<BUTTON TYPE="sumbit" CLASS="headersearchbutton"><IMG SRC="/static/s.svg" ALT="Search"></BUTTON>
						</FORM></LI>
				</UL>
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
	sb.WriteString("<FORM CLASS=\"newblogpost\">\n")
	sb.WriteString(fmt.Sprintf("<LABEL FOR=\"title\">Title</LABEL>\n"))
	sb.WriteString(fmt.Sprintf("<INPUT CLASS=\"newblogposttitle\" ID=\"title\" VALUE=\"%s\"><BR>\n", title))
	sb.WriteString(fmt.Sprintf("<LABEL FOR=\"shorttitle\">Short Title</LABEL>\n"))
	sb.WriteString(fmt.Sprintf("<INPUT CLASS=\"newblogposttitle\" ID=\"shorttitle\" VALUE=\"%s\">\n", shortTitle))
	//Write out the body
	sb.WriteString("<TEXTAREA ID=\"body\" ROWS=\"25\">")
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
