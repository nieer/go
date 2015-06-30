/*
	时间工具包
	时间格式为"2015-06-30"
	by nieer 2015-06-30 18:04:24
*/

package timeFormat

import (
	"fmt"
	"math"
	"strconv"
	"time"
)

var today time.Time

//获取是当周的第几天
func getWeekDay(s time.Time) int {
	w := fmt.Sprintf("%d", s.Weekday())
	b, _ := strconv.Atoi(w)
	return b
}

//获取是当月的第几月
func getYearMonth(s time.Time) int {
	w := fmt.Sprintf("%d", s.Month())
	b, _ := strconv.Atoi(w)
	return b
}

//获取是当月的第几天
func getMonthday(t time.Time) int { return t.Day() }

//初始化当天日期
//t的格式为 "2006-01-02"
func Init(t string) {
	str_time, err := time.Parse("2006-01-02", t)
	if err == nil {
		today = str_time
	} else {
		panic("today format is wrong!")
	}
}

func Today() string {
	return today.Format("2006-01-02")
}

//获取前一天的日期
func Yesterday() string {
	yesterday := today.AddDate(0, 0, -1)
	return yesterday.Format("2006-01-02")
}

//获取当周周一的日期
func ThisWeekStart() string {
	w := getWeekDay(today)
	weekStart := today.AddDate(0, 0, -(w - 1))
	return weekStart.Format("2006-01-02")
}

//获取当周周末的日期
func ThisWeekEnd() string {
	w := getWeekDay(today)
	weekEnd := today.AddDate(0, 0, 7-w)
	return weekEnd.Format("2006-01-02")
}

//获取上周周一的日期
func LastWeekStart() string {
	w := getWeekDay(today)
	weekStart := today.AddDate(0, 0, -(w-1)-7)
	return weekStart.Format("2006-01-02")
}

//获取上周周末的日期
func LastWeekEnd() string {
	w := getWeekDay(today)
	weekEnd := today.AddDate(0, 0, -w)
	return weekEnd.Format("2006-01-02")
}

//获取当月第一天的日期
func ThisMonthStart() string {
	m := getMonthday(today)
	monthStart := today.AddDate(0, 0, -m+1)
	return monthStart.Format("2006-01-02")
}

//获取当月最后一天的日期
func ThisMonthEnd() string {
	nextMonth := today.AddDate(0, 1, 0)
	m := getMonthday(nextMonth)
	monthStart := nextMonth.AddDate(0, 0, -m)
	return monthStart.Format("2006-01-02")
}

//获取上个月第一天的日期
func LastMonthStart() string {
	lastMonth := today.AddDate(0, -1, 0)
	m := getMonthday(lastMonth)
	monthStart := lastMonth.AddDate(0, 0, -m+1)
	return monthStart.Format("2006-01-02")
}

//获取上月最后一天的日期
func LastMonthEnd() string {
	m := getMonthday(today)
	monthEnd := today.AddDate(0, 0, -m)
	return monthEnd.Format("2006-01-02")
}

//获取当季度第一天的日期
func ThisSeasonStart() string {
	m := getYearMonth(today)
	s := math.Ceil(float64(m)/3)*3 - 2
	the_time := time.Date(today.Year(), time.Month(s), 1, 0, 0, 0, 0, time.Local)
	t := the_time.Unix()
	r := time.Unix(t, 0).Format("2006-01-02")
	return r
}

//获取当季度最后一天的日期
func ThisSeasonEnd() string {
	ym := getYearMonth(today)
	s := math.Ceil(float64(ym)/3) * 3
	tmp := time.Date(today.Year(), time.Month(s), 1, 0, 0, 0, 0, time.Local)
	nextMonth := tmp.AddDate(0, 1, 0)
	m := getMonthday(nextMonth)
	monthEnd := nextMonth.AddDate(0, 0, -m)
	return monthEnd.Format("2006-01-02")
}

//获取日同比日期
func DayTb() string {
	dayTb := today.AddDate(0, 0, -7)
	return dayTb.Format("2006-01-02")
}

//获取周同比日期（4周前周末）
func WeekTb() string {
	thisWeekEnd := ThisWeekEnd()
	str_time, _ := time.Parse("2006-01-02", thisWeekEnd)
	WeekTb := str_time.AddDate(0, 0, -7*4)
	return WeekTb.Format("2006-01-02")
}

//获取月同比日期
func MonthTb() string {
	year, month, _ := today.UTC().Date()
	tmp := time.Date(year-1, month, 1, 0, 0, 0, 0, time.Local)
	nextMonth := tmp.AddDate(0, 1, 0)
	m := getMonthday(nextMonth)
	monthEnd := nextMonth.AddDate(0, 0, -m)
	return monthEnd.Format("2006-01-02")
}

//获取季度同比日期
func SeasonTb() string {
	thisSeasonEnd := ThisSeasonEnd()
	str_time, _ := time.Parse("2006-01-02", thisSeasonEnd)
	year, month, _ := str_time.UTC().Date()
	tmp := time.Date(year-1, month, 1, 0, 0, 0, 0, time.Local)
	nextMonth := tmp.AddDate(0, 1, 0)
	m := getMonthday(nextMonth)
	monthEnd := nextMonth.AddDate(0, 0, -m)
	return monthEnd.Format("2006-01-02")
}
