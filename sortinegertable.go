package piscine

func SortIntegerTable(table []int) {
	var b int
	var c bool
	for !c {
		c = true
		for i := 0; i < len(table); i++ {
			if i < len(table)-1 {
				if table[i] > table[i+1] {
					b = table[i]
					table[i] = table[i+1]
					table[i+1] = b
					c = false
				}
			}
		}
	}
}
