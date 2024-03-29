package sqlxplus

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

// Logger logger
type Logger interface {
	// Printf 打印
	Printf(ctx context.Context, query string, args ...interface{})
	// Println 打印
	Println(ctx context.Context, args ...interface{})
	// Errorf 错误
	Errorf(ctx context.Context, format string, args ...interface{})
	// Errorln 错误
	Errorln(ctx context.Context, args ...interface{})
}

type sqlString interface {
	String() string
}

func InterfaceToString(src interface{}) string {
	if src == nil {
		return ""
	}
	switch v := src.(type) {
	case string:
		return fmt.Sprintf("'%s'", v)
	case sqlString:
		return fmt.Sprintf("'%s'", v.String())
	case uint8, uint16, uint32, uint64, int, int8, int32, int64, float32, float64:
		return fmt.Sprint(v)
	case bool:
		if v {
			return "true"
		} else {
			return "false"
		}
	case sql.NullBool:
		if v.Valid {
			if v.Bool {
				return "true"
			} else {
				return "false"
			}
		} else {
			return "null"
		}
	case sql.NullByte:
		if v.Valid {
			return fmt.Sprint(v.Byte)
		} else {
			return "null"
		}
	case sql.NullFloat64:
		if v.Valid {
			return fmt.Sprint(v.Float64)
		} else {
			return "null"
		}
	case sql.NullInt16:
		if v.Valid {
			return fmt.Sprint(v.Int16)
		} else {
			return "null"
		}
	case sql.NullInt32:
		if v.Valid {
			return fmt.Sprint(v.Int32)
		} else {
			return "null"
		}
	case sql.NullInt64:
		if v.Valid {
			return fmt.Sprint(v.Int64)
		} else {
			return "null"
		}
	case sql.NullString:
		if v.Valid {
			return v.String
		} else {
			return "null"
		}
	case sql.NullTime:
		if v.Valid {
			return v.Time.Format("2006-01-02 15:04:05")
		} else {
			return "null"
		}
	}
	data, err := json.Marshal(src)
	if err != nil {
		return ""
	}
	return fmt.Sprintf("'%s'", string(data))
}

func StringIndex(str string, s rune) (indexs []int) {
	strRunes := []rune(str)
	strLen := len(strRunes)
	for i := 0; i < strLen; i++ {
		u := strRunes[i]
		if u == rune(s) {
			indexs = append(indexs, i)
		}
	}
	return
}

func StringIndexReplace(str string, indexs []int, args []interface{}) string {
	if len(indexs) != len(args) {
		return str
	}
	strRunes := []rune(str)
	var b strings.Builder
	for i := 0; i < len(strRunes); i++ {
		replace := false
		for j := 0; j < len(indexs); j++ {
			if i == indexs[j] {
				replace = true
				b.WriteString(InterfaceToString(args[j]))
			}
		}
		if !replace {
			b.WriteRune(strRunes[i])
		}
	}
	return b.String()
}

// StdLogger ...
type StdLogger struct {
}

// Printf 打印
func (StdLogger) Printf(ctx context.Context, query string, args ...interface{}) {
	indexs := StringIndex(query, '?')
	query = StringIndexReplace(query, indexs, args)
	log.Printf("[sqlx] %s", query)
}

// Println 打印
func (StdLogger) Println(ctx context.Context, args ...interface{}) {
	nArgs := []interface{}{
		"[sqlx]",
	}
	nArgs = append(nArgs, args...)
	log.Println(nArgs...)
}

// Errorf 错误
func (StdLogger) Errorf(ctx context.Context, format string, args ...interface{}) {
	log.Printf("[sqlx-error] "+format, args...)
}

// Errorln 错误
func (StdLogger) Errorln(ctx context.Context, args ...interface{}) {
	nArgs := []interface{}{
		"[sqlx-error]",
	}
	nArgs = append(nArgs, args...)
	log.Println(nArgs...)
}
