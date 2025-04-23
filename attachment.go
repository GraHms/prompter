package prompter

// Attachment represents an image (or any binary) to send alongside the prompt.
type Attachment struct {
	Name string // e.g. "input_image"
	URL  string // remote URL, or…
	// Data []byte // …or raw image bytes, if you prefer
	// Mime string // e.g. "image/png"
}
