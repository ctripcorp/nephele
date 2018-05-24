# Nephele团队代码风格指南。

## 背景

Nephele是一套企业级的图片解决方案。Go是Nephele团队主要使用的编程语言。Go语言简单，强大。其官方提供了大量值得参考遵循的规范。但Go始终是一门相对年轻，受众相对小的语言，会在一部分场景因缺乏相应约束造成代码可读性低下，形成学习壁垒，阻碍团队效率的进一步提升。

## 目标

一部分的目标不会有具体实例，但它们将作为一种理念被保留。

**与go fmt tool无冲突**

遵循该指南的代码经由go fmt tool格式化之后将依然遵循该指南。

**更在乎阅读者的体验**

可以要求开发者完成一些额外的编码工作，这些工作对于功能的完成而言可能是不必要的，但有利于阅读者的学习。

**代码拒绝小聪明**

将提倡直白的编码方式。

**注释拒绝浪漫**

将提倡严肃呆板的注释风格。

**清晰的作用范围**

可以通过制定命名的约定，限制变量和函数的调用与声明，达到可以快速判断变量函数出处的目的。

**最小化对于接口的冲击**

现有工程的改造需要一个循序渐进的过程。
该指南提供的规则将尽可能达到这样的效果：
不遵循该指南的代码被更改为遵循该指南之后，其调用方的代码结构无需变动。

**顺从不可避免的性能优化**

一些引发质变的性能优化可能引入部分含蓄的代码，这些代码是被允许并鼓励的。

## 目录

* [Import Package](#import-package)

* [Package](#package)

* [Indent](#indent)

* [Blank line](#blank-line)

* [Naming](#naming)

* [Variables](#variables)

* [Functions](#functions)

* [Lock](#lock)


## Import Package

**For example:**
```go
    import (
        "time"
        "net/http"
    )
```

**Not:**
```go
    import "time"
    import "net/http"
```

## Package

一个包下面必须有一个与包同名的文件。为了方便，这里称之为主文件(master file)，其他的与包不同名的文件，这里称之为分支文件(branch file)。

## Indent

缩进请务必使用tab。

## Blank line

在变量(variable)，方法(function)，接口(interface)，结构体(struct)之间插入空行。

**For example:**
```go
    var a int

    func b() {
        ...
    }

    type c struct {
        ...
    }
```

**Not:**
```go
    var a int
    func b() {
        ...
    }
    type c struct {
        ...
    }
```

## Naming

包名全部小写。
其他使用驼峰命名法，并有一些额外的要求：

**尽可能少的介词:**
```go
    func SetUsername() {
        ...
    }

    func (name *Username) String() string {
        ...
    }
```

**Not:**
```go
    func SetNameOfUser() {
        ...
    }

    func (name *Username) ToString() string {
        ...
    }
```

**既可以视为两个单词又可以视为一个单词的视为一个单词，以谷歌翻译为准:**
```go
    func SetUsername() {
        ...
    }

    func SetClientCode() {
        ...
    }
```

**Not:**
```go
    func SetUserName() {
        ...
    }

    func SetClientcode() {
        ...
    }
```

**动词在名词之前:**
```go
    func SetName() {
        ...
    }
```

**Not:**
```go
    func NameSet() {
        ...
    }
```

形参使用完整词，函数内可以使用简写变量:

**Example:**
```go
    func SetUsername(username string) {
        u := username
        ...
    }

    func SetClientcode(clientCode string) {
        cc := clientCode
        ...
    }
```

**Not:**
```go
    func SetUsername(u string) {
        ...
    }

    func SetClientcode(cc string) {
        ...
    }
```

字母变量尽可能使用原变量名首字母组合，避免由类型演变的简写:

**Not:**
```go
    func SetUsername(username string) {
        var s string
        ...
    }

    func SetUsername(username string) {
        var str string
        ...
    }

```

面对类型转换或数据解析时需要区分类型或格式的，将类型或格式完整名作为后缀，
这里建议一个数据都是在同一类型或格式下参与业务计算，该类型或格式下的数据变量名可以没有后缀：

**Example:**
```go
    //var usernameBytes []byte
    username := string(usernameBytes)
    //do something with username

    //var dataJson []byte
    //var data Data
    //dataJson := json.Marshal(data)
    json.Unmarshal(dataJson, &data)
    //do something with data
```

## Variables

禁止在分支文件中声明公用变量。

**For example:**
```go
    package foo
    //foo.go

    var A int
```

**Banned:**
```go
    package foo
    //goo.go

    var A int
```

在分支文件中, 全局私有变量名必须带上与分之文件名相同的前缀。该前缀不视为修饰词，前缀之后的单词首字母需要大写。

**For example:**
```go
    package foo
    //goo.go

    var gooA int
    const gooB = 1
```

**Not:**
```go
    package foo
    //goo.go

    var a int
    const b = 1
```

**Not:**
```go
    package foo
    //goo.go

    var gooa int
    const goob = 1
```

在实践中我们会发现, 分支文件命名的好坏和代码可读性强弱有着相当大的关系。可以说分之文件名是需要着重设计的。

## Functions

非方法的公有函数只能在主文件中定义。

**Allowed:**
```go
    package foo
    //foo.go

    func A() {
        ...
    }
```

**Banned:**
```go
    package foo
    //goo.go

    func B() {
        ...
    }
```

在分支文件中, 非方法的私有函数名必须带上与分之文件名相同的前缀。该前缀不视为修饰词，前缀之后的单词首字母需要大写。
原本在分支文件中定义的函数，通常会以如下形式再现。

**For example:**
```go
    package foo
    //goo.go

    type goo struct {
        a int
    }
    
    const gooA = 0
    const gooB = 1

    var gooC int = 2
    var gooD int = 3

    var gooInstance *goo = &goo{}

    func (g goo) step1() {
        ...
    }

    func (g goo) step2() {
        ...
    }

    func (g goo) do() {
        g.step1()
        g.step2()
    }
```

**or:**
```go
    package foo
    //goo.go

    func gooStep1() {
        ...
    }

    func gooStep2() {
        ...
    }

    func gooDo() {
        gooStep1()
        gooStep2()
    }
```

**Not:**
```go
    package foo
    //goo.go

    func step1() {
        ...
    }

    func step2() {
        ...
    }

    func do() {
        step1()
        step2()
    }
```

## Lock

用go sdk提供的sync包。别用channel。

**Example:**
```go
    var fooLock sync.Mutex

    func foo() {
        fooLock.Lock()
        ...
        fooLock.Unlock()
    }
```

**Not:**
```go
    var fooLock chan int = make(chan int, 1)

    func foo() {
        fooLock<-0
        ...
        <-fooLock
    }
```
