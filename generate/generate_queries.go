package main

import (
	"fmt"
	"os"
)

func main() {
	// 创建或覆盖queries.go文件
	file, err := os.Create("queries.go")
	if err != nil {
		fmt.Println("Error creating queries.go file:", err)
		return
	}
	defer file.Close()
	// 将生成的代码写入文件
	_, err = file.WriteString(queriesTemplate)
	if err != nil {
		fmt.Println("Error writing to queries.go file:", err)
		return
	}
	fmt.Println("Queries generated successfully.")
}
