package consts_bytecodes

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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
	NO_ARG = iota // 这不是一个有效的操作码，用于表示无操作数
	LOAD_VALUE
	STORE_NAME
	LOAD_NAME
	POP_TOP
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

func ParseLineBytecode(line string) Bytecode { //(Bytecode, error)
	// TODO: 解析字节码指令
	fmt.Println("生成了字节码: ", line)
	return NewBytecode(NO_ARG, 0)
}

func ParseBytecodeFromString(str string) ([]Bytecode, error) {
	bytecodes := make([]Bytecode, 0)

	// 使用strings.NewReader将字符串转换为io.Reader接口
	reader := strings.NewReader(str)

	// 创建bufio.Scanner来按行扫描
	scanner := bufio.NewScanner(reader)

	// 使用Scan方法按行遍历
	for scanner.Scan() {
		line := scanner.Text() // 获取当前行
		bytecodes = append(bytecodes, ParseLineBytecode(line))

	}

	// 检查遍历过程中是否有错误发生
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	return bytecodes, nil
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

	case POP_TOP:
		return "POP_TOP"

	default:
		return "Unknown Opcode"
	}
}
