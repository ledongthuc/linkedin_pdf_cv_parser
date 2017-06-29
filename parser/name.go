package parser

import (
	"strings"

	"github.com/ledongthuc/pdf"
)

// ### Checker ###
func (p *Parser) isName(text pdf.Text, sentenceIndex int) bool {
	if (text.FontSize == 20 && p.Seeker < IndexName) || (text.S == p.ResumeProfile.Name) {
		return true
	}
	return false
}

// ### Parser ###
func (p *Parser) parseName(text string) string {
	return strings.TrimSpace(text)
}
