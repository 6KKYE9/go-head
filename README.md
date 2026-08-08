# go-head

复制粘贴改格式改到手酸？这玩意儿一行就搞定。

显示文件或标准输入的开头若干行（`-n`）或若干字节（`-b`）。

## 安装

```bash
go build -o go-head.exe
```

## 用法

```bash
go-head -n 5 file.txt     # 前 5 行
go-head -b 100 file.txt   # 前 100 字节（优先级高于 -n）
echo "hello" | go-head -n 3
```

## 说明

零依赖纯 Go。字节模式和行模式二选一，给了 `-b` 就走字节。
