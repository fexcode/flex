package runBytecode

import (
	// "fmt"
	bc "flex/bytecodes"
	stk "flex/stack"
	"fmt"
)

// RunBytecode 运行字节码
func RunBytecode(bytecodes []bc.Bytecode) {
	stack := stk.NewStack()
	variables := make(map[string]interface{})

	// 遍历字节码，逐个执行操作
	for _, bytecode := range bytecodes {

		switch bytecode.Opcode {
		case bc.LOAD_VALUE:
			// 将一个常量值加载到堆栈上。
			// 参数：常量值的索引或值。
			// 作用：将常量值压入栈顶。
			/*
				注意: LOAD_VALUE多个数时,会这样
				RUNBC:  LOAD_VALUE [20]
				STACK:  [[10 20] [20]]
			*/
			stack.Push(bytecode.Arg)

		case bc.STORE_NAME:
			// 将堆栈顶部的值存储到变量中。
			// 参数：变量名的索引。
			// 作用：将栈顶的值存储到变量中,并弹出栈顶的值。
			variables[bytecode.Arg.(string)], _ = stack.Pop()

		case bc.LOAD_NAME:
			// 将变量的值加载到堆栈上。
			// 参数：变量名的索引。
			// 作用：将变量的值压入栈顶。
			stack.Push(variables[bytecode.Arg.(string)])
			

		default:
			panic("Invalid operation code")
		}
		fmt.Print("RANBC: \t", bc.Bytecode2String(bytecode.Opcode),"\t", bytecode.Arg, "\n")
		fmt.Print("STACK: \t", stack.Data(), "\n")
		fmt.Print("VARS: \t", variables, "\n")
		fmt.Print("\n")

	}
}

func Test() {
	bytecodes := []bc.Bytecode{
		bc.NewBytecode(bc.LOAD_VALUE, 10),
		bc.NewBytecode(bc.STORE_NAME, "x"),
		bc.NewBytecode(bc.LOAD_VALUE, 20),
		bc.NewBytecode(bc.STORE_NAME, "y"),
		bc.NewBytecode(bc.LOAD_NAME, "x"),
		bc.NewBytecode(bc.LOAD_NAME, "y"),
		bc.NewBytecode(bc.LOAD_NAME, "x"),
	}
	RunBytecode(bytecodes)
}
