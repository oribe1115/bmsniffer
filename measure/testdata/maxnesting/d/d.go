package d

func d() {
	var i interface{}
	switch i := i.(type) {
	case int:
		var j interface{}
		switch j := j.(type) {
		case int:
		case string:
		default:
		}
	case string:
		var j interface{}
		switch j := j.(type) {
		case int:
		case string:
			var k interface{}
			switch k := k.(type) {
			case int:
			case string:
			default:
			}
		default:
		}
	default:
	}
}
