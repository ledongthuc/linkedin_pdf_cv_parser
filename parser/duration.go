package parser

import "strings"

type Duration struct {
	From string `json:"from"`
	To   string `json:"to"`
	Time string `json:"time"`
}

// ### Checker ###
func (p *Parser) isDurationTime(text string) bool {
	checkingText := strings.ToLower(strings.TrimSpace(text))
	if strings.Contains(checkingText, "\n") {
		return false
	}
	if !strings.HasPrefix(checkingText, "(") ||
		!strings.HasSuffix(checkingText, ")") {
		return false
	}

	if !strings.ContainsAny(checkingText, "year month day") {
		return false
	}

	return true
}

// ### Parser ###
func (p *Parser) parseDuration(text string) (result Duration) {
	parts := strings.Split(text, " - ")
	if len(parts) > 0 {
		result.From = strings.TrimSpace(parts[0])
	}
	if len(parts) > 1 {
		result.To = strings.TrimSpace(parts[1])
	}
	return result
}
