package dailyexcise

import (
	"strconv"
	"strings"
)

// TwoSum 1.两数之和
func TwoSum(nums []int, target int) []int {
	hash := make(map[int]int)
	for i, k := range nums {
		if p, ok := hash[target - k]; ok {
			return []int{i, p}
		}
		hash[k] = i
	}
	return nil
}

// RomanToInt 9.回文数
func RomanToInt(s string) int {
	hash := map[string]int {"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}
	ans := hash[string(s[len(s)-1])]
	for i := len(s) - 2; i >=0; i-- {
		if hash[string(s[i])] < hash[string(s[i+1])] {
			ans -= hash[string(s[i])]
		} else {
			ans += hash[string(s[i])]
		}
	}
	return ans
}

// isPalindrome 13.罗马数字转整数
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	strX := strconv.Itoa(x)
	for i, j := 0, len(strX)-1; i < len(strX)/2; i, j = i+1, j-1 {
		if strX[i] == strX[j] {
			continue
		} else {
			return false
		}
	}
	return true
}

// longestCommonPrefix 14.最长公共前缀
func longestCommonPrefix(strs []string) string {
	ans := ""
	minLength := 200
	for _, v := range strs {
		if len(v) < minLength {
			minLength = len(v)
		}
	}
	for i := 0; i < minLength; i++ {
		stand := string(strs[0][i])
		for j := 1; j < len(strs); j++ {
			if string(strs[j][i]) != stand {
				return ans
			}
		}
		ans += stand
	}
	return ans
}

// isValid 20.有效的括号

