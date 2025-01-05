package utils

import (
	"encoding/json"
	"fmt"
	"github.com/yanshicheng/kube-onec/common/k8swrapper/core"
)

// JSONToMap 将 JSON 字符串转为 map[string]interface{}
func JSONToMap(jsonStr string) (map[string]interface{}, error) {
	var result map[string]interface{}
	err := json.Unmarshal([]byte(jsonStr), &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON to map: %w", err)
	}
	return result, nil
}

// MapStringToJSON 将 map[string]string 转为 JSON 字符串
func MapStringToJSON(data map[string]string) (string, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", fmt.Errorf("failed to marshal map[string]string to JSON: %w", err)
	}
	return string(jsonData), nil
}

// JSONToMapString 将 JSON 字符串转为 map[string]string
func JSONToMapString(jsonStr string) (map[string]string, error) {
	var result map[string]string
	err := json.Unmarshal([]byte(jsonStr), &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON to map[string]string: %w", err)
	}
	return result, nil
}

func TaintsToJSON(taints []core.Taint) (string, error) {
	jsonData, err := json.Marshal(taints)
	if err != nil {
		return "", fmt.Errorf("failed to marshal taints: %w", err)
	}
	return string(jsonData), nil
}

func JSONToTaints(jsonData string) ([]core.Taint, error) {
	var taints []core.Taint
	err := json.Unmarshal([]byte(jsonData), &taints)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON to taints: %w", err)
	}
	return taints, nil
}
