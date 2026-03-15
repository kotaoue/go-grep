package grep

import "regexp"

// Find returns lines that match the given pattern.
func Find(lines []string, pattern string) ([]string, error) {
	re, err := regexp.Compile(pattern)
	if err != nil {
		return nil, err
	}

	var result []string
	for _, line := range lines {
		if re.MatchString(line) {
			result = append(result, line)
		}
	}
	return result, nil
}
