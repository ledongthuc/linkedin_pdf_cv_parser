package parser

import (
	"strings"

	"github.com/ledongthuc/pdf"
)

type Experience struct {
	Duration    Duration `json:"duration"`
	Position    string   `json:"position"`
	CompanyName string   `json:"company_name"`
	Description string   `json:"description"`
}

// ### Checker ###
func (p *Parser) isBeginExperience(text pdf.Text, sentenceIndex int) bool {
	if strings.ToLower(text.S) == "experience\n" && p.Seeker >= IndexEmail {
		return true
	}
	return false
}

func (p *Parser) isTitleExperience(text pdf.Text, sentenceIndex int) bool {
	if p.Seeker == IndexBeginExperience || p.Seeker == IndexContentExperience {
		return true
	}
	return false
}

func (p *Parser) isContentExperience(text pdf.Text, sentenceIndex int) bool {
	if p.Seeker == IndexTitleExperience {
		return true
	}
	return false
}

// ### Parser ###
func (p *Parser) parseTitleExperiences(text string) Experience {
	result := Experience{}
	parts := strings.Split(text, " at ")

	if len(parts) > 0 {
		result.Position = strings.TrimSpace(parts[0])
	}

	if len(parts) > 1 {
		result.CompanyName = strings.TrimSpace(parts[1])
	}

	return result
}

func (p *Parser) parseContentExperiences(text string) (result Experience) {
	parts := strings.Split(text, "\n")
	if len(parts) > 0 {
		result.Duration = p.parseDuration(parts[0])
	}
	if len(parts) <= 1 {
		return
	}

	var contentParts []string
	if p.isDurationTime(parts[1]) {
		result.Duration.Time = strings.Trim(parts[1], "()")
		if len(parts) > 2 {
			contentParts = parts[2 : len(parts)-1]
		}
	} else {
		contentParts = parts[1 : len(parts)-1]
	}

	for _, part := range contentParts {
		result.Description += strings.TrimSpace(part) + ". "
	}
	return result
}
