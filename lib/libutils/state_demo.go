package libutils

import "fmt"

type statusDemo struct { //
	Undefined int `v:"0" d:"未定义"`
	Valid     int `v:"1" d:"生效"`
	Invalid   int `v:"2" d:"无效"`
}

type listItem struct {
	Value int    `json:"value"`
	Desc  string `json:"desc"`
}

func demo() {
	var (
		StatusDemo   statusDemo
		StatusMap    = make(map[int]string, 0)
		StatusRevMap = make(map[string]int, 0)
		List         = make([]listItem, 0)

		//StatusList []SliceItem
	)

	_ = InitStateX(&StatusDemo)
	_ = InitMap(&StatusDemo, StatusMap)
	_ = InitInverseMap(&StatusDemo, StatusRevMap)
	_ = InitSlice(&StatusDemo, &List, &listItem{})
	StatusList, _ := GetStateList(&StatusDemo)
	fmt.Println(StatusList)
}
