package romanToInt

func RomanToInt(s string) int {

	var mappingList = map[string]int{
		"I":  1,
		"IV": 4,
		"V":  5,
		"IX": 9,
		"X":  10,
		"XL": 40,
		"L":  50,
		"XC": 90,
		"C":  100,
		"CD": 400,
		"D":  500,
		"CM": 900,
		"M":  1000,
	}

	result, label := 0, ""
	for i := 0; i < len(s); {
		if i+1 < len(s) {
			label = s[i : i+2]
			if unit, ok := mappingList[label]; ok {
				result += unit
				i += 2
				continue
			}
		}

		label = s[i : i+1]
		unit := mappingList[label]
		result += unit
		i++
	}
	return result
}
