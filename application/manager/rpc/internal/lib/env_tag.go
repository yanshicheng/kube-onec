package lib

// 集群环境标签定义
const (
	EnvTagTest             = 1 // 开发环境
	EnvTagDevelopment      = 2 // 测试环境
	EnvTagProduction       = 3 // 生产环境
	EnvTagPreRelease       = 4 // 预发布环境
	EnvTagDisasterRecovery = 5 // 灾备环境
)

// EnvTagToString 将环境标签转为字符串
func EnvTagToString(envTag int64) string {
	switch envTag {
	case EnvTagTest:
		return "开发环境"
	case EnvTagDevelopment:
		return "测试环境"
	case EnvTagProduction:
		return "生产环境"
	case EnvTagPreRelease:
		return "预发布环境"
	case EnvTagDisasterRecovery:
		return "灾备环境"
	default:
		return "未知环境"
	}
}
