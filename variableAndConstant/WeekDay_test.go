package ist

import "testing"

const (
	// 定义 常量方法
	MonDay = 1 + iota
	TuesDay
	WednesDay
)
func TestWeekDay(t *testing.T)  {
	t.Log(MonDay)
	t.Log(TuesDay)
	t.Log(WednesDay)
}