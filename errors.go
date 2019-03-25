typedef ErrorNumber int

typedef Error struct {
	number	ErrorNumber
	text	string
}



func (e *Error) Error() string {
	return fmt.Sprintf("E%06d %s.", e.number, e.text)
}

func UnknownFormatterError (formatter string) Error {
	return &Error{
}


