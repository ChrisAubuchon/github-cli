package commands

func stringInSlice(s string, sl []string) bool {
	for _, t := range sl {
		if t == s {
			return true
		}
	}

	return false
}
