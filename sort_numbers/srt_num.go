package sort_numbers

func SortNums(numbers []int) {
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
