package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// 为所有题目批量添加ACM模式文件
func main() {
	rootDir := "."
	
	// 获取所有题目目录
	dirs, err := ioutil.ReadDir(rootDir)
	if err != nil {
		fmt.Printf("Error reading directory: %v\n", err)
		return
	}
	
	// 匹配题目目录的正则表达式（如：0001_two_sum）
	dirRegex := regexp.MustCompile(`^\d{4}_.*$`)
	
	for _, dir := range dirs {
		if dir.IsDir() && dirRegex.MatchString(dir.Name()) {
			dirPath := filepath.Join(rootDir, dir.Name())
			addACMModeForProblem(dirPath, dir.Name())
		}
	}
	
	fmt.Println("ACM模式文件添加完成！")
}

func addACMModeForProblem(dirPath, dirName string) {
	fmt.Printf("正在处理: %s\n", dirName)
	
	// 检查是否已有ACM模式文件
	acmGoFile := filepath.Join(dirPath, "solution_acm.go")
	if _, err := os.Stat(acmGoFile); err == nil {
		fmt.Printf("  - ACM模式文件已存在，跳过: %s\n", dirName)
		return
	}
	
	// 读取现有的solution.go文件
	solutionFile := filepath.Join(dirPath, "solution.go")
	content, err := ioutil.ReadFile(solutionFile)
	if err != nil {
		fmt.Printf("  - 无法读取solution.go文件: %v\n", err)
		return
	}
	
	// 解析题目信息
	problemInfo := parseProblemInfo(string(content))
	
	// 生成ACM模式文件
	acmContent := generateACMModeFile(dirName, problemInfo)
	
	// 写入ACM模式文件
	err = ioutil.WriteFile(acmGoFile, []byte(acmContent), 0644)
	if err != nil {
		fmt.Printf("  - 写入ACM模式文件失败: %v\n", err)
		return
	}
	
	// 生成ACM模式测试文件
	testContent := generateACMTestFile()
	testFile := filepath.Join(dirPath, "solution_acm_test.go")
	err = ioutil.WriteFile(testFile, []byte(testContent), 0644)
	if err != nil {
		fmt.Printf("  - 写入ACM测试文件失败: %v\n", err)
		return
	}
	
	fmt.Printf("  - 成功添加ACM模式文件: %s\n", dirName)
}

// 解析题目信息
func parseProblemInfo(content string) map[string]string {
	info := make(map[string]string)
	
	lines := strings.Split(content, "\n")
	for _, line := range lines {
		if strings.Contains(line, "题目：") {
			// 提取题目名称
			parts := strings.Split(line, "题目：")
			if len(parts) > 1 {
				info["title"] = strings.TrimSpace(parts[1])
			}
		} else if strings.Contains(line, "链接：") {
			// 提取题目链接
			parts := strings.Split(line, "链接：")
			if len(parts) > 1 {
				info["link"] = strings.TrimSpace(parts[1])
			}
		} else if strings.Contains(line, "难度：") {
			// 提取题目难度
			parts := strings.Split(line, "难度：")
			if len(parts) > 1 {
				info["difficulty"] = strings.TrimSpace(parts[1])
			}
		}
	}
	
	return info
}

// 生成ACM模式文件内容
func generateACMModeFile(dirName string, problemInfo map[string]string) string {
	title := problemInfo["title"]
	if title == "" {
		title = "未知题目"
	}
	
	// 根据题目类型生成不同的ACM模式实现
	var implementation string
	switch {
	case strings.Contains(dirName, "sum"):
		// 和相关问题
		implementation = generateSumProblemACM()
	case strings.Contains(dirName, "list"):
		// 链表相关问题
		implementation = generateListProblemACM()
	case strings.Contains(dirName, "tree"):
		// 树相关问题
		implementation = generateTreeProblemACM()
	case strings.Contains(dirName, "parentheses"):
		// 括号相关问题
		implementation = generateParenthesesProblemACM()
	default:
		// 默认模板
		implementation = generateDefaultACM()
	}
	
	return fmt.Sprintf(`package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 题目：%s (ACM模式)
// 输入格式：
// 根据具体题目要求定义输入格式
// 输出格式：
// 根据具体题目要求定义输出格式

// 思路：
// 根据原solution.go中的思路实现ACM版本

// 时间复杂度：
// 空间复杂度：

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	
	// 读取输入
	// 根据具体题目调整输入处理逻辑
	for scanner.Scan() {
		line := scanner.Text()
		// 处理输入并调用解题函数
		// 输出结果
		break // 示例：只处理一行输入
	}
}

%s
`, title, implementation)
}

// 生成和相关问题的ACM实现
func generateSumProblemACM() string {
	return `// 示例：两数之和类型的实现
// 根据具体题目修改函数签名和逻辑
func solve(nums []int, target int) []int {
	// 实现具体算法
	return nil
}`
}

// 生成链表相关问题的ACM实现
func generateListProblemACM() string {
	return `type ListNode struct {
	Val  int
	Next *ListNode
}

// 示例：链表相关问题的实现
// 根据具体题目修改函数签名和逻辑
func solve(head *ListNode) *ListNode {
	// 实现具体算法
	return head
}`
}

// 生成树相关问题的ACM实现
func generateTreeProblemACM() string {
	return `type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 示例：树相关问题的实现
// 根据具体题目修改函数签名和逻辑
func solve(root *TreeNode) int {
	// 实现具体算法
	return 0
}`
}

// 生成括号相关问题的ACM实现
func generateParenthesesProblemACM() string {
	return `// 示例：括号相关问题的实现
// 根据具体题目修改函数签名和逻辑
func solve(s string) bool {
	// 实现具体算法
	return true
}`
}

// 生成默认ACM实现
func generateDefaultACM() string {
	return `// 根据具体题目实现
// 此处应根据原solution.go中的具体函数进行调整
func solve(input interface{}) interface{} {
	// 实现具体算法
	return nil
}`
}

// 生成ACM测试文件内容
func generateACMTestFile() string {
	return `package main

import "testing"

func TestACMMode(t *testing.T) {
	// 根据具体题目实现测试用例
	// 这里可以测试各个解题函数的逻辑
	t.Log("ACM模式测试")
}
`
}