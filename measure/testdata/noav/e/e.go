package e

func e() (int, error) {
	f := func() (int, error) {
		return 0, nil
	}

	first, err := f()
	second, err := f()

	return first + second, err
}
