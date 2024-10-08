package models

import (
	"auraskin/pkg/utils"
)

type Category struct {
	CategoryID   string `json:"category_id"`
	CategoryName string `json:"category_name"`
}

// ToMap converts a Category object to a map
func (c *Category) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"category_id":   c.CategoryID,
		"category_name": c.CategoryName,
	}
}

// FromMap converts a map to a Category object
func (c *Category) FromMap(data map[string]interface{}) (*Category, error) {
	categoryID := utils.GetString(data, "category_id")
	categoryName := utils.GetString(data, "category_name")

	return &Category{
		CategoryID:   categoryID,  
		CategoryName: categoryName,
	}, nil
}
