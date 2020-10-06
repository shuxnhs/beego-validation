package formatter

import (
	"strconv"

	beegoContext "github.com/astaxie/beego/context"
)

type intFormatter struct{}

func (i *intFormatter) FormatterType() string {
	return ValidTypeInt
}

func (i *intFormatter) Parse(ctx *beegoContext.Context, rule Rule) interface{} {
	key := rule["name"].(string)
	def, ok := rule["default"]
	str := ctx.Input.Query(key)
	if len(str) == 0 {
		if !ok && rule["required"].(bool) {
			ApiError(ctx, GetParamRequireError(key).Error())
		}
		return def.(int)
	}
	value, err := strconv.Atoi(str)
	if err != nil {
		ApiError(ctx, err.Error())
	}
	return value
}
