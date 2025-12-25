package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type User struct {
	ID       int    `json:"user_id"`
	Username string `json:"username"`
	IsAdmin  bool   `json:"is_admin"`
	Password string `json:"-"`
}

func test2() {
	u := User{
		ID:       1,
		Username: "even lemon",
		IsAdmin:  true,
		Password: "123456",
	}

	fmt.Printf("%v\n", u)
	jsonData, err := json.Marshal(u)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(string(jsonData))
	}

	jsonStr := `{"username": "even lemon", "is_admin": true, "user_id": 2}`
	var u2 User
	err = json.Unmarshal([]byte(jsonStr), &u2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("%#v\n", u2)
	}

	dir := "data"
	fileNmae := "config.txt"
	fullPath := filepath.Join(dir, fileNmae)
	fmt.Println(fullPath)

	content := []byte("Hello, World!的内容")
	// 先创建文件夹（如果不存在）
	err = os.MkdirAll(dir, 0755)
	if err != nil {
		fmt.Println("Error creating directory:", err)
	} else {
		err = os.WriteFile(fullPath, content, 0644)
	}
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("File written successfully")
	}

	readDta, err := os.ReadFile(fullPath)
	if err != nil {
		fmt.Println("Error reading file:", err)
	} else {
		fmt.Println(string(readDta))
	}
	os.Remove(fullPath)
	os.RemoveAll(dir)
}
