
# Prompter

`prompter` is an idiomatic Go library for building LLaMA- and Chat-style prompts in a fluent, role-delimited way. It lets you compose complex, multi-turn conversations with predefined roles (`system`, `user`, `assistant`) or any custom roles you like, plus built-in list formatting utilities.

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



## Complex Multi-Turn Dialog Example

Below is a realistic, multi-turn travel-planning dialog. We define a system prompt, a custom validator role to enforce date and budget formats, then simulate three back-and-forths between user and assistant—ending with the model’s next response.

```go
package main

import (
  "fmt"
  "github.com/grahms/prompter"
)

func main() {
  // Step 1: Define travel parameters
  // Step 2: Ask follow-up questions
  // Step 3: Provide an initial itinerary
  // Step 4: Refine with culinary suggestions

  prompt := prompter.New().
    System(
      "You are a professional travel-planning assistant. " +
      "You propose detailed multi-day itineraries (flights, hotels, activities) " +
      "and can refine suggestions on request.",
    ).
    Role(prompter.Role("validator"),
      "Ensure all dates use YYYY-MM-DD format and budget is a positive integer.",
    ).
    User("I want to plan a trip to Paris.").
    Assistant("Sure! What are your travel dates and budget?").
    User("Dates: 2025-05-10 to 2025-05-20; Departing from New York City; Budget: 1500 USD.").
    Assistant(
      "Great! Here is a high-level itinerary:\n" +
      "1. Flight NYC→CDG on 2025-05-10, Economy.\n" +
      "2. Hotel: Hôtel Le Régent (5 nights).\n" +
      "3. Day 1: Louvre + Seine river cruise.\n" +
      "… (days 2-10 itinerary)…\n",
    ).
    User("Can you add local culinary experiences on Day 3 and Day 5?").
    Assistant("").
    Build()

  fmt.Println(prompt)
}
```

This produces the following raw prompt:

```
<|system|>
You are a professional travel-planning assistant. You propose detailed multi-day itineraries (flights, hotels, activities) and can refine suggestions on request.

<|validator|>
Ensure all dates use YYYY-MM-DD format and budget is a positive integer.

<|user|>
I want to plan a trip to Paris.

<|assistant|>
Sure! What are your travel dates and budget?

<|user|>
Dates: 2025-05-10 to 2025-05-20; Departing from New York City; Budget: 1500 USD.

<|assistant|>
Great! Here is a high-level itinerary:
1. Flight NYC→CDG on 2025-05-10, Economy.
2. Hotel: Hôtel Le Régent (5 nights).
3. Day 1: Louvre + Seine river cruise.
… (days 2-10 itinerary)…

<|user|>
Can you add local culinary experiences on Day 3 and Day 5?

<|assistant|>

```

At this point, sending that prompt to your LLaMA-style or Chat API will generate a detailed response for the final `<|assistant|>` block, incorporating all prior context and the validator’s directive.

---

## Why Roles & Structured Dialogs?

1. **Clarity & Modularity**  
   Each `<|role|>` block isolates intent:
  - `<|system|>` for global rules
  - `<|user|>` for actual input
  - `<|assistant|>` marks reply points
  - `<|validator|>` (custom) for format checks

2. **Model Alignment**  
   Chat-trained models expect these delimiters. They respect `<|system|>` instructions above everything else, parse `<|user|>` as queries, and know to generate at `<|assistant|>`.

3. **Maintainability**  
   To insert a new validation step, simply add:
   ```go
   Role("validator", "Ensure currency is USD.")
   ```
   No fragile string splicing required.

---

## Testing

BDD-style GoConvey tests are included:

```bash
go get github.com/smartystreets/goconvey
go test ./...
```

---

## License

Released under MIT. See [LICENSE](LICENSE).
