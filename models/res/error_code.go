package res

type ErrorCode int

const (
	SettingError      ErrorCode = 1001 //系统错误
	MarkdownhtmlError ErrorCode = 1101 //markdown转换错误
)

var (
	ErrorMap = map[ErrorCode]string{
		SettingError:      "系统错误",
		MarkdownhtmlError: "markdown转换错误",
	}
)
