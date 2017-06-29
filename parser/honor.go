package parser

import (
	"strings"

	"github.com/ledongthuc/pdf"
)

type Honor struct {
	Name string `json:"name"`
}

// ### Checker ###
func (p *Parser) isBeginHonor(text pdf.Text, sentenceIndex int) bool {
	if strings.ToLower(text.S) == "honors and awards\n" && p.Seeker >= IndexEmail {
		return true
	}
	return false
}

func (p *Parser) isHonor(text pdf.Text, sentenceIndex int) bool {
	if p.Seeker == IndexBeginHonor || p.Seeker == IndexHonor {
		return true
	}
	return false
}

// ### Parser ###
func (p *Parser) parseHonor(text string) (honor Honor) {
	honor.Name = strings.TrimSpace(text)
	return honor
}
