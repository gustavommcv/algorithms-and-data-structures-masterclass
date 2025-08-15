package naiveStringSearch

func naiveStringSearch(str string, target string) int {
	count := 0

	for i := range str {
		for j := range target {
			if str[i+j] != target[j] {
				break
			}

			if j == len(target)-1 {
				count++
			}
		}
	}

	return count
}
