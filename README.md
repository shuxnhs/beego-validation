[Chinese](README_CN.md)
# beego-validation

This is a BaseController and Validation for [Beego](https://github.com/astaxie/beego) framework.

It uses [govalidator](https://github.com/asaskevich/govalidator) to support the request param validaton. It also provides baseController that other handler struct can combine this BaseController.


## Usage

Download and install using [go module](https://blog.golang.org/using-go-modules):

```sh
export GO111MODULE=on
go get -u github.com/shuxnhs/beego-validation
```

Import it in your code:

```go
import (
    validation "github.com/shuxnhs/beego-validation"
)
```


## Example

Please see [the example Controller](example/example.go) and you can use `Declarative parameter validation` to validator the request data.

[embedmd]:# (example/example.go go)
```go
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
```

## explain