package main

import "fmt"

// Wagon представляет вагон с лампочкой
type Wagon struct {
	lightOn bool
}

// toggleLight переключает лампочку
func (w *Wagon) toggleLight() {
	w.lightOn = !w.lightOn
}

func countWagons(train []Wagon) int {
	if len(train) == 0 {
		return 0
	}

	// Устанавливаем начальное состояние вагона 0 как метку
	startState := train[0].lightOn
	steps := 0
	current := 0

	// Выключаем лампочку в первом вагоне как стартовую метку
	if startState {
		train[0].toggleLight()
	}
	startState = train[0].lightOn // Новая метка — выключено

	for {
		steps++
		current = steps % len(train) // Кольцевое перемещение
		// Если вернулись в вагон 0 с начальным состоянием
		if current == 0 && train[0].lightOn == startState {
			break
		}
	}

	return steps
}

func main() {
	// train := []Wagon{
	// 	{false},
	// 	{true},  // Вагон 0
	// 	{false}, // Вагон 1
	// 	{true},  // Вагон 2
	// 	{true},
	// 	{false}, // Вагон 3
	// 	{false}, // Вагон 3
	// 	{false}, // Вагон 3
	// 	{true},
	// }
	// fmt.Println("Количество вагонов:", countWagons(train)) // 4

	fmt.Println(5 % 4)
}
