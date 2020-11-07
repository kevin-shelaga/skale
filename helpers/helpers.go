package helpers

//ProcessFlags returns values for the given flag
func ProcessFlags(args []string, flagToFind string) []string {
	var result []string
	for i := 0; i < len(args); i++ {
		if args[i] == flagToFind {
			result = append(result, args[i+1])
		}
	}
	return result
}
