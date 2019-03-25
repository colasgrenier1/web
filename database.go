
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

//Remove Comment
fn (db Database) RemoveBlogPostComment() {
	
}

//Like Comment
fn (db Database) LikeBlogPostComment() {
	
}

//Unlike Comment
fn (db Database) Unlike
