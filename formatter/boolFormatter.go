package formatter

import (
	"strconv"

	beegoContext "github.com/astaxie/beego/context"
)

type boolFormatter struct{}

func (b *boolFormatter) FormatterType() string {
	return ValidTypeBool
}

func (b *boolFormatter) Parse(ctx *beegoContext.Context, rule Rule) interface{} {
	key := rule["name"].(string)
	def, ok := rule["default"]
	str := ctx.Input.Query(key)
	if len(str) == 0 {
		if !ok && rule["required"].(bool) {
			ApiError(ctx, GetParamRequireError(key).Error())
		}
		return def.(bool)
	}
	value, err := strconv.ParseBool(str)
	if err != nil {
		ApiError(ctx, err.Error())
	}
	return value
}
