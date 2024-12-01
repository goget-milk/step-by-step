package main

import "fmt"

/*
Цифровой корень натурального числа — это цифра, полученная в результате итеративного процесса суммирования цифр,
на каждой итерации которого для подсчета суммы цифр берут результат, полученный на предыдущей итерации.
Этот процесс повторяется до тех пор, пока не будет получена одна цифра.

Например цифровой корень 65536 это 7 , потому что 6+5+5+3+6=25 и 2+5=7 .

По данному числу определите его цифровой корень.
*/

func main() {

	var num, sum, res int
	fmt.Scan(&num)
	for i := 100000000; i >= 1; i /= 10 {
		if num/i == 0 {
			continue
		}
		sum += num / i % 10
	}
	for i := 100; i >= 1; i /= 10 {
		if sum/i == 0 {
			continue
		}
		res += sum / i % 10
	}
	fmt.Println(res)

}

/*
Альтернативное решение

func main()  {
	var a int
	fmt.Scan(&a)
	fmt.Println((a-1)%9+1)
}
*/

/*
Альтернфтивное решение

func main() {
  var a,num int
  fmt.Scan(&a)
  for a>9{
    num = a%10
    a /= 10
    a = a + num
  }
  fmt.Print(a)
}
*/