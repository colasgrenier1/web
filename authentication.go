
import (
	"encoding/csv"
	"encoding/base64"
	"time"
)


type Session struct {
	id		integer		//session id
	active		boolean
	user		int		//user id
	address		string		//ip address, etc.
	hash		string		//cookie value
	last		Time
}

//Gets the URL
fn GetURLEncodedSessionID(id integer) {
}

//Manages sessions
type SessionManager {
	csvFile		csv.Writer
	database	Database
	sessions	[]Session
	timeout		integer //number of seconds between each request until termination of account.
}

//Login: returns cookie value
fn (db SessionManager) Login (username string, password string, address string) {
	//We check
	id, 
}

//Access: returns the username (or nil) and the cookie value to be set (or nil)
fn (db SessionManager) Access (address string, hash string) {
	
}
