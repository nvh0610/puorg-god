package resp

var MsgFlags = map[int]string{}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[500000]
}
