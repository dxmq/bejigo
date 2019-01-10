package syserrors

type UnLoginError struct {
	UnKnowError
}

func (c UnLoginError) Code() int {
	return 1001
}

func (c UnLoginError) Error() string {
	return "请登录系统！"
}
