package view

import (
	"strings"

	"github.com/Z00mZE/rss-news-tgbot/internal/app/bot/domain/model"
)

var (
	replacer = strings.NewReplacer(
		"-",
		"\\-",
		"_",
		"\\_",
		"*",
		"\\*",
		"[",
		"\\[",
		"]",
		"\\]",
		"(",
		"\\(",
		")",
		"\\)",
		"~",
		"\\~",
		"`",
		"\\`",
		">",
		"\\>",
		"#",
		"\\#",
		"+",
		"\\+",
		"=",
		"\\=",
		"|",
		"\\|",
		"{",
		"\\{",
		"}",
		"\\}",
		".",
		"\\.",
		"!",
		"\\!",
	)
)

func ArticleRender(data model.Article) string {
	var buff = strings.Builder{}
	buff.WriteString(`*` + replacer.Replace(data.Title) + "*\n")
	buff.WriteString(replacer.Replace(data.Summary) + "\n\n")
	buff.WriteString(replacer.Replace(data.Link))

	return buff.String()
}
