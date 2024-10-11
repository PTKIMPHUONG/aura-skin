package models

import (
	"errors"
	"time"
	"auraskin/pkg/utils"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j/dbtype"
)

type Sale struct {
	SaleID      string  `json:"sale_id"`
	DateStart   string  `json:"date_start"`
	DateEnd     string  `json:"date_end"`
	PercentSale float64 `json:"percent_sale"`
	Description string  `json:"description"`
	IsActive    bool    `json:"is_active"`
}

func (s *Sale) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"sale_id":      s.SaleID,
		"date_start":   s.DateStart,
		"date_end":     s.DateEnd,
		"percent_sale": s.PercentSale,
		"description":  s.Description,
		"is_active":    s.IsActive,
	}
}

func (s *Sale) FromMap(data map[string]interface{}) (*Sale, error) {
	var dateStart, dateEnd string

	tryParseDate := func(value string, formats ...string) (string, error) {
		for _, format := range formats {
			if t, err := time.Parse(format, value); err == nil {
				return t.Format("2006-01-02"), nil 
			}
		}
		return "", errors.New("invalid date format")
	}

	// Xử lý trường date_start
	if val, ok := data["date_start"]; ok {
		switch v := val.(type) {
		case string:
			parsedDate, err := tryParseDate(v, "2006-01-02", time.RFC3339)
			if err != nil {
				return nil, errors.New("invalid date_start format")
			}
			dateStart = parsedDate
		case dbtype.Date:
			dateStart = v.String()
		default:
			dateStart = ""
		}
	}

	// Xử lý trường date_end
	if val, ok := data["date_end"]; ok {
		switch v := val.(type) {
		case string:
			parsedDate, err := tryParseDate(v, "2006-01-02", time.RFC3339)
			if err != nil {
				return nil, errors.New("invalid date_end format")
			}
			dateEnd = parsedDate
		case dbtype.Date:
			dateEnd = v.String()
		default:
			dateEnd = ""
		}
	}

	return &Sale{
		SaleID:      utils.GetString(data, "sale_id"),
		DateStart:   dateStart,
		DateEnd:     dateEnd,
		PercentSale: utils.GetFloat64(data, "percent_sale"),
		Description: utils.GetString(data, "description"),
		IsActive:    utils.GetBool(data, "is_active"),
	}, nil
}