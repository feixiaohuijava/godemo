package pkg

// FilterStruct FilterStuct filter需要的基础结构体
//field             等于                  exact
//field__contains   包含                  contains
//field__gte        大于或等于             gte
//field__gt         大于                   gt
//field__lte        小于或等于             lte
//field__lt         小于                   lt
//field__in         在数组中存在            in
//field__between    介于                   between
type FilterStruct struct {
	Model  interface{}         `json:"model"`
	Fields map[string][]string `json:"fields"`
	Order  []string            `json:"order"`
}
