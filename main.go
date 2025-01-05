package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/yanshicheng/kube-onec/common/k8swrapper/manager"
	"github.com/yanshicheng/kube-onec/utils"
	"k8s.io/client-go/rest"
)

const (
	token  = "eyJhbGciOiJSUzI1NiIsImtpZCI6IlhkcFBpMnJZeDF4UE5MdWphUGY5RzJFYURWMjdJd0V2OG5YekpQNV9sMjQifQ.eyJpc3MiOiJrdWJlcm5ldGVzL3NlcnZpY2VhY2NvdW50Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9uYW1lc3BhY2UiOiJkZWZhdWx0Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9zZWNyZXQubmFtZSI6Im15LXBlcm1hbmVudC10b2tlbi1zYS10b2tlbi1zZWNyZXQiLCJrdWJlcm5ldGVzLmlvL3NlcnZpY2VhY2NvdW50L3NlcnZpY2UtYWNjb3VudC5uYW1lIjoibXktcGVybWFuZW50LXRva2VuLXNhIiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9zZXJ2aWNlLWFjY291bnQudWlkIjoiZmYwZjY2NmYtN2FhYy00MzE1LWJlZTUtZTg3MWM0MDQ4NWM4Iiwic3ViIjoic3lzdGVtOnNlcnZpY2VhY2NvdW50OmRlZmF1bHQ6bXktcGVybWFuZW50LXRva2VuLXNhIn0.sw-fkYEZbeQOQI90oOSxl8guuE6ySr8nxd8xasYHNr7pTKFizU050NQgxpzfnKxLrnf3QiF6PILhfEIyjM5ChjcCHlMFNJa8wP3cKNm4QjiBal8N4JbfJCcVufSyJv9ZOMN37eFQY3Ao22yGsFbQLur6h7fhhAnJF38ix31diKW-CUufvW_F6g3K7W99pBQQdJTsTR90jPC23eK6hdaXHBdtjagAQ_XNNvbllBxeIpjJ4VWq_jHw0ZWzyO5W6l_gGb2jdUbLmxotpb99at-cR640xmrsiTxqur8VIwaJOGdVsqxx9h1RkjoDA6_II-NB1Sv_5Z-IK1C3LY3g8-h4eA"
	ipaddr = "https://172.16.1.21:6443"
)

func main() {
	// 创建配置
	config := &rest.Config{
		Host:        ipaddr,
		BearerToken: token,
		TLSClientConfig: rest.TLSClientConfig{
			Insecure: true,
		},
	}
	ctx := context.Background()
	// 初始化 Kubernetes 客户端
	manager := manager.OnecK8sClientManager{}
	client, err := manager.GetOrCreateOnecK8sClient(ctx, "default", config)

	if err != nil {
		fmt.Printf("创建客户端失败: %v\n", err)
		return
	}

	client.GetNodes().AddLabel("ik8s-master-01", "test", "test")
	fmt.Println(client.GetCluster().GetClusterInfo())
	info, err := client.GetNodes().GetAllNodesInfo()
	if err != nil {
		return
	}
	res, err := json.MarshalIndent(info, "", "  ")
	if err != nil {
		fmt.Println("JSON 序列化失败:", err)
		return
	}
	fmt.Println(string(res))
	fmt.Println(utils.GenerateRandomID())
	fmt.Println(len("3e78158e-ccb6-42da-8a16-41649ebfc6ba"))
	fmt.Println(len("1650d290-f40d-4078-93bc-57ff1a4edb1a"))
	fmt.Println("nodeUid", len("eb34c08d-b2df-47ca-947f-d2dc2a8209da"))
	fmt.Println(len(token))
}
