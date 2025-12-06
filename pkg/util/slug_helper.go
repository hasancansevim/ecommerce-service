package util

import (
	"strings"

	"github.com/google/uuid"
	"github.com/gosimple/slug"
)

func GenerateSlug(text string) string {
	slug.CustomSub = map[string]string{
		"ı": "i", "İ": "i", "ş": "s", "Ş": "s", "ğ": "g", "Ğ": "g",
		"ü": "u", "Ü": "u", "ö": "o", "Ö": "o", "ç": "c", "Ç": "c",
	}
	return slug.Make(text)
}

func GenerateUniqueSlug(text string) string {
	s := GenerateSlug(text)
	randomPart := strings.Split(uuid.New().String(), "-")[0]
	return s + "-" + randomPart
}
