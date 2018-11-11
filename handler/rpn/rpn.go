package rpn

import (
	"container/list"
	"strconv"
	"strings"
)

// Calculate returns arithmetic result
func Calculate(formula *string) string {
	separatedFormula := strings.Fields(*formula)
	if isValidBracket(separatedFormula) == false {
		return "SYNTAX ERROR"
	}

	queue := list.New()

	transform(queue, separatedFormula)

	result := resolve(queue)

	return result
}

// Checks validation of brackets
func isValidBracket(separatedFormula []string) bool {
	var bracketCount = 0
	for _, str := range separatedFormula {
		if str == "(" {
			bracketCount++
		} else if str == ")" {
			bracketCount--
		}
	}

	if bracketCount != 0 {
		return false
	} else {
		return true
	}
}

// Transforms Infix notation to Postfix notation (Reverse Polish Notation, RPN)
func transform(queue *list.List, separatedFormula []string) {
	stack := list.New()
	for _, str := range separatedFormula {
		rankVal := rank(&str)
		if str == "(" {
			stack.PushBack(str)
		} else if str == ")" {
			for el := stack.Back(); el != nil; el = el.Prev() {
				temp := el.Value.(string)
				if temp == "(" {
					stack.Remove(el)
					break
				}
				queue.PushBack(temp)
				stack.Remove(el)
			}
		} else if str == "+" || str == "-" || str == "*" || str == "/" {
			for el := stack.Back(); el != nil; el = el.Prev() {
				tempValue := el.Value.(string)
				tempRankVal := rank(&tempValue)
				if rankVal < tempRankVal {
					break
				}
				temp := stack.Remove(el).(string)
				queue.PushBack(temp)
			}
			stack.PushBack(str)
		} else { // Digits
			queue.PushBack(str)
		}
	}

	for el := stack.Back(); el != nil; el = el.Prev() {
		if el.Value.(string) != "(" {
			queue.PushBack(el.Value.(string))
		}
	}
}

// Resolve calculates RPN formatted formula
func resolve(queue *list.List) string {
	stack := list.New()
	var firstValue = ""
	var secondValue = ""
	var result = ""

	for el := queue.Front(); el != nil; el = el.Next() {
		str := el.Value.(string)
		if str == "+" || str == "-" || str == "*" || str == "/" {
			if stack.Back() != nil {
				firstValue = stack.Remove(stack.Back()).(string)
			} else {
				break
			}

			if stack.Back() != nil {
				secondValue = stack.Remove(stack.Back()).(string)
			} else {
				break
			}

			val1, _ := strconv.ParseFloat(firstValue, 64)
			val2, _ := strconv.ParseFloat(secondValue, 64)

			if str == "+" {
				result = strconv.FormatFloat(val2+val1, 'g', 10, 64)
			} else if str == "-" {
				result = strconv.FormatFloat(val2-val1, 'g', 10, 64)
			} else if str == "*" {
				result = strconv.FormatFloat(val2*val1, 'g', 10, 64)
			} else if str == "/" {
				if val1 == 0 {
					return "INFINITY"
				}
				result = strconv.FormatFloat(val2/val1, 'g', 10, 64)
			}
			stack.PushBack(result)
		} else {
			stack.PushBack(str)
		}
	}
	return stack.Back().Value.(string)
}

// Returns each priority of operators
func rank(str *string) int {
	if *str == "(" || *str == ")" {
		return 3
	} else if *str == "+" || *str == "-" {
		return 2
	} else if *str == "*" || *str == "/" {
		return 1
	} else { // Digits
		return 0
	}
}
