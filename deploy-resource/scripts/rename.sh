#!/bin/bash

# 定义目标目录, 将目录下的所有文件重命名， 文件名结尾添加 ".tpl"
DIR=$(PWD)"/cmd/kitctl/internal/cmd/templates/skeleton"

# 确保目标目录存在
if [ -d "$DIR" ]; then
  # 使用 find 命令递归查找所有文件并重命名
  find "$DIR" -type f | while read -r FILE; do
    # 检查文件是否以 .tpl 结尾
    if [[ "$FILE" != *.tpl ]]; then
      # 获取文件的目录路径
      DIRNAME=$(dirname "$FILE")
      # 获取文件名
      BASENAME=$(basename "$FILE")
      if [ $BASENAME == "docs.go" ]; then
        #echo "跳过：$FILE"
        continue # 跳过当前迭代
      fi
      # 新文件名
      NEWNAME="${BASENAME}.tpl"
      # 重命名文件
      mv "$FILE" "$DIRNAME/$NEWNAME"
    else
      echo "文件已以 .tpl 结尾，跳过：$FILE"
    fi
  done
  echo "所有文件已重命名并添加 .tpl 后缀"
else
  echo "目录不存在：$DIR"
fi
