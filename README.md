Catcher [![GoDoc](https://godoc.org/github.com/xlab/catcher?status.svg)](https://godoc.org/github.com/xlab/catcher)
=======

A package to gracefully handle crap software. From the creator of [closer](https://github.com/xlab/closer).

See [example](example/main.go).

```go
func main() {
    defer catcher.Catch(
        catcher.RecvWrite(os.Stderr, true),
    )

    if err := safeCall(); err != nil {
        log.Println("[ERR] safeCall failed with:", err)
    }

    suspiciousFunc()
}

// suspiciousFunc will definitely panic. Usually this kind of functions
// panic only on Saturdays or holidays, but for test simplicity this one
// will panic 100% of the time.
func suspiciousFunc() {
    panic("sorry pls")
}
```

Result:

```
$ go run main.go

caught error: catch me
main.go:21: [ERR] safeCall failed with: catch me
caught panic: sorry pls from suspiciousFunc

stacktrace: panic
/usr/local/go/src/runtime/asm_amd64.s:479 (0x4f4cc)
    call32: CALLFN(·call32, 32)
/usr/local/go/src/runtime/panic.go:458 (0x26fc3)
    gopanic: reflectcall(nil, unsafe.Pointer(d.fn), deferArgs(d), uint32(d.siz), uint32(d.siz))
/Users/xlab/Documents/dev/go/src/github.com/xlab/catcher/example/main.go:28 (0x22fd)
    suspiciousFunc: panic("sorry")
/Users/xlab/Documents/dev/go/src/github.com/xlab/catcher/example/main.go:24 (0x2186)
    main: suspiciousFunc()
/usr/local/go/src/runtime/proc.go:183 (0x28c74)
    main: main_main()
/usr/local/go/src/runtime/asm_amd64.s:2059 (0x51f31)
    goexit: BYTE    $0x90   // NOP
```

## License

MIT
