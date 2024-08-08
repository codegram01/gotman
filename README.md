```
              #   ___                            
    -*~*-     #  <_*_>        ()_()          
    (o o)     #  (o o)        (o o)          
ooO--(_)--Ooo-8---(_)--OooooO--`o'--OooooO
```
# Gotman - A CLI tool for make Http request

Gotman is an CLI API Client write in Go.
It like Postman or Curl alternative 

[Video when i create gotman](https://youtu.be/-oDIp2BBs6w)

# Feature 
- Make api request (http request) like Postman
- Work in CLI, so you can test your api when ssh in Server 
- Enable load file config, so type less than 

# Benefit this projects 
- Make a tool (write X from scratch)
- Deep understand Go:
    + http Client, Request, Response
    + bytes buffer 
    + io Reader, Writer interface
    + How get user input, get user config file

## Usage
```sh
gotman 
    -u [Url] 
    -m [Method]
    -h [Header]
    -b [Body]
    -f [File_Config]
```

Examples: 
[example.gotman.sh](https://github.com/codegram01/gotman/blob/main/example/gotman.sh)

## Install
Prerequisites: [Install Golang](https://go.dev/doc/install) 

### Option 1: Install Binary
```sh
go install github.com/codegram01/gotman@latest
```

### Option 2: Install and build from source
```sh
git clone https://github.com/codegram01/gotman.git

cd gotman

# Run code: 
go run . 

# Build 
go build .
```

# Contribute 
You are welcome to contribute this project. 
Help me upgrade or test it. 

Open issue or create fork and pull request to contribute 