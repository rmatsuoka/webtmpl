package must

func Nil(err error) {
	if err != nil {
		panic(err)
	}
}

func Do[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}
