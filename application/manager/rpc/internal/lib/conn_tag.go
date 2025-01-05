package lib

import "strings"

const (
	TOKEN      = 1
	KUBECONFIG = 2
	AGENT      = 3
	OTHER      = 4
)

func GetConnTag(conn string) int64 {
	// conn 转为小写
	switch strings.ToLower(conn) {
	case "token":
		return TOKEN
	case "kubeconfig":
		return KUBECONFIG
	case "agent":
		return AGENT
	default:
		return OTHER
	}
}

func GetConnTagName(conn int64) string {
	switch conn {
	case TOKEN:
		return "token"
	case KUBECONFIG:
		return "kubeconfig"
	case AGENT:
		return "agent"
	default:
		return "other"
	}
}
