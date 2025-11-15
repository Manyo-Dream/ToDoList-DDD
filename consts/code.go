package consts

const (
	Success       = 200
	InvalidParams = 400
	Error         = 500
)

var MsgFlags = map[int]string{
	Success:       "操作成功",
	InvalidParams: "请求参数错误",
	Error:         "操作失败",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[Error]
}