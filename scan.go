package sqlxplus

import "github.com/jmoiron/sqlx"

// MapScan 解决sqlx字节数组问题
func MapScan(r sqlx.ColScanner, dest map[string]interface{}) (err error) {
	if err = sqlx.MapScan(r, dest); err != nil {
		return
	}
	for k, v := range dest {
		if _, ok := v.([]byte); ok {
			dest[k] = string(v.([]byte))
		}
	}
	return
}
