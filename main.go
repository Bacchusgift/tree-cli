package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
)

// 定义命令行参数
type Options struct {
	MaxDepth   int
	ShowHidden bool
	DirOnly    bool
	FullPath   bool
	NoIndent   bool
	OutputFile string
	Color      string
}

// 定义树状结构的绘制字符
const (
	Pipe     = "│   "
	LastPipe = "└── "
	MidPipe  = "├── "
	Space    = "    "
)

func main() {
	// 解析命令行参数
	opts := parseFlags()

	// 获取要遍历的目录路径
	dir := "."
	if len(flag.Args()) > 0 {
		dir = flag.Args()[0]
	}

	// 获取绝对路径
	absPath, err := filepath.Abs(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "错误: %v\n", err)
		os.Exit(1)
	}

	// 准备输出
	var output *os.File
	if opts.OutputFile != "" {
		var err error
		output, err = os.Create(opts.OutputFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "无法创建输出文件: %v\n", err)
			os.Exit(1)
		}
		defer output.Close()
	} else {
		output = os.Stdout
	}

	// 打印目录名（根节点用红色）
	red := color.New(color.FgRed).SprintFunc()
	if opts.FullPath {
		fmt.Fprintln(output, red(absPath))
	} else {
		fmt.Fprintln(output, red(filepath.Base(absPath)))
	}

	// 开始遍历目录
	count := &Counter{}
	err = walkDir(absPath, "", output, opts, 0, count)
	if err != nil {
		fmt.Fprintf(os.Stderr, "错误: %v\n", err)
		os.Exit(1)
	}

	// 打印统计信息
	fmt.Fprintf(output, "\n%d 个目录, %d 个文件\n", count.Dirs, count.Files)
}

// 解析命令行参数
func parseFlags() *Options {
	opts := &Options{}

	flag.IntVar(&opts.MaxDepth, "L", -1, "设置最大显示深度")
	flag.BoolVar(&opts.ShowHidden, "a", false, "显示所有文件，包括隐藏文件")
	flag.BoolVar(&opts.DirOnly, "d", false, "仅列出目录")
	flag.BoolVar(&opts.FullPath, "f", false, "显示完整路径")
	flag.BoolVar(&opts.NoIndent, "i", false, "不缩进，平铺显示")
	flag.StringVar(&opts.OutputFile, "o", "", "输出到文件")
	flag.StringVar(&opts.Color, "c", "green", "设置输出颜色")

	flag.Parse()
	return opts
}

// 计数器结构体
type Counter struct {
	Dirs  int
	Files int
}

// 遍历目录的函数
func walkDir(path, indent string, output *os.File, opts *Options, depth int, count *Counter) error {
	// 检查深度限制
	if opts.MaxDepth >= 0 && depth > opts.MaxDepth {
		return nil
	}

	// 读取目录内容
	entries, err := os.ReadDir(path)
	if err != nil {
		return err
	}

	// 过滤隐藏文件
	if !opts.ShowHidden {
		filtered := make([]os.DirEntry, 0, len(entries))
		for _, entry := range entries {
			if !strings.HasPrefix(entry.Name(), ".") {
				filtered = append(filtered, entry)
			}
		}
		entries = filtered
	}

	// 根据颜色选项设置颜色函数
	var colorFunc func(a ...interface{}) string
	switch opts.Color {
	case "red":
		colorFunc = color.New(color.FgRed).SprintFunc()
	case "blue":
		colorFunc = color.New(color.FgBlue).SprintFunc()
	case "yellow":
		colorFunc = color.New(color.FgYellow).SprintFunc()
	case "magenta":
		colorFunc = color.New(color.FgMagenta).SprintFunc()
	case "cyan":
		colorFunc = color.New(color.FgCyan).SprintFunc()
	default:
		colorFunc = color.New(color.FgGreen).SprintFunc()
	}

	// 遍历目录内容
	for i, entry := range entries {
		isLast := i == len(entries)-1

		// 构建当前项的完整路径
		entryPath := filepath.Join(path, entry.Name())

		// 确定是否为目录
		isDir := entry.IsDir()

		// 如果只显示目录且当前项不是目录，则跳过
		if opts.DirOnly && !isDir {
			continue
		}

		// 构建显示行
		var line string
		if opts.NoIndent {
			// 平铺模式
			if opts.FullPath {
				line = entryPath
			} else {
				line = entry.Name()
			}
		} else {
			// 树状模式
			prefix := ""
			if isLast {
				prefix = indent + LastPipe
			} else {
				prefix = indent + MidPipe
			}

			line = prefix + entry.Name()
		}

		// 输出当前行（根据节点类型设置颜色）
		if isDir {
			// 目录使用普通颜色
			fmt.Fprintln(output, line)
		} else {
			// 文件（叶子节点）使用指定颜色
			fmt.Fprintln(output, colorFunc(line))
		}

		// 更新计数器
		if isDir {
			count.Dirs++

			// 计算下一级的缩进
			nextIndent := indent
			if !opts.NoIndent {
				if isLast {
					nextIndent += Space
				} else {
					nextIndent += Pipe
				}
			}

			// 递归遍历子目录
			err := walkDir(entryPath, nextIndent, output, opts, depth+1, count)
			if err != nil {
				return err
			}
		} else {
			count.Files++
		}
	}

	return nil
}