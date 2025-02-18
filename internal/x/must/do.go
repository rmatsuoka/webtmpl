package must

func PanicIf(err error) {
	if err != nil {
		panic(err)
	}
}

func Do[T any](v T, err error) T {
	PanicIf(err)
	return v
}
