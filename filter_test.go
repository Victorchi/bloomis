package bloomis

import (
	"fmt"
)

func ExampleKBits() {
	f := NewFilter("test", 335477044, 23)
	fmt.Println(f.bits([]byte("hello")))
	// Output:
	// [93375935 300021868 195851017 67019906 298326099 169494988 40663877 271970070 143138959 38968108 245614041 116782930 12612079 219258012 115087161 321733094 192901983 88731132 295377065 191206214 62375103 269021036 164850185]
}