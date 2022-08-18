package service

import "strings"

func IsValid(message string) bool {
	strand := strings.ReplaceAll(message, "\r\n", "\n")
	for _, v := range strand {
		if (v < 32 || v > 126) && v != 10 {
			return true
		}
	}
	return false
}
