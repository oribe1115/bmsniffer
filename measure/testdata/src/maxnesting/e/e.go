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
		select {
		case n := <-c2:
		}
	case m := <-c2:
	}
}
