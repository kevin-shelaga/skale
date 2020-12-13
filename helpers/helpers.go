package helpers

//ProcessFlags returns values for the given flag
func ProcessFlags(args []string, flagToFind string) []string {
	var result []string
	for i := 0; i < len(args); i++ {
		if args[i] == flagToFind || args[i] == "-"+flagToFind {
			if !contains(result, args[i+1]) {
				result = append(result, args[i+1])
			}
		}
	}
	return result
}

//IsDryRun returns true/false if d flag is present
func IsDryRun(args []string) bool {
	var result bool = false
	for i := 0; i < len(args); i++ {
		if args[i] == "d" || args[i] == "-d" || args[i] == "--dry-run" {
			result = true
			break
		}
	}
	return result
}

//AllNamespaces returns true/false if A flag is present
func allNamespaces(args []string) bool {
	var result bool = false
	for i := 0; i < len(args); i++ {
		if args[i] == "A" || args[i] == "-A" || args[i] == "--all-namespaces" {
			result = true
			break
		}
	}
	return result
}

//GetNamespaces return the list of -n args namespaces or all as ""
func GetNamespaces(args []string) []string {
	var namespaces []string
	allNamespaces := allNamespaces(args)
	if !allNamespaces {
		namespaces = ProcessFlags(args, "n")
		namespaces = append(namespaces, "-namspace")
	} else {
		namespaces = append(namespaces, "")
	}
	return namespaces
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
