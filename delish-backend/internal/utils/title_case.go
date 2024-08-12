package utils

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// Convert a string to TitleCase
func ToTitleCase(s string) string {
	caser := cases.Title(language.English)
	return caser.String(s)
}