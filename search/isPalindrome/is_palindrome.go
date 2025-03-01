package isPalindrome

/*
给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。

回文数
是指正序（从左向右）和倒序（从右向左）读都是一样的整数。

例如，121 是回文，而 123 不是。

示例 1：
输入：x = 121
输出：true

示例 2：
输入：x = -121
输出：false
解释：从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。

示例 3：
输入：x = 10
输出：false
解释：从右向左读, 为 01 。因此它不是一个回文数。

提示：
-2的31次方 <= x <= 2的31次方 - 1

进阶：你能不将整数转为字符串来解决这个问题吗？
*/
func IsPalindrome(x int) bool {
	if x < 0 || (x != 0 && x%10 == 0) { // 小于0或者为只有一个数
		return false
	}
	num, res := x, 0
	for num > 0 {
		res = res*10 + num%10
		num = num / 10
	}
	return x == res
}
