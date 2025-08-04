package main

import (
	"fmt"
)

func modify(userMap map[string]map[string]string, name string) {
	if user, exists := userMap[name]; exists {
		user["passwd"] = "888888"
	} else {
		userMap[name] = make(map[string]string)
		userMap[name]["passwd"] = "888888"
	}
	if userMap[name] != nil {
		userMap[name]["nickname"] = "ModifiedName"
	} else {
		userMap[name] = make(map[string]string)
		userMap[name]["nickname"] = "ModifiedName"
	}

}

func main() {
	userMap := make(map[string]map[string]string)
	userMap["user1"] = make(map[string]string)
	userMap["user1"]["nickname"] = "Alice"
	userMap["user1"]["passwd"] = "111111"

	userMap["user2"] = make(map[string]string)
	userMap["user2"]["nickname"] = "Bob"
	userMap["user2"]["passwd"] = "222222"
	fmt.Println("User Map:", userMap)
	modify(userMap, "user1")
	fmt.Println("Modified User Map:", userMap)
}
