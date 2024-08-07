package runBytecode

import (
	// "fmt"
	bc "flex/consts/bytecodes"
	errs "flex/consts/errors"
	stk "flex/types"
	"fmt"
)

// RunBytecodes 运行字节码
func RunBytecodes(bytecodes []bc.Bytecode) {
	stack := stk.NewStack()
	variables := make(map[string]interface{})

	// 遍历字节码，逐个执行操作
	for _, bytecode := range bytecodes {

		switch bytecode.Opcode {
		case bc.LOAD_VALUE:
			// 将一个常量值加载到堆栈上。
			// 参数：常量值的索引或值。
			// 作用：将常量值压入栈顶。
			stack.Push(bytecode.Arg)

		case bc.STORE_NAME:
			// 将堆栈顶部的值存储到变量中。
			// 参数：变量名的索引。
			// 作用：将栈顶的值存储到变量中,不会弹出栈顶的值。
			// 注意：变量名必须是字符串。
			// 注意：变量名不能以数字开头。
			// 注意：变量名为关键字.
			// 注意: 储存后栈顶的值不会被弹出。
			strArg, ok := bytecode.Arg.(string)
			if ok && CheckName(strArg) {
				variables[strArg], _ = stack.Peek()
			} else {
				errs.RaiseError(errs.NameCannotBeUsedError)
			}

		case bc.LOAD_NAME:
			// 将变量的值加载到堆栈上。
			// 参数：变量名。
			// 作用：将变量的值压入栈顶。
			// 注意: 如果变量不存在会报错.
			varName, _ := bytecode.Arg.(string)
			varValue, ok := variables[varName]
			if ok {
				stack.Push(varValue)
			} else {
				errs.RaiseError(errs.NameNotDefinedError)
			}

		case bc.POP_TOP:
			// 弹出栈顶的值。
			// 作用：弹出栈顶的值，不会做任何操作.
			// 这里建议传入一个 NO_ARG 常量作为参数，表示不用参数.
			_, error := stack.Pop()
			if error != nil {
				errs.RaiseError(errs.StackEmptyButPopError)
			}

		default:
			errs.RaiseError(errs.InvalidBytecodeError)
		}
		fmt.Print("RANBC: \t", bc.Bytecode2String(bytecode.Opcode), "\t", bytecode.Arg, "\n")
		fmt.Print("STACK: \t", stack.Data(), "\n")
		fmt.Print("VARS: \t", variables, "\n")
		fmt.Print("\n")

	}
}

func Test() {
	bytecodes := []bc.Bytecode{
		bc.NewBytecode(bc.LOAD_VALUE, 10),
		bc.NewBytecode(bc.STORE_NAME, "x"),
		bc.NewBytecode(bc.LOAD_VALUE, "hello world"),
		bc.NewBytecode(bc.STORE_NAME, "hello"),
		bc.NewBytecode(bc.POP_TOP, bc.NO_ARG),
	}
	RunBytecodes(bytecodes)

	bc.ParseBytecodeFromString(`
	LOAD_VALUE 10
	STORE_NAME x
	LOAD_VALUE "hello world"
	`)
}
