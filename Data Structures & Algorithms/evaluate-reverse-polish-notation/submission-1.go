func evalRPN(tokens []string) int {
	stack := []int{}
	for _, token := range tokens {
		switch token {
			case "+" :
				operand1 := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				operand2 := stack[len(stack)-1]
				stack = stack[:len(stack)-1]

				stack = append(stack, operand1+operand2)
			case "-" :
				operand1 := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				operand2 := stack[len(stack)-1]
				stack = stack[:len(stack)-1]

				stack = append(stack, operand2-operand1)
			case "*" :
				operand1 := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				operand2 := stack[len(stack)-1]
				stack = stack[:len(stack)-1]

				stack = append(stack, operand1*operand2)
			case "/" :
				operand1 := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				operand2 := stack[len(stack)-1]
				stack = stack[:len(stack)-1]

				if operand1 != 0 {
					stack = append(stack, operand2 / operand1)
				}
				
			default:
				number,_ := strconv.Atoi(token)
				stack = append(stack, number)
		}
	}

	return stack[len(stack)-1]
}
