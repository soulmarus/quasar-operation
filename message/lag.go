package message

import "strings"

// func getMessage(messages ...[]string) (msg string) {
// 	finalMessage := ""
// 	seen := make(map[string]bool)

// 	for _, message := range messages {
// 		for _, item := range message {
// 			if item != "" {
// 				if seen[item] {
// 					continue
// 				}
// 				finalMessage += item + " "
// 				seen[item] = true
// 			}
// 		}
// 	}

// 	return strings.TrimSpace(finalMessage)
// }

func GetMessage(messages ...[]string) (msg string) {
	finalMessage := ""
	seen := make(map[string]bool)
	maxLen := determineMaxLength(messages...)

	for i := 0; i < maxLen; i++ {
		for _, message := range messages {
			if i < len(message) {
				if message[i] != "" {
					if seen[message[i]] {
						continue
					}
					finalMessage += message[i] + " "
					seen[message[i]] = true
				}
			}
		}
	}
	return strings.TrimSpace(finalMessage)
}

func determineMaxLength(messages ...[]string) int {
	maxLen := 0
	for _, message := range messages {
		len := len(message)
		if maxLen < len {
			maxLen = len
		}
	}
	return maxLen
}
