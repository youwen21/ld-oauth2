package model

import (
	"database/sql/driver"
	"fmt"
	"time"
)

const (
	//timeFormat = "`2006-01-02 15:04:05`"
	timeFormat = "2006-01-02 15:04:05"
	dateFormat = "2006-01-02"
)

// DateTime 自定义日期时间格式
type DateTime time.Time

// Date
type Date time.Time

// UnmarshalJSON json转struct
func (t *DateTime) UnmarshalJSON(data []byte) error {
	// Ignore null, like in the main JSON package.
	stringData := string(data)
	if len(stringData) == 0 || stringData == "null" || stringData == "0000-00-00 00:00:00" {
		*t = DateTime(time.Time{})
		return nil
	}

	// Fractional seconds are handled implicitly by Parse.
	time, err := time.Parse(`"`+timeFormat+`"`, stringData)
	*t = DateTime(time)

	return err
}

// MarshalJSON 转json 接口实现
func (t DateTime) MarshalJSON() ([]byte, error) {
	if time.Time(t).IsZero() {
		// return []byte("\"\""), nil
		return []byte("\"0000-00-00 00:00:00\""), nil
	}

	//do your serializing here
	stamp := fmt.Sprintf("\"%s\"", time.Time(t).Format(timeFormat))
	return []byte(stamp), nil
}

// String 转string 接口实现
func (t DateTime) String() string {
	dt := time.Time(t)
	if dt.IsZero() {
		return "\"0000-00-00 00:00:00\""
	}

	if y := dt.Year(); y < 0 || y >= 10000 {
		// RFC 3339 is clear that years are 4 digits exactly.
		// See golang.org/issue/4556#c15 for more discussion.
		return "\"0000-00-00 00:00:00\""
	}

	stamp := fmt.Sprintf("\"%s\"", time.Time(t).Format(timeFormat))
	return stamp
}

// https://gorm.io/docs/data_types.html
//Scan gorm receive
//func (t *DateTime) Scan(value interface{}) error {
//	byt, ok := value.([]byte)
//	if !ok {
//		*t = DateTime(time.Time{})
//		return errors.New(fmt.Sprint("convert failed, value:", value))
//	}
//
//	stringTime := string(byt)
//
//	if stringTime == "0000-00-00 00:00:00" {
//		*t = DateTime(time.Time{})
//		return nil
//	}
//
//	ttime, err := time.ParseInLocation(timeFormat, stringTime, time.UTC)
//	// ttime, err := time.ParseInLocation(`"`+timeFormat+`"`, stringTime, time.UTC)
//	*t = DateTime(ttime)
//	return err
//}

// Value gorm save
// Value return json value, implement driver.Valuer interface
func (t DateTime) Value() (driver.Value, error) {
	ttime := time.Time(t)
	if ttime.IsZero() {
		return nil, nil
	}

	//do your serializing here
	stamp := fmt.Sprintf("%s", ttime.Format(timeFormat))
	return stamp, nil
}

// UnmarshalJSON json转struct
func (t *Date) UnmarshalJSON(data []byte) error {
	// Ignore null, like in the main JSON package.
	stringData := string(data)
	if len(stringData) == 0 || stringData == "null" || stringData == "0000-00-00" {
		*t = Date(time.Time{})
		return nil
	}

	// Fractional seconds are handled implicitly by Parse.
	time, err := time.Parse(`"`+dateFormat+`"`, stringData)
	*t = Date(time)

	return err
}

// MarshalJSON 转json 接口实现
func (t Date) MarshalJSON() ([]byte, error) {
	if time.Time(t).IsZero() {
		// return []byte("\"\""), nil
		return []byte("\"0000-00-00\""), nil
	}

	//do your serializing here
	stamp := fmt.Sprintf("\"%s\"", time.Time(t).Format(dateFormat))
	return []byte(stamp), nil
}

// String 转string 接口实现
func (t Date) String() string {
	dt := time.Time(t)
	if dt.IsZero() {
		return "\"0000-00-00\""
	}

	if y := dt.Year(); y < 0 || y >= 10000 {
		// RFC 3339 is clear that years are 4 digits exactly.
		// See golang.org/issue/4556#c15 for more discussion.
		return "\"0000-00-00\""
	}

	stamp := fmt.Sprintf("\"%s\"", time.Time(t).Format(dateFormat))
	return stamp
}

// https://gorm.io/docs/data_types.html
//Scan gorm receive
//func (t *Date) Scan(value interface{}) error {
//	byt, ok := value.([]byte)
//	if !ok {
//		*t = Date(time.Time{})
//		return errors.New(fmt.Sprint("convert failed, value:", value))
//	}
//
//	stringTime := string(byt)
//
//	if stringTime == "0000-00-00" {
//		*t = Date(time.Time{})
//		return nil
//	}
//
//	ttime, err := time.ParseInLocation(dateFormat, stringTime, time.UTC)
//	// ttime, err := time.ParseInLocation(`"`+timeFormat+`"`, stringTime, time.UTC)
//	*t = Date(ttime)
//	return err
//}

// Value gorm save
// Value return json value, implement driver.Valuer interface
func (t Date) Value() (driver.Value, error) {
	ttime := time.Time(t)
	if ttime.IsZero() {
		return "", nil
	}

	//do your serializing here
	stamp := fmt.Sprintf("%s", ttime.Format(dateFormat))
	return stamp, nil
}
