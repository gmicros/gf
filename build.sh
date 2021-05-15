#! /bin/bash

go tool compile fg.go
go tool compile ./fractal_gen/fractal_gen.go

go tool link fg.o
