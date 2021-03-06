# go-design-pattern

## 简单工厂模式( Simple Factory Pattern )
```bash
➜  simple_factory (main) ✗ go test -v *.go
=== RUN   TestCacheFactory_Create
redis data: v1
mem data: m1
--- PASS: TestCacheFactory_Create (0.00s)
PASS
ok  	command-line-arguments	0.574s
```
### 优点
 实现了解耦
 
### 缺点
 违背了"开闭原则"

### 适合
 创建的对象比较少 
 
 
 ------
- 参考：
 1. [图说设计模式](https://design-patterns.readthedocs.io/zh_CN/latest/index.html)