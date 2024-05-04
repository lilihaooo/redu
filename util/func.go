package util

import (
	"redu/global"
	"time"
)

func GetOffset(currentPage, pageSize int) int {
	offset := (currentPage - 1) * pageSize
	if offset < 0 {
		return 0
	}
	return offset
}

// TimeParse layout "2006-01-02 15:04:05"; "2006-01-02";
func TimeParse(timeStr string, layout string) (t time.Time) {
	// 解析时间
	t, err := time.Parse(layout, timeStr)
	if err != nil {
		global.Logrus.Error(err)
		return time.Time{}
	}
	return t
}

// TimeToString layout "2006-01-02 15:04:05"; "2006-01-02"
func TimeToString(t time.Time, layout string) string {
	return t.Format(layout)
}
