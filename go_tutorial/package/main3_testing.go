package mylib

import "testing"

var Debug bool = true

//テスト用の関数を作る
func TestAverage(t *testing.T) {
	//Debugがtrueの場合テストをスキップする
	if Debug {
		t.Skip("Skip Reason")
	}

	//Averageの実行結果が3でない場合はエラーを出力する
	v := Average([]int{1, 2, 3, 4, 5})
	if v != 3 {
		t.Error("Expected 3, got", v)
	}
}