package chapter_5

func SpiralOrdering(data [][]int) []int {
	var res []int
	for offset := 0; offset < len(data[0])/2+1; offset++ {
		for i := offset; i < len(data[offset])-offset; i++ {
			res = append(res, data[offset][i])
		}
		for i := 1 + offset; i < len(data[offset])-offset; i++ {
			res = append(res, data[i][len(data[offset])-1-offset])
		}
		for i := len(data[offset]) - offset - 2; i > offset; i-- {
			res = append(res, data[len(data[offset])-1-offset][i])
		}
		for i := len(data[offset]) - offset - 1; i > offset; i-- {
			res = append(res, data[i][offset])
		}
	}
	return res
}
