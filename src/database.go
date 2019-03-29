package main

import (
	"database/sql"
	"github.com/mattn/go-sqlite3"
)

db.NewTransaction()

// Database operation object.
//
// All operations must take place within a transaction.
//
//
type Database struct {
	db *sql.DB
}

//Connect to the database server.
func (db Database) Connect(address string, port int, username string, password string, name string) error {

}

//Get a transaction context from the database.
func (db *Database) Begin() (*Transaction, error) {
	tx, err := db.BeginTx()
	return &Transaction{}
}



//
// Transaction object: used to do
//
type Transaction struct {
	tx *sql.Tx
}

func (tx *Transaction) Exec (query string, args ...Interface{}) (sql.Result, error) {
	return tx.tx.Exec(query, args...)
}

func (tx *Transaction) Query(query string, args ...Interface{}) (*sql.Rows, error) {
	return tx.tx.Query(query, args...)
}

func (tx *Transaction) Query(query string, args ...Interface{}) *sql.Row {
	return tx.tx.QueryRow(query, args...)
}



//
// B L O G   P O S T S
//

func (tx *Transaction) DoesBlogPostExist(post int) error {
	tmp int
	row := db.countBlogPostStmt.QueryRow(post)
	err := row.Scan(&tmp)
	if err {
		return DatabaseError()
	}
	if tmp == 0 {
		return BlogPostNotFoundError()
	} else {
		return nil
	}
}

func (db *Database) GetBlogPostNumber(year int, month int, shorttitle string) (int, error) {
	//Return value
	number int
	//Fetch from the database
	rows, err := db.getBlogPostFromDateStmt.Query(year, month, shorttitle)
	if err != nil {
		return 0, DatabaseError()
	}
	//Check the result
	if len(rows) == 0 {
		return 0, BlogPostNotFoundError()
	} else {
		//Good result
		rows[0].Scan(&number)
		return number, nil
	}
}

//Raw insert: formatting required beforehands (is Draft and not Deleted)
//Returns post id, error
func (db Database) InsertBlogPost(
		deleted bool,
		draft bool,
		revision int,
		author int,
		created time.Time,
		modified time.Time,
		year int,
		month int,
		shortTitle string,
		title string,
		content string,
		likeable bool,
		hidelikes bool,
		commentable bool,
		hidecomments bool
	) (int, error, []Message) {


}

func (db *Database) CreateBlogPost() (int, error) {
	db.Exec("INSERT INTO BLOGPOSTS() RETURNING NUMBER")
	db.Exec("INSERT INTO BLOGPOSTSOURCES(BLOGPOST, REVISION, FORMATTER, FORMATTERVERSION, SUCCESSFUL, CONTENT, LOG, TIMESTAMP, USERID)")
}

func (db *Database) EditBlogPost() (error) {

}

//Remove draft flag
func (db *Database) PublishBlogPost(blogPost id) (error, []Message) {

}

//Mark post as deleted
func (db *Database) DeleteBlogPost() (error, []Message) {
	db.Exec("UPDATE BLOGPOSTS SET DELETED=TRUE WHERE NUMBER=?")
}

func (db *Database) DisableBlogPostCommenting(post int, userid int) {
	db.Exec("UPDATE BLOGPOSTS SET COMMENTABLE=FALSE WHERE NUMBER=?")
}

func (db *Database) HideBlogPostComments(post int, userid int) {
	db.Exec("UPDATE BLOGPOSTS SET HIDECOMMENTS=TRUE WHERE NUMBER=?")
}

//Prevent new likes
func (db *Database) DisableBlogPostLiking(post int, userid int) {
	db.Exec("UPDATE BLOGPOSTS SET LIKEABLE=FALSE WHERE NUMBER=?")
}

//Hide existing likes
func (db *Database) HideBlogPostLikes(post int, userid int) {
	db.Exec("UPDATE BLOGPOSTS SET HIDELIKES=TRUE WHERE NUMBER=?")
}


//
//  B L O G   P O S T   L I K E S
//

//Get the number of likes on a post
func (db *Database) CountBlogPostLikes(post int) (int, error, []Message) {
	//We first check that the post exists
}

//Like Post
func (db Database) LikeBlogPost() (error, []Message) {

}

//Unlike Post
func (db *Database) UnlikeBlogPost() {

}



//
//  B L O G   P O S T   C O M M E N T S
//

//Raw insert, returns comment id
func (db *Database) insertBlogPostComment() (int, error, []Message) {

}

//Raw insert, returns comment revision id
func (db *Database) insertBlogPostCommentSource() (int, error, []Message) {

}

//Comment Post, returns absolute comment id
func (db Database) CommentBlogPost() (int, error, []Message) {

}

//Update Comment
func (db Database) EditBlogPostComment() (error, []Message) {

}

//Remove Comment (username needed for access control)
func (db Database) RemoveBlogPostComment(comment int, userid int) (error, []Message) {

}

//
//   B L O G   P O S T   C O M M E N T   L I K E S
//

//Like Comment
func (db Database) LikeBlogPostComment(comment int, userid int) (error, []Message) {

}

//Unlike Comment
func (db Database) UnlikeBlogPostComment (comment int, userid int) (error, []Message) {

}



//
// U S E R   P R O F I L E S
//

//Create user profile, return user id.
func (tx *Transaction) CreateUserProfile (userName string, firstName string, middleNames string, lastName string, email string) (int, error, []Message) {

}

//Check user/password, returns user id
func (tx *Transaction) CheckAuthentication (username string, password string) (int, error, []Message) {

}

//Change password
func (tx *Transaction) ChangePassword (target id, newPassword string, userid int) (error, []Message) {

}

//Change user name (userid for access control)
func (tx *Transaction) ChangeUserName (target id, newUserName string, userid id) (error, []Message) {

}

//Change user names
func (tx *Transaction) ChangeUserNames (target id, newFirstName string, newMiddleNames string, newLastName string, userid int) (error, []Message) {

}

//Make the user <target> a moderator.
func (tx *Transaction) MakeUserModerator(target int, changer int) {

}

//
// These functions act on a database as they are not transactional.
//

//Get a blog post.
//
//Returns BLOGPOST {
//	"id" : integer number
//	"path" : /year/month/shorttitle
//	"authorid": author id
//	"author": author name
//	"created" create date&time
//	"modified" : datetime or nil
//	"canedit" : whether the current user can edit
//	"ncomments" number of comments
//	"nlikes" number of likes
//	"canlike" whether we can like
//	"cancomment" whether we can comment
//	"comments" {
//		"id" number of the comment (database)
//		"seq" sequence number of comment (always in order)
//		"userid" userif of creator
//		"username" username of creator
//		"inreplyto" sequence number of comment in reply to
//		"inreplytouserid" userid of the user in reply to
//		"inreplytouser" username of user which is being replied to
//		"nreplies" the number of replies
//		"nlikes" the number of likes
//		"canlike"
//		"canedit"
//		"canreply"
//		"replies" {
//			"id" unique id of the reply
//			"seq" seq number of the reply
//			"username" username of the replier
//}
//}
//}
//set includecomments to false to not get the comments (only ncomments)
func (db *Database) GetBlogPost(number int, includeComments boolean) (error) {

	db.QueryRow("SELECT ")

}
