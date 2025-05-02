package validators

import "regexp"

var urlRegex = regexp.MustCompile(`^(https?://)?([\w\-]+\.)+[\w\-]+(/[\w\-./?%&=]*)?$`)

func IsValidURL(url string) bool {
	return urlRegex.MatchString(url)
}
