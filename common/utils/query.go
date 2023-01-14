package utils

import (
	"fmt"
	"gym-app/app/model"

	"gorm.io/gorm"
)

func CreateSearchProgramQuery(params *model.SearchRequest, tx *gorm.DB) *gorm.DB {
	if params.SortBy == "" {
		params.SortBy = "date"
		params.Order = "asc"
	}
  if params.GroupName != "" {
    tx.Where("group_name = ?", params.GroupName)
  } else {
    tx.Where("group_name = ?", "UDARNIK")
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
		text := "%" + params.Text + "%"
		tx.Where("text LIKE ?", text)
	}
	tx.Order(fmt.Sprintf("%s %s", params.SortBy, params.Order))
	return tx
}
