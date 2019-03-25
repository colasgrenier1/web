
import (
	"database/sql"
	"github.com/mattn/go-sqlite3"
)

type Database struct {
	filename string
	db sql.DB
}

fn (db Database) Connect(filename string) {
	
}

//
// B L O G   P O S T S
//

//Raw insert: formatting required beforehands.
fn (db Database) InsertBlogPost() {
	.
}

//Format and insert
fn (db Database) CreateDraftBlogPost() {
	
}

//Remove draft flag
fn (db Database) PublishBlogPost() {
		
}

//Mark post as deleted
fn (db Database) DeleteBlogPost() {
	
}

//Like Post
fn (db Database) LikeBlogPost() {
	
}

//Comment Post
fn (db Database) CommentBlogPost() {
	
}

//Update Comment
fn (db Database) EditBlogPostComment() {
	
}

//Remove Comment (username needed for access control)
fn (db Database) RemoveBlogPostComment(comment int, userid int) {
	
}

//Like Comment
fn (db Database) LikeBlogPostComment(comment int, userid int) {
	
}

//Unlike Comment
fn (db Database) UnlikeBlogPostComment (comment int, userid int) {
	
}



//
// U S E R   P R O F I L E S
//

//Create user profile, return user id.
fn (db Database) CreateUserProfile (userName string, firstName string, middleNames string, lastName string, email string) int {
	
}

//Change user name (userid for access control)
fn (Db Database) ChangeUserName (target id, newUserName string, userid id) {
	
}

//Change user names
fn (db Database) ChangeUserNames (target id, newFirstName string, newMiddleNames string, newLastName string, userid int) {
	
}
