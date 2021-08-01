package utils

import (
	"fmt"
	"gorm.io/gorm"
	"gym-app/app/model"
)

func CreateSearchProgramQuery(params *model.SearchRequest, tx *gorm.DB) *gorm.DB {
	if params.Limit == 0 {
		params.Limit = 20
	}
	if params.SortBy == "" {
		params.SortBy = "date"
		params.Order = "asc"
	}
	tx.Order(fmt.Sprintf("%s %s", params.SortBy, params.Order))
	if params.StartDate != "" {
		startDate := StartOfTheDay(ParseDate(params.StartDate))
		tx.Where("date >= ?", startDate)
	}
	if params.EndDate != "" {
		endDate := EndOfTheDay(ParseDate(params.EndDate))
		tx.Where("date <= ?", endDate)
	}
	if params.Text != "" {
		text := "%"+params.Text+"%"
		tx.Where("text LIKE ?", text)
	}
	tx.Order(fmt.Sprintf("%s %s", params.SortBy, params.Order)).Limit(params.Limit)
	return tx
}
