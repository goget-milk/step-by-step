package main

import "fmt"

/*
Идёт k-я секунда суток. Определите, сколько целых часов h и целых минут m прошло с начала суток. Например, если
k=13257=3*3600+40*60+57,
то h=3 и m=40.
*/

func main() {

	var sec int
	fmt.Scan(&sec)
	secInHour := 3600
	secInMin := 60
	hour := sec / secInHour
	min := (sec - hour*3600) / secInMin
	fmt.Printf("It is %d hours %d minutes.", hour, min)

}

/*
Альтернативное решение

func main() {
  var seconds int
  fmt.Scan(&seconds)
  fmt.Print("It is ", seconds/3600, " hours ", seconds/60%60, " minutes.")
}
*/
