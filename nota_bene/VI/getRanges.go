package vi

func GetRanges(n []int, target int) []int {
	hi, lo := len(n)-1, 0

	k := 0
	for lt, rt := 0, len(n)-1; lt < rt; {
		k++
		i := (rt + lt) / 2
		if n[i] == target && (i == 0 || n[i-1] != target) {
			lo = i
			break
		}
		if n[i] < target {
			lt = i
		} else {
			rt = i
		}
	}

	k = 0
	for lt, rt := 0, len(n)-1; lt < rt; {
		k++
		i := (rt + lt) / 2
		if n[i] == target && (i == len(n)-1 || n[i+1] != target) {
			hi = i
			break
		}
		if n[i] > target {
			rt = i
		} else {
			lt = i
		}
	}
	return []int{lo, hi}
}

func NaiveGetRanges(n []int, target int) []int {
	low, hi := 0, len(n)-1
	for n[low] != target || n[hi] != target {
		if n[low] < target {
			low++
		}
		if n[hi] > target {
			hi--
		}
	}
	return []int{low, hi}
}
