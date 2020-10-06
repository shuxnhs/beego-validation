package validation

import (
	"strings"

	"github.com/asaskevich/govalidator"
	"github.com/astaxie/beego"
	"github.com/shuxnhs/beego-validation/formatter"
)

type BaseController struct {
	response *formatter.APIResponse

	beego.Controller
}

func (b *BaseController) Rules(ruleMap map[string]formatter.Rule) map[string]interface{} {
	b.response = formatter.NewAPIResponse()
	inputMap := make(map[string]interface{})
	mapTemplate := make(map[string]interface{})
	for param, rule := range ruleMap {
		if err := formatter.CheckRule(rule); err != nil {
			b.ApiError(503, err.Error())
		}
		queryName := rule["name"].(string)
		switch rule["type"] {
		case formatter.ValidTypeInt:
			inputMap[param] = formatter.IntFormatter.Parse(b.Ctx, rule)
		case formatter.ValidTypeBool:
			inputMap[param] = formatter.BoolFormatter.Parse(b.Ctx, rule)
		case formatter.ValidTypeFloat:
			inputMap[param] = formatter.FloatFormatter.Parse(b.Ctx, rule)
		case formatter.ValidTypeMap:
			inputMap[param] = b.GetMap(queryName)
		case formatter.ValidTypeString:
			inputMap[param] = formatter.StringFormatter.Parse(b.Ctx, rule)
		case formatter.ValidTypeStrings:
			inputMap[param] = b.GetStrings(queryName)
		}
		mapTemplate[param] = rule["rule"]
	}
	if _, err := govalidator.ValidateMap(inputMap, mapTemplate); err != nil {
		b.ApiError(400, err.Error())
	}
	return inputMap
}

func (b *BaseController) GetMap(key string) map[string]string {
	dicts := make(map[string]string)
	for k, v := range b.Input() {
		if i := strings.IndexByte(k, '['); i >= 1 && k[0:i] == key {
			if j := strings.IndexByte(k[i+1:], ']'); j >= 1 {
				dicts[k[i+1:][:j]] = v[0]
			}
		}
	}
	return dicts
}

/**
 * 异常返回
 */
func (b *BaseController) ApiError(ret int, msg string) {
	b.Data["json"] = b.response.Error(ret, msg)
	b.ServeJSON()
	b.StopRun()
}

/**
 * 异常返回，带数据
 */
func (b *BaseController) ApiErrorData(ret int, msg string, data interface{}) {
	b.Data["json"] = b.response.Error(ret, msg).SetData(b.response.Data.SetData(data))
	b.ServeJSON()
	b.StopRun()
}

/**
 * 业务失败返回
 */
func (b *BaseController) ApiFail(code int, msg string) {
	b.Data["json"] = b.response.SetData(b.response.Data.SetCode(code).SetMsg(msg))
	b.ServeJSON()
	b.StopRun()
}

/**
 * 业务失败返回，带数据
 */
func (b *BaseController) ApiFailData(code int, msg string, data interface{}) {
	b.Data["json"] = b.response.SetData(b.response.Data.SetCode(code).SetMsg(msg).SetData(data))
	b.ServeJSON()
	b.StopRun()
}

/**
 * 业务成功返回
 */
func (b *BaseController) ApiSuccess(msg string) {
	b.Data["json"] = b.response.SetData(b.response.Data.SetMsg(msg))
	b.ServeJSON()
	b.StopRun()
}

/**
 * 业务成功返回，带数据
 */
func (b *BaseController) ApiSuccessData(msg string, data interface{}) {
	b.Data["json"] = b.response.SetData(b.response.Data.SetMsg(msg).SetData(data))
	b.ServeJSON()
	b.StopRun()
}

// 重写ServeJSON，兼容jsonp
func (b *BaseController) ServeJSON(encoding ...bool) {
	callback := b.GetString("callback")
	if callback == "" {
		b.Controller.ServeJSON(encoding...)
	} else {
		b.Controller.ServeJSONP()
	}
}
