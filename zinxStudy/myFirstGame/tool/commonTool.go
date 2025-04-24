package tool

/**
 * @Author: lcy
 * @Description: 常用工具
 * @File: commonTool
 * @Version: 1.0.0
 * @Date: 2022/5/15 16:09
 */

// 判断切片中是否包含某个元素
func Contains[T comparable](s []T, target T) bool {
	for _, v := range s {
		if v == target {
			return true
		}
	}
	return false
}
