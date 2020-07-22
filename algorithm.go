package main

import (
	"fmt"
	"strconv"
	"strings"
)

//1
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	l := len(nums)
	for i := 0; i < l; i++ {

		c := target - nums[i]
		if v, ok := m[c]; ok {
			return []int{v, i}
		}
		m[nums[i]] = i
	}
	return nil
}

//97. 交错字符串
func isInterleave(s1 string, s2 string, s3 string) bool {

	return false
}

//7. 整数反转

func reverseInt(i int) int {
	var nums int

	for i != 0 {

		temp := i % 10
		nums = nums*10 + temp
		i = i / 10

		if i > 1<<31-1 || i < -1<<31 {
			return 0
		}

	}
	return nums
}

//9. 回文数
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	s := strconv.Itoa(x)
	j := len(s)
	for i := 0; i < j/2; i++ {
		if s[i] != s[j-i-1] {
			return false
		}
	}

	return true

}

//13. 罗马数字转整数
func romanToInt(s string) int {
	res := 0
	m := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	last := 0
	for i := len(s) - 1; i >= 0; i-- {
		temp := m[s[i]]

		sign := 1
		if temp < last {
			sign = -1
		}
		res += sign * temp
		last = temp
	}
	return res
}

//14. 最长公共前缀
func longestCommonPrefix(strs []string) string {

	if len(strs) == 0 {
		return ""
	}
	ans := strs[0]

	for i, v := range strs {
		if v == "" {
			return ""
		}
		if i == 0 {
			continue
		}

		minLen := len(ans)
		l1 := len(v)
		if minLen > l1 {
			minLen = l1
		}

		for j := minLen; j > 0; j-- {
			if strings.HasPrefix(v, ans[:j]) {
				ans = ans[:j]
				break
			}
			ans = ans[:j]
			if j == 1 {
				return ""
			}
		}
	}

	return ans
}

//20. 有效的括号
func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}
	if len(s)%2 == 1 {
		return false
	}
	m := map[string]string{")": "(", "]": "[", "}": "{"}
	//m := map[byte]byte{'(': ')', '[': ']', '{': '}'}

	var stack []string

	for i := 0; i < len(s); i++ {
		if len(stack) == 0 {
			stack = append(stack, string(s[i]))
		} else {
			if stack[len(stack)-1] == m[string(s[i])] {
				stack = stack[:len(stack)-1]

			} else {
				stack = append(stack, string(s[i]))
			}
		}
	}

	if len(stack) == 0 {
		return true
	}
	return false
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	//fmt.Println("l1 is",l1)
	//fmt.Println("l2 is",l2)
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	var res *ListNode

	if l1.Val >= l2.Val {
		res = l2
		res.Next = mergeTwoLists(l1, l2.Next)

	} else {
		res = l1
		res.Next = mergeTwoLists(l1.Next, l2)
	}
	return res
}

//26. 删除排序数组中的重复项
func removeDuplicates(nums []int) int {

	if len(nums) == 0 {
		return 0
	}
	left, right := 0, 1

	for ; right < len(nums); right++ {
		if nums[left] == nums[right] {
			continue
		}
		left++
		nums[left] = nums[right]
	}
	fmt.Println(nums[:left+1])
	return left + 1
}

//27. 移除元素
func removeElement(nums []int, val int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}
	i := 0
	for ; i < len(nums); {
		if nums[i] == val {
			nums = append(nums[:i], nums[i+1:]...)
			//fmt.Println(nums)
			continue
		}
		i++
	}
	//fmt.Println(nums)
	//fmt.Println(len(nums))
	return i
}

//实现 strStr()

func strStr(haystack string, needle string) int {
	m := len(haystack)
	n := len(needle)
	if m == 0 || m < n {
		return -1
	}
	if n == 0 {
		return 0
	}

	for i := m - n; i >= 0; i++ {
		if haystack[i:] == needle {
			return i
		}
	}
	return -1
}

func searchInsert(nums []int, target int) int {

	if len(nums) == 0 || nums[0] > target {
		return 0
	}

	i, j := 0, 1
	for ; i < j && j < len(nums); j++ {
		if nums[i] < target && nums[j] > target {

			return j
		}
		i++

	}
	return j
}

