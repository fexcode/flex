# ERRORS

## NameCannotBeUsedError
如果[变量名不可用](../runBytecode/CHECKERS.md#变量名规定)(包含非法字符或是关键字), <br>
抛出该错误.

## NameNotDefinedError
如果变量名未定义,却尝试使用它,<br>
抛出该错误.

## StackEmptyButPopError
如果栈为空,却尝试弹出元素,<br>
抛出该错误.

## InvalidBytecodeError
如果遇到未定义的[字节码](../consts/BYTECODES.md#字节码), <br>
抛出该错误.