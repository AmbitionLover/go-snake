## 日志

记录日志使用 zap

安装：

```shell
go get -u go.uber.org/zap
```

#### 分级归档

1. 按文件大小

文件大小切割借助了第三方库 [lumberjack](https://github.com/natefinch/lumberjack) ，可以选择文件的大小来切割

```go
func  sizeDivisionWriter(filename string) io.Writer {
	return &lumberjack.Logger{
		Filename:   filename, // 日志文件的位置
		MaxSize:    10,   // 在进行切割之前，日志文件的最大大小（以MB为单位）
		MaxBackups: 200,  // 保留旧文件的最大个数
		MaxAge:     30,   // 保留旧文件的最大天数
		Compress:   true, // 是否压缩/归档旧文件
	}
}
```

2. 按时间切割

按时间切割借助到了第三方的库[file-rotatelogs](https://github.com/lestrrat/go-file-rotatelogs)

```go
func timeDivisionWriter(filename string) io.Writer {
	hook, err := rotatelogs.New(
        strings.Replace(filename, ".log", "", -1)+"_%Y-%m-%d.log",
        rotatelogs.WithLinkName(filename),
        rotatelogs.WithMaxAge(time.Duration(int64(24*time.Hour)*int64(30))),
        rotatelogs.WithRotationTime(time.Hour*24),
	)
	if err != nil {
		panic(err)
	}
	return hook
}
```