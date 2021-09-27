package countAndSay

import "fmt"

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}

	last := countAndSay(n-1)

	result := ""
	i, count:=0, 0
	for ; i<len(last); i++  {
		if i == 0 {
			count++
			continue
		}

		if last[i] == last[i-1] {
			count++
		} else {
			result = fmt.Sprintf("%s%d%c",  result, count, last[i-1])
			count = 1
		}
	}

	if count > 0 {
		result = fmt.Sprintf("%s%d%c",  result, count, last[i-1])
	}
	return result
}