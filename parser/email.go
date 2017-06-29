package parser

import (
	"strings"

	"github.com/ledongthuc/pdf"
)

// ### Checker ###
func (p *Parser) isEmail(text pdf.Text, sentenceIndex int) bool {
	if text.FontSize == 13 && p.Seeker < IndexEmail {
		return true
	}
	return false
}

// ### Parser ###
func (p *Parser) parseEmail(text string) string {
	texts := strings.Split(text, "\n")
	if len(texts) > 1 {
		return texts[1]
	}
	return ""
}
