# sqlxplus
golang sqlx wrap logger

# Usage
```bash
go get github.com/nilorg/sqlxplus
```
# Import
```bash
import "github.com/nilorg/sqlxplus"
```

# Example
```go
import (
    "github.com/jmoiron/sqlx"
    "github.com/nilorg/sqlxplus"
)

func main() {
	driverName := "mysql"
	dataSourceName := "xxx"
	xdb, err := sqlx.Connect(driverName, dataSourceName)
	if err != nil {
		log.Fatalln(err)
	}
	xdbpuls := &sqlxplus.DB{sqlxDB: xdb, log: &StdLogger{}}
    var result map[string]interface{}
    err = sqlx.GetContext(ctx, xdbpuls.DB(), &result, "select * from user where id = ?", 1)
    if err != nil {
		log.Fatalln(err)
	}
}



```