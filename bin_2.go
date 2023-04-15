// реализован бин поиск по ответу, макс. кол-во требуемых изменений r = k/2, где k - желаемая длина подотрезка.
// l = 0. Вспомогательная функция для бин поиска принимает строку, количество возможных изменений,
// возвращает максимально возможную длину подстроки, состоящей из одинаковых элементов, после изменений.

package main

import "fmt"

func main() {
	s := "0011100"
	n := 7
	k := 4

	fmt.Println(MinChanges(s, n, k))
}

// принимает строку, ее длину и длину желаемого подотрезка. Внутри функции происходит бин поиск по ответу.
func MinChanges(s string, n, k int) int {
	ans := 0

	l := 0
	r := k / 2

	for l <= r {
		mid := (l + r) >> 1

		if LongestSubstr(s, mid) >= k {
			r = mid - 1
			ans = mid
		} else {
			l = mid + 1
		}
	}

	return ans
}

// принимает строку, количество возможных изменений, возвращает максимально возможную длину подстроки, состоящей из одинаковых элементов
func LongestSubstr(s string, changes int) int {
	cntZero := 0
	cntOne := 0

	l := 0
	ans := 0

	for i := 0; i < len(s); i++ {

		if s[i] == '0' {
			cntZero++
		} else {
			cntOne++
		}

		for cntOne > changes && cntZero > changes {
			if s[l] == '0' {
				cntZero--
			} else {
				cntOne--
			}

			l++
		}

		if cntOne+cntZero > ans {
			ans = cntOne + cntZero
		}

	}

	return ans
}
