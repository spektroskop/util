package util

func MapStrings(xs []string, f func(s string) string) (r []string) {
	for _, s := range xs {
		r = append(r, f(s))
	}

	return
}

func HasString(s string, xs []string) bool {
	for _, x := range xs {
		if s == x {
			return true
		}
	}

	return false
}
