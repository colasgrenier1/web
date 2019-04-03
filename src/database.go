package main

import (
	"database/sql"
	"fmt"
	"time"
	_ "github.com/lib/pq"
)

// Database operation object.
//
// All operations must take place within a transaction.
//
//
type Database struct {
	db *sql.DB
}

// Transaction object: used to do
type Transaction struct {
	tx *sql.Tx
}

//Connect to the database server.
func (db *Database) Connect(address string, port int, username string, password string, name string) error {
	d, err := sql.Open("postgres", fmt.Sprintf("postgres://%s:%s@%s/%s", username, password, address, name))
	if err != nil {
		return err
	}
	db.db = d
	return nil
}

//Get a transaction context from the database.
func (db *Database) Begin() (*Transaction, error) {
	t, err := db.db.Begin()
	return &Transaction{tx: t}, err
}

func (tx *Transaction) Exec(query string, args ...interface{}) (sql.Result, error) {
	return tx.tx.Exec(query, args...)
}

func (tx *Transaction) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return tx.tx.Query(query, args...)
}

func (tx *Transaction) QueryRow(query string, args ...interface{}) *sql.Row {
	return tx.tx.QueryRow(query, args...)
}

func (tx *Transaction) Commit() error {
	return tx.tx.Commit()
}

func (tx *Transaction) Rollback() error {
	return tx.tx.Rollback()
}



//
// B L O G   P O S T S
//

func (tx *Transaction) DoesBlogPostExist(post int) error {
	var tmp int
	row := tx.QueryRow("SELECT COUNT(*) FROM BLOGPOSTS WHERE NUMBER=$1", post)
	err := row.Scan(&tmp)
	if err != nil {
		return DatabaseError(err)
	}
	if tmp == 0 {
		return BlogPostNotFoundError()
	} else {
		return nil
	}
}

func (tx *Transaction) GetBlogPostNumber(year int, month int, shorttitle string) (int, error) {
	//Return value
	var number int
	var n int = 0
	//Fetch from the database
	rows, err := tx.Query("SELECT NUMBER FROM BLOGPOSTS WHERE YEAR=$1 AND MONTH=$2 AND SHORTTITLE=$3", year, month, shorttitle)
	if err != nil {
		return 0, DatabaseError(err)
	}
	for rows.Next() {
		n += 1
		err = rows.Scan(&number)
		if  err != nil {
			return 0, DatabaseError(err)
		}
	}
	if n != 1 {
		return 0, RecordNotUniqueError()
	}
	return number, nil
}

