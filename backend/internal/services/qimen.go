package services

import (
	"time"

	"github.com/qimen/backend/internal/models"
	"github.com/qimen/backend/pkg/utils"
)

// QimenService 奇門遁甲服務
type QimenService struct{}

// NewQimenService 創建奇門遁甲服務實例
func NewQimenService() *QimenService {
	return &QimenService{}
}

// CalculateChart 計算奇門盤
func (s *QimenService) CalculateChart(userID uint, dateTime time.Time) (*models.QimenChart, error) {
	// 1. 確定節氣
	solarTerm := utils.GetSolarTerm(dateTime)

	// 2. 確定陰陽遁
	isYang := solarTerm.IsYang

	// 3. 確定局數
	juNumber := solarTerm.JuNumber

	// 4. 計算干支
	yearGanZhi := utils.CalculateYearGanZhi(dateTime)
	monthGanZhi := utils.CalculateMonthGanZhi(dateTime)
	dayGanZhi := utils.CalculateDayGanZhi(dateTime)
	hourGanZhi := utils.CalculateHourGanZhi(dateTime)

	// 5. 排布九宮
	chart := &models.QimenChart{
		UserID:   userID,
		DateTime: dateTime,
		IsYang:   isYang,
		JuNumber: juNumber,
	}

	// 6. 排布三奇六儀
	s.arrangeThreeStrangeSixRituals(chart)

	// 7. 排布九星
	s.arrangeNineStars(chart)

	// 8. 排布八門
	s.arrangeEightDoors(chart)

	// 9. 排布八神
	s.arrangeEightGods(chart)

	// 10. 設置干支信息
	chart.Palace1.EarthlyBranch = yearGanZhi[1:2]
	chart.Palace1.HeavenlyStem = yearGanZhi[0:1]
	chart.Palace2.EarthlyBranch = monthGanZhi[1:2]
	chart.Palace2.HeavenlyStem = monthGanZhi[0:1]
	chart.Palace3.EarthlyBranch = dayGanZhi[1:2]
	chart.Palace3.HeavenlyStem = dayGanZhi[0:1]
	chart.Palace4.EarthlyBranch = hourGanZhi[1:2]
	chart.Palace4.HeavenlyStem = hourGanZhi[0:1]

	return chart, nil
}

// arrangeThreeStrangeSixRituals 排布三奇六儀
func (s *QimenService) arrangeThreeStrangeSixRituals(chart *models.QimenChart) {
	// 三奇六儀的順序
	threeStrangeSixRituals := []string{"戊", "己", "庚", "辛", "壬", "癸", "丁", "丙", "乙"}

	// 根據局數確定起始位置
	startIndex := (chart.JuNumber - 1) * 3

	// 根據陰陽遁確定排布方向
	if chart.IsYang {
		// 陽遁順時針排布
		for i := 0; i < 9; i++ {
			index := (startIndex + i) % 9
			switch i {
			case 0:
				chart.Palace1.HiddenStem = threeStrangeSixRituals[index]
			case 1:
				chart.Palace2.HiddenStem = threeStrangeSixRituals[index]
			case 2:
				chart.Palace3.HiddenStem = threeStrangeSixRituals[index]
			case 3:
				chart.Palace4.HiddenStem = threeStrangeSixRituals[index]
			case 4:
				chart.Palace5.HiddenStem = threeStrangeSixRituals[index]
			case 5:
				chart.Palace6.HiddenStem = threeStrangeSixRituals[index]
			case 6:
				chart.Palace7.HiddenStem = threeStrangeSixRituals[index]
			case 7:
				chart.Palace8.HiddenStem = threeStrangeSixRituals[index]
			case 8:
				chart.Palace9.HiddenStem = threeStrangeSixRituals[index]
			}
		}
	} else {
		// 陰遁逆時針排布
		for i := 0; i < 9; i++ {
			index := (startIndex - i + 9) % 9
			switch i {
			case 0:
				chart.Palace1.HiddenStem = threeStrangeSixRituals[index]
			case 1:
				chart.Palace2.HiddenStem = threeStrangeSixRituals[index]
			case 2:
				chart.Palace3.HiddenStem = threeStrangeSixRituals[index]
			case 3:
				chart.Palace4.HiddenStem = threeStrangeSixRituals[index]
			case 4:
				chart.Palace5.HiddenStem = threeStrangeSixRituals[index]
			case 5:
				chart.Palace6.HiddenStem = threeStrangeSixRituals[index]
			case 6:
				chart.Palace7.HiddenStem = threeStrangeSixRituals[index]
			case 7:
				chart.Palace8.HiddenStem = threeStrangeSixRituals[index]
			case 8:
				chart.Palace9.HiddenStem = threeStrangeSixRituals[index]
			}
		}
	}
}

