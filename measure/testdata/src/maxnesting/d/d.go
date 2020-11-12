package d

func d() {
	var i interface{}
	switch i := i.(type) {
	case int:
		i++
		var j interface{}
		switch j := j.(type) {
		case int:
			j++
		case string:
			j += ""
		default:
			j = nil
		}
	case string:
		var j interface{}
		switch j := j.(type) {
		case int:
			j++
		case string:
			var k interface{}
			switch k := k.(type) {
			case int:
				k++
			case string:
				k += ""
			default:
				k = nil
			}
		default:
			j = nil
		}
	default:
		i = nil
	}
}
