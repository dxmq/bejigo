package syserrors

type Error404 struct {
	UnKnowError // 继承于UnknowError结构体
}

// 相当重写UnKnowError 中的Code()和Error()方法
func (c Error404) Code() int {
	return 1002
}

func (c Error404) Error() string {
	return "非法访问"
}