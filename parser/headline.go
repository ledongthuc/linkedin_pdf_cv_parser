package parser

import (
	"strings"

	"github.com/ledongthuc/pdf"
)

// ### Checker ###
func (p *Parser) isHeadline(text pdf.Text, sentenceIndex int) bool {
	if text.FontSize == 13 && p.Seeker < IndexHeadline {
		return true
	}
	return false
}

// ### Parser ###
func (p *Parser) parseHeadline(text string) string {
	texts := strings.Split(text, "\n")
	if len(texts) > 0 {
		return texts[0]
	}
	return ""
}
