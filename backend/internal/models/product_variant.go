package models

import (
	"auraskin/pkg/utils"
	"errors"
)

type ProductVariant struct {
	VariantID         string   `json:"variant_id"`
	VariantName       string   `json:"variant_name"`
	Size              string   `json:"size"`
	Color             string   `json:"color"`
	Price             float64  `json:"price"`
	StockQuantity     int      `json:"stock_quantity"`
	Thumbnail         string   `json:"thumbnail"`
	DescriptionImages []string `json:"descriptionImages"`
	IsActive          bool     `json:"is_active"`
}

func (pv *ProductVariant) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"variant_id":        pv.VariantID,
		"variant_name":      pv.VariantName,
		"size":              pv.Size,
		"color":             pv.Color,
		"price":             pv.Price,
		"stock_quantity":    pv.StockQuantity,
		"thumbnail":         pv.Thumbnail,
		"descriptionImages": pv.DescriptionImages,
		"is_active":         pv.IsActive,
	}
}

func (pv *ProductVariant) FromMap(data map[string]interface{}) (*ProductVariant, error) {
	price := utils.GetFloat64(data, "price")
	if price <= 0.0 {
		return nil, errors.New("invalid or missing price")
	}

	stockQuantity := utils.GetInt(data, "stock_quantity")

	return &ProductVariant{
		VariantID:         utils.GetString(data, "variant_id"),
		VariantName:       utils.GetString(data, "variant_name"),
		Size:              utils.GetString(data, "size"),
		Color:             utils.GetString(data, "color"),
		Price:             price,
		StockQuantity:     stockQuantity,
		DescriptionImages: utils.GetStringArray(data, "descriptionImages"),
		IsActive:          utils.GetBool(data, "is_active"),
	}, nil
}
