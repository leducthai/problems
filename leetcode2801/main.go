package main

import (
	"fmt"
	"math/big"
	"strconv"
)

func countSteppingNumbers(low string, high string) int {
	mod := int(10e8 + 7)
	dim := new(big.Int)
	dim.SetString(low, 10)
	dim.Sub(dim, big.NewInt(1))
	low = dim.String()
	memo := make(map[int]map[int]int)
	lh := len(high)
	ll := len(low)

	for i := 0; i < 10; i++ {
		memo[i] = map[int]int{0: 1}
	}

	for i := 1; i < lh; i++ {
		for j := 0; j < 10; j++ {
			if j == 9 {
				memo[j][i] = memo[j-1][i-1]
			} else if j == 0 {
				memo[j][i] = memo[j+1][i-1]
			} else {
				memo[j][i] = (memo[j-1][i-1] + memo[j+1][i-1]) % mod
			}
		}
	}

	// steps from 0 to high
	fnum, _ := strconv.Atoi(high[0:1])
	hstep := find_steps(memo, high, 1, fnum)
	for i := 1; i < fnum; i++ {
		hstep = (hstep + memo[i][lh-1]) % mod
	}
	for i := lh - 2; i >= 0; i-- {
		for j := 1; j < 10; j++ {
			hstep = (hstep + memo[j][i]) % mod
		}
	}

	// steps from 0 to low
	fnum, _ = strconv.Atoi(low[0:1])
	if fnum == 0 {
		return hstep % mod
	}
	lstep := find_steps(memo, low, 1, fnum)
	for i := 1; i < fnum; i++ {
		lstep = (lstep + memo[i][ll-1]) % mod
	}
	for i := ll - 2; i >= 0; i-- {
		for j := 1; j < 10; j++ {
			lstep = (lstep + memo[j][i]) % mod
		}
	}

	return (hstep + mod - lstep) % mod
}

func find_steps(mm map[int]map[int]int, tar string, pos int, pre int) int {
	l := len(tar)

	mod := int(10e8 + 7)
	if pos == l {
		return 1
	}

	num_at_pos, _ := strconv.Atoi(tar[pos : pos+1])

	rt := 0

	if num_at_pos > pre+1 {
		rt += mm[pre+1][l-pos-1]
	} else if num_at_pos == pre+1 {
		rt += find_steps(mm, tar, pos+1, num_at_pos)

	}
	if num_at_pos > pre-1 && pre-1 >= 0 {
		rt += mm[pre-1][l-pos-1]
	} else if num_at_pos == pre-1 && pre-1 >= 0 {
		rt += find_steps(mm, tar, pos+1, num_at_pos)

	}

	return rt % mod
}

func main() {
	l := "999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999"
	h := "9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999"

	answer := countSteppingNumbers(l, h)
	fmt.Println(answer)
}
