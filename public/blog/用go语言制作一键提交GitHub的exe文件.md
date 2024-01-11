点击返回[🔗我的博客文章目录](https://percheung.github.io/#/toc)
* 目录
{:toc}
<div onclick="window.scrollTo({ top: 0, behavior: 'smooth' });" style="background-color: transparent; position: fixed; bottom: 20px; right: 40px; padding: 10px 10px 5px 10px; cursor: pointer; z-index: 10; border-radius: 13%; box-shadow: 0.5px 3px 7px rgba(0, 0, 0, 0.3);"><img src="https://percheung.github.io/blogImg/backTop.png" alt="TOP" style="background-color: transparent; width: 30px;"></div>

# 用go语言<img src="https://percheung.github.io/blogImg/golang_.png" width="50px" alt="C" />制作一键提交GitHub的exe文件

## 1.代码如下

```go
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	fmt.Println("--- Git拉取 ---")
	runCommand("git", "pull")

	fmt.Println("--- Git添加 ---")
	runCommand("git", "add", ".")

	fmt.Println("--- Git提交 ---")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("请输入提交信息: ")
	commitMessage, _ := reader.ReadString('\n')
	runCommand("git", "commit", "-m", commitMessage)

	fmt.Println("--- Git推送 ---")
	runCommand("git", "push", "origin", "main")

	fmt.Println("--- 完成 ---")
	fmt.Print("请按任意键继续...")
	reader.ReadString('\n')
}

// 执行命令并打印输出
func runCommand(command string, args ...string) {
	cmd := exec.Command(command, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
```

> 我的思路是，先拉取，然后添加到本地仓库，然后提交。

## 2.制作成exe文件

> go语言打包命令

```bash
go build -o github-commit.exe github-commit.go
```

## 3.使用

Windows系统下，使用管理员身份双击执行`github-commit.exe`即可。