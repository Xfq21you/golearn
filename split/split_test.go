package split

import "testing"

/*
	func TestSplit(t *testing.T) {
		got := Split("a:b:c", ":")
		want := []string{"a", "b", "c"}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("excepted:%v, got:%v", want, got)
		}
	}

	func TestMoreSplit(t *testing.T) {
		got := Split("abcd", "bc")
		want := []string{"a", "d"}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("except:%v, got:%v", want, got)
		}
	}

	func TestSplit(t *testing.T) {
		type test struct {
			input string
			sep   string
			want  []string
		}
		tests := map[string]test{ // 测试用例使用map存储
			"simple":      {input: "a:b:c", sep: ":", want: []string{"a", "b", "c"}},
			"wrong sep":   {input: "a:b:c", sep: ",", want: []string{"a:b:c"}},
			"more sep":    {input: "abcd", sep: "bc", want: []string{"a", "d"}},
			"leading sep": {input: "枯藤老树昏鸦", sep: "老", want: []string{"枯藤", "树昏鸦"}},
		}
		for name, tc := range tests {
			t.Run(name, func(t *testing.T) {
				got := Split(tc.input, tc.sep)
				if !reflect.DeepEqual(got, tc.want) {
					t.Errorf("excepted:%#v, got:%#v", tc.want, got)
				}
			})
		}
	}
func ExampleSplit() {
	fmt.Println(split.Split("a:b:c", ":"))
	fmt.Println(split.Split("枯藤老树昏鸦", "老"))
	// Output:
	// [a b c]
	// [ 枯藤 树昏鸦]
}
*/
func BenchmarkSplit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Split("枯藤老树昏鸦", "老")
	}
}