// arrangeNineStars 排布九星
func (s *QimenService) arrangeNineStars(chart *models.QimenChart) {
	// 九星的順序
	nineStars := []string{"天蓬", "天芮", "天沖", "天輔", "天禽", "天心", "天柱", "天任", "天英"}

	// 根據局數確定起始位置
	startIndex := (chart.JuNumber - 1) * 3

	// 根據陰陽遁確定排布方向
	if chart.IsYang {
		// 陽遁順時針排布
		for i := 0; i < 9; i++ {
			index := (startIndex + i) % 9
			switch i {
			case 0:
				chart.Palace1.Star = nineStars[index]
			case 1:
				chart.Palace2.Star = nineStars[index]
			case 2:
				chart.Palace3.Star = nineStars[index]
			case 3:
				chart.Palace4.Star = nineStars[index]
			case 4:
				chart.Palace5.Star = nineStars[index]
			case 5:
				chart.Palace6.Star = nineStars[index]
			case 6:
				chart.Palace7.Star = nineStars[index]
			case 7:
				chart.Palace8.Star = nineStars[index]
			case 8:
				chart.Palace9.Star = nineStars[index]
			}
		}
	} else {
		// 陰遁逆時針排布
		for i := 0; i < 9; i++ {
			index := (startIndex - i + 9) % 9
			switch i {
			case 0:
				chart.Palace1.Star = nineStars[index]
			case 1:
				chart.Palace2.Star = nineStars[index]
			case 2:
				chart.Palace3.Star = nineStars[index]
			case 3:
				chart.Palace4.Star = nineStars[index]
			case 4:
				chart.Palace5.Star = nineStars[index]
			case 5:
				chart.Palace6.Star = nineStars[index]
			case 6:
				chart.Palace7.Star = nineStars[index]
			case 7:
				chart.Palace8.Star = nineStars[index]
			case 8:
				chart.Palace9.Star = nineStars[index]
			}
		}
	}
}

// arrangeEightDoors 排布八門
func (s *QimenService) arrangeEightDoors(chart *models.QimenChart) {
	// 八門的順序
	eightDoors := []string{"休", "生", "傷", "杜", "景", "死", "驚", "開"}

	// 根據局數確定起始位置
	startIndex := (chart.JuNumber - 1) * 3

	// 根據陰陽遁確定排布方向
	if chart.IsYang {
		// 陽遁順時針排布
		for i := 0; i < 8; i++ {
			index := (startIndex + i) % 8
			switch i {
			case 0:
				chart.Palace1.Door = eightDoors[index]
			case 1:
				chart.Palace2.Door = eightDoors[index]
			case 2:
				chart.Palace3.Door = eightDoors[index]
			case 3:
				chart.Palace4.Door = eightDoors[index]
			case 4:
				chart.Palace6.Door = eightDoors[index]
			case 5:
				chart.Palace7.Door = eightDoors[index]
			case 6:
				chart.Palace8.Door = eightDoors[index]
			case 7:
				chart.Palace9.Door = eightDoors[index]
			}
		}
	} else {
		// 陰遁逆時針排布
		for i := 0; i < 8; i++ {
			index := (startIndex - i + 8) % 8
			switch i {
			case 0:
				chart.Palace1.Door = eightDoors[index]
			case 1:
				chart.Palace2.Door = eightDoors[index]
			case 2:
				chart.Palace3.Door = eightDoors[index]
			case 3:
				chart.Palace4.Door = eightDoors[index]
			case 4:
				chart.Palace6.Door = eightDoors[index]
			case 5:
				chart.Palace7.Door = eightDoors[index]
			case 6:
				chart.Palace8.Door = eightDoors[index]
			case 7:
				chart.Palace9.Door = eightDoors[index]
			}
		}
	}
}

// arrangeEightGods 排布八神
func (s *QimenService) arrangeEightGods(chart *models.QimenChart) {
	// 八神的順序
	eightGods := []string{"值符", "騰蛇", "太陰", "六合", "白虎", "玄武", "九地", "九天"}

	// 根據局數確定起始位置
	startIndex := (chart.JuNumber - 1) * 3

	// 根據陰陽遁確定排布方向
	if chart.IsYang {
		// 陽遁順時針排布
		for i := 0; i < 8; i++ {
			index := (startIndex + i) % 8
			switch i {
			case 0:
				chart.Palace1.God = eightGods[index]
			case 1:
				chart.Palace2.God = eightGods[index]
			case 2:
				chart.Palace3.God = eightGods[index]
			case 3:
				chart.Palace4.God = eightGods[index]
			case 4:
				chart.Palace6.God = eightGods[index]
			case 5:
				chart.Palace7.God = eightGods[index]
			case 6:
				chart.Palace8.God = eightGods[index]
			case 7:
				chart.Palace9.God = eightGods[index]
			}
		}
	} else {
		// 陰遁逆時針排布
		for i := 0; i < 8; i++ {
			index := (startIndex - i + 8) % 8
			switch i {
			case 0:
				chart.Palace1.God = eightGods[index]
			case 1:
				chart.Palace2.God = eightGods[index]
			case 2:
				chart.Palace3.God = eightGods[index]
			case 3:
				chart.Palace4.God = eightGods[index]
			case 4:
				chart.Palace6.God = eightGods[index]
			case 5:
				chart.Palace7.God = eightGods[index]
			case 6:
				chart.Palace8.God = eightGods[index]
			case 7:
				chart.Palace9.God = eightGods[index]
			}
		}
	}
}
