package cmd

import (
	"os"
	"strings"

	"github.com/olekukonko/tablewriter"
)

func output(m metaInfo) error {
	width := 60
	desc := formatLongText(m.description[0], width)
	keywords := strings.Join(m.keywords, ",")
	data := [][]string{
		{"url", m.url},
		{"title", m.title[0]},
		{"description", desc},
		{"keywords", keywords},
		{"canonical", m.canonical[0]},
		{"alternate", m.alternate[0]},
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "Sign"})
	table.SetAutoWrapText(false)
	for _, v := range data {
		table.Append(v)
	}

	table.Render()

	return nil
}

// 長い文字列を適切に表示するため、一定の文字数ごとに改行を入れる
func formatLongText(text string, interval int) string {
	splited := strings.Split(text, "")
	var result string
	for k, v := range splited {
		result += v
		if (k+1)%interval == 0 {
			result += "\n"
		}
	}
	return result
}
