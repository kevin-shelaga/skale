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

//IsDryRun returns true/false if d flag is present
func IsDryRun(args []string) bool {
	var result bool = false
	for i := 0; i < len(args); i++ {
		if args[i] == "d" {
			result = true
			break
		}
	}
	return result
}
