package main

/*
	表达式计算
    遍历表达式遇到（不处理
*/
func stackCompute(expressionString string) int {
	var (
		valueStack stackLink
		opStack    stackLink
	)
	for _, v := range expressionString {
		if v <= '9' && v >= '0' {
			valueStack.push()
		}
	}
}

func main() {

}
