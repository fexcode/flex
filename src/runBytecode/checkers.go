package runBytecode

func CheckName(name string) bool {
	/* 
	变量名不以数字开头，且只能包含字母、数字、下划线,
	变量名不能包含空格

	Arg:
		name: 变量名

	Return:
		true: 变量名合法
		false: 变量名不合法
	*/

	// 变量名不能为空
	if len(name) == 0 {
		return false
	}

	// 变量名不能以数字开头
	if name[0] >= '0' && name[0] <= '9' {
		return false
	}

	// 变量名只能包含字母、数字、下划线
	for i := 0; i < len(name); i++ {
		// 变量名不能包含空格
		if name[i] == ' ' {
			return false
		}

		// 变量名只能包含字母、数字、下划线
		if name[i] == '_' {
			continue
		}

		// 变量名只能包含字母、数字、下划线
		if (name[i] >= 'a' && name[i] <= 'z') || (name[i] >= 'A' && name[i] <= 'Z') || (name[i] >= '0' && name[i] <= '9') {
			continue
		}
		return false
	}

	// 变量名合法
	return true
}

