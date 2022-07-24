package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var (
		s string
		a string
		b string
	)
	fmt.Scanln(&s)
	flag := 0
	for i := 0; i < len(s); i++ {
		if s[i] != '=' && flag == 0 {
			a += string(s[i])
			continue
		}
		flag = 1
		if s[i] == '=' {
			continue
		}
		b += string(s[i])
	}
	if b[0] != '-' && b[0] != '+' {
		b = "+" + b
	}
	if a[0] != '-' && a[0] != '+' {
		a = "+" + a
	}
	for _, str := range b {
		st := string(str)
		if st == "-" {
			a += "+"
		} else if st == "+" {
			a += "-"
		} else {
			a += st
		}
	}
	a = strings.Replace(a, "-", "+-", -1)
	var k rune
	for _, k = range a {
		if k >= 'a' && k <= 'z' {
			break
		}
	}
	ks := string(k)
	strArr := strings.Split(a, "+")
	var (
		aArr []string
		arr  []string
	)
	for _, va := range strArr {
		if strings.Contains(va, ks) {
			aArr = append(aArr, va)
		} else {
			arr = append(arr, va)
		}
	}
	res := 0
	for _, va := range arr {
		n, _ := strconv.Atoi(va)
		res += n
	}
	res = -res
	bs := 0
	for _, va := range aArr {
		if va == ks {
			bs += 1
		} else if va == "-"+ks {
			bs += -1
		} else {
			va = strings.TrimRight(va, ks)
			n, _ := strconv.Atoi(va)
			bs += n
		}
	}
	ans := float64(res) / float64(bs)
	if ans == 0 {
		ans = 0
	}
	fmt.Printf("%s=%.3f\n", ks, ans)
}
