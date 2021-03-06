# go-design-pattern

## factory
```bash
➜  factory go test -v *.go
=== RUN   TestCacheFactory_Create
redis data: v1
mem data: m1
--- PASS: TestCacheFactory_Create (0.00s)
PASS
ok  	command-line-arguments	0.619s
```
### 优点
 实现了解耦
 
### 缺点
 违背了"开闭原则"

### 适合
 创建的对象比较少 