package student

func Compare(a, b string) int {
	ar := []rune(a)
	br := []rune(b)
	for i := 0; i < len(ar) && i < len(br); i++ {
		if ar[i] != br[i] {
			if ar[i] < br[i] {
				return -1
			}
			return 1
		}
	}
	if len(ar) < len(br) {
		return -1
	}
	if len(ar) > len(br) {
		return 1
	}
	return 0
}
