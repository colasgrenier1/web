

import (
	"database/sql"
	"github.com/mattn/go-sqlite3"
)

type Database struct {
	filename string
	db sql.DB
	insertBlogPostStmt db.Stmt
	publishBlogPostStmt db.Stmt
	insertBlogPostSourceStmt db.Stmt
}

fn (db Database) Initialize() {
		
}

fn (db Database) Connect(filename string) {
	
}

//
// B L O G   P O S T S
//

//Raw insert: formatting required beforehands.
//Returns post id, error
fn (db Database) InsertBlogPost() (int, error) {
	
}

//Raw insert: returns absolute revision index
func(db *Database) InsertBlogPostSource() (int, error) {
	
}

//Format and insert: returns post id
fn (db *Database) CreateDraftBlogPost() (int, error) {
	
}

//Remove draft flag
fn (db *Database) PublishBlogPost() (error) {
		
}

//Mark post as deleted
fn (db *Database) DeleteBlogPost() (error) {
	
}

//Like Post
fn (db Database) LikeBlogPost() (error) {
	
}

//Raw insert, returns comment id
fn (db *Database) insertBlogPostComment() (int, error) {
	
}

//Raw insert, returns comment revision id
fn (db *Database) insertBlogPostCommentSource() (int, error) {
	
}

//Comment Post, returns absolute comment id
fn (db Database) CommentBlogPost() (int, error) {
	
}

//Update Comment
fn (db Database) EditBlogPostComment() error {
	
}

//Remove Comment (username needed for access control)
fn (db Database) RemoveBlogPostComment(comment int, userid int) error {
	
}

//Like Comment
fn (db Database) LikeBlogPostComment(comment int, userid int) error {
	
}

//Unlike Comment
fn (db Database) UnlikeBlogPostComment (comment int, userid int) error {
	
}



//
// U S E R   P R O F I L E S
//

//Create user profile, return user id.
fn (db *Database) CreateUserProfile (userName string, firstName string, middleNames string, lastName string, email string) (int, error) {
	
}

//Check user/password, 
fn (db *Database) CheckAuthentication (username string, password string) (boolean, error) {

}

//Change password
fn (db *Database) ChangePassword (target id, newPassword string, userid int) error {
	
}

//Change user name (userid for access control)
fn (db *Database) ChangeUserName (target id, newUserName string, userid id) error {
	
}

//Change user names
fn (db *Database) ChangeUserNames (target id, newFirstName string, newMiddleNames string, newLastName string, userid int) error {
	
}
