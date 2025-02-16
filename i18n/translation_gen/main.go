// go run main.go ..\translations\zh.json ..\consts.go
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strings"
)

// toPascalCase 将字符串的首字母转换为大写
func toPascalCase(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

// formatKey 按照 . 分割，然后每部分转换为大驼峰，再拼接
func formatKey(key string) string {
	parts := strings.Split(key, ".")
	for i, part := range parts {
		parts[i] = toPascalCase(part)
	}
	return strings.Join(parts, "")
}

func main() {
	// 检查命令行参数，需要提供输入 JSON 文件和输出 Go 文件
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run main.go input.json output.go")
		os.Exit(1)
	}

	inputFilename := os.Args[1]
	outputFilename := os.Args[2]

	data, err := ioutil.ReadFile(inputFilename)
	if err != nil {
		log.Fatalf("读取文件错误: %v", err)
	}

	// 假设 JSON 格式为 map[string]string
	var m map[string]string
	if err := json.Unmarshal(data, &m); err != nil {
		log.Fatalf("解析 JSON 错误: %v", err)
	}

	// 收集所有 key 并排序，使输出顺序稳定
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// 创建输出文件
	f, err := os.Create(outputFilename)
	if err != nil {
		log.Fatalf("创建输出文件错误: %v", err)
	}
	defer f.Close()

	// 写入生成的 Go 代码到文件
	fmt.Fprintln(f, "package i18n\n")
	fmt.Fprintln(f, "const (")

	for _, k := range keys {
		constName := formatKey(k)
		value := m[k]
		fmt.Fprintf(f, "\t// %s\n", value)
		fmt.Fprintf(f, "\t %s Key = \"%s\"\n", constName, k)
	}

	fmt.Fprintln(f, ")")
}
