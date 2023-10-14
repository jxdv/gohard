# gohard

### Description

gohard (golang harden) is a partial rewrite of [grapheneX](https://github.com/grapheneX/grapheneX), which I also try to actively maintain.
The goal was to get rid of all third party libs, and only have a CLI app.

I started learning Golang recently, so I thought this would be a good practice.

### Build

64-bit linux
```commandline
GOOS=linux GOARCH=amd64 go build -o bin/ src/main.go
```

64-bit windows
```commandline
GOOS=windows GOARCH=amd64 go build -o bin/ src/main.go
```

### Usage

Use SSH hardening modules in linux:
```commandline
sudo -E bin/main --service=ssh
```

### License

[GPL](LICENSE)
