package formatter

import (
	"strconv"

	beegoContext "github.com/astaxie/beego/context"
)

type floatFormatter struct{}

func (f *floatFormatter) FormatterType() string {
	return ValidTypeFloat
}

func (f *floatFormatter) Parse(ctx *beegoContext.Context, rule Rule) interface{} {
	key := rule["name"].(string)
	def, ok := rule["default"]
	str := ctx.Input.Query(key)
	if len(str) == 0 {
		if !ok && rule["required"].(bool) {
			ApiError(ctx, GetParamRequireError(key).Error())
		}
		return def.(float64)
	}
	value, err := strconv.ParseFloat(str, 64)
	if err != nil {
		ApiError(ctx, err.Error())
	}
	return value
}
