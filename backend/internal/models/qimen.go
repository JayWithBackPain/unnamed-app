package models

import "time"

// QimenChart 奇門盤結構
type QimenChart struct {
	ID       uint      `json:"id" gorm:"primaryKey"`
	UserID   uint      `json:"user_id" gorm:"not null"`
	DateTime time.Time `json:"date_time" gorm:"not null"` // 起盤時間
	IsYang   bool      `json:"is_yang" gorm:"not null"`   // 是否陽遁
	JuNumber int       `json:"ju_number" gorm:"not null"` // 局數
	User     User      `json:"user" gorm:"foreignKey:UserID"`

	// 九宮格數據
	Palace1 Palace `json:"palace_1" gorm:"embedded"` // 坎宮
	Palace2 Palace `json:"palace_2" gorm:"embedded"` // 坤宮
	Palace3 Palace `json:"palace_3" gorm:"embedded"` // 震宮
	Palace4 Palace `json:"palace_4" gorm:"embedded"` // 巽宮
	Palace5 Palace `json:"palace_5" gorm:"embedded"` // 中宮
	Palace6 Palace `json:"palace_6" gorm:"embedded"` // 乾宮
	Palace7 Palace `json:"palace_7" gorm:"embedded"` // 兌宮
	Palace8 Palace `json:"palace_8" gorm:"embedded"` // 艮宮
	Palace9 Palace `json:"palace_9" gorm:"embedded"` // 離宮
}

// Palace 宮位結構
type Palace struct {
	EarthlyBranch string `json:"earthly_branch"` // 地支
	HeavenlyStem  string `json:"heavenly_stem"`  // 天干
	Star          string `json:"star"`           // 九星
	Door          string `json:"door"`           // 八門
	God           string `json:"god"`            // 八神
	HiddenStem    string `json:"hidden_stem"`    // 暗干
}

// SolarTerm 節氣結構
type SolarTerm struct {
	Name      string    `json:"name"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	JuNumber  int       `json:"ju_number"` // 局數
	IsYang    bool      `json:"is_yang"`   // 是否陽遁
}
