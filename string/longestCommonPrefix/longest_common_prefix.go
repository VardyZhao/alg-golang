package longestCommonPrefix

/*
编写一个函数来查找字符串数组中的最长公共前缀。

如果不存在公共前缀，返回空字符串 ""。

示例 1：
输入：strs = ["flower","flow","flight"]
输出："fl"

示例 2：
输入：strs = ["dog","racecar","car"]
输出：""
解释：输入不存在公共前缀。

提示：

1 <= strs.length <= 200
0 <= strs[i].length <= 200
strs[i] 仅由小写英文字母组成
*/
func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	prefix := strs[0]
	index := 0
	for i := 1; i < len(strs); i++ {
		strlen := len(strs[i])
		index = 0
		for index < len(prefix) {
			if index >= strlen {
				break
			}
			if prefix[index] != strs[i][index] {
				break
			} else {
				index++
			}
		}
		prefix = prefix[:index]
		if prefix == "" {
			return ""
		}
	}
	return prefix[:index]
}
