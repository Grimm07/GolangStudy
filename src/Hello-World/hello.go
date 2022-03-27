package Hello

// simple hello world program
func Hello(name string) (ret string) {
	if name == "" {
		ret = "Hello World!"
	} else {
		ret = "Hello " + name + "!"
	}
	return
}
