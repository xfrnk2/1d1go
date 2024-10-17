package main

import (
	"bufio"
	"fmt"
	"os"
)

func exec(cmd int, i int, x int, trains []int) {

	switch cmd {
	case 1:
		trains[i] |= 1 << x
	case 2:
		trains[i] &= ^(1 << x)
	case 3:
		trains[i] <<= 1
		trains[i] &= ((1 << 21) - 1)
	case 4:
		trains[i] >>= 1
		trains[i] &= ^1
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var N, M int
	fmt.Fscanln(reader, &N, &M)

	trains := make([]int, N+1)

	for i := 0; i < M; i++ {
		var a, b, c int
		a = 0
		b = 0
		c = 0

		fmt.Fscanln(reader, &a, &b, &c)
		exec(a, b, c, trains)
	}

	vis := make([]bool, 1<<21)
	res := 0
	for i := 1; i <= N; i++ {
		if !vis[trains[i]] {
			res++
		}
		vis[trains[i]] = true
	}

	fmt.Println(res)
}
