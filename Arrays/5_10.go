package chapter_5

func PermuteArrayElements(array, perms []int) []int {
	for i := range array {
		array[i], array[perms[i]] = array[perms[i]], array[i]
		perms[i], perms[perms[i]] = perms[perms[i]], perms[i]
	}
	return array
}
