package hrule

import "fmt"

type BoolExpr struct{}

func (e BoolExpr) OpPriority(token string) int {
	opPriority := map[string]int{
		"&": 1,
		"|": 1,
	}

	return opPriority[token]
}

func (e BoolExpr) IsOperator(token string) bool {
	return token == "&" || token == "|"
}

func (e BoolExpr) IsParenthesis(token string) bool {
	return token == "(" || token == ")"
}

func (e BoolExpr) Tokenizer(expr string) []string {
	l := len(expr)
	i := 0
	var tokens []string
	for i < l {
		for expr[i] == ' ' {
			i++
		}
		if e.IsOperator(expr[i : i+1]) {
			tokens = append(tokens, expr[i:i+1])
			i++
		} else if e.IsParenthesis(expr[i : i+1]) {
			tokens = append(tokens, expr[i:i+1])
			i++
		} else {
			j := i
			for j < l && !e.IsOperator(expr[j:j+1]) && !e.IsParenthesis(expr[j:j+1]) {
				j++
			}
			tokens = append(tokens, expr[i:j])
			i = j
		}
	}

	return tokens
}

func (e BoolExpr) ToPolish(tokens []string) ([]string, error) {
	var polish []string
	var stack []string
	for _, token := range tokens {
		if e.IsOperator(token) {
			for len(stack) != 0 && e.IsOperator(stack[len(stack)-1]) && e.OpPriority(stack[len(stack)-1]) >= e.OpPriority(token) {
				polish = append(polish, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, token)
		} else if token == "(" {
			stack = append(stack, token)
		} else if token == ")" {
			for len(stack) != 0 && stack[len(stack)-1] != "(" {
				polish = append(polish, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			if len(stack) == 0 {
				return nil, fmt.Errorf("parenthesis not match")
			}
			stack = stack[:len(stack)-1]
		} else {
			polish = append(polish, token)
		}
	}
	for len(stack) != 0 {
		polish = append(polish, stack[len(stack)-1])
		stack = stack[:len(stack)-1]
	}

	return polish, nil
}
