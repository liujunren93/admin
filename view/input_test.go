package view

import (
	"fmt"
	"github.com/liujunren93/admin/view/component/checkbox"
	"github.com/liujunren93/admin/view/component/datePicker"
	"testing"
	"time"
)

func TestInput_checkbox(t *testing.T) {

	//fmt.Println(checkbox.GetGroup("test","请选择test",1,true,
	//	checkbox.GetCheckbox("aaa",1),
	//	checkbox.GetCheckbox("bbb",2)))
	//tt := []time.Time{time.Now(),time.Now()}
	//getRange := datePicker.GetRange("time", "请选择经营时间", tt, true)
	//datePicker.GetPicker("time","请选择经营时间",time.Now(),true,"")


	html := checkbox.NewGroup("do", "", "干什么", 1, true, checkbox.NewCheckbox("学习", 2),
		checkbox.NewCheckbox("打游戏", 1), checkbox.NewCheckbox("玩", 3)).Html()
	fmt.Println(html)
}

func TestInput_datepicker(t *testing.T) {
	fmt.Println(	datePicker.NewSimple("time","query.time","放假时间",time.Now(),true).Html())
	fmt.Println(	datePicker.NewRanger("time","","放假时间",[]time.Time{time.Now(),time.Now()},true).Html())
}

