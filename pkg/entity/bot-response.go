package entity

import "strings"

type BotResponse struct {
	StreamName   string
	StreamStart  string
	StreamEnd    string
	CategoryName string
	IsRecurring  string
}

func (resp *BotResponse) ToString() string {
	return strings.Join([]string{"Name of stream:", resp.StreamName,
		"Stream start:", resp.StreamStart, "Stream end:", resp.StreamEnd,
		"Category name:", resp.CategoryName, "Is recurring:", resp.IsRecurring}, "")
}
