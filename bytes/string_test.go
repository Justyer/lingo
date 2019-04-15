package bytes

import (
	"fmt"
	"testing"
)

var (
	g  []string
	gi []interface{}
)

func init() {
	for i := 0; i < 100000000; i++ {
		g = append(g, "11111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111")
	}
	for i := 0; i < 100000000; i++ {
		gi = append(gi, "11111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111")
	}
}

func TestJoin(t *testing.T) {
	x := Join(g...)

	fmt.Println("rlt:", x)
}

func BenchmarkJoin(b *testing.B) {
	b.StopTimer()
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		Join(g...)
	}
}

func BenchmarkJoinObj(b *testing.B) {
	b.StopTimer()
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		JoinObj(gi...)
	}
}

func BenchmarkJoin2(b *testing.B) {
	b.StopTimer()
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("%s%s%s", g[0], g[1], g[2])
	}
}

func BenchmarkJoin3(b *testing.B) {
	b.StopTimer()
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		var x string
		for i := 0; i < len(g); i++ {
			x = x + g[i]
		}
	}
}

// 基准测试结论：
// 在字符串数量极少(不超过3个)并且每个字符串很短(1 2个)的情况下，加号拼接最快
// 其余情况下均为buffer最快，平均性能是fmt.Sprintf的近三倍
// 在字符串很多，每个字符串很长的情况下，加号拼接的性能无法忍受(字符串数量为1000时，所需时间是fmt.Sprintf的30W倍))
