# Tree CLI

一个用Go语言编写的命令行工具，用于以树状结构显示目录内容，类似于Unix/Linux的tree命令。

## 功能

- 递归显示目录结构
- 支持显示/隐藏隐藏文件
- 可以限制显示深度
- 可以只显示目录
- 支持显示完整路径
- 支持输出到文件

## 安装

```bash
# 克隆仓库
git clone https://github.com/youzi/tree-cli.git
cd tree-cli

# 编译安装
go build -o tree

# 可选：将编译好的二进制文件移动到PATH中的目录
mv tree /usr/local/bin/
```

## 使用方法

```bash
# 显示当前目录结构
tree

# 显示指定目录结构
tree /path/to/directory

# 显示包括隐藏文件
tree -a

# 限制显示深度为2层
tree -L 2

# 只显示目录
tree -d

# 显示完整路径
tree -f

# 平铺显示（不缩进）
tree -i

# 输出到文件
tree -o output.txt
```

## 命令行选项

- `-L` 设置最大显示深度
- `-a` 显示所有文件，包括隐藏文件
- `-d` 仅列出目录
- `-f` 显示完整路径
- `-i` 不缩进，平铺显示
- `-o` 输出到文件

## 许可证

MIT