package libutils

import (
	"github.com/spf13/cast"
	"strconv"
	"strings"
)

const maxInt = int(^uint(0) >> 1)

// 转换逗号分割的整数字符串为整数数组
func SplitToIntList(strCommas string, sep string) []int {
	var idList []int
	for _, i := range strings.Split(strCommas, sep) {
		i = strings.Trim(i, " \t")
		if i == "" {
			continue
		}
		id, err := strconv.Atoi(i)
		if err != nil {
			continue
		}
		idList = append(idList, id)
	}
	return idList
}

func IntSliceJoin(elems []int, sep string) string {
	strSlice := make([]string, len(elems), len(elems))

	for i, v := range elems {
		strSlice[i] = cast.ToString(v)
	}

	return strings.Join(strSlice, sep)
}
