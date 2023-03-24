package main

import "fmt"

func main() {
	groupCity := map[int][]string{
		10:   []string{"Село", "Деревня", "ПГТ"},  // города с населением 10-99 тыс. человек
		100:  []string{"Город", "Станица"},        // города с населением 100-999 тыс. человек
		1000: []string{"Мегаполис", "Миллионник"}, // города с населением 1000 тыс. человек и более
	}

	cityPopulation := map[string]int{
		"Город":     600,
		"Станица":   102,
		"Село":      30,
		"Мегаполис": 1000,
	}
	for z, _ := range cityPopulation {
		isInMap := false
		for _, v := range groupCity[100] {
			if z == v {
				isInMap = true
				break 
			}
		}
		if isInMap == false {
			delete(cityPopulation, z)
		} 
	}
	fmt.Println(cityPopulation)
}

/*
 * Группировка городов по численности населения в тысячах человек
 * от 10 до 100, от 100 до 1000 и более 1000:
 * groupCity map[int][]string{
 *	 10: []string{...},
 *	 100: []string{...},
 *	 1000: []string{...},
 * }
 *
 * Население городов в тысячах человек:
 * cityPopulation map[string]int{...}
 */
