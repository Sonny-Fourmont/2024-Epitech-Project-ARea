package services

import "fmt"

func LatestMailAction(userID string, this string) []string {
	var result []string
	fmt.Println("Executing Latest Mail Action user_id: ", userID, " this: ", this)
	return result
}
