/*
	Не использовал явно 2 указателя. Идея решения: фиксируем индекс i и двигаем его. Каждый раз, когда следующий элемент равен
	текущему (s[i+1]==s[i]), инкрементируем i и внутренний счетчик cnt, который обозначает, сколько было последовательно одинаковых символов.
	Когда следующий элемент отличается, записываем символ и итоговый count.

	Не понял тут двух указателей. Может быть, тут можно обозначить зафиксированный индекс i за l, и когда следующий элемент будет равен,
	оставлять l на месте, а вправо двигать r.

*/

package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := "abbcdddccaaa"
	n := 12

	fmt.Println(StringPack(s, n)) // вывод a1b2c1d3c2a3
}

// функция распаковки строки
func StringPack(s string, n int) string {

	ans := ""

	for i := 0; i < n; i++ {

		// cnt сразу один, т.к. текущий символ уже встречается один раз
		cnt := 1

		// проверка, чтобы избежать выхода индекса за пределы строки.
		if i < n-1 {
			// пока следующий символ равен, прибавляем его счетчик и двигаем i
			for s[i+1] == s[i] {

				cnt++
				i++

				// если i уже стоит на последнем элементе, то выходим из цикла. Сделано, чтобы избежать выхода индекса за пределы.
				if i == n-1 {
					break
				}
			}
		}

		// функция перевода из int в string. Перевожу count в string
		cntStr := strconv.Itoa(cnt)

		// к ответу прибавляю символ и его count
		ans += string(s[i]) + cntStr
	}

	return ans
}
