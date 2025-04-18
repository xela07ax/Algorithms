package sort_numbers

func SortMyArtifice(numbers []int) {
	// необходимо расставить по местам только n – 2 числа,
	// но если придет на вход 2 или 3 числа,
	// то сортировка не сможет начаться, потому делаем n - 1
	for globalNext := 1; globalNext < len(numbers); globalNext++ {
		var previous int = numbers[0]
		for i := 0; i < len(numbers)-1; i++ { // необходимо сравнить только n – 1 раз
			next := numbers[i+1]
			if previous > next {
				numbers[i] = previous
				previous = next
			} else {
				numbers[i] = next
			}
		}
		numbers[len(numbers)-1] = previous
	}
	return
}

func SortBuble(numbers []int) {
	for globalNext := 1; globalNext < len(numbers); globalNext++ { //n – 1 операций
		// надо менять местами от большего к меньшему
		for i := 0; i < len(numbers)-1; i++ { // необходимо сравнить только n – 1 раз
			if numbers[i] < numbers[i+1] {
				numbers[i], numbers[i+1] = numbers[i+1], numbers[i]
			}
		}
	}
	return
}

func SortBlock(numbers []int) {
	a := make([]int, 101) // инициализируется равной 0
	var i, j, t int
	for i, t = range numbers { // считываем каждое число в переменную t
		a[t]++ // заполняем массив
	}
	newArr := make([]int, 0, len(numbers))
	for i = 100; i >= 0; i-- { // поочередно оцениваем a[0]~a[10]
		for j = 1; j <= a[i]; j++ {
			newArr = append(newArr, i) // выводим результаты
		}
	}
	for i, t = range newArr {
		numbers[i] = t
	}
}
