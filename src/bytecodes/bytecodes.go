package code2bytecode

// Opcode 定义了字节码的操作类型
type Opcode int

// Bytecode 定义字节码的结构
type Bytecode struct {
	Opcode Opcode      // 操作码
	Arg    interface{} // 操作数
}

//定义字节码的常量
/*
详见 Flex/docs/code2bytecode/BYTECODES.md

LOAD_VALUE - 将一个常量加载到堆栈上。
	参数：常量的索引。
STORE_NAME - 将堆栈顶部的值存储到一个变量中。
	参数：变量名的索引。
LOAD_NAME - 将一个变量的值加载到堆栈上。
	参数：变量名的索引。
ADD - 从堆栈中弹出两个值，将它们相加，然后将结果推回堆栈。
SUB - 从堆栈中弹出两个值，将第二个值从第一个值中减去，然后将结果推回堆栈。
MUL - 从堆栈中弹出两个值，将它们相乘，然后将结果推回堆栈。
DIV - 从堆栈中弹出两个值，将第一个值除以第二个值，然后将结果推回堆栈。
JUMP - 无条件跳转到指定的字节码索引。
JUMP_IF_FALSE - 如果堆栈顶部的值为假，则跳转到指定的字节码索引。
CALL_FUNCTION - 调用一个函数。
	参数：函数名的索引，以及传递给函数的参数数量。
RETURN_VALUE - 从当前函数返回堆栈顶部的值。
*/
const (
	LOAD_VALUE = iota
	STORE_NAME
	LOAD_NAME
	ADD
	SUB
	MUL
	DIV
	JUMP
	JUMP_IF_FALSE
	CALL_FUNCTION
	RETURN_VALUE
)

// NewBytecode 创建一个新的Bytecode实例
/*
Args:
	opcode - 字节码的操作类型
	operands - 字节码的操作数(多个)
Returns:
	一个新的Bytecode实例
	{opcode, [operands]}
*/
func NewBytecode(opcode Opcode, operand interface{}) Bytecode {
	return Bytecode{
		Opcode: opcode,
		Arg:    operand,
	}
}

// 字节码常量转换为字节码指令字符串
func Bytecode2String(opcode Opcode) string {
	switch opcode {
	case LOAD_VALUE:
		return "LOAD_VALUE"

	case STORE_NAME:
		return "STORE_NAME"

	case LOAD_NAME:
		return "LOAD_NAME"

	case ADD:
		return "ADD"

	case SUB:
		return "SUB"

	case MUL:
		return "MUL"

	case DIV:
		return "DIV"

	case JUMP:
		return "JUMP"

	case JUMP_IF_FALSE:
		return "JUMP_IF_FALSE"

	case CALL_FUNCTION:
		return "CALL_FUNCTION"

	case RETURN_VALUE:
		return "RETURN_VALUE"

	default:
		return "Unknown Opcode"
	}
}
