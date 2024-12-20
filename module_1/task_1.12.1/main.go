package main

import "fmt"

/* 
На первом этапе на стандартный ввод подается 10 целых положительных чисел, которые должны быть записаны в порядке ввода в массив из 10 элементов. 
Тип чисел, входящих в массив, должен соответствовать минимально возможному целому беззнаковому числу. 
Имя массива который вы должны сами создать workArray (условие обязательное). 
Для чтения из стандартного ввода уже импортирован пакет fmt. 
На втором этапе на стандартный ввод подаются еще 3 пары чисел - индексы элементов этого массива, 
которые требуется поменять местами (если такая пара чисел 3 и 7, значит в массиве элемент с 3 индексом нужно поменять местами с элементом, индекс которого 7).

Элементы полученного массива должны быть выведены через пробел на стандартный вывод. 
Далее автоматически будет проведена проверка используемых типов, результат которой будет добавлен к вашему ответу.
*/

func main(){ 
	
	workArray := [10]uint8{}
	
	for i := 0; i < 10; i++ {
		var num uint8
		fmt.Scan(&num)
		workArray[i] = num
	}
	
	for i := 0; i < 6; i+=2 {
		var fIdx, sIdx uint8
		fmt.Scan(&fIdx, &sIdx)
		workArray[fIdx], workArray[sIdx] = workArray[sIdx], workArray[fIdx]
	}
	
	for _, num := range workArray {
		fmt.Print(num, " ")
	}
}

/*
Альтернативное решение

func main(){
	var workArray [10] uint8
	var in, in2 uint8
	
	for i,_ := range workArray{
	    fmt.Scan(&workArray[i])
	    }
	
	for i:=0;i<3;i++{
	    fmt.Scan(&in, &in2)
	    workArray[in ],workArray[in2]=workArray[in2], workArray[in]
	    }

	for _,i := range workArray{
	    fmt.Print(i, " ",)}
}
*/