// //Raw insert: formatting required beforehands (is Draft and not Deleted)
// //Returns post id, error
// func (tx *Transaction) InsertBlogPost(
// 		deleted bool,
// 		draft bool,
// 		revision int,
// 		author int,
// 		created time.Time,
// 		modified time.Time,
// 		year int,
// 		month int,
// 		shortTitle string,
// 		title string,
// 		content string,
// 		likeable bool,
// 		hidelikes bool,
// 		commentable bool,
// 		hidecomments bool) (int, error) {
// }
//
// //func (tx *Transaction) CreateBlogPost() (int, error) {
// //	tx.Exec("INSERT INTO BLOGPOSTS() RETURNING NUMBER")
// //	tx.Exec("INSERT INTO BLOGPOSTSOURCES(BLOGPOST, REVISION, FORMATTER, FORMATTERVERSION, SUCCESSFUL, CONTENT, LOG, TIMESTAMP, USERID)")
// //}
//
// func (tx *Transaction) EditBlogPost() (error) {
// }
//
// //Remove draft flag
// func (tx *Transaction) PublishBlogPost(blogPost id) (error) {
// 	return nil
// }
//
// //Mark post as deleted
// func (tx *Transaction) DeleteBlogPost() (error) {
// 	tx.Exec("UPDATE BLOGPOSTS SET DELETED=TRUE WHERE NUMBER=?")
// }
//
// func (tx *Transaction) DisableBlogPostCommenting(post int, userid int) {
// 	tx.Exec("UPDATE BLOGPOSTS SET COMMENTABLE=FALSE WHERE NUMBER=?")
// }
//
// func (tx *Transaction) HideBlogPostComments(post int, userid int) {
// 	tx.Exec("UPDATE BLOGPOSTS SET HIDECOMMENTS=TRUE WHERE NUMBER=?")
// }
//
// //Prevent new likes
// func (tx *Transaction) DisableBlogPostLiking(post int, userid int) {
// 	tx.Exec("UPDATE BLOGPOSTS SET LIKEABLE=FALSE WHERE NUMBER=?")
// }
//
// //Hide existing likes
// func (tx *Transaction) HideBlogPostLikes(post int, userid int) {
// 	tx.Exec("UPDATE BLOGPOSTS SET HIDELIKES=TRUE WHERE NUMBER=?")
// }
//
//
// //
// //  B L O G   P O S T   L I K E S
// //
//
// //Get the number of likes on a post
// func (tx *Transaction) CountBlogPostLikes(post int) (int, error) {
// 	//We first check that the post exists
// 	return nil
// }
//
// //Like Post
// func (tx *Transaction) LikeBlogPost() (error) {
//
// }
//
// //Unlike Post
// func (tx *Transaction) UnlikeBlogPost() {
//
// }
//
//
//
// //
// //  B L O G   P O S T   C O M M E N T S
// //
//
// //Raw insert, returns comment id
// func (tx *Transaction) insertBlogPostComment() (int, error) {
//
// }
//
// //Raw insert, returns comment revision id
// func (tx *Transaction) insertBlogPostCommentSource() (int, error) {
//
// }
//
// //Comment Post, returns absolute comment id
// func (tx *Transaction) CommentBlogPost() (int, error) {
//
// }
//
// //Update Comment
// func (tx *Transaction) EditBlogPostComment() (error, ) {
//
// }
//
// //Remove Comment (username needed for access control)
// func (tx *Transaction) RemoveBlogPostComment(comment int, userid int) (error) {
//
// }
//
// //
// //   B L O G   P O S T   C O M M E N T   L I K E S
// //
//
// //Like Comment
// func (tx *Transaction) LikeBlogPostComment(comment int, userid int) (error) {
//
// }
//
// //Unlike Comment
// func (tx *Transaction) UnlikeBlogPostComment (comment int, userid int) (error) {
//
// }
//
//
//
// //
// // U S E R   P R O F I L E S
// //
//
// //Create user profile, return user id.
// func (tx *Transaction) CreateUserProfile (userName string, firstName string, middleNames string, lastName string, email string) (int, error) {
//
// }
//
// //Check user/password, returns user id
// func (tx *Transaction) CheckAuthentication (username string, password string) (int, error) {
//
// }
//
// //Change password
// func (tx *Transaction) ChangePassword (target id, newPassword string, userid int) (error) {
//
// }
//
// //Change user name (userid for access control)
// func (tx *Transaction) ChangeUserName (target id, newUserName string, userid id) (error) {
//
// }
//
// //Change user names
// func (tx *Transaction) ChangeUserNames (target id, newFirstName string, newMiddleNames string, newLastName string, userid int) (error) {
//
// }
//
// //Make the user <target> a moderator.
// func (tx *Transaction) MakeUserModerator(target int, changer int) {
//
// }
//
//
// These functions act on a database as they are not transactional.
//

//Get the user
func (tx *Transaction) GetBlogPostBody(post int) (title string, authorusername string, author string, created time.Time, modified time.Time, body string, err error) {
	row := tx.QueryRow("SELECT TITLE, USERNAME, FIRSTNAME || ' ' || LASTNAME, BLOGPOSTS.CREATED, MODIFIED, CONTENT FROM BLOGPOSTS LEFT JOIN USERS ON AUTHOR=USERS.NUMBER WHERE BLOGPOSTS.NUMBER=$1", post)
	e := row.Scan(&title, &authorusername, &author, &created, &modified, &body)
	if e != nil {
		err = DatabaseError(e)
	} else {
		err = nil
	}
	return
}

// //Get a blog post.
// //
// //Returns BLOGPOST {
// //	"id" : integer number
// //	"path" : /year/month/shorttitle
// //	"authorid": author id
// //	"author": author name
// //	"created" create date&time
// //	"modified" : datetime or nil
// //	"canedit" : whether the current user can edit
// //	"ncomments" number of comments
// //	"nlikes" number of likes
// //	"canlike" whether we can like
// //	"cancomment" whether we can comment
// //	"comments" {
// //		"id" number of the comment (database)
// //		"seq" sequence number of comment (always in order)
// //		"userid" userif of creator
// //		"username" username of creator
// //		"inreplyto" sequence number of comment in reply to
// //		"inreplytouserid" userid of the user in reply to
// //		"inreplytouser" username of user which is being replied to
// //		"nreplies" the number of replies
// //		"nlikes" the number of likes
// //		"canlike"
// //		"canedit"
// //		"canreply"
// //		"replies" {
// //			"id" unique id of the reply
// //			"seq" seq number of the reply
// //			"username" username of the replier
// //}
// //}
// //}
// //set includecomments to false to not get the comments (only ncomments)
// func (db *Database) GetBlogPost(number int, includeComments boolean) (error) {
//
// 	db.QueryRow("SELECT ")
//
// }
