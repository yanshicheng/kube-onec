package namespace

import (
	"fmt"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// PodConfig 定义 Pod 的所有配置项
type PodConfig struct {
	Name                      string                            // Pod 名称
	Namespace                 string                            // 命名空间
	Image                     string                            // 容器镜像
	ImagePullPolicy           corev1.PullPolicy                 // 镜像拉取策略
	Labels                    map[string]string                 // 标签
	Annotations               map[string]string                 // 注解
	Resources                 corev1.ResourceRequirements       // 资源限制与请求
	HostNetwork               bool                              // 是否启用主机网络模式
	HostPID                   bool                              // 是否启用主机 PID 模式
	HostIPC                   bool                              // 是否启用主机 IPC 模式
	Ports                     []corev1.ContainerPort            // 容器端口设置
	VolumeMounts              []VolumeMountConfig               // 容器内挂载存储卷
	Environment               []corev1.EnvVar                   // 环境变量
	LivenessProbe             *corev1.Probe                     // 存活检查配置
	ReadinessProbe            *corev1.Probe                     // 就绪检查配置
	StartupProbe              *corev1.Probe                     // 启动检查配置
	NodeSelector              map[string]string                 // 节点选择器
	Affinity                  *corev1.Affinity                  // 亲和性设置
	Tolerations               []corev1.Toleration               // 污点容忍
	RestartPolicy             corev1.RestartPolicy              // 重启策略
	TerminationGrace          int64                             // 强制终止 Pod 的时间（秒）
	Privileged                bool                              // 是否启用特权模式
	SecurityContext           *corev1.SecurityContext           // 容器的安全上下文
	ServiceAccount            string                            // 服务账户名称
	ImagePullSecrets          []string                          // 镜像拉取 Secret
	DNSPolicy                 corev1.DNSPolicy                  // DNS 策略
	HostAliases               []corev1.HostAlias                // 主机别名
	PreStop                   *corev1.Handler                   // 生命周期钩子：PreStop
	PostStart                 *corev1.Handler                   // 生命周期钩子：PostStart
	RuntimeClassName          *string                           // Pod 的运行时类名称
	PriorityClassName         string                            // 优先级类名称
	TopologySpreadConstraints []corev1.TopologySpreadConstraint // 拓扑扩展约束
	ContainerLogPath          string                            // 自定义容器日志路径
	Overhead                  corev1.ResourceList               // Pod 的额外资源开销
	EphemeralContainers       []corev1.EphemeralContainer       // 调试用的临时容器
}

// VolumeMountConfig 定义存储卷挂载的配置，包括 subPath
type VolumeMountConfig struct {
	Name      string        // 存储卷名称
	MountPath string        // 挂载路径
	SubPath   string        // 子路径（可选）
	ReadOnly  bool          // 是否只读
	Volume    corev1.Volume // 对应的存储卷
}

// createPod 函数动态支持所有功能点，包括 subPath 和存储卷挂载
func createPod(config PodConfig) *corev1.Pod {
	// 构造存储卷列表
	var volumes []corev1.Volume

	// 动态添加存储卷
	for _, vm := range config.VolumeMounts {
		volumes = append(volumes, vm.Volume)
	}

	// 构造容器
	container := corev1.Container{
		Name:            config.Name,
		Image:           config.Image,
		ImagePullPolicy: config.ImagePullPolicy,
		Command:         []string{}, // 用户可以动态传入
		Args:            []string{}, // 用户可以动态传入
		Ports:           config.Ports,
		Resources:       config.Resources,
		Env:             config.Environment,
		SecurityContext: config.SecurityContext,
		LivenessProbe:   config.LivenessProbe,
		ReadinessProbe:  config.ReadinessProbe,
		StartupProbe:    config.StartupProbe,
		Lifecycle: &corev1.Lifecycle{
			PostStart: config.PostStart,
			PreStop:   config.PreStop,
		},
	}

	// 动态挂载存储卷（支持 subPath）
	for _, vm := range config.VolumeMounts {
		container.VolumeMounts = append(container.VolumeMounts, corev1.VolumeMount{
			Name:      vm.Name,
			MountPath: vm.MountPath,
			SubPath:   vm.SubPath, // 动态支持 subPath
			ReadOnly:  vm.ReadOnly,
		})
	}

	// 构造 Pod
	pod := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:        config.Name,
			Namespace:   config.Namespace,
			Labels:      config.Labels,
			Annotations: config.Annotations,
		},
		Spec: corev1.PodSpec{
			Containers:                    []corev1.Container{container},
			Volumes:                       volumes,
			HostNetwork:                   config.HostNetwork,
			HostPID:                       config.HostPID,
			HostIPC:                       config.HostIPC,
			NodeSelector:                  config.NodeSelector,
			Affinity:                      config.Affinity,
			Tolerations:                   config.Tolerations,
			RestartPolicy:                 config.RestartPolicy,
			TerminationGracePeriodSeconds: &config.TerminationGrace,
			ServiceAccountName:            config.ServiceAccount,
			ImagePullSecrets:              convertImagePullSecrets(config.ImagePullSecrets),
			DNSPolicy:                     config.DNSPolicy,
			HostAliases:                   config.HostAliases,
			RuntimeClassName:              config.RuntimeClassName,
			PriorityClassName:             config.PriorityClassName,
			TopologySpreadConstraints:     config.TopologySpreadConstraints,
			Overhead:                      config.Overhead,
		},
	}

	// 添加调试容器
	if len(config.EphemeralContainers) > 0 {
		pod.Spec.EphemeralContainers = config.EphemeralContainers
	}

	return pod
}

// convertImagePullSecrets 将字符串转换为 LocalObjectReference
func convertImagePullSecrets(secrets []string) []corev1.LocalObjectReference {
	var refs []corev1.LocalObjectReference
	for _, secret := range secrets {
		refs = append(refs, corev1.LocalObjectReference{Name: secret})
	}
	return refs
}

// 示例：构造完整的 Pod 配置
func main() {
	// 定义一个 ConfigMap 存储卷
	configMapVolume := corev1.Volume{
		Name: "config-volume",
		VolumeSource: corev1.VolumeSource{
			ConfigMap: &corev1.ConfigMapVolumeSource{
				LocalObjectReference: corev1.LocalObjectReference{Name: "my-config"},
			},
		},
	}

	// 定义一个 PVC 存储卷
	pvcVolume := corev1.Volume{
		Name: "pvc-volume",
		VolumeSource: corev1.VolumeSource{
			PersistentVolumeClaim: &corev1.PersistentVolumeClaimVolumeSource{
				ClaimName: "my-pvc",
			},
		},
	}

	// 定义 Pod 配置
	podConfig := PodConfig{
		Name:            "example-pod",
		Namespace:       "default",
		Image:           "nginx:1.21",
		ImagePullPolicy: corev1.PullIfNotPresent,
		Labels:          map[string]string{"app": "example"},
		Annotations:     map[string]string{"description": "Example Pod"},
		VolumeMounts: []VolumeMountConfig{
			{
				Name:      "config-volume",
				MountPath: "/etc/config",
				SubPath:   "app-config",
				Volume:    configMapVolume,
			},
			{
				Name:      "pvc-volume",
				MountPath: "/data",
				SubPath:   "user-data",
				Volume:    pvcVolume,
			},
		},
		Environment: []corev1.EnvVar{
			{Name: "ENV_VAR", Value: "value"},
		},
		RestartPolicy: corev1.RestartPolicyAlways,
	}

	// 创建 Pod
	pod := createPod(podConfig)

	// 打印 Pod 详情
	fmt.Printf("Pod: %+v\n", pod)
}
