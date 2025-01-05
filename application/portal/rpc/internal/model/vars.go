package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var ErrNotFound = sqlx.ErrNotFound

type Operator string
type Direction string

const (
	OpEqual          Operator = "="
	OpNotEqual       Operator = "<>"
	OpGreaterThan    Operator = ">"
	OpLessThan       Operator = "<"
	OpGreaterOrEqual Operator = ">="
	OpLessOrEqual    Operator = "<="
	OpLike           Operator = "LIKE"
	OpIn             Operator = "IN"
	// 根据需要添加其他操作符

	DirAsc  Direction = "ASC"
	DirDesc Direction = "DESC"
)

type (
	FilterCondition struct {
		Field    string      // 字段名
		Operator Operator    // 操作符，使用自定义类型
		Value    interface{} // 值
	}
	SortCondition struct {
		Field     string    // 字段名
		Direction Direction // 排序方向，使用自定义类型
	}
)

// 获取 Organizations 结构体的字段映射
// getFieldMap 函数接收任意结构体类型，返回结构体字段名与数据库列名的映射关系
// dataType: 任意结构体类型的实例，例如 Organizations{}
// 返回值: map[string]string，键为结构体字段名，值为数据库列名（由 `db` 标签指定）
/* func getFieldMap(dataType interface{}) map[string]string {
	fieldMap := make(map[string]string)
	t := reflect.TypeOf(dataType)

	// 如果传入的是指针类型，获取其所指向的元素类型
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	// 遍历结构体的所有字段
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		structFieldName := field.Name
		dbTag := field.Tag.Get("db")
		if dbTag == "" {
			// 如果没有设置 `db` 标签，默认使用结构体字段名（注意大小写）
			dbTag = structFieldName
		}
		fieldMap[structFieldName] = dbTag
	}
	return fieldMap
}
*/
// 操作符校验函数
func isValidOperator(op Operator) bool {
	switch op {
	case OpEqual, OpNotEqual, OpGreaterThan, OpLessThan, OpGreaterOrEqual, OpLessOrEqual, OpLike, OpIn:
		return true
	default:
		return false
	}
}

// 排序方向校验函数
func isValidDirection(dir Direction) bool {
	switch dir {
	case DirAsc, DirDesc:
		return true
	default:
		return false
	}
}

// 辅助函数：检查值是否为零值
func isZeroValue(value any) bool {
	switch v := value.(type) {
	case string:
		return v == ""
	case int, int32, int64, uint, uint32, uint64:
		return v == 0
	default:
		return value == nil
	}
}
