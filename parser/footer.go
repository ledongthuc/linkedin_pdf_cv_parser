package parser

import (
	"strings"

	"github.com/ledongthuc/pdf"
)

// ### Checker ###
func (p *Parser) isFooter(text pdf.Text, sentenceIndex int) bool {
	if strings.TrimSpace(strings.ToLower(text.S)) == strings.TrimSpace(strings.ToLower(p.ResumeProfile.Name)) {
		return true
	}
	return false
}
