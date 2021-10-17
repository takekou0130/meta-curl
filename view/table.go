package view

import (
	"os"
	"strings"

	"github.com/olekukonko/tablewriter"
)

type tableRenderer struct {
	width int
	table *tablewriter.Table
}

func NewTableRenderer() *tableRenderer {
	width := 60
	table := tablewriter.NewWriter(os.Stdout)
	return &tableRenderer{width: width, table: table}
}

func (tr *tableRenderer) Render(m MetaInfo) error {
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
