package gin_req

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"strconv"
	"strings"
)

func IsAjax(c *gin.Context) bool {
	if strings.EqualFold(c.Request.Header.Get("X-Requested-With"), "XMLHttpRequest") {
		return true
	}

	return false
}

func ParamInt(c *gin.Context, key string) (int, error) {
	v := c.Param(key)
	return strconv.Atoi(v)
}

func ParamStrList(c *gin.Context, key string) ([]string, error) {
	v := c.Param(key)

	return strings.Split(v, ","), nil
}

func ParamIntList(c *gin.Context, key string) ([]int, error) {
	v := c.Param(key)

	strS := strings.Split(v, ",")
	return StringSliceToIntSlice(strS)
}

func RestParams(c *gin.Context) (map[string]interface{}, error) {
	if len(c.Params) == 0 {
		return nil, errors.New("no params")
	}

	var params = make(map[string]interface{})
	for key, values := range c.Request.Form {
		if len(values) > 1 {
			return nil, fmt.Errorf("key value lg 1, key:%v", key)
		}
		params[key] = values[0]
	}

	return params, nil
}

func PureRestParams(c *gin.Context) (map[string]interface{}, error) {
	params, err := RestParams(c)
	if err != nil {
		return params, err
	}

	for k, _ := range params {
		switch k {
		case "page", "psize", "orderBy", "orderDirect", "rest_method_xyz":
			delete(params, k)
		}
	}

	return params, nil
}

func GetCurrentUserId(c *gin.Context, userKey string) int {
	userId, ok := c.Get(userKey)
	if !ok {
		return 0
	}
	return cast.ToInt(userId)
}

func StringSliceToIntSlice(s []string) ([]int, error) {

	intS := make([]int, len(s), len(s))
	var err error
	for idx, v := range s {
		intV, err1 := strconv.Atoi(v)
		if err == nil && err1 != nil {
			err = err1
		}
		intS[idx] = intV
	}

	return intS, err
}

func StringSliceToInt64Slice(s []string) ([]int64, error) {
	intS := make([]int64, len(s), len(s))
	var err error
	for idx, v := range s {
		intS[idx] = cast.ToInt64(v)
	}

	return intS, err
}
