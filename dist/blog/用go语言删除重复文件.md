点击返回[🔗我的博客文章目录](https://percheung.github.io/#/toc)
* 目录
{:toc}
<div onclick="window.scrollTo({top:0,behavior:'smooth'});" style="background-color:white;position:fixed;bottom:20px;right:40px;padding:10px 10px 5px 10px;cursor:pointer;z-index:10;border-radius:13%;box-shadow:0.5px 3px 7px rgba(0,0,0,0.3);"><img src="https://percheung.github.io/blogImg/backTop.png" alt="TOP" style="background-color:white;width:30px;"></div>

# 用go<img src="https://percheung.github.io/blogImg/golang_.png" width="50px" alt="C" />语言删除重复文件

## 需求：将同级别目录（只有一层的目录，没子目录）下的重复文件删除

```go
package main

import (
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	files := make(map[string]string)
	duplicates := make(map[string]bool)

	// 获取当前目录下的所有文件
	fileList, err := ioutil.ReadDir(".")
	if err != nil {
		fmt.Println("读取目录时出错:", err)
		return
	}

	// 遍历所有文件
	for _, file := range fileList {
		if !file.IsDir() {
			filePath := file.Name()
			content, err := ioutil.ReadFile(filePath)
			if err != nil {
				fmt.Println("读取文件时出错:", err)
				continue
			}
			hash := sha256.Sum256(content)
			hashString := fmt.Sprintf("%x", hash)

			// 检查是否有重复文件
			if existingFile, ok := files[hashString]; ok {
				fmt.Printf("发现重复文件: %s 和 %s\n", filePath, existingFile)
				duplicates[filePath] = true
			} else {
				files[hashString] = filePath
			}
		}
	}

	// 删除重复文件
	for file := range duplicates {
		err := os.Remove(file)
		if err != nil {
			fmt.Println("删除文件时出错:", err)
		} else {
			fmt.Println("已删除重复文件:", file)
		}
	}
}
```

## 打包成exe文件

```bash
go build -o 删除重复文件.exe main.go
```

## 使用

只需将`删除重复文件.exe`放到想要删掉重复文件的目录下，双击运行就会删掉重复的文件了。