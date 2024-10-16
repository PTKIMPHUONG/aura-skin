package models

import (
	"auraskin/pkg/utils"
	"errors"
	"fmt"
	"time"

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

func (s *Sale) ToResponseMap() map[string]interface{} {
	return map[string]interface{}{
		"sale_id":      s.SaleID,
		"date_start":   s.DateStart,
		"date_end":     s.DateEnd,
		"percent_sale": fmt.Sprintf("%.0f%%", s.PercentSale*100), // Display as percentage (e.g., "20%")
		"description":  s.Description,
		"is_active":    s.IsActive,
	}
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

	tryParseDate := func(value interface{}) (string, error) {
		switch v := value.(type) {
		case string:
			parsedDate, err := time.Parse("2006-01-02", v)
			if err != nil {
				return "", err
			}
			return parsedDate.Format("2006-01-02"), nil
		case dbtype.Date:
			return v.String(), nil
		default:
			return "", errors.New("invalid date format")
		}
	}

	//date_start
	if val, ok := data["date_start"]; ok {
		parsedDate, err := tryParseDate(val)
		if err != nil {
			return nil, err
		}
		dateStart = parsedDate
	}

	//date_end
	if val, ok := data["date_end"]; ok {
		parsedDate, err := tryParseDate(val)
		if err != nil {
			return nil, err
		}
		dateEnd = parsedDate
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