package main

import "fmt"

func main() {
	m := map[string]interface{}{
		"b": true,
		"i": 42,
		"f": 3.14,
		"s": []byte{20, 13, 7, 9}[1:3],
		"m": map[int]string{1: "i", 5: "v"},
	}

	for k, v := range m {
		fmt.Printf("%s: %v\n", k, v)

		switch v.(type) {
		case int:
			fmt.Printf("\t%#x\n", v.(int))
		case []byte:
			s := v.([]byte)
			fmt.Println("\t", len(s), cap(s))
		}
	}
}
