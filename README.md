# go-09 Go version too new
go.mod declares `go 1.99` (never satisfiable), so ANY real toolchain fails with a version error.
For a realistic case, edit it to a version newer than the build environment\x27s Go (e.g. 1.26).
