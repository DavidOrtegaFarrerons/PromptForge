package domain

import "time"

type PromptID string
type Prompt struct {
	promptID  PromptID
	ownerID   string
	title     string
	template  PromptTemplate
	tags      []string
	createdAt time.Time
	updatedAt time.Time
}

func NewPrompt(promptID PromptID, ownerID string, title string, promptTemplate PromptTemplate, tags []string, createdAt time.Time, updatedAt time.Time) (Prompt, error) {
	if promptID == "" {
		return Prompt{}, ErrPromptIDEmpty
	}

	if ownerID == "" {
		return Prompt{}, ErrOwnerIdEmpty
	}

	if len(title) < 4 {
		return Prompt{}, ErrTitleMinLength
	}

	if len(title) > 64 {
		return Prompt{}, ErrTitleMaxLength
	}

	if len(tags) != 0 {
		seen := make(map[string]bool)
		uniqueTags := make([]string, 0)

		for _, p := range tags {
			if !seen[p] {
				seen[p] = true
				uniqueTags = append(uniqueTags, p)
			}
		}

		tags = uniqueTags
	}

	p := Prompt{
		promptID:  promptID,
		ownerID:   ownerID,
		title:     title,
		template:  promptTemplate,
		tags:      tags,
		createdAt: createdAt,
		updatedAt: updatedAt,
	}

	return p, nil
}

func (p Prompt) PromptID() PromptID {
	return p.promptID
}

func (p Prompt) OwnerID() string {
	return p.ownerID
}

func (p Prompt) Title() string {
	return p.title
}

func (p Prompt) Template() PromptTemplate {
	return p.template
}

func (p Prompt) Tags() []string {
	return p.tags
}

func (p Prompt) CreatedAt() time.Time {
	return p.createdAt
}

func (p Prompt) UpdatedAt() time.Time {
	return p.updatedAt
}
