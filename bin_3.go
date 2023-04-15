/*
	Использую бин поиск по ответу. Задача аналогичная той, которую разбирали на лекции. Бинарю значение mult и ищу кол-во произведений i*j,
	которое меньше mid. Сравниваю кол-во этих произведений с k. Последнее кол-во произведений, которое удовлетворяет >=k будет как раз равно k и это будет равнятся k-му произведению.
*/

package main

import "fmt"

func main() {
	n := 5
	k := 9

	fmt.Println(FindMult(n, k))
}

func FindMult(n, k int) int {

	l := 0
	r := 10000000000 + 1

	ans := 0

	for l <= r {
		mid := (l + r) >> 1

		if Num(n, mid) >= k {
			r = mid - 1
			ans = mid
		} else {
			l = mid + 1
		}
	}

	return ans
}

// принимает число n и аргумент mult. Возвращает, сколько есть пар i*j <= mult, где 1 <= i,j <= n
func Num(n, mult int) int {
	l := 1
	r := n

	ans := 0

	for r >= 1 {
		for r*l <= mult {
			l++
		}

		ans += l - 1

		r--
	}

	return ans
}
