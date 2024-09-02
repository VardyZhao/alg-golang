package addBinary

/*
给你两个二进制字符串 a 和 b ，以二进制字符串的形式返回它们的和。



示例 1：
输入:a = "11", b = "1"
输出："100"

示例 2：
输入：a = "1010", b = "1011"
输出："10101"


提示：

1 <= a.length, b.length <= 104
a 和 b 仅由字符 '0' 或 '1' 组成
字符串如果不是 "0" ，就不含前导零
*/

func AddBinary(a string, b string) string {
	res := ""
	pre := "0"
	aLen, bLen := len(a), len(b)
	aIndex, bIndex := aLen-1, bLen-1
	for aIndex >= 0 || bIndex >= 0 {
		tmpA := '0'
		if aIndex >= 0 {
			tmpA = int32(a[aIndex])
		}
		tmpB := '0'
		if bIndex >= 0 {
			tmpB = int32(b[bIndex])
		}

		if tmpA == '1' && tmpB == '1' {
			if pre == "0" {
				res = "0" + res
			} else {
				res = "1" + res
			}
			pre = "1"
		} else if tmpA == '1' || tmpB == '1' {
			if pre == "1" {
				res = "0" + res
			} else {
				res = "1" + res
				pre = "0"
			}
		} else {
			if pre == "1" {
				res = "1" + res
			} else {
				res = "0" + res
			}
			pre = "0"
		}
		aIndex--
		bIndex--
	}
	if pre == "1" {
		res = pre + res
	}
	return res
}
