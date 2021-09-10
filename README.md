

## module

Go 1.11 和 1.12 初步包含了对模块的支持，Go 的 新依赖管理系统 使依赖版本信息明确且易于管理。

### 启用go module

```shell
go env -w GO111MODULE=auto
```



### 创建module

```
# 创建模块文件夹
$ mkdir hello && cd hello
# 生成模块
$ go mod init example.com/hello
$ ll 
total 8
drwxr-xr-x  3 yangsx  staff   96 Jul 30 10:42 .
drwxr-xr-x  6 yangsx  staff  192 Jul 30 10:41 ..
-rw-r--r--  1 yangsx  staff   34 Jul 30 10:42 go.mod
```

生成的go.mod信息如下：

```
// module的名称
module example.com/hello
// go语言环境的版本
go 1.16
```

### 创建package

在go中，一个文件夹下，只能拥有一个包，但是可以拥有同个包的多个go源码文件。通常情况模块根目录用来定义main.go文件，用作module的启动入口，其他包要在根目录下创建同名文件夹。

创建`sayhello` package：

```
$ mkdir sayhello && cd sayhello
$ touch sayhello.go
```

在package下增加golang源码文件`sayhello.go`，并定义一个方法：

```
package hello

import "fmt" // 导入系统模块

func SayHello(name string) {
  // 方法大写字母代表此方法是公开的
  fmt.Println("Hello" + name)
}
```

### 创建模块入口main

```
# 切换路径到模块根目录
$ pwd
/Users/yangsx/GoLandProjects/notes-golang/hello
```







