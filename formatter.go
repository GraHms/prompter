package prompter

import "strings"

// FormatBulletList returns a bullet-pointed list.
// If bullet is empty, defaults to "-".
func FormatBulletList(bullet string, items []string) string {
	if bullet == "" {
		bullet = "-"
	}
	// prefix each item with “<bullet> ” and join with newline
	for i, it := range items {
		items[i] = bullet + " " + it
	}
	return strings.Join(items, "\n")
}

// FormatBulletListDefault is a shorthand using "-" as the bullet.
func FormatBulletListDefault(items []string) string {
	return FormatBulletList("-", items)
}
