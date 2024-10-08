package models

import (
	"errors"
	"time"
	"auraskin/pkg/utils"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j/dbtype"
)

type Supplier struct {
	SupplierID    string `json:"supplier_id"`
	SupplierName  string `json:"supplier_name"`
	SupplierEmail string `json:"supplier_email"`
	SupplierPhone string `json:"supplier_phone"`
	DefaultImage  string `json:"default_image"`
	CreatedAt     string `json:"created_at"`
}

func (s *Supplier) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"supplier_id":    s.SupplierID,
		"supplier_name":  s.SupplierName,
		"supplier_email": s.SupplierEmail,
		"supplier_phone": s.SupplierPhone,
		"default_image":  s.DefaultImage,
		"created_at":     s.CreatedAt,
	}
}

func (s *Supplier) FromMap(data map[string]interface{}) (*Supplier, error) {
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

	return &Supplier{
		SupplierID:    utils.GetString(data, "supplier_id"),
		SupplierName:  utils.GetString(data, "supplier_name"),
		SupplierEmail: utils.GetString(data, "supplier_email"),
		SupplierPhone: utils.GetString(data, "supplier_phone"),
		DefaultImage:  utils.GetString(data, "default_image"),
		CreatedAt:     createdAt,
	}, nil
}
