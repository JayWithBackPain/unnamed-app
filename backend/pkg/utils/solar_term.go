package utils

import (
	"time"
)

// SolarTermInfo 節氣信息
type SolarTermInfo struct {
	Name      string
	StartTime time.Time
	EndTime   time.Time
	JuNumber  int
	IsYang    bool
}

var solarTerms = []SolarTermInfo{
	{
		Name:      "立春",
		StartTime: time.Date(2024, 2, 4, 4, 27, 0, 0, time.Local),
		EndTime:   time.Date(2024, 2, 19, 0, 13, 0, 0, time.Local),
		JuNumber:  1,
		IsYang:    true,
	},
	{
		Name:      "雨水",
		StartTime: time.Date(2024, 2, 19, 0, 13, 0, 0, time.Local),
		EndTime:   time.Date(2024, 3, 5, 10, 23, 0, 0, time.Local),
		JuNumber:  2,
		IsYang:    true,
	},
	{
		Name:      "驚蟄",
		StartTime: time.Date(2024, 3, 5, 10, 23, 0, 0, time.Local),
		EndTime:   time.Date(2024, 3, 20, 11, 6, 0, 0, time.Local),
		JuNumber:  3,
		IsYang:    true,
	},
	{
		Name:      "春分",
		StartTime: time.Date(2024, 3, 20, 11, 6, 0, 0, time.Local),
		EndTime:   time.Date(2024, 4, 4, 15, 2, 0, 0, time.Local),
		JuNumber:  4,
		IsYang:    true,
	},
	{
		Name:      "清明",
		StartTime: time.Date(2024, 4, 4, 15, 2, 0, 0, time.Local),
		EndTime:   time.Date(2024, 4, 20, 2, 1, 0, 0, time.Local),
		JuNumber:  5,
		IsYang:    true,
	},
	{
		Name:      "穀雨",
		StartTime: time.Date(2024, 4, 20, 2, 1, 0, 0, time.Local),
		EndTime:   time.Date(2024, 5, 5, 8, 10, 0, 0, time.Local),
		JuNumber:  6,
		IsYang:    true,
	},
	{
		Name:      "立夏",
		StartTime: time.Date(2024, 5, 5, 8, 10, 0, 0, time.Local),
		EndTime:   time.Date(2024, 5, 21, 1, 0, 0, 0, time.Local),
		JuNumber:  7,
		IsYang:    true,
	},
	{
		Name:      "小滿",
		StartTime: time.Date(2024, 5, 21, 1, 0, 0, 0, time.Local),
		EndTime:   time.Date(2024, 6, 5, 12, 10, 0, 0, time.Local),
		JuNumber:  8,
		IsYang:    true,
	},
	{
		Name:      "芒種",
		StartTime: time.Date(2024, 6, 5, 12, 10, 0, 0, time.Local),
		EndTime:   time.Date(2024, 6, 21, 4, 51, 0, 0, time.Local),
		JuNumber:  9,
		IsYang:    true,
	},
	{
		Name:      "夏至",
		StartTime: time.Date(2024, 6, 21, 4, 51, 0, 0, time.Local),
		EndTime:   time.Date(2024, 7, 7, 2, 20, 0, 0, time.Local),
		JuNumber:  1,
		IsYang:    false,
	},
	{
		Name:      "小暑",
		StartTime: time.Date(2024, 7, 7, 2, 20, 0, 0, time.Local),
		EndTime:   time.Date(2024, 7, 23, 9, 50, 0, 0, time.Local),
		JuNumber:  2,
		IsYang:    false,
	},
	{
		Name:      "大暑",
		StartTime: time.Date(2024, 7, 23, 9, 50, 0, 0, time.Local),
		EndTime:   time.Date(2024, 8, 7, 16, 20, 0, 0, time.Local),
		JuNumber:  3,
		IsYang:    false,
	},
	{
		Name:      "立秋",
		StartTime: time.Date(2024, 8, 7, 16, 20, 0, 0, time.Local),
		EndTime:   time.Date(2024, 8, 23, 1, 50, 0, 0, time.Local),
		JuNumber:  4,
		IsYang:    false,
	},
	{
		Name:      "處暑",
		StartTime: time.Date(2024, 8, 23, 1, 50, 0, 0, time.Local),
		EndTime:   time.Date(2024, 9, 7, 19, 20, 0, 0, time.Local),
		JuNumber:  5,
		IsYang:    false,
	},
	{
		Name:      "白露",
		StartTime: time.Date(2024, 9, 7, 19, 20, 0, 0, time.Local),
		EndTime:   time.Date(2024, 9, 23, 4, 50, 0, 0, time.Local),
		JuNumber:  6,
		IsYang:    false,
	},
	{
		Name:      "秋分",
		StartTime: time.Date(2024, 9, 23, 4, 50, 0, 0, time.Local),
		EndTime:   time.Date(2024, 10, 8, 10, 50, 0, 0, time.Local),
		JuNumber:  7,
		IsYang:    false,
	},
	{
		Name:      "寒露",
		StartTime: time.Date(2024, 10, 8, 10, 50, 0, 0, time.Local),
		EndTime:   time.Date(2024, 10, 23, 13, 50, 0, 0, time.Local),
		JuNumber:  8,
		IsYang:    false,
	},
	{
		Name:      "霜降",
		StartTime: time.Date(2024, 10, 23, 13, 50, 0, 0, time.Local),
		EndTime:   time.Date(2024, 11, 7, 13, 50, 0, 0, time.Local),
		JuNumber:  9,
		IsYang:    false,
	},
	{
		Name:      "立冬",
		StartTime: time.Date(2024, 11, 7, 13, 50, 0, 0, time.Local),
		EndTime:   time.Date(2024, 11, 22, 11, 20, 0, 0, time.Local),
		JuNumber:  1,
		IsYang:    false,
	},
	{
		Name:      "小雪",
		StartTime: time.Date(2024, 11, 22, 11, 20, 0, 0, time.Local),
		EndTime:   time.Date(2024, 12, 7, 6, 50, 0, 0, time.Local),
		JuNumber:  2,
		IsYang:    false,
	},
	{
		Name:      "大雪",
		StartTime: time.Date(2024, 12, 7, 6, 50, 0, 0, time.Local),
		EndTime:   time.Date(2024, 12, 21, 23, 50, 0, 0, time.Local),
		JuNumber:  3,
		IsYang:    false,
	},
	{
		Name:      "冬至",
		StartTime: time.Date(2024, 12, 21, 23, 50, 0, 0, time.Local),
		EndTime:   time.Date(2025, 1, 5, 17, 50, 0, 0, time.Local),
		JuNumber:  4,
		IsYang:    false,
	},
	{
		Name:      "小寒",
		StartTime: time.Date(2025, 1, 5, 17, 50, 0, 0, time.Local),
		EndTime:   time.Date(2025, 1, 20, 11, 20, 0, 0, time.Local),
		JuNumber:  5,
		IsYang:    false,
	},
	{
		Name:      "大寒",
		StartTime: time.Date(2025, 1, 20, 11, 20, 0, 0, time.Local),
		EndTime:   time.Date(2025, 2, 4, 4, 27, 0, 0, time.Local),
		JuNumber:  6,
		IsYang:    false,
	},
}

// GetSolarTerm 獲取指定日期的節氣信息
func GetSolarTerm(date time.Time) *SolarTermInfo {
	for _, term := range solarTerms {
		if date.After(term.StartTime) && date.Before(term.EndTime) {
			return &term
		}
	}
	// 如果找不到對應的節氣，返回最近的節氣
	// 這裡簡單處理，返回第一個節氣
	return &solarTerms[0]
}
