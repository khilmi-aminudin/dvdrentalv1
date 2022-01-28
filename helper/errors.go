package helper

func PanicIfError(err error) {
	if err != nil {
		// panic(err.Error())
		LoggerInit().Panic(err)
	}
}

func FatalError(err error) {
	if err != nil {
		LoggerInit().Fatal(err)
	}
}
