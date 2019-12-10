package hrule

import (
	"regexp"
	"strings"
)

type RuleExpr struct{}

var tokenRegex = regexp.MustCompile(`^([a-zA-Z][a-zA-Z0-9]*)(\((.*?)\))?$`)

func (e RuleExpr) Tokenizer(expr string) (string, string) {
	expr = strings.Trim(expr, " ")
	if strings.HasPrefix(expr, ">=") || strings.HasPrefix(expr, "<=") || strings.HasPrefix(expr, "==") {
		return expr[:2], expr[2:]
	}
	if strings.HasPrefix(expr, ">") || strings.HasPrefix(expr, "<") || strings.HasPrefix(expr, "=") {
		return expr[:1], expr[1:]
	}

	idx := strings.Index(expr, " ")
	if idx == -1 {
		return expr, ""
	}
	return expr[:idx], expr[idx+1:]
}
