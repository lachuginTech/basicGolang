package main

import (
	"errors"
	"fmt"
	"math"
)

const IMTPower = 2

func main() {
	fmt.Println("Программа для расчета индекса массы тела")

	for {
		weight, height := getUserInput()

		bmi, err := calculateIMT(weight, height)
		if err != nil {
			fmt.Println(err)
			continue
		}

		outputResult(bmi)

		if !checkRepeatCalculation() {
			break
		}
	}
}

// outputResult выводит результат расчета BMI и классифицирует его
func outputResult(bmi float64) {
	fmt.Printf("Ваш индекс массы тела: %.1f", bmi)
	switch {
	case bmi < 16:
		fmt.Println(" - Выраженный вес недостаточен")
	case bmi >= 16 && bmi < 18.5:
		fmt.Println(" - Недостаточная масса тела")
	case bmi >= 18.5 && bmi < 25:
		fmt.Println(" - Нормальная масса тела")
	case bmi >= 25 && bmi < 30:
		fmt.Println(" - Избыточная масса тела")
	case bmi >= 30:
		fmt.Println(" - Ожирение")
	}
}

// calculateIMT рассчитывает индекс массы тела (BMI)
func calculateIMT(weight float64, height float64) (float64, error) {
	if height <= 0 || weight <= 0 {
		return 0, errors.New("Некорректный вес или рост")
	}
	bmi := weight / math.Pow(height/100, IMTPower)
	return bmi, nil
}

// getUserInput запрашивает ввод данных у пользователя и валидирует их
func getUserInput() (float64, float64) {
	var height, weight float64

	for {
		fmt.Print("Введите рост в сантиметрах: ")
		fmt.Scan(&height)
		fmt.Print("Введите вес в килограммах: ")
		fmt.Scan(&weight)

		if height > 0 && weight > 0 {
			break
		} else {
			fmt.Println("Рост и вес должны быть положительными числами. Повторите ввод.")
		}
	}

	return weight, height
}

// checkRepeatCalculation проверяет, хочет ли пользователь повторить расчет
func checkRepeatCalculation() bool {
	var repeatCalculation string
	fmt.Print("Повторить расчет? (y/n): ")
	fmt.Scan(&repeatCalculation)

	return repeatCalculation == "y" || repeatCalculation == "Y"
}
