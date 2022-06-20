package util

func GetStringFromArrayOrDefault(args []string, index int, d string) string {
	if len(args) > index {
		return args[index]
	}
	return d
}
