//
// Error type definition
//

package main

import (
	"fmt"
)

type Error struct {
	number	int
	text	string
}

func (e *Error) Error() string {
	return fmt.Sprintf("E%04d %s.", e.number, e.text)
}

type Message struct {
	number	int
	text	string
}

func (m *Message) Message() string {
	return fmt.Sprintf("I%04d %s.", m.number, m.text)
}







//
// Global Error
//

func ErrorUnknown (err error) error {
	return &Error{1, fmt.Sprintf("UNKNOWN ERROR '%s'", err.Error())}
}

func DatabaseError (err error) error {
	return &Error{5, fmt.Sprintf("DATABASE ERROR '%s'", err.Error())}
}

func RecordNotUniqueError() error {
	return &Error{6, "MULTIPLE RECORDS RETURNED WHERE ONLY ONE IS EXPECTED"}
}

func NotImplementedError() error {
	return &Error{9, "NOT IMPLEMENTED"}
}


//
// Session errors/messages PREFIX=10
//

func NotLoggedInError() error {
	return &Error{10, "NOT AUTHENTICATED"}
}

func InvalidCredentialsError() error {
	return &Error{11, "INVALID CREDENTIALS"}
}

func SessionExpiredError() error {
	return &Error{13, "SESSION EXPIRED"}
}

func AccountLockedError() error {
	return &Error{14, "ACCOUNT LOCKED"}
}

func AccountLockWarning(num int) *Message {
	return &Message{16, "ACCOUNT WILL BE LOCKED AFTER %d FURTHER UNSUCCESSFUL LOG-IN ATTEMPTS"}
}

func PreviousBadAttempsMessage(num int) *Message {
	return &Message{18, "%d PREVIOUS BAD LOG-IN ATTEMPTS"}
}

func PasswordWillExpireSoonMessage() *Message {
	return &Message{19, "PASSWORD WILL SOON EXPIRE"}
}

//
// Blog post errors
//

func BlogPostNotFoundError () error {
	return &Error{102, "BLOG POST NOT FOUND"}
}

func InvalidBlogPathError (path string) error {
	return &Error{}
}

func NotAuthorizedToViewBlogPost () error {
	return &Error{104, "NOT AUTHORIZED TO VIEW BLOG POST"}
}

func NotAuthorizedToCommentOnBlogPostError() error {
	return &Error{120, "NOT AUTHORIZED TO COMMENT ON BLOG POST"}
}

func AlreadyLikedError() error {
	return &Error{141, "ALREADY LIKED"}
}

func NotLikedError () error {
	return &Error{142, "NOT LIKED"}
}

func CannotLikeOwnError () error {
	return &Error{145, "CANNOT LIKE OWN"}
}

//
// Formatter Errors
//

func UnknownFormatterError (formatter string) error {
	return &Error{}
}

func FormatterError (error string) error {
	return &Error{}
}

//
// Search Errors
//

func SearchResultsTooLargeStripped() *Message {
	return &Message{2011, "SEARCH RESULTS TOO LARGE SO STRIPPED"}
}

func SearchTooLongStrippedMessage () *Message {
	return &Message{2019, "SEARCH TOO LONG EXTRA CHARACTERS REMOVED"}
}

func SearchEscapeCharactersRemovedMessage () *Message {
	return &Message{2021, "ESCAPE CHARACTERS IN SEARCH STRIPPED"}
}
