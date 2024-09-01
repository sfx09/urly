package internal

import "net/url"

func IsValidUrl(rawUrl string) bool {
	_, err := url.Parse(rawUrl)
	return err == nil
}
