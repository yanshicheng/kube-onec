package kubeutils

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// MetadataType 定义元数据类型，用于区分 Labels 和 Annotations
type MetadataType string

const (
	Labels      MetadataType = "labels"
	Annotations MetadataType = "annotations"
)

// Option 是一个用于修改 Kubernetes 对象的通用函数类型
type Option func(metav1.Object)

// WithMetadata 通用 Option 函数，用于自定义 Labels 或 Annotations
func WithMetadata(metaType MetadataType, data map[string]string) Option {
	return func(obj metav1.Object) {
		switch metaType {
		case Labels:
			existing := obj.GetLabels()
			if existing == nil {
				existing = make(map[string]string)
			}
			for key, value := range data {
				existing[key] = value
			}
			obj.SetLabels(existing)
		case Annotations:
			existing := obj.GetAnnotations()
			if existing == nil {
				existing = make(map[string]string)
			}
			for key, value := range data {
				existing[key] = value
			}
			obj.SetAnnotations(existing)
		}
	}
}

// WithLabels 封装 WithMetadata，用于 Labels
func WithLabels(labels map[string]string) Option {
	return WithMetadata(Labels, labels)
}

// WithAnnotations 封装 WithMetadata，用于 Annotations
func WithAnnotations(annotations map[string]string) Option {
	return WithMetadata(Annotations, annotations)
}
