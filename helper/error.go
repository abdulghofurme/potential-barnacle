package helper

func PanicIfErrof(err error) {
	if err != nil {
		panic(err)
	}
}
