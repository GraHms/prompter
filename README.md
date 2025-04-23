# Prompter

`prompter` is an idiomatic Go library for building LLaMA- and Chat-style prompts in a fluent, role-delimited way.  
It supports predefined roles (`system`, `user`, `assistant`), arbitrary custom roles, built-in list formatting, and a simple multimodal attachment API for visual workflows.

---

## Installation

```bash
go get github.com/grahms/prompter
```

---

## Quick Start

```go
import (
  "fmt"
  "github.com/grahms/prompter"
)

func main() {
  prompt := prompter.New().
    System("You are a culinary assistant.").
    User("What can I cook tonight with eggs, spinach and cheese?").
    Assistant("").
    Build()

  fmt.Println(prompt)
}
```

---

## Text-Only API


## Complex Multi-Turn Dialog Example

```go
prompt := prompter.New().
  System("You are a professional travel-planning assistant.").
  Role("validator", "Ensure dates use YYYY-MM-DD and budget is positive integer.").
  User("I want to plan a trip to Paris.").
  Assistant("Sure! What are your dates and budget?").
  User("2025-05-10 to 2025-05-20; NY→CDG; Budget: 1500 USD.").
  Assistant("Here’s a high-level itinerary:\n1. …").
  User("Add local culinary experiences on Day 3 and Day 5.").
  Assistant("").
  Build()
```

Sent to your LLM, the model will generate the next `<|assistant|>` response in context.

---

## Multimodal / Visual Prompts

Most LLaMA-style endpoints accept **text only**, but you can prepare a combined payload for a multimodal API:

1. **Attachment**: an image URL or raw bytes.
2. **VisualPromptBuilder**: reuse the same role API and add images.

```go
vb := prompter.NewVisualBuilder().
  System("You are an image captioning assistant. Generate a concise caption.").
  AddImageURL("input_image", "https://…/dog.jpg").
  Assistant("")

payload, err := vb.BuildJSON()
```

```json
{
  "text": "<|system|>\nYou are an image captioning assistant...\n\n<|assistant|>\n\n\n",
  "attachments": [
    { "Name": "input_image", "URL": "https://…/dog.jpg" }
  ]
}
```

> **Note**:
> - **LLaMA itself is text-only** and will ignore attachments if sent to its endpoint.
> - To use images, send this payload to a **multimodal** or **vision-capable** API (e.g. Titan Vision, Claude+Vision).

---

## Why Roles & Structured Dialogs?

1. **Separation of Concerns**  
   Each `<|role|>` block isolates intent—easy to read, extend, and debug.

2. **Model Alignment**  
   Chat-trained models expect `<|system|>`, `<|user|>`, `<|assistant|>` delimiters.

3. **Maintainability**  
   Insert new steps via:
   ```go
   builder.Role("validator", "Ensure currency is USD.")
   ```
   No fragile text splicing.

---

## Testing

BDD-style tests with [GoConvey](https://github.com/smartystreets/goconvey):

```bash
go get github.com/smartystreets/goconvey
go test ./...
```

---

## Contributing

1. Fork & clone
2. Branch (`git checkout -b feature/XYZ`)
3. `go mod tidy` & `gofmt -s -w .`
4. Update `README.md` if needed
5. PR against `main`

See [CODE_OF_CONDUCT.md](./CODE_OF_CONDUCT.md) and [CONTRIBUTING.md](./CONTRIBUTING.md).

---

## License

MIT — see [LICENSE](./LICENSE) for details.
