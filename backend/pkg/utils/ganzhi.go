package utils

import (
	"time"
)

var (
	heavenlyStems   = []string{"甲", "乙", "丙", "丁", "戊", "己", "庚", "辛", "壬", "癸"}
	earthlyBranches = []string{"子", "丑", "寅", "卯", "辰", "巳", "午", "未", "申", "酉", "戌", "亥"}
	zhiHours        = map[string][]int{
		"子": {23, 0, 1},   // 23:00-01:00
		"丑": {1, 2, 3},    // 01:00-03:00
		"寅": {3, 4, 5},    // 03:00-05:00
		"卯": {5, 6, 7},    // 05:00-07:00
		"辰": {7, 8, 9},    // 07:00-09:00
		"巳": {9, 10, 11},  // 09:00-11:00
		"午": {11, 12, 13}, // 11:00-13:00
		"未": {13, 14, 15}, // 13:00-15:00
		"申": {15, 16, 17}, // 15:00-17:00
		"酉": {17, 18, 19}, // 17:00-19:00
		"戌": {19, 20, 21}, // 19:00-21:00
		"亥": {21, 22, 23}, // 21:00-23:00
	}
)

// GetZhiHour 獲取時辰
func GetZhiHour(hour int) string {
	for zhi, hours := range zhiHours {
		for _, h := range hours {
			if hour == h {
				return zhi
			}
		}
	}
	return "子" // 默認返回子時
}

// CalculateYearGanZhi 計算年干支
func CalculateYearGanZhi(date time.Time) string {
	year := date.Year()
	stemIndex := (year - 4) % 10
	branchIndex := (year - 4) % 12
	return heavenlyStems[stemIndex] + earthlyBranches[branchIndex]
}

// CalculateMonthGanZhi 計算月干支
func CalculateMonthGanZhi(date time.Time) string {
	year := date.Year()
	month := int(date.Month())

	// 計算年干
	yearStemIndex := (year - 4) % 10

	// 月干 = 年干 * 2 + 月數 - 1
	stemIndex := (yearStemIndex*2 + month - 1) % 10

	// 月支 = 月數 + 1
	branchIndex := (month + 1) % 12

	return heavenlyStems[stemIndex] + earthlyBranches[branchIndex]
}

// CalculateDayGanZhi 計算日干支
func CalculateDayGanZhi(date time.Time) string {
	// 計算從 1900 年 1 月 31 日（甲子日）開始的天數
	baseDate := time.Date(1900, 1, 31, 0, 0, 0, 0, time.Local)
	days := int(date.Sub(baseDate).Hours() / 24)

	stemIndex := days % 10
	branchIndex := days % 12

	return heavenlyStems[stemIndex] + earthlyBranches[branchIndex]
}

// CalculateHourGanZhi 計算時干支
func CalculateHourGanZhi(date time.Time) string {
	hour := date.Hour()

	// 計算日干
	baseDate := time.Date(1900, 1, 31, 0, 0, 0, 0, time.Local)
	days := int(date.Sub(baseDate).Hours() / 24)
	dayStemIndex := days % 10

	// 獲取時辰
	zhi := GetZhiHour(hour)
	branchIndex := 0
	for i, b := range earthlyBranches {
		if b == zhi {
			branchIndex = i
			break
		}
	}

	// 時干 = 日干 * 2 + 時支序號
	stemIndex := (dayStemIndex*2 + branchIndex) % 10

	return heavenlyStems[stemIndex] + earthlyBranches[branchIndex]
}
