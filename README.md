## 定义错误的方式
```go
var ErrBusiness = bizerr.New("usecase.business.error.code")
```

## 使用
可以直接使用定义的Error变量作为错误，如果需要增加params：
```go
err := ErrBusiness.WithParam("param1", "param2")
```

在错误传递时可以使用`fmt.Errorf` 的 Wrap Verb `%w` 追加错误信息：
```go
err := ErrNoPermission.WithParam("param1", "param2")
newErr := fmt.Errorf("%w\n%s is required!", err, "root")
```
可以反复使用`fmt.Errorf`Wrap错误，这样在代码中每一层处理错误时都可以追加信息。

> ⚠️注意
> `%w` 的 `%` 前不能有任何字符出现，应该以 `%w` 开头，在后面追加信息。
> 如果一定要在前面插入信息，应该在 `%w` 前增加一个空格，比如：

不要这样：
```go
err := ErrNoPermission.WithParam("param1", "param2")
newErr := fmt.Errorf("error:%w\n%s is required!", err, "root")
```
但增加了空格后是可以的(不建议)：
```go
err := ErrNoPermission.WithParam("param1", "param2")
newErr := fmt.Errorf("error: %w\n%s is required!", err, "root")
```

## 处理
直接打印会包含Wrap的每一层的错误信息：
```go
fmt.Print(err)
```

有两种方式可以提取原始的错误码信息并进行值对比：

1. 使用 `errors.Is` 函数
```go
if errors.Is(err, ErrNoPermission) {
    fmt.Printf("biz error\ncode: %s\nparams: %+v\n",
        ErrNoPermission, bizerr.ExtractParams(err))
} else if errors.Is(err, ErrFolderNameConflict) {
    fmt.Printf("biz error: %s\n", err)
} else if err != nil {
    fmt.Println("internal error")
}
```

2. 使用 `errors.As` 函数
```go
var bizErr bizerr.Error
if errors.As(err, &bizErr) {
    switch bizErr {
    case ErrNoPermission:
        fmt.Printf("biz error\ncode: %s\nparams: %+v\n",
            bizErr, bizerr.ExtractParams(err))
    case ErrNetworkUnreachable:
        fmt.Printf("biz error: %s\n", err)
    case ErrFolderNameConflict:
        fmt.Printf("biz error: %s\n", err)
    default:
        fmt.Println("internal error")
    }
}
```

而error中的Params信息可以使用 `bizerr.ExtractParams` 函数提取，上面的代码已经演示了如何使用，需要注意的是，传入的err不是经`errors.As`提取出来的bizerr.Error, 而是原始的`err`。