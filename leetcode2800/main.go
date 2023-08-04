package main

import "fmt"

func minimumString(a string, b string, c string) string {
	answer := ""
	perm([]string{a, b, c}, merger, 0, &answer)
	return answer
}

func perm(a []string, f func([]string, *string), i int, s *string) {
	if i > len(a) {
		f(a, s)
		return
	}
	perm(a, f, i+1, s)
	for j := i + 1; j < len(a); j++ {
		a[i], a[j] = a[j], a[i]
		perm(a, f, i+1, s)
		a[i], a[j] = a[j], a[i]
	}
}

func merger(ar []string, as *string) {
	rt := ar[0]

	l := bi_search(ar[0], ar[1])
	rt += ar[1][l:]

	l = bi_search(rt, ar[2])
	rt += ar[2][l:]

	if (len(rt) == len(*as) && rt < *as) || len(rt) < len(*as) || *as == "" {
		*as = rt
	}
}

func bi_search(a string, b string) int {
	lea := len(a)
	leb := len(b)
	m := 0
	if lea < leb {
		m = lea
	} else {
		m = leb
	}

	if lea > leb {
		for i := leb; i <= lea; i++ {
			if a[i-leb:i] == b {
				return leb
			}
		}
	}

	for a[lea-m:] != b[:m] {
		m -= 1
	}

	return m
}

func main() {
	a := "ab"
	b := "a"
	c := "c"

	st := minimumString(a, b, c)

	fmt.Println(st)
	fmt.Println("abc" < "abac")
}
