package main

//package main
//
//import (
//	"fmt"
//	"time"
//)
//
//func main() {
//	// 获取当前时间
//	now := time.Now()
//
//	// 格式化当前时间为 年-月-日 格式
//	formattedDate := now.Format("2006-01-02")
//	fmt.Println("当前时间 (年-月-日):", formattedDate)
//
//	// 将 年-月-日 字符串解析回时间对象
//	parsedTime, err := time.Parse("2006-01-02", formattedDate)
//	if err != nil {
//		fmt.Println("时间解析错误:", err)
//		return
//	}
//
//	// 转换为时间戳（秒级别）
//	timestamp := parsedTime.Unix()
//	fmt.Println("时间戳 (秒):", timestamp)
//
//	// 假设一个时间戳（秒级别）
//	timestamp1 := int64(1718236800)
//
//	// 将时间戳转换为 time.Time 类型
//	convertedTime := time.Unix(timestamp1, 0)
//
//	// 格式化为 年-月-日 格式
//	formattedDate2 := convertedTime.Format("2006-01-02")
//
//	// 输出结果
//	fmt.Println("转换后的时间 (年-月-日):", formattedDate2)
//}
