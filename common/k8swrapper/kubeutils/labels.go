package kubeutils

// 通用的 Labels 和 Annotations
var (
	defaultLabels = map[string]string{
		"project": "ikubeops",
	}

	defaultAnnotations = map[string]string{
		"official_website_address": "https://www.ikubeops.com",
		"author":                   "yanshicheng",
	}
)
