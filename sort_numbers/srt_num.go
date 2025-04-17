package sort_numbers

func SortNums(numbers []int) {
	for _, _ = range numbers {
		var previous int
		for i, _ := range numbers {
			if i == len(numbers)-1 {
				// если дошли до предела, то уходим
				break
			} else if i == 0 {
				previous = numbers[i]
			}
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

//TIP See GoLand help at <a href="https://www.jetbrains.com/help/go/">jetbrains.com/help/go/</a>.
// Also, you can try interactive lessons for GoLand by selecting 'Help | Learn IDE Features' from the main menu.