func isValid(s string) bool {
	n := len(s)
	if n % 2 == 1 {
		return false
	}
	pairs := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	stack := []byte{}
	for i := 0; i < n; i++ {
		if _, ok := pairs[s[i]]; ok {  // 出现反括号
			if len(stack) == 0 || stack[len(stack)-1] != pairs[s[i]] {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0
}

// 21 todo

// removeDuplicates 26.删除有序数组中的重复项
func removeDuplicates(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	//双指针
	slow, fast := 1, 1
	for fast = 1; fast < length; fast++ {
		if nums[fast] != nums[fast-1] {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}

// removeElement 27.移除元素
func removeElement(nums []int, val int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	fast, slow := 0, 0
	for fast = 0; fast < length; fast++ {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}

// strStr 28.实现strStr()
func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	if len(haystack) < len(needle) {
		return -1
	}
	if len(haystack) == len(needle) {
		if haystack == needle {
			return 0
		} else {
			return -1
		}
	}
	for i := 0; i <= len(haystack)-len(needle); i++ {
		if haystack[i:i+len(needle)] == needle {
			return i
		}
	}
	return -1
}

// searchInsert 35.搜索插入位置
func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums) - 1
	for left <= right {
		mid := (right-left)>>1 + left
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}

// findCenter 1791.找出星型图的中心节点
func findCenter(edges [][]int) int {
	if edges[0][0] == edges[1][0] || edges[0][0] == edges[1][1]{
		return edges[0][0]
	} else {
		return edges[0][1]
	}
}

// lengthOfLastWord 58.末尾单词的长度
func lengthOfLastWord(s string) int {
	res := 0
	i := 0
	for i = len(s)-1; i >= 0; i-- {
		if s[i] == ' '{
			continue
		} else {
			break
		}
	}
	for j := i; j >= 0; j-- {
		if s[j] != ' ' {
			res ++
		} else {
			return res
		}
	}
	return res
}

// plusOne 66.加一
func plusOne(digits []int) []int {
	if len(digits) == 1 && digits[0] == 0 {
		digits[0] = 1
		return digits
	}
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i] += 1
		if digits[i] != 10 {
			return digits
		} else {
			digits[i] = 0
		}
	}
	if digits[0] == 0 {
		digits[0] = 1
		digits = append(digits,  0)
	}
	return digits
}

// addBinary 67.二进制求和
func addBinary(a string, b string) string {
	flag := 0       //进位记录
	ans := ""
	lena, lenb := len(a) - 1, len(b) - 1
	for lena >=0 || lenb>= 0 {
		stra, strb := 0, 0
		if lena >= 0 {
			stra = int(a[lena] - '0')
		}
		if lenb >= 0 {
			strb = int(b[lenb] - '0')
		}
		sum := stra + strb + flag   // 对应位相加再加上上一位的进位
		switch sum {
		case 0: flag = 0; ans = "0" + ans
		case 1: flag = 0; ans = "1" + ans
		case 2: flag = 1; ans = "0" + ans
		case 3: flag = 1; ans = "1" + ans
		}
		lena--
		lenb--
	}
	if flag == 1 {
		ans = "1" + ans   //第一位如果有进位, 就再进位一次
	}
	return ans
}

// mySqrt 69.x的平方根
func mySqrt(x int) int {
	if x == 0 {
		return 0
	}
	left, right := 0, x
	for left <= right {
		mid := (right - left) >> 1 + left
		if mid * mid > x {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return right
}

// merge 88.合并两个有序数组  方法一: 从前往后插入, 需要中间数组store来保存当前结果, 因为直接存入num1可能会被覆盖
//func merge(nums1 []int, m int, nums2 []int, n int)  {   //双指针
//	i, j := 0, 0
//	store := make([]int, 0, m+n)
//	for {
//		if i == m {
//			store = append(store, nums2[j:]...)
//			break
//		}
//		if j == n {
//			store = append(store, nums1[i:]...)
//			break
//		}
//		if nums1[i] > nums2[j] {
//			store = append(store, nums2[j])
//			j++
//		} else {
//			store = append(store, nums1[i])
//			i++
//		}
//	}
//	for i = 0; i < m+n; i++ {
//		nums1[i] = store[i]
//	}
//}

// merge 88.合并两个有序数组  方法二: 从后往前插入, 不用担心未使用的元素被覆盖, 比较大的元素插到队列尾部
func merge(nums1 []int, m int, nums2 []int, n int)  {   //双指针
	for i, j, k := m-1, n-1, m+n-1; i >= 0 || j >= 0; k-- {
		temp := 0
		if i == -1 {
			temp = nums2[j]
			j--
		} else if j == -1 {
			temp = nums1[i]
			i--
		} else if nums1[i] > nums2[j] {
			temp = nums1[i]
			i--
		} else {
			temp = nums2[j]
			j--
		}
		nums1[k] = temp
	}
}

//isPalindrome 125.验证回文串
// asicc码对照表   48-57 -> "0-9"  65-90 -> "A-Z"   97-122 -> "a-z"
func IsPalindrome(s string) bool { //双指针
	if len(s) <= 1 {
		return true
	}
	left, right := 0, len(s) - 1
	s = strings.ToLower(s)
	for left < right{
		for !((s[left] >= '0' && s[left] <= '9') || (s[left] >= 'a' && s[left] <= 'z')) {
			left ++
			if left == len(s) - 1 {
				return true
			}
		}
		for !((s[right] >= '0' && s[right] <= '9') || (s[right] >= 'a' && s[right] <= 'z')) {
			right --
			if right == left {
				return true
			}
		}
		if s[left] == s[right] {
			left ++
			right --
			continue
		} else {
			return false
		}
	}
	return true
}

//func IsPalindrome(s string) bool { //先清洗再比较, 需要一个store数组， 此方法也能過， 但是时间空间开销较大
//	store := ""
//	s = strings.ToLower(s)
//	for i:= 0; i< len(s); i++ {
//		if (s[i] >= '0' && s[i] <= '9') || (s[i] >= 'a' && s[i] <= 'z') {
//			store += string(s[i])
//		}
//	}
//	if len(store) <= 1 {
//		return true
//	}
//	length := len(store)
//	for j := 0; j <= length/2; j++ {
//		if store[j] != store[length-j-1] {
//			return false
//		}
//	}
//	return true
//}

// singleNumber 136.只出现一次的数字
//交换律：a ^ b ^ c <=> a ^ c ^ b
//
//任何数于0异或为任何数 0 ^ n => n
//
//相同的数异或为0: n ^ n => 0
//
//var a = [2,3,2,4,4]
//
//2 ^ 3 ^ 2 ^ 4 ^ 4等价于 2 ^ 2 ^ 4 ^ 4 ^ 3 => 0 ^ 0 ^3 => 3
func singleNumber(nums []int) int {  // 数组中所有数做异或运算即可得结果
	res := 0
	for _, k := range nums {
		res ^= k
	}
	return res
}

// hammingWeight 191.数字1的数量
// n & (n-1) 这个操作是算法中常见的，作用是消除数字 n 的二进制表示中的最后一个 1, 不断消去1来计数
func hammingWeight(num uint32) int {
	res := 0
	for num != 0 {
		num = num & (num-1)
		res++
	}
	return res
}

// convertToTitle 16.excel表列名称
func convertToTitle(columnNumber int) string {
	ans := ""
	for columnNumber > 0 {
		columnNumber -= 1
		ans = string(columnNumber % 26 + 'A') + ans
		columnNumber /= 26
	}
	return ans
}

