package must

func Do0(err error) {
	if err != nil {
		panic(err)
	}
}

func Do[T any](v T, err error) T {
	Do0(err)
	return v
}
