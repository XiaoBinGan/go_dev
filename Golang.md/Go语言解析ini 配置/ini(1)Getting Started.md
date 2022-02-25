[Getting Started](https://github.com/go-ini/ini)
We will go through a very simple example to illustrate how to get started.

First of all, create two files (my.ini and main.go) under the directory of your choice, let’s say we choose /tmp/ini.
```
$ mkdir -p /tmp/ini
$ cd /tmp/ini
$ touch my.ini main.go
$ tree .
.
├── main.go
└── my.ini
```
0 directories, 2 files
Now, we put some content into the my.ini file (partially take from Grafana).
```ini
# possible values : production, development
app_mode = development

[paths]
# Path to where grafana can store temp files, sessions, and the sqlite3 db (if that is used)
data = /home/git/grafana

[server]
# Protocol (http or https)
protocol = http

# The http port  to use
http_port = 9999

# Redirect to correct domain if host header does not match domain
# Prevents DNS rebinding attacks
enforce_domain = true
```
Great, let’s start writing some code in main.go to manipulate this file.
```go
package main

import (
    "fmt"
    "os"

    "gopkg.in/ini.v1"
)

func main() {
    cfg, err := ini.Load("my.ini")
    if err != nil {
        fmt.Printf("Fail to read file: %v", err)
        os.Exit(1)
    }

    // Classic read of values, default section can be represented as empty string
    fmt.Println("App Mode:", cfg.Section("").Key("app_mode").String())
    fmt.Println("Data Path:", cfg.Section("paths").Key("data").String())

    // Let's do some candidate value limitation
    fmt.Println("Server Protocol:",
        cfg.Section("server").Key("protocol").In("http", []string{"http", "https"}))
    // Value read that is not in candidates will be discarded and fall back to given default value
    fmt.Println("Email Protocol:",
        cfg.Section("server").Key("protocol").In("smtp", []string{"imap", "smtp"}))

    // Try out auto-type conversion
    fmt.Printf("Port Number: (%[1]T) %[1]d\n", cfg.Section("server").Key("http_port").MustInt(9999))
    fmt.Printf("Enforce Domain: (%[1]T) %[1]v\n", cfg.Section("server").Key("enforce_domain").MustBool(false))
    
    // Now, make some changes and save it
    cfg.Section("").Key("app_mode").SetValue("production")
    cfg.SaveTo("my.ini.local")
}
```
Almost there, let’s run this program and check the output.
```go
$ go run main.go
App Mode: development
Data Path: /home/git/grafana
Server Protocol: http
Email Protocol: smtp
Port Number: (int) 9999
Enforce Domain: (bool) true

$ cat my.ini.local
# possible values : production, development
app_mode = production

[paths]
# Path to where grafana can store temp files, sessions, and the sqlite3 db (if that is used)
data = /home/git/grafana
...
```
