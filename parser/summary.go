package parser

import (
	"strings"

	"github.com/ledongthuc/pdf"
)

// ### Checker ###
func (p *Parser) isBeginSummary(text pdf.Text, sentenceIndex int) bool {
	if strings.ToLower(text.S) == "summary\n" && p.Seeker >= IndexEmail {
		return true
	}
	return false
}

// ### Parser ###
func (p *Parser) isSummary(text pdf.Text, sentenceIndex int) bool {
	if p.Seeker == IndexBeginSummary {
		return true
	}
	return false
}

func (p *Parser) parseSummary(text string) string {
	return strings.TrimSpace(text)
}
