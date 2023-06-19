package IDL

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Format 标准格式
const Format = "2006-01-02 15:04:05"

// SimpleFormat 添加简易格式，比如 Y-m-d
const SimpleFormat = "2006-01-02"

// EmptyTime 空时间，各种数据库，以及下游可能会返回该时间
const EmptyTime = "0000-00-00 00:00:00"

const (
	//秒级别时间戳最大值
	maxTimestampSeconds = 2147483647
)

// Time 时间格式
//
// 时区：
// 		1.当使用Time解析时间字符串:"Y-m-d H:i:s"的时候，都会自动转换为北京时间!
// 		2.统一使用CST时区,以便和上下游统一
//
// json解析:
//       1.在json结构体中使用 Time 会自动进行格式转换，生成对应的时间；比如mysql、rpc、redis的返回值
// 		 2.默认支持：Y-m-d H:i:s格式，更多格式参考 UnmarshalJSON 以及对应的单测
//
// json编码：
//       1.Time的数据在进行marshal的时候会自动调用 Format 格式进行设置
//
// 空时间 EmptyTime:
//       1.golang的time包不支持解析和编码 EmptyTime 格式的数据，当解析的时候会生成 0001-01-01 00:00:00的时间
//         但是我们的业务系统明显不期望有这种转换，所以 Time 包会将空时间数据，继续特殊处理为 EmptyTime
//       2.如果需要判断空时间，使用 IsZero()
type Time struct {
	innerTime time.Time
}

// NowTime 获取当前时间的格式
func NowTime() *Time {
	return &Time{innerTime: time.Now()}
}

func NewTime(t time.Time) *Time {
	return &Time{
		innerTime: t,
	}
}

// NewFromTimestamp 基于时间戳构建
func NewFromTimestamp(timestamp int64) *Time {
	return &Time{
		innerTime: time.Unix(timestamp, 0),
	}
}

// NewTimeFromString 基于字符串格式构建，必须是： Y-m-d H:i:s
func NewTimeFromString(t string) (*Time, error) {
	if t == EmptyTime {
		return &Time{}, nil
	}
	loc, _ := time.LoadLocation("Asia/Shanghai")
	tm, err := time.ParseInLocation(Format, t, loc)
	if err != nil {
		return nil, err
	}
	return &Time{innerTime: tm}, nil
}

// UnmarshalJSON
//支持格式如下：
//   1. Y-m-d H:i:s 格式的时间字符串,将会按照北京时区解析
//   2. 秒级别时间戳,字符串/数字："1650445202" / 1650445202
//   3. 毫秒级时间戳,字符串/数字： "1655976744820" / 1655976744820
//   4. 毫秒级时间戳,字符串： "1655976744820"
//   5. 科学计数法格式,字符串/数字： "1.657886194E9",1.657886194E9
//   6. Y-m-d 格式， 将会按照北京时区解析
//所以在底层做兼容
func (t *Time) UnmarshalJSON(data []byte) (err error) {
	s := strings.Trim(string(data), "\"")
	//时间空值情况，有些时间的默认值是"{}"，会导致解码失败
	if s == "{}" || s == "[]" {
		return nil
	}
	if s == "null" {
		t.innerTime = time.Time{}
		return fmt.Errorf("data is nil")
	}

	//如果是y-m-d格式
	if strings.ContainsAny(s, "-: ") {
		//特殊处理空时间
		if s == EmptyTime {
			t.innerTime = time.Time{}
			return nil
		}
		//按照北京时区解析
		loc, _ := time.LoadLocation("Asia/Shanghai")
		if strings.Contains(s, ":") {
			t.innerTime, err = time.ParseInLocation(Format, s, loc)
		} else {
			t.innerTime, err = time.ParseInLocation(SimpleFormat, s, loc)
		}
		return
	}

	//否则按照时间戳处理
	timestamp, err := strconv.ParseFloat(s, 64)
	if err != nil {
		//如果兜底逻辑处理失败，则不做处理
		return nil
		//return errors.Wrapf(err, "timestamp format data should be valid number")
	}

	//毫秒级时间戳改为秒级
	//bugfix:不能用字符串长度判断是秒级别，还是毫秒级别，不然如果是秒级字符串，但是是科学计数法，就会判断错误。
	//case: "1.657886194E9"
	if timestamp > maxTimestampSeconds {
		timestamp /= 1000
	}

	t.innerTime = time.Unix(int64(timestamp), 0)

	return
}

func (t Time) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, t.String())), nil
}

//转换格式，生成的时间格式为: Y-m-d H:i:s
func (t Time) String() string {
	if t.IsZero() {
		return EmptyTime
	}
	return t.innerTime.Format(Format)
}

// Deprecated: StringByFormat  不建议直接使用本方法，因为无法处理 EmptyTime 的情况，更多信息参考 Time 的注释
func (t Time) StringByFormat(format string) string {
	return t.innerTime.Format(format)
}

// Timestamp 生成timestamp
func (t Time) Timestamp() int64 {
	if t.IsZero() {
		return 0
	}
	return t.innerTime.Unix()
}

// FormatYmd 按照Ymd格式格式化
//withSeparator 是否带有分割线，如果为true，则格式为2022-03-18
//如果为false，不带分割线，比如20220318
func (t Time) FormatYmd(withSeparator bool) string {
	//兼容EmptyTime
	if t.IsZero() {
		if withSeparator {
			return "0000-00-00"
		}
		return "00000000"
	}
	if withSeparator {
		return t.innerTime.Format("2006-01-02")
	}
	return t.innerTime.Format("20060102")
}

// Add 时间增量
func (t Time) Add(second time.Duration) time.Time {
	return t.innerTime.Add(time.Second * second)
}

// Year 返回时间的年份，YYYY格式
func (t Time) Year() int {
	//兼容 EmptyTime
	if t.IsZero() {
		return 0
	}
	return t.innerTime.Year()
}

// GetRawTime 获取原始time.Time类型的值
func (t Time) GetRawTime() time.Time {
	return t.innerTime
}

// IsZero 判断时间是否是0，也就是时间戳的开始
// 如果在json解析的时候，时间格式是 EmptyTime ，那么可能需要判断这种情况
func (t Time) IsZero() bool {
	return t.innerTime.IsZero()
}
