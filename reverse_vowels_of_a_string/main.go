package main

func reverseVowels(s string) string {
	m := make(map[rune]struct{}, 10)
	m['a'] = struct{}{}
	m['i'] = struct{}{}
	m['u'] = struct{}{}
	m['e'] = struct{}{}
	m['o'] = struct{}{}
	m['A'] = struct{}{}
	m['I'] = struct{}{}
	m['U'] = struct{}{}
	m['E'] = struct{}{}
	m['O'] = struct{}{}
	var (
		vowelsi []int
		arr     []rune
	)
	for i, r := range s {
		arr = append(arr, r)
		if _, ok := m[r]; ok {
			vowelsi = append(vowelsi, i)
		}
	}
	for i := 0; i < len(vowelsi)/2; i++ {
		arr[vowelsi[i]], arr[vowelsi[len(vowelsi)-i-1]] = arr[vowelsi[len(vowelsi)-i-1]], arr[vowelsi[i]]
	}
	return string(arr)
}

//Input: "leet code"
//Output: "leot cede"
