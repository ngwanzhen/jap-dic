package main

import (
	"github.com/bregydoc/gtranslate"
)

func translate(s string, lang string) string {
	// fmt.Println(s)
	// text := "Hello World"
	text := s

	if lang == "ja" {
		translated, err := gtranslate.TranslateWithParams(
			text,
			gtranslate.TranslationParams{
				From: "ja",
				To:   "en",
			},
		)
		return translated
		if err != nil {
			panic(err)
		}
	} else {
		translated, err := gtranslate.TranslateWithParams(
			text,
			gtranslate.TranslationParams{
				From: "en",
				To:   "ja",
			},
		)
		return translated
		if err != nil {
			panic(err)
		}
	}
	return "translated"

	// fmt.Printf("en: %s | ja: %s \n", text, translated)

	// en: Hello World | ja: こんにちは世界
}
