# Tree CLI

一个用Go语言实现的命令行目录树工具，类似于Unix的`tree`命令。

## 功能
- 可自定义深度的目录树显示
- 显示隐藏文件选项
- 仅列出目录
- 显示完整路径
- 平铺（无缩进）显示模式
- 结果输出到文件
- 可自定义颜色的彩色输出

## 命令行选项
- `-L` 设置最大显示深度
- `-a` 显示所有文件，包括隐藏文件
- `-d` 仅列出目录
- `-f` 显示完整路径
- `-i` 平铺显示（无缩进）
- `-o` 输出到文件
- `-c` 设置输出颜色（可选：red, blue, green, yellow, magenta, cyan）

## 使用方法
```bash
# 显示当前目录结构
./tree-cli

# 显示最多2层深度
./tree-cli -L 2

# 显示包括隐藏文件
./tree-cli -a

# 仅列出目录
./tree-cli -d

# 显示完整路径
./tree-cli -f

# 平铺显示（无缩进）
./tree-cli -i

# 输出到文件
./tree-cli -o output.txt

# 使用自定义颜色显示
./tree-cli -c blue
```

## 构建
确保已安装Go，然后运行：
```bash
go build -o tree-cli main.go
```

## 许可证
MIT
