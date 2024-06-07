package res

import (
	"blog_server/global"
	"encoding/json"
	"os"
)

type ErrorCode int

const file = "models/res/error_code.json"

type ErrorMap map[ErrorCode]string

func InitErrorCode() {
	data, err := os.ReadFile(file)
	if err != nil {
		global.LOG.Error(err)
		return
	}
	var errMap = ErrorMap{}
	err = json.Unmarshal(data, &errMap)
	if err != nil {
		global.LOG.Error(err)
		return
	}
	ErrorMaps = errMap
}

var (
	ErrorMaps map[ErrorCode]string
)
