```
go test -bench . -cpuprofile cpu.out
```


```
go tool pprof cpu.out
web
quit
```

* 代码覆盖
* 使用go test 获取代码覆盖报告
* 使用go tool cover 查看代码覆盖报告


# 性能测试

* test.B 的使用
* 例： 使用pprof优化性能

