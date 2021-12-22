package helper

import (
	"errors"
	"time"
)

var TimeZone, _ = time.LoadLocation("Asia/Shanghai")

var (
	InitMonth          string = "2006-01"
	InitDate           string = "2006-01-02"
	InitDateTime       string = "2006-01-02 15:04:05"
	InitDateMinute     string = "2006-01-02 15:04"
	InitDateHour       string = "2006-01-02 15"
	InitMonthDayDate   string = "01月02日"
	InitSmallYearMonth string = "0601"
)

// FormatTime 格式化时间
func FormatTime(time time.Time) string {
	return time.Format("2006-01-02 15:04:05")
}

// UnixTime 将时间转化为毫秒数
func UnixTime(t time.Time) int64 {
	return t.UnixNano() / 1000000
}

func GetDateWithWeekdayBetween(format, from, to string) (map[string]int, error) {
	if from > to {
		return nil, errors.New("from time greater than to time")
	}
	fromTime, err := time.ParseInLocation(format, from, TimeZone)
	if err != nil {
		return nil, err
	}

	toTime, err := time.ParseInLocation(format, to, TimeZone)
	if err != nil {
		return nil, err
	}

	res := make(map[string]int)
	for begin := fromTime; toTime.Sub(begin).Nanoseconds() >= 0; begin = begin.AddDate(0, 0, 1) {
		res[begin.Format(format)] = int(begin.Weekday())
	}
	return res, nil
}

// 验证时间格式是否符合规则
func CheckTime(inputTime string, format string) bool {
	unixtime, err := time.Parse(format, inputTime)
	if err != nil {
		return false
	}
	formatTime := unixtime.Format(format)
	if formatTime != inputTime {
		return false
	}
	return true
}

/**
获取星期
*/
func GetZHWeekday(date string) string {

	var res string
	var weekday string

	dates, _ := time.Parse(InitDate, date)
	weekday = dates.Weekday().String()

	switch weekday {
	case "Monday":
		res = "周一"
		break
	case "Tuesday":
		res = "周二"
		break
	case "Wednesday":
		res = "周三"
		break
	case "Thursday":
		res = "周四"
		break
	case "Friday":
		res = "周五"
		break
	case "Saturday":
		res = "周六"
		break
	default:
		res = "周日"
		break
	}

	return res
}

//获取相差时间
func GetHourDiffer(startTime, endTime, layout string) float32 {

	var hour float32
	t1, err := time.ParseInLocation(layout, startTime, TimeZone)
	t2, err := time.ParseInLocation(layout, endTime, TimeZone)

	if err == nil && t1.Before(t2) {
		diff := t2.Unix() - t1.Unix() //
		hour = float32(diff) / 3600
		return hour
	} else {
		return hour
	}
}

//获取传入的时间所在月份的第一天，即某月第一天的0点。如传入time.Now(), 返回当前月份的第一天0点时间。
func StartOfMonth(d time.Time) time.Time {
	d = d.AddDate(0, 0, -d.Day()+1)
	return GetZeroTime(d)
}

//获取传入的时间所在月份的最后一天，即某月最后一天的0点。如传入time.Now(), 返回当前月份的最后一天0点时间。
func EndOfMonth(d time.Time) time.Time {
	return StartOfMonth(d).AddDate(0, 1, 0).Add(-1)
}

//获取某一天的0点时间
func GetZeroTime(d time.Time) time.Time {
	return time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, d.Location())
}

// 时间转换
func TimeParse(layout, value string) time.Time {
	timeP, _ := time.ParseInLocation(layout, value, TimeZone)
	return timeP
}
