# 常量
## NO_ARGS
- *功能* : 空参数常量.
- *用途* : 作为参数的占位符.

# 字节码
## LOAD_VALUE
- *功能* : 将一个常量值加载到堆栈顶.
- *参数* : 常量的值.
- *格式*
    ```flex-bytecode
    LOAD_VALUE <value>
    ```
- *示例* :
    ```flex-bytecode
    LOAD_VALUE 10
    ```
    此时堆栈顶为常量值 10.


## STORE_NAME
- *功能* : 将堆栈顶部的值存储到变量中.
- *参数* : 变量名.
- *格式*
    ```flex-bytecode
    STORE_NAME <name>
    ```
- *示例* :
    ```flex-bytecode
    LOAD_VALUE 10
    STORE_NAME 2
    ```
    此时将堆栈顶部的值 10 存储到变量 x 中.
- *注意*: 储存后栈顶的值不会被弹出.
- *error* : 
  - [NameCannotBeUsedError](./ERRORS.md#namecannotbeusederror) 在[变量名不合法](../runBytecode/CHECKERS.md#变量名规定)时报错.


## LOAD_NAME
- *功能* : 从变量中加载值到堆栈顶.
- *参数* : 变量名.
- *格式*
    ```flex-bytecode
    LOAD_NAME <name>
    ```
- *示例* :
    ```flex-bytecode
    LOAD_NAME x
    ```
    此时将变量 x 中的值加载到堆栈顶.
- *注意*: 如果变量不存在会报错.
- *error* : [NameNotDefinedError](./ERRORS.md#namenotdefinederror) 在变量不存在却被引用时报错.

## POP_TOP
- *功能* : 弹出栈顶的值.
- *参数* : 无, 但是建议传入一个 [NO_ARGS  常量](#no_args)
- *格式*
    ```flex-bytecode
    POP_TOP NO_ARGS
    ```
- *示例* :
    ```flex-bytecode
    LOAD_VALUE  10
    POP_TOP     NO_ARGS
    ```
    之前栈顶为常量值 10, 但是此时栈顶的值已经被弹出.
- *error* : [StackEmptyButPopError]( ./ERRORS.md#stackemptybutpoperror) 在栈为空却尝试弹出时报错.

## 未定义的字节码
- *error* : [InvalidBytecodeError](./ERRORS.md#invalidbytecoderror) 遇到未定义的字节码报错.