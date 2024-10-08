package models

import (
	"errors"
	"time"
	"auraskin/pkg/utils"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j/dbtype"
)

type Order struct {
	OrderID       string  `json:"order_id"`
	Country       string  `json:"country"`
	DeliveryFee   float64 `json:"delivery_fee"`
	AddressLine   string  `json:"address_line"`
	Province      string  `json:"province"`
	TotalAmount   float64 `json:"total_amount"`
	District      string  `json:"district"`
	Ward          string  `json:"ward"`
	RecipientName string  `json:"recipient_name"`
	ContactNumber string  `json:"contact_number"`
	Status        string  `json:"status"`
	CreatedAt     string  `json:"created_at"`
}

func (o *Order) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"order_id":        o.OrderID,
		"country":         o.Country,
		"delivery_fee":    o.DeliveryFee,
		"address_line":    o.AddressLine,
		"province":        o.Province,
		"total_amount":    o.TotalAmount,
		"district":        o.District,
		"ward":            o.Ward,
		"recipient_name":  o.RecipientName,
		"contact_number":  o.ContactNumber,
		"status":          o.Status,
		"created_at":      o.CreatedAt,
	}
}

func (o *Order) FromMap(data map[string]interface{}) (*Order, error) {
	var createdAt string

	// Hàm helper để thử các định dạng thời gian khác nhau
	tryParseDate := func(value string, formats ...string) (string, error) {
		for _, format := range formats {
			if t, err := time.Parse(format, value); err == nil {
				return t.Format("2006-01-02"), nil // Chuẩn hóa về định dạng YYYY-MM-DD
			}
		}
		return "", errors.New("invalid date format")
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

	return &Order{
		OrderID:       utils.GetString(data, "order_id"),
		Country:       utils.GetString(data, "country"),
		DeliveryFee:   utils.GetFloat64(data, "delivery_fee"),
		AddressLine:   utils.GetString(data, "address_line"),
		Province:      utils.GetString(data, "province"),
		TotalAmount:   utils.GetFloat64(data, "total_amount"),
		District:      utils.GetString(data, "district"),
		Ward:          utils.GetString(data, "ward"),
		RecipientName: utils.GetString(data, "recipient_name"),
		ContactNumber: utils.GetString(data, "contact_number"),
		Status:        utils.GetString(data, "status"),
		CreatedAt:     createdAt,
	}, nil
}