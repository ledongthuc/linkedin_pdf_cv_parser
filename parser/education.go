package parser

import (
	"strings"

	"github.com/ledongthuc/pdf"
)

type Education struct {
	Duration Duration `json:"duration"`
	Name     string   `json:"name"`
	Degree   string   `json:"degree"`
	Field    string   `json:"field"`
	Activity string   `json:"activity"`
}

// ### Checker ###
func (p *Parser) isBeginEducation(text pdf.Text, sentenceIndex int) bool {
	if strings.ToLower(text.S) == "education\n" && p.Seeker >= IndexEmail {
		return true
	}
	return false
}

func (p *Parser) isEducationName(text pdf.Text, sentenceIndex int) bool {
	if p.Seeker == IndexBeginEducation ||
		(p.Seeker == IndexEducationDescription && !strings.HasPrefix(text.S, "Activities and Societies:")) ||
		p.Seeker == IndexEducationActivity {
		return true
	}
	return false
}

func (p *Parser) isEducationDescription(text pdf.Text, sentenceIndex int) bool {
	if p.Seeker == IndexEducationName {
		return true
	}
	return false
}

func (p *Parser) isBeginEducationActivity(text pdf.Text, sentenceIndex int) bool {
	if p.Seeker == IndexEducationDescription && strings.HasPrefix(text.S, "Activities and Societies:") {
		return true
	}
	return false
}

func (p *Parser) isEducationActivity(text pdf.Text, sentenceIndex int) bool {
	if p.Seeker == IndexBeginEducationActivity {
		return true
	}
	return false
}

// ### Parser ###
func (p *Parser) parseEducationName(text string) (result Education) {
	result.Name = strings.TrimSpace(text)
	return result
}

func (p *Parser) parseEducationDescription(text string) (result Education) {
	parts := strings.Split(text, ",")
	partsLength := len(parts)

	if partsLength == 0 {
		return
	}

	// Master's degree, HR, 2007 - 2007
	if partsLength >= 3 {
		result.Degree = strings.TrimSpace(parts[partsLength-3])
		result.Field = strings.TrimSpace(parts[partsLength-2])
		result.Duration = p.parseDuration(parts[partsLength-1])
		return
	}

	// Master's degree, 2007 - 2007
	if partsLength == 2 && strings.Contains(parts[1], " - ") {
		result.Degree = strings.TrimSpace(parts[0])
		result.Duration = p.parseDuration(parts[1])
		return
	}

	// Master's degree, HR
	if partsLength == 2 && !strings.Contains(parts[1], " - ") {
		result.Degree = strings.TrimSpace(parts[0])
		result.Field = strings.TrimSpace(parts[1])
		return
	}

	// Master's degree
	if !strings.Contains(parts[0], " - ") {
		result.Degree = strings.TrimSpace(parts[0])
	}

	// 2007 - 2007
	result.Duration = p.parseDuration(parts[0])

	result.Name = strings.TrimSpace(text)
	return result
}

func (p *Parser) parseEducationActivity(text string) string {
	return strings.TrimSpace(strings.Replace(text, "Activities and Societies: ", "", -1))
}