//38. 外观数列

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	if n == 2 {
		return "11"
	}

	ans := countAndSay(n - 1)
	count := 1
	r := ""

	for i := 1; i < len(ans); i++ {
		if ans[i-1] == ans[i] {
			count++
			continue
		} else {

			r += strconv.Itoa(count) + string(ans[i])
			count = 1
		}
		if i == len(ans)-1 {
			r += strconv.Itoa(count) + string(ans[i])
			ans = r
			return ans
		}

	}

	return ans
}

//53. 最大子序和
func maxSubArray(nums []int) int {

	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	sum := nums[0]
	res := 0
	for _, v := range nums {
		if sum < 0 {
			sum = v
		} else {
			sum += v
		}
		res = max(res, sum)
	}

	return res
}
func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
func maxSubArrayKMP(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	res, sum := nums[0], nums[0]
	for _, v := range nums {

		if res > 0 {
			res += v
		} else {
			res = v
		}

		if res > sum {
			sum = res
		}
	}

	return sum
}

func lengthOfLastWord(s string) int {
	if len(s) == 0 {
		return 0
	}
	ss := strings.Split(s, " ")
	return len(ss[len(ss)-1])
}

//66. 加一

func plusOne(digits []int) []int {

	flag := true
	for i := len(digits) - 1; i >= 0; i-- {

		if flag {
			digits[i] += 1

			if digits[i] == 10 {
				digits[i] = 0
				flag = true

				if i == 0 {
					digits = append([]int{1}, digits...)
				}
				continue
			}

			flag = false
		}

	}

	return digits
}

//67. 二进制求和

func addBinary(a string, b string) string {
	la := len(a)
	lb := len(b)
	length := la
	if length < lb {
		length = lb
	}
	r := make([]byte, length)
	var extra byte
	for i := 0; i < length; i++ {
		var dib byte
		ai, bi := len(a)-1-i, len(b)-1-i
		if ai >= 0 && bi >= 0 {
			dib = a[ai] - '0' + b[bi] - '0'
		} else if ai >= 0 {
			dib = a[ai] - '0'
		} else if bi >= 0 {
			dib = b[bi] - '0'
		}
		dib += extra
		if dib > 1 {
			extra = 1
		} else {
			extra = 0
		}
		r[length-1-i] = dib%2 + '0'
	}
	if extra == 1 {
		r = append([]byte{'1'}, r...)
	}
	return string(r)
}

func getAbs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}

//69. x 的平方根
func mySqrt(x float64) float64 {

	z := 1.0

	if x < 0 {
		return 0
	} else if x == 0 {
		return 0
	} else {
		for getAbs(z*z-x) > 1e-6 {
			z = (z + x/z) / 2
		}
		return z
	}
}

//83. 删除排序链表中的重复元素

//Definition for singly-linked list.

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	cur := head

	if cur.Next != nil {
		if cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}

	return head
}

//88. 合并两个有序数组

func merge(nums1 []int, m int, nums2 []int, n int) {
	temp := make([]int, m)
	copy(temp, nums1)

	t, j := 0, 0
	for i := 0; i < len(nums1); i++ {

		if t >= len(temp) {
			nums1[i] = nums2[j]
			j++
			continue
		}

		if j >= len(nums2) {

			nums1[i] = temp[t]
			t++
			continue
		}

		if nums2[j] <= temp[t] {
			nums1[i] = nums2[j]
			j++
		} else {
			nums1[i] = temp[t]
			t++
		}
	}
}

//100. 相同的树
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {

	if p == nil && q == nil {
		return false
	}
	if p == nil || q == nil || p.Val != q.Val {
		return true
	}
	return isSameTree(p.Right, q.Right) && isSameTree(p.Left, q.Left)
}

func main() {

	n1 := []int{1, 2, 3, 0, 0, 0}
	n2 := []int{2, 5, 6}
	m := 3
	n := 3
	fmt.Println(n1)
	merge(n1, m, n2, n)
	fmt.Println(n1)
}
