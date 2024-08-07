package errors

type flexErrorType string

const (
	NameNotDefinedError   flexErrorType = "该变量名不存在"
	NameCannotBeUsedError flexErrorType = "变量名不可用"
	InvalidBytecodeError  flexErrorType = "无效的字节码"
	StackEmptyButPopError flexErrorType = "栈为空, 不可弹出元素"
)

func RaiseError(err flexErrorType) {
	panic(err)
}
