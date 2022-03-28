package commons

import (
	"fmt"
	"iris-demo-new/config"
	"iris-demo-new/slog"
	"strconv"
	"time"
)

const TIME_LAYOUT = "2006-01-02 15:04:05"
const TIME_DAY_LAYOUT = "02"

func SetLocation() *time.Location {
	local, err := time.LoadLocation(config.Setting.App.Timezone)
	if err == nil {
		return local
	}
	slog.Errorf(" timestamp err:%s", err.Error())
	return time.Local
}

// time to string time
func TimeToString(t *time.Time) string {
	return t.In(SetLocation()).Format(TIME_LAYOUT)
}

// get time corresponding day
func GetDay(t *time.Time) string {
	return t.In(SetLocation()).Format(TIME_DAY_LAYOUT)
}

// get current month's the last day
func GetCurrentMonthsLastDay() string {
	nowTime := time.Now()
	currentYear,currentMonth,_ := nowTime.Date()
	startTime := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, SetLocation())
	endTime := startTime.AddDate(0, 1, 0).Add(-time.Second)
	lastDay := GetDay(&endTime)
	return lastDay
}

func StringToTime(timeString string) time.Time {
	t, err := time.ParseInLocation(TIME_LAYOUT, timeString, SetLocation())
	if err != nil {
		slog.Error(err.Error())
	}
	return t
}

// if float64 number's decimal bit more than 0, then ensure int part can increase 1
func IncreaseFiveToLastIntBit(number float64) float64 {
	if number <= 0 {
		return 0
	}
	res := number
	// float64 format to no decimal string
	noDecimalStr := fmt.Sprintf("%.f", number * 100)
	fmt.Println(noDecimalStr)
	// get the last number
	lastNumber,_ := strconv.Atoi(noDecimalStr[len(noDecimalStr) - 1:])
	flag := false
	lastSecondNumber := 0
	// get the last second number
	if len(noDecimalStr) >= 2 {
		lastSecondNumber,_ = strconv.Atoi(noDecimalStr[len(noDecimalStr) - 2:len(noDecimalStr) - 1])
		if lastSecondNumber > 0 && lastSecondNumber <= 5 {
			flag = true
		}
	}
	if lastSecondNumber == 0 && lastNumber > 0 {
		flag = true
	}
	if flag {
		floatNumber,_ := strconv.ParseFloat(noDecimalStr,64)
		res = (floatNumber + float64(50)) / float64(100)
	}

	return res
}

// if float64 number's decimal bit more than 0, then ensure int part can increase 1
func IncreaseFiveToLastIntBitForFourDecimal(number float64) float64 {
	if number == 0 {
		return 0
	}
	res := number
	// float64 format to no decimal string
	noDecimalStr := fmt.Sprintf("%.f", number * 10000)
	fmt.Println(noDecimalStr)
	// get the last number
	lastNumber,_ := strconv.Atoi(noDecimalStr[len(noDecimalStr) - 1:])
	flag := false
	var lastSecondNumber,lastThirdNumber,lastFourNumber int
	// get the last four number
	if len(noDecimalStr) >= 4 {
		lastFourNumber,_ = strconv.Atoi(noDecimalStr[len(noDecimalStr) - 4:len(noDecimalStr) - 3])
		if lastFourNumber > 0 && lastFourNumber <= 5 {
			flag = true
		}
	}
	// get the last third number
	if len(noDecimalStr) >= 3 {
		lastThirdNumber,_ = strconv.Atoi(noDecimalStr[len(noDecimalStr) - 3:len(noDecimalStr) - 2])
	}
	// get the last second number
	if len(noDecimalStr) >= 2 {
		lastSecondNumber,_ = strconv.Atoi(noDecimalStr[len(noDecimalStr) - 2:len(noDecimalStr) - 1])
	}
	if lastFourNumber == 0 && (lastNumber > 0 || lastSecondNumber > 0 || lastThirdNumber > 0) {
		flag = true
	}
	if flag {
		floatNumber,_ := strconv.ParseFloat(noDecimalStr,64)
		res = (floatNumber + float64(5000)) / float64(10000)
	}

	return res
}

// append method, will copy a new slice2 params first, so will not change slice1 param's value.
func AppendSlice(slice1 []int, number int) []int {
	slice2 := append(slice1, number)
	return slice2
}

func ChangeSlice(slice1 []int, number int) []int {
	slice1[0] = number
	fmt.Printf("%p \n", &slice1)
	return slice1
}

func ChangePointer(pointer1 *struct{
	Number int
}, number int) {
	pointer1.Number = number
}

func TestSlice(sl []string){
	fmt.Printf("%v, %p, %p\n",sl, &sl,sl)   // [a b c], 0xc0004fc220, 0xc0007242d0
	sl[0] = "aa"
	fmt.Printf("%v, %p, %p\n",sl, &sl,sl)   // [aa b c], 0xc0004fc220, 0xc0007242d0
	sl = append(sl, "d")
	fmt.Printf("%v, %p, %p\n",sl, &sl,sl)   // [aa b c d], 0xc0004fc220, 0xc00041a180
	s2 := append(sl, "d")
	fmt.Printf("%v, %p, %p\n",s2, &s2,s2)   // [aa b c d d], 0xc0004fc340, 0xc00041a180
}
