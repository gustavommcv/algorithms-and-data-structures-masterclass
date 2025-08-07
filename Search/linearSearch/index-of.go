package index_of

func indexOf[T comparable](arr []T, input T) int {
	for i, element := range arr {
		if element == input {
			return i
		}
	}

	return -1
}
