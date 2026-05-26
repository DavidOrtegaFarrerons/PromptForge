package domain

import "regexp"

type PromptTemplate struct {
	content      string
	placeholders []string
}

var regexPromptPlaceholders = regexp.MustCompile("{{([A-Za-z0-9_-]+)}}")

func NewPromptTemplate(content string) (PromptTemplate, error) {
	if len(content) == 0 {
		return PromptTemplate{}, ErrContentEmpty
	}

	placeholders := regexPromptPlaceholders.FindAllStringSubmatch(content, -1)
	if len(placeholders) == 0 {
		return PromptTemplate{content: content, placeholders: make([]string, 0)}, nil
	}

	seen := make(map[string]bool)
	uniquePlaceholders := make([]string, 0)

	for _, p := range placeholders {
		if !seen[p[1]] {
			seen[p[1]] = true
			uniquePlaceholders = append(uniquePlaceholders, p[1])
		}
	}

	return PromptTemplate{content: content, placeholders: uniquePlaceholders}, nil
}

func (pt PromptTemplate) Content() string {
	return pt.content
}

func (pt PromptTemplate) Placeholders() []string {
	return pt.placeholders
}
