package e

func e() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		c1 <- ""
	}()

	go func() {
		c2 <- ""
	}()

	select {
	case m := <-c1:
		m += ""
		select {
		case n := <-c2:
			n += ""
		}
	case m := <-c2:
		m += ""
	}
}
