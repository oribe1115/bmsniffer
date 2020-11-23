package e

func e() {
	i := 0

	func() {
		if i == 0 {
			i++
		} else {
			i--
		}
	}()
}
