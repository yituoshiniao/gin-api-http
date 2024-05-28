#!/bin/bash

# 设置 LC_ALL 以避免非法字节序列错误
export LC_ALL=C

# 定义要替换的目录
target_dir=$(PWD)"/cmd/kitctl/internal/cmd/templates/skeleton"

# 使用 find 命令递归查找所有文件，并使用 sed 进行替换
find "$target_dir" -type f -exec sed -i '' 's|github.com/yituoshiniao/gin-api-http|{{ .Mod }}|g' {} +

find "$target_dir" -type f -exec sed -i '' 's|github.com/yituoshiniao/kit|gitlab.abc.net/cs-server2/kit|g' {} +

# 替换应用名称
find "$target_dir" -type f -exec sed -i '' 's|gin-api-http|{{ .AppName }}|g' {} +
find "$target_dir" -type f -exec sed -i '' 's|goodsCenterLogic|{{ .AppName }}|g' {} +




# 将  abcdefg/cad 替换成  {{ .Mod }}
#find "$target_dir" -type f -exec sed -i '' 's|abcdefg/cad|{{ .Mod }}|g' {} +

echo "替换完成"