package utils

import "url-shortener/models"


func AlreadyExists(value string) bool {
	links, err := models.GetLinks()
	if(err != nil) {
		panic(err)
	}

	for i := 0; i < len(links); i++ {
		if links[i].ShortLink == value {
			return true
		}
	}
	return false
}
