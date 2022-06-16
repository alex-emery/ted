package util

import markdown "github.com/MichaelMure/go-term-markdown"

func MarkdownToText(in string, width int) string {
	result := markdown.Render(string(in), width, 1)
	return string(result)
}
