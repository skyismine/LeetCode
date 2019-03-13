# pprof 性能分析工具使用方法

### 在代码中加入 pprof 统计代码

```go
func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}
	//other operation
}
```

### 运行程序时添加参数使数据输出到文件中

```bash
$ ./leetcode --cpuprofile=twosum.prof
```

### 使用 go tool pprof 工具分析

```bash
$ go tool leetcode twosum.prof
(pprof) top
```

