package main

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	var length int
	for i := 0; i < len(strs[0]); i++ {
		var ok int
		for j := 1; j < len(strs); j++ {
			k := strs[j]
			if len(k)-1 < i || strs[0][i:i+1] != k[i:i+1] {
				break
			}
			ok++
		}
		if ok != len(strs)-1 {
			break
		}
		length++
	}
	if length == 0 {
		return ""
	}
	return strs[0][0:length]
}
