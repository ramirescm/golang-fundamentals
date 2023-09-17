package main

func Repeat(value string, total int) string {
	var rep string
	for i := 1; i <= total; i++ {
		rep += value
	}
	return rep
}
