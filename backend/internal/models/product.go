package models

import (
	"errors"
	"time"
	"auraskin/pkg/utils"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j/dbtype"
)

type Product struct {
	ProductID       string  `json:"product_id"`
	ProductName     string  `json:"product_name"`
	Description     string  `json:"description"`
	Features        string  `json:"features"`
	Origin          string  `json:"origin"`
	ManufacturedIn  string  `json:"manufactured_in"`
	Usage           string  `json:"usage"`
	DefaultPrice    float64 `json:"default_price"`
	Capacity        string  `json:"capacity"`
	Ingredients     string  `json:"ingredients"`
	DefaultImage    string  `json:"default_image"`
	Storage         string  `json:"storage"`
	ExpirationDate  string  `json:"expiration_date"`
	CreatedAt       string  `json:"created_at"`
	TargetCustomers string  `json:"target_customers"`
}

func (p *Product) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"product_id":       p.ProductID,
		"product_name":     p.ProductName,
		"description":      p.Description,
		"features":         p.Features,
		"origin":           p.Origin,
		"manufactured_in":  p.ManufacturedIn,
		"usage":            p.Usage,
		"default_price":    p.DefaultPrice,
		"capacity":         p.Capacity,
		"ingredients":      p.Ingredients,
		"default_image":    p.DefaultImage,
		"storage":          p.Storage,
		"expiration_date":  p.ExpirationDate,
		"created_at":       p.CreatedAt,
		"target_customers": p.TargetCustomers,
	}
}

func (p *Product) FromMap(data map[string]interface{}) (*Product, error) {
	var expirationDate, createdAt string

	// Hàm helper để thử các định dạng thời gian khác nhau
	tryParseDate := func(value string, formats ...string) (string, error) {
		for _, format := range formats {
			if t, err := time.Parse(format, value); err == nil {
				return t.Format("2006-01-02"), nil // Chuẩn hóa về định dạng YYYY-MM-DD
			}
		}
		return "", errors.New("invalid date format")
	}

	// Xử lý trường expiration_date
	if val, ok := data["expiration_date"]; ok {
		switch v := val.(type) {
		case string:
			parsedDate, err := tryParseDate(v, "2006-01-02", time.RFC3339)
			if err != nil {
				return nil, errors.New("invalid expiration date format")
			}
			expirationDate = parsedDate
		case dbtype.Date:
			expirationDate = v.String()
		default:
			expirationDate = ""
		}
	}

	// Xử lý trường created_at
	if val, ok := data["created_at"]; ok {
		switch v := val.(type) {
		case string:
			parsedDate, err := tryParseDate(v, "2006-01-02", time.RFC3339)
			if err != nil {
				return nil, errors.New("invalid created_at format")
			}
			createdAt = parsedDate
		case dbtype.Date:
			createdAt = v.String()
		default:
			createdAt = ""
		}
	}

	// Trả về Product với các trường đã được xử lý
	return &Product{
		ProductID:       utils.GetString(data, "product_id"),
		ProductName:     utils.GetString(data, "product_name"),
		Description:     utils.GetString(data, "description"),
		Features:        utils.GetString(data, "features"),
		Origin:          utils.GetString(data, "origin"),
		ManufacturedIn:  utils.GetString(data, "manufactured_in"),
		Usage:           utils.GetString(data, "usage"),
		DefaultPrice:    utils.GetFloat64(data, "default_price"),
		Capacity:        utils.GetString(data, "capacity"),
		Ingredients:     utils.GetString(data, "ingredients"),
		DefaultImage:    utils.GetString(data, "default_image"),
		Storage:         utils.GetString(data, "storage"),
		ExpirationDate:  expirationDate,
		CreatedAt:       createdAt,
		TargetCustomers: utils.GetString(data, "target_customers"),
	}, nil
}

