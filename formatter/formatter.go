package formatter

import (
	"errors"
	"reflect"

	"github.com/astaxie/beego"
	beegoContext "github.com/astaxie/beego/context"
)

type Formatter interface {
	FormatterType() string
	Parse(bc *beegoContext.Context, rule Rule) interface{}
}

func CheckRule(rule Rule) error {
	def, ok := rule["default"]
	if ok && reflect.TypeOf(def).String() != rule["type"].(string) {
		return errors.New("Rule Is Illegal ")
	}
	return nil
}

func GetParamRequireError(key string) error {
	return errors.New("Param " + key + " Is Required ")
}

func ApiError(bc *beegoContext.Context, errMsg string) {
	res := NewAPIResponse()
	_ = bc.Output.JSON(res.SetRet(400).SetMsg(errMsg), true, false)
	panic(beego.ErrAbort)
}
