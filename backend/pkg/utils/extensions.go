package utils

import (
	"encoding/json"
	"strconv"
)

func StructToMap(input interface{}) (map[string]interface{}, error) {
	// Chuyển struct thành JSON
	jsonData, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	// Giải mã JSON thành map[string]interface{}
	var result map[string]interface{}
	if err := json.Unmarshal(jsonData, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func GetString(data map[string]interface{}, key string) string {
	if val, ok := data[key]; ok {
		if str, ok := val.(string); ok {
			return str
		}
	}
	return ""
}

func GetStringArray(data map[string]interface{}, key string) []string {
	if val, ok := data[key]; ok {
		if array, ok := val.([]interface{}); ok {
			strArray := make([]string, len(array))
			for i, v := range array {
				if str, ok := v.(string); ok {
					strArray[i] = str
				}
			}
			return strArray
		}
	}
	return []string{}
}

func GetBool(data map[string]interface{}, key string) bool {
	if val, ok := data[key]; ok {
		if b, ok := val.(bool); ok {
			return b
		}
	}
	return false
}

func GetFloat64(data map[string]interface{}, key string) float64 {
	if val, ok := data[key]; ok {
		if f, ok := val.(float64); ok {
			return f
		} else if i, ok := val.(int); ok { // Trường hợp dữ liệu float64 lưu dưới dạng int
			return float64(i)
		} else if s, ok := val.(string); ok { // Trường hợp dữ liệu float64 lưu dưới dạng string
			if parsedFloat, err := strconv.ParseFloat(s, 64); err == nil {
				return parsedFloat
			}
		}
	}
	return 0.0
}

func GetInt(data map[string]interface{}, key string) int {
	if val, ok := data[key]; ok {
		if i, ok := val.(int); ok {
			return i
		} else if f, ok := val.(float64); ok { // Trường hợp dữ liệu int lưu dưới dạng float64
			return int(f)
		} else if s, ok := val.(string); ok { // Trường hợp dữ liệu int lưu dưới dạng string
			if parsedInt, err := strconv.Atoi(s); err == nil {
				return parsedInt
			}
		}
	}
	return 0
}

func GetStringSlice(data map[string]interface{}, key string) []string {
	if val, ok := data[key]; ok {
		if strSlice, ok := val.([]string); ok {
			return strSlice
		}
		if interfaceSlice, ok := val.([]interface{}); ok {
			var result []string
			for _, v := range interfaceSlice {
				if str, ok := v.(string); ok {
					result = append(result, str)
				}
			}
			return result
		}
	}
	return []string{}
}

