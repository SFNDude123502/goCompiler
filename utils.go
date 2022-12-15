package main

func contains(list []string, val string) bool {

	for _, v := range list {
		if v == val {
			return true
		}
	}

	return false
}

func merge(x []string, y []string) []string {
	for _, v := range x {
		if !contains(y, v) {
			y = append(y, v)
		}
	}
	return y
}
