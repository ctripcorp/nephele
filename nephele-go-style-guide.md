# Nephele Go Style Guide

This style guide outlines the coding conventions of Nephele team.

## Background

Nephele is a complete solution to enterprise multimedia needs. Go is one of the main development language used by most of Nephele team members. The very go language is simple, powerful and probably a most officially restricted language, yet still used in a relatively smaller range. Thus in some common cases, we still lack rules and conventions intended to prevent obstacles of team efficiency and barriers to learning.

## Goals of the Style Guide

**No conflict with go fmt tool**

**Optimize for the reader, not the writer**

**Avoid tricky realization**

**Be mindful of scale**

**No impact on interface**

**Concede to optimization when necessary**

## Table of Contents

* [How to Import Package](#how-to-import-package)

* [Indent](#indent)

* [Blank line](#blank-line)

* [Naming](#naming)

* [Variables](#variables)

* [Functions](#functions)

* [Lock](#lock)


## How to Import Package

**For example:**
```go
    import (
        time_p "time"
        http_p "net/http"
    )
```

**Not:**
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

## Indent

Indentation MUST use tabs.

## Blank line

Blank line between variable, function, interface and struct.

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

Basically we use Camel-Case.

ALSO PROBABLY

**No preposition:**
```go
    func getUsername() {
        ...
    }


    func GetUsername() {
        ...
    }

    func (name *Username) TryGet() string {
        ...
    }

    func (name *Username) String() string {
        ...
    }
```

**Not:**
```go
    func GetNameOfUser() {
        ...
    }

    func GetTheNameOfUser() {
        ...
    }

    func (name *Username) TryToGet() string {
        ...
    }

    func (name *Username) ToString() string {
        ...
    }
```

**Adorned nouns and verbs regarded as one:**
```go
    func GetUsername() {
        ...
    }

    func QuicklygetUsername() {
    }
```

**Not:**
```go
    func QuicklyGetUserName() {
        ...
    }
```

**Verb before noun:**
```go
    func GetName() {
        ...
    }
```

**Not:**
```go
    func NameGet() {
        ...
    }
```

## Variables

A suffixed filename is needed for global private variable or const.

**For example:**
```go
    //foo.go

    var a_foo int
    const b_foo = 1

    func c() {
        ...
    }
```

**Not:**
```go
    //foo.go

    var a int
    const b = 1

    func c() {
        ...
    }
```

## Functions

We name the file whose name is the same as its package name "the core file" and the others "the branch file".
Global functions are only allowed in core files.

**Allowed:**
```go
    package foo
    //foo.go

    func A() {
        ...
    }

    func a() {
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

    func b() {
        ...
    }
```

So in branch file, we integrate functions into a single struct, PROBABLY a struct whose name is the same as the branch file name.

**For example:**
```go
    package foo
    //goo.go

    type goo struct {
        a int
    }
    
    const a_goo = 0
    const b_goo = 1

    var c_goo int = 2
    var d_goo int = 3

    var instance_goo *goo = &goo{}

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

**Not:**
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

## Lock

Use sync package instead of channel.


**Example:**
```go
    var lock_foo sync_p.Mutex

    func foo() {
        lock_foo.Lock()
        ...
        lock_foo.Unlock()
    }
```


**Not:**
```go
    var lock_foo chan int = make(chan int, 1)

    func foo() {
        lock_foo<-0
        ...
        <-lock_foo
    }
```

