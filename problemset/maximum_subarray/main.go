package main

// ## 問題
// 整数の配列として`nums`が与えられた場合、最大の和を持つ連続した部分配列を探してその和を返す。
//
// ## 解答
// - 整数の配列のnumsの先頭を最大値maxと現在値curに記録する。
// - 先頭の値を除いたnumsをループする。
//   - 現在値curが0未満であればnumsの値vを現在値curに上書きする。
//   - 現在値curが0以上であればnumsの値vを現在値curに加算する。
//   - 現在値curが最大値maxより大きければ最大値maxに現在値curを上書きする。
//   - ここまでの処理を全てのnumsの末尾の値まで適用する。
// - 最後に最大値maxを返す。

func maxSubArray(nums []int) int {
	max := nums[0]
	cur := nums[0]
	for _, v := range nums[1:] {
		if cur < 0 {
			cur = v
		} else {
			cur += v
		}
		if max < cur {
			max = cur
		}
	}
	return max
}
