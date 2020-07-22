package main

//70. 爬楼梯
func climbStairs(n int) int {

	var count *int
	Climb(0, n, count)
	return *count
}

func Climb(current, total int, count *int) {

	if current == total {
		*count++
		return
	}

	if current+1 <= total {
		Climb(current+1, total, count)
	}
	if current+2 <= total {
		Climb(current+2, total, count)
	}
}

func climb(n int) int {

	seq := []int{1, 2, 3}

	if n < 3 {
		return seq[n-1]
	}
	for i := 3; i < n; i++ {
		seq = append(seq, seq[i-1]+seq[i-2])
	}
	return seq[len(seq)-1]
}

//func main() {
//
//	fmt.Println(climb(4))
//}