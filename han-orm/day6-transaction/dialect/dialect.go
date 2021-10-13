package dialect

import "reflect"

var dialectMap = map[string]Dialect{}

//数据库操作的基础接口 mysql sqlite3都需要继承这个接口
type Dialect interface {
	DataTypeOf(typ reflect.Value) string  //用于将 Go 语言的类型转换为该数据库的数据类型
	TableExistSQL(tableName string) (string, []interface{}) // 返回某个表名是否存的sql语句，参数是表名（Table）。
}

func RegisterDialect(name string, dialect Dialect) {
	dialectMap[name] = dialect
}

func GetDialect(name string) (dialect Dialect, ok bool) {
	dialect, ok = dialectMap[name]
	return
}