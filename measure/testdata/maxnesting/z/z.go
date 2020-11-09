package z

func z() {
	hoge := func() {
		if 1 == 1 {

		} else if 1 == 2 {
			var i interface{}
			switch i := i.(type) {
			case int:
				switch i {
				case 0:
					c1 := make(chan string)
					c2 := make(chan string)
					go func() {
						c1 <- ""
					}()
					go func() {
						c2 <- ""
					}()
					for i = 0; i < 2; i++ {
						select {
						case m := <-c1:
							go func() {}()
						case m := <-c2:
						}
					}
				case 1:
				default:
				}
			case string:
			default:
			}

		} else {

		}
	}

	hoge()
}
