package main

func balancedBrackets(bracketsToBalance []rune) bool {

	brackets := [][2]rune{
		{'[',']'},
		{'{', '}'},
		{'(',')'},
	}

	stack := []rune{}

	for i := 0; i < len(bracketsToBalance); i++ {
		for _, bracket := range brackets {
			if bracket[0] == bracketsToBalance[i] {
				// add an opening bracket to the stack
				stack = append(stack, bracketsToBalance[i])
				break
			}
		}
		for _, bracket := range brackets {
			if bracket[1] == bracketsToBalance[i] {
				// closing bracket
				if bracket[0] != stack[len(stack)-1] {
					// the closing bracket doesnt match
					return false
				} else {
					// the closing bracket does match
					stack = stack[:len(stack)-1]
				}
			}
		}
	}

	return true
}
