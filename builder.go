package prompter

import (
	"fmt"
	"strings"
)

// message holds a role and its associated content.
type message struct {
	role    Role
	content string
}

// PromptBuilder builds a role-delimited prompt in a fluent style.
type PromptBuilder struct {
	msgs []message
}

// New creates and returns a new PromptBuilder.
func New() *PromptBuilder {
	return &PromptBuilder{msgs: make([]message, 0, 4)}
}

// Add appends a block with any custom Role.
func (b *PromptBuilder) Add(role Role, content string) *PromptBuilder {
	b.msgs = append(b.msgs, message{role: role, content: content})
	return b
}

// System adds a "system" instruction block.
func (b *PromptBuilder) System(text string) *PromptBuilder {
	return b.Add(RoleSystem, text)
}

// User adds a "user" input block.
func (b *PromptBuilder) User(text string) *PromptBuilder {
	return b.Add(RoleUser, text)
}

// Assistant adds an "assistant" block marking where generation starts.
func (b *PromptBuilder) Assistant(text string) *PromptBuilder {
	return b.Add(RoleAssistant, text)
}

// Role is an alias for Add(), letting you insert any custom Role.
func (b *PromptBuilder) Role(role Role, text string) *PromptBuilder {
	return b.Add(role, text)
}

// Build constructs the final prompt string with <|role|> delimiters.
func (b *PromptBuilder) Build() string {
	var sb strings.Builder
	for _, m := range b.msgs {
		sb.WriteString("<|")
		sb.WriteString(string(m.role))
		sb.WriteString("|>\n")
		sb.WriteString(m.content)
		sb.WriteString("\n\n")
	}
	return sb.String()
}

// List appends a titled bullet-list block under the given role.
// e.g. .List(RoleUser, "Ingredients", []string{"eggs","flour"})
func (b *PromptBuilder) List(role Role, title string, items []string) *PromptBuilder {
	content := fmt.Sprintf("%s:\n%s", title, FormatBulletListDefault(items))
	return b.Add(role, content)
}

// UserList is a convenience for List(RoleUser, …).
func (b *PromptBuilder) UserList(title string, items []string) *PromptBuilder {
	return b.List(RoleUser, title, items)
}

// SystemList is a convenience for List(RoleSystem, …).
func (b *PromptBuilder) SystemList(title string, items []string) *PromptBuilder {
	return b.List(RoleSystem, title, items)
}

// AssistantList is a convenience for List(RoleAssistant, …).
func (b *PromptBuilder) AssistantList(title string, items []string) *PromptBuilder {
	return b.List(RoleAssistant, title, items)
}
