package implementation

import (
	"errors"
	"fmt"
	"math"
	"regexp"
	"strconv"
)

// Stack data structure
type Stack []string

// IsEmpty : check if stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push : push a new value onto the stack
func (s *Stack) Push(str string) {
	*s = append(*s, str) // Simply append the new value to the end of the stack
}

// Pop : remove and return top element of stack. Return false if stack is empty.
func (s *Stack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	} else {
		index := len(*s) - 1   // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		*s = (*s)[:index]      // Remove it from the stack by slicing it off.
		return element, true
	}
}

func CalculatePostfixExpression(postfixExpression string) (float64, error) {
	if postfixExpression == "" {
		return 0, errors.New("Postfix expression string cannot be empty")
	}

	var operandStack Stack
	var tokenList = regexp.MustCompile(`\s+`).Split(postfixExpression, -1)

	for i := 0; i < len(tokenList); i++ {
		_, err := strconv.ParseFloat(tokenList[i], 64) // try to parse token into integer

		// if operand is integer
		if err == nil {
			operandStack.Push(tokenList[i])
		} else {
			token := tokenList[i]

			if token == "+" || token == "-" || token == "*" || token == "/" || token == "^" {
				operand2, hasMore2 := operandStack.Pop()
				operand1, hasMore1 := operandStack.Pop()

				if !hasMore2 && !hasMore1 {
					return 0, errors.New("Wrong postfix expression")
				}

				leftOperand, err1 := strconv.ParseFloat(operand1, 64)
				if err1 != nil {
					return 0, errors.New("Wrong postfix expression: left operand is not an number")
				}

				rightOperand, err2 := strconv.ParseFloat(operand2, 64)
				if err2 != nil {
					return 0, errors.New("Wrong postfix expression: right operand is not an number")
				}

				switch token {
				case "+":
					operandStack.Push(fmt.Sprintf("%f", leftOperand + rightOperand))
				case "-":
					operandStack.Push(fmt.Sprintf("%f", leftOperand - rightOperand))
				case "*":
					operandStack.Push(fmt.Sprintf("%f", leftOperand * rightOperand))
				case "/":
					operandStack.Push(fmt.Sprintf("%f", leftOperand / rightOperand))
				case "^":
					operandStack.Push(fmt.Sprintf("%f", math.Pow(leftOperand, rightOperand)))
				default:
					return 0, errors.New("Postfix expression contains invalid token")
				}
			} else {
				return 0, errors.New("Postfix expression contains invalid token")
			}
		}
	}

	result, hasMore := operandStack.Pop()
	if !hasMore || !operandStack.IsEmpty() {
		return 0, errors.New("Wrong postfix expression")
	}

	resultInt, err := strconv.ParseFloat(result, 64)

	return resultInt, err
}
