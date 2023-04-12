package main

import "fmt"



type TaskStruct struct { // объявляем структуру
	On    bool
	Ammo  int
	Power int
}
// метод принято именовать по первой букве названия структуру, для которой он реализован
  // если TaskStruct, то t

func (t *TaskStruct) Shoot() bool {
	if !t.On { // если On выключен, то возвращаем false по заданию
		return false
	}
	if t.Ammo > 0 { // проверяем наличие Ammo и убавляем на 1
		t.Ammo--
		return true
	}
	return false
}

func (t *TaskStruct) RideBike() bool {
	if !t.On {
		return false 
	}
	if t.Power > 0 {
		t.Power--  // все аналогично, только теперь проверяем Power
		return true
	}
	return false
}
func main() {
    testStruct := new(TaskStruct) // по заданию нужно дать указатель на структуру (можно через & и через new(), здесь работает через new, так как структура)
	// скобку в конце не ставим по задднию
// вот решение
}