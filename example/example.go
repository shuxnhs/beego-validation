package main

import (
	validation "github.com/shuxnhs/beego-validation"
	"github.com/shuxnhs/beego-validation/formatter"
)

type ExampleController struct {
	validation.BaseController
}

func (e *ExampleController) Ping() {
	ruleMap := map[string]formatter.Rule{
		"objectId": {"name": "object_id", "type": formatter.ValidTypeInt, "required": true, "rule": "", "default": 123},
		"objectName": {"name": "object_name", "type": formatter.ValidTypeString, "required": true,
			"rule": "length(1|10),in(string1|string2|...|stringN)"},
	}
	paramMap := e.Rules(ruleMap)
	e.ApiSuccessData("success", paramMap["objectId"])
}
