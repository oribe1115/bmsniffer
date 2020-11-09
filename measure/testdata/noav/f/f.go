package f

import "fmt"

func f() {
	var j interface{}
	switch j := j.(type) {
	case int:
		fmt.Println(j)
	case float32:
		fmt.Println(j)
	case string:
		fmt.Println(j)
	}
}
