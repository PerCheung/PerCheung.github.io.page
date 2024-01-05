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

	fmt.Println("--- 更新目录 ---")

	// 获取当前文件夹路径
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// 拼接目标文件夹路径
	blogDir := filepath.Join(currentDir, "blog")

	// 切换到目标文件夹
	err = os.Chdir(blogDir)
	if err != nil {
		log.Fatal(err)
	}

	// 执行 toc.exe
	runCommand(filepath.Join(blogDir, "toc.exe"))

	// 切换回原始文件夹
	err = os.Chdir(currentDir)
	if err != nil {
		log.Fatal(err)
	}

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