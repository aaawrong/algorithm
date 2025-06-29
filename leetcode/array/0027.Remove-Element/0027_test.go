package array

import (
	"fmt"
	"testing"
)

type question27 struct {
	para27
	ans27
}

// para 是参数
// one 代表第一个参数
type para27 struct {
	one []int
	two int
}

// ans 是答案
// one 代表第一个答案
type ans27 struct {
	one int
}

func Test_Problem27(t *testing.T) {
	qs := []question27{
		{
			para27{[]int{0, 1, 0, 3, 0, 12}, 0},
			ans27{3},
		},

		//{
		//	para27{[]int{0, 1, 0, 3, 0, 0, 0, 0, 1, 12}, 0},
		//	ans27{4},
		//},
	}

	fmt.Printf("------------------------Leetcode Problem 27------------------------\n")

	for _, q := range qs {
		fmt.Println(q.para27.one)
		_, p := q.ans27, q.para27
		fmt.Printf("【input】:%v    【output】:%v\n", p.one, removeElement(p.one, p.two))
	}
	fmt.Printf("\n\n\n")
}
