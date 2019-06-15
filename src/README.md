# Introduction

# How to
Compile udistance.go like this

Install dependancy
```
go get github.com/stianeikeland/go-rpio
```

Run build

```
env GOOS=linux GOARCH=arm GOARM=6 go build udistance.go
```
Copy and execute inside your raspberry the file "bin/arm6/led" if you're using raspberryi pi zero

And for raspberryi pi 3 "bin/arm7/led"

```
./udistance
```

Happy blinking 