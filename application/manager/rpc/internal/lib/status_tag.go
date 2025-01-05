package lib

func StatusTagName(id int64) string {
	switch id {
	case 0:
		return "正常"

	default:
		return "异常"

	}
}
