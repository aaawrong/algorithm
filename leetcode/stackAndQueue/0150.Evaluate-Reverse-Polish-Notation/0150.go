package stackAndQueue

import "strconv"

// 逆波兰表达式求值
// 力扣：
// https://leetcode.cn/problems/evaluate-reverse-polish-notation/description/
// 随想录：

// 逆波兰表达式的计算规则如下：
//
// 从左往右遍历表达式的每个元素
// 遇到操作数（数字）时，直接压入栈中
// 遇到运算符时，从栈顶弹出两个操作数进行运算，并将结果压回栈中
// 重复上述过程，直到表达式扫描完毕
// 最后栈中剩下的唯一元素就是表达式的计算结果

// 正常波兰表达式则是从右到左入栈

// 时间复杂度：O(n)
// 空间复杂度：O(n)
func evalRPN(tokens []string) int {
	var stack []int
	for _, token := range tokens {
		val, err := strconv.Atoi(token)
		if err == nil {
			stack = append(stack, val)
		} else { // err不为空说明是运算符
			num1, num2 := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			switch token {
			case "+":
				stack = append(stack, num1+num2)
			case "-":
				stack = append(stack, num1-num2)
			case "*":
				stack = append(stack, num1*num2)
			case "/":
				stack = append(stack, num1/num2)
			}
		}
	}
	return stack[0]
}
