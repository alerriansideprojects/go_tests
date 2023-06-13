package hello

func Hello(name, language string) string {
	if name == "" {
		name = "Go!"
	}

	return getHelloPrefix(language) + name
}

func getHelloPrefix(language string) (prefix string) {
	switch language {
	case "Spanish":
		prefix = "Hola, "
	case "French":
		prefix = "Bonjour, "
	case "German":
		prefix = "Hallo, "
	default:
		prefix = "Hello, "
	}

	return prefix
}