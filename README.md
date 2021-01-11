
# qr-decoder-cli

Simple CLI tool for decoding qrcodes from jpeg,png images. Uses <https://github.com/liyue201/goqr> library.

### Install
Depends upon  <https://github.com/liyue201/goqr> 


So install it first via `go get github.com/liyue201/goqr`.
Then you can simply run the script file directly.

### Usage
#### 1. Single file

`go run qr-decoder-cli/main.go --input=</path/to/image.jpg>`
#### 2. Multiple files via GLOB pattern

Both are valid

`go run qr-decoder-cli/main.go --input="</path/to/*.png>"`

or

`go run qr-decoder-cli/main.go --input "</path/to/*.png>"`

  

#### 3. If decoded data is a url , use --status to check the HTTP Status

`go run qr-decoder-cli/main.go --input=</path/to/image.jpg> --status`

  

`go run qr-decoder-cli/main.go --input="</path/to/*.jpg>" --status`

  

### Help
```

-i, --input string
Image file path. Can also be a glob pattern like /home/user/*.jpeg

-status, --status
Perform an HTTP GET request to the decoded url

```