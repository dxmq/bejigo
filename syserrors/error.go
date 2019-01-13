package syserrors

type Error interface {
	Code() int
	Error() string
	ReasonError() error
}

func New(msg string, reason error) Error {
	return UnKnowError{msg, reason}
}
