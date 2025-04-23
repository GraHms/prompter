package prompter

import "encoding/json"

// VisualPrompt packages text + attachments into a single payload.
type VisualPrompt struct {
	Text        string       `json:"text"`
	Attachments []Attachment `json:"attachments,omitempty"`
}

// VisualPromptBuilder lets you compose both text (via PromptBuilder) and images.
type VisualPromptBuilder struct {
	prompt *PromptBuilder
	imgs   []Attachment
}

// NewVisualBuilder creates a fresh VisualPromptBuilder.
func NewVisualBuilder() *VisualPromptBuilder {
	return &VisualPromptBuilder{
		prompt: New(),
		imgs:   nil,
	}
}

// System delegates to the underlying PromptBuilder.
func (v *VisualPromptBuilder) System(txt string) *VisualPromptBuilder {
	v.prompt.System(txt)
	return v
}

// User delegates to the underlying PromptBuilder.
func (v *VisualPromptBuilder) User(txt string) *VisualPromptBuilder {
	v.prompt.User(txt)
	return v
}

// Assistant delegates to the underlying PromptBuilder.
func (v *VisualPromptBuilder) Assistant(txt string) *VisualPromptBuilder {
	v.prompt.Assistant(txt)
	return v
}

// Role delegates to the underlying PromptBuilder.
func (v *VisualPromptBuilder) Role(role Role, txt string) *VisualPromptBuilder {
	v.prompt.Role(role, txt)
	return v
}

// AddImageURL registers an image by its URL.
func (v *VisualPromptBuilder) AddImageURL(name, url string) *VisualPromptBuilder {
	v.imgs = append(v.imgs, Attachment{Name: name, URL: url})
	return v
}

// Build packages everything into a VisualPrompt.
func (v *VisualPromptBuilder) Build() VisualPrompt {
	return VisualPrompt{
		Text:        v.prompt.Build(),
		Attachments: v.imgs,
	}
}

// BuildJSON returns the JSON-encoded payload ready for your multimodal API.
func (v *VisualPromptBuilder) BuildJSON() ([]byte, error) {
	vp := v.Build()
	return json.Marshal(vp)
}
