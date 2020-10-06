package formatter

import (
	beegoContext "github.com/astaxie/beego/context"
)

type stringFormatter struct{}

func (s *stringFormatter) FormatterType() string {
	return ValidTypeString
}

func (s *stringFormatter) Parse(ctx *beegoContext.Context, rule Rule) interface{} {
	key := rule["name"].(string)
	def, ok := rule["default"]
	str := ctx.Input.Query(key)
	if len(str) == 0 {
		if !ok && rule["required"].(bool) {
			ApiError(ctx, GetParamRequireError(key).Error())
		} else if ok {
			return def.(string)
		}
		return ""
	}
	return str
}
