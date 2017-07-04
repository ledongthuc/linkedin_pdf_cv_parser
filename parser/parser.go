package parser

import (
	"fmt"
	"strings"

	"github.com/ledongthuc/pdf"
)

type PageIndex int

type Parser struct {
	Seeker        PageIndex
	ResumeProfile ResumeProfile
}

type ResumeProfile struct {
	Name        string       `json:"name"`
	Headline    string       `json:"headline"`
	Email       string       `json:"email"`
	Summary     string       `json:"summary"`
	Experiences []Experience `json:"experiences"`
	Educations  []Education  `json:"educations""`
	Honors      []Honor      `json:"honors"`
}

const (
	IndexNone PageIndex = iota
	IndexName
	IndexHeadline
	IndexEmail
	IndexBeginSummary
	IndexSummary
	IndexBeginExperience
	IndexTitleExperience
	IndexContentExperience
	IndexBeginEducation
	IndexEducationName
	IndexEducationDescription
	IndexBeginEducationActivity
	IndexEducationActivity
	IndexBeginHonor
	IndexHonor
	IndexEnd
)

func ParsePDF(path string) (ResumeProfile, error) {
	r, err := pdf.Open(path)
	if err != nil {
		return ResumeProfile{}, err
	}
	totalPage := r.NumPage()

	var parser Parser
	var sentenceIndex = 0
	var lastTextStyle pdf.Text

	for pageIndex := 1; pageIndex <= totalPage; pageIndex++ {
		p := r.Page(pageIndex)
		if p.V.IsNull() {
			continue
		}
		texts := p.Content().Text
		for _, text := range texts {
			if isPageNumber(text) {
				continue
			}

			if isSameSentence(text, lastTextStyle) {
				lastTextStyle.S = lastTextStyle.S + text.S
			} else {
				if !isEmptyText(lastTextStyle.S) {
					// fmt.Printf("Debug: %+v\n", lastTextStyle)
					parser.parseData(lastTextStyle, sentenceIndex)
					sentenceIndex++
				}
				lastTextStyle = text
			}
		}
	}
	return parser.ResumeProfile, nil
}

func (p *Parser) parseData(text pdf.Text, sentenceIndex int) {
	if p.isFooter(text, sentenceIndex) {
		p.Seeker = IndexEnd
		return
	}
	if p.isName(text, sentenceIndex) {
		p.ResumeProfile.Name = p.parseName(text.S)
		p.Seeker = IndexName
	}
	if p.isHeadline(text, sentenceIndex) {
		p.ResumeProfile.Headline = p.parseHeadline(text.S)
		p.Seeker = IndexHeadline
	}
	if p.isEmail(text, sentenceIndex) {
		p.ResumeProfile.Email = p.parseEmail(text.S)
		p.Seeker = IndexEmail
	}
	if p.isBeginSummary(text, sentenceIndex) {
		p.Seeker = IndexBeginSummary
		return
	}
	if p.isBeginExperience(text, sentenceIndex) {
		p.Seeker = IndexBeginExperience
		return
	}
	if p.isBeginEducation(text, sentenceIndex) {
		p.Seeker = IndexBeginEducation
		return
	}
	if p.isBeginHonor(text, sentenceIndex) {
		p.Seeker = IndexBeginHonor
		return
	}
	if p.isSummary(text, sentenceIndex) {
		p.ResumeProfile.Summary += p.parseSummary(text.S)
		p.Seeker = IndexSummary
		return
	}
	if p.isTitleExperience(text, sentenceIndex) {
		p.ResumeProfile.Experiences = append(p.ResumeProfile.Experiences, p.parseTitleExperiences(text.S))
		p.Seeker = IndexTitleExperience
		return
	}
	if p.isContentExperience(text, sentenceIndex) {
		experience := p.parseContentExperiences(text.S)
		updatingExperience := p.ResumeProfile.Experiences[len(p.ResumeProfile.Experiences)-1]
		p.ResumeProfile.Experiences = p.ResumeProfile.Experiences[:len(p.ResumeProfile.Experiences)-1]
		updatingExperience.Duration = experience.Duration
		updatingExperience.Description = experience.Description
		p.ResumeProfile.Experiences = append(p.ResumeProfile.Experiences, updatingExperience)
		p.Seeker = IndexContentExperience
		return
	}
	fmt.Println(" - " + text.Font)
	fmt.Printf(" - %f\n", text.FontSize)
	fmt.Println(" - " + text.S)
	if p.isEducationName(text, sentenceIndex) {
		p.ResumeProfile.Educations = append(p.ResumeProfile.Educations, p.parseEducationName(text.S))
		p.Seeker = IndexEducationName
		return
	}
	if p.isEducationDescription(text, sentenceIndex) {
		education := p.parseEducationDescription(text.S)
		updatingEducation := p.ResumeProfile.Educations[len(p.ResumeProfile.Educations)-1]
		p.ResumeProfile.Educations = p.ResumeProfile.Educations[:len(p.ResumeProfile.Educations)-1]
		updatingEducation.Degree = education.Degree
		updatingEducation.Field = education.Field
		updatingEducation.Duration = education.Duration
		p.ResumeProfile.Educations = append(p.ResumeProfile.Educations, updatingEducation)
		p.Seeker = IndexEducationDescription
		return
	}
	if p.isBeginEducationActivity(text, sentenceIndex) {
		p.Seeker = IndexBeginEducationActivity
		return
	}
	if p.isEducationActivity(text, sentenceIndex) {
		updatingEducation := p.ResumeProfile.Educations[len(p.ResumeProfile.Educations)-1]
		p.ResumeProfile.Educations = p.ResumeProfile.Educations[:len(p.ResumeProfile.Educations)-1]
		updatingEducation.Activity = p.parseEducationActivity(text.S)
		p.ResumeProfile.Educations = append(p.ResumeProfile.Educations, updatingEducation)
		p.Seeker = IndexEducationActivity
		return
	}
	if p.isHonor(text, sentenceIndex) {
		p.ResumeProfile.Honors = append(p.ResumeProfile.Honors, p.parseHonor(text.S))
		p.Seeker = IndexHonor
		return
	}
}

func isPageNumber(text pdf.Text) bool {
	if text.FontSize == 10 &&
		text.X > 500 {
		return true
	}
	return false
}

func isSameSentence(text1, text2 pdf.Text) bool {
	if text1.Font == text2.Font &&
		text1.FontSize == text2.FontSize {
		return true
	}
	return false
}

func isEmptyText(text string) bool {
	return strings.TrimSpace(text) == ""
}
