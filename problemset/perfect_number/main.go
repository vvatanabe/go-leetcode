package main

func checkPerfectNumber(num int) bool {
	var pn int
	for i := 1; i*i < num; i++ {
		if num%i == 0 {
			pn = pn + i + (num / i)
		}
	}
	return pn-num == num
}

// time over
func checkPerfectNumber2(num int) bool {
	var pn int
	for i := 1; i < num; i++ {
		if num%i != 0 {
			continue
		}
		pn += i
		if pn > num {
			return false
		}
	}
	return pn == num
}
