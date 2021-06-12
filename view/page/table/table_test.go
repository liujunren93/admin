package table

import (
	"fmt"
	"github.com/liujunren93/admin/view/component/checkbox"
	"github.com/liujunren93/admin/view/component/datePicker"

	"testing"
	"time"
)

func Test_buildSearchForm(t *testing.T) {
	form := buildSearchForm(
		checkbox.NewGroup("do", "queryParam.do", "干什么", 1, true, checkbox.NewCheckbox("学习", 2),
			checkbox.NewCheckbox("打游戏", 1), checkbox.NewCheckbox("玩", 3)),
		datePicker.NewRanger("time", "queryParam.time", "放假时间", []time.Time{time.Now(), time.Now()}, true),
		datePicker.NewRanger("time", "queryParam.time", "放假时间", []time.Time{time.Now(), time.Now()}, true),
		datePicker.NewRanger("time", "queryParam.time", "放假时间", []time.Time{time.Now(), time.Now()}, true))
	fmt.Println(form)

}
