package hrule

import (
	"fmt"
	"reflect"
)

var boolExpr = BoolExpr{}

func NewCond(expr string, t reflect.Type) (*Cond, error) {
	// (>=3 & <=4) | (>=7 & <=8)
	polish, err := boolExpr.ToPolish(boolExpr.Tokenizer(expr))
	if err != nil {
		return nil, err
	}

	var stack []*Cond
	for _, token := range polish {
		if boolExpr.IsOperator(token) {
			if len(stack) == 0 {
				return nil, fmt.Errorf("miss operand")
			}
			right := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				return nil, fmt.Errorf("miss operand")
			}
			left := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			cond := &Cond{
				token: token,
				left:  left,
				right: right,
			}
			stack = append(stack, cond)
		} else {
			rule, err := NewRule(token, t)
			if err != nil {
				return nil, err
			}
			cond := &Cond{
				token: token,
				rule:  rule,
			}
			stack = append(stack, cond)
		}
	}

	if len(stack) != 1 {
		return nil, fmt.Errorf("miss operator")
	}

	return stack[0], nil
}

type Cond struct {
	token string
	left  *Cond
	right *Cond
	rule  Rule
}

func (c *Cond) Evaluate(v interface{}) bool {
	if c.left == nil && c.right == nil {
		return c.rule(v)
	}

	var lb, rb bool
	if c.left != nil {
		lb = c.left.Evaluate(v)
		if c.token == "&" && !lb {
			return lb
		}
		if c.token == "|" && lb {
			return lb
		}
	}

	if c.right != nil {
		rb = c.right.Evaluate(v)
		if c.token == "&" && !rb {
			return rb
		}
		if c.token == "|" && rb {
			return rb
		}
	}

	if c.token == "&" {
		return true
	}
	return false
}
