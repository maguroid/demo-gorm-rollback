# gorm tx rollback demo

# how to reproduce

1. start database

```bash
make db/up
```

2. run go script

```bash
go run main.go
```

will print

```bash
go run main.go
> 2024/01/30 10:02:04 foo created
> 2024/01/30 10:02:04 2nd transaction returned error: should rollback
> 2024/01/30 10:02:04 foos: [{ID:1 Description:foo}]
```

you can also check the database record via phpmyadmin running at http://localhost:8080
