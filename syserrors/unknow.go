package syserrors

type UnKnowError struct {
	msg string
	reason error
}

func (c UnKnowError)Code() int {
	return 1000
}

func (c UnKnowError)Error() string {
	if len(c.msg) == 0 {
		return "未知错误"
	}
	return c.msg
}

func (c UnKnowError)ReasonError() error {
	return c.reason
}
