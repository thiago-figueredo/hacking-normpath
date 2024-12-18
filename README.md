# Motivation

Hey guys, i have tried to bypass python 3.12 builtin http file server (`python3 -m http.server <port> `).
I was intrigued that we cannot read a file that is outside the base path of the server. 
If the server has a base path like `/home/foo/bar`, we can only access files or directories inside it, and /home/foo/bar/../../../etc/passwd it's redirect to `/home/foo/bar`.
I've read the python 3.12 source code and it seems it's using a function called normpath from a posix module. 
I've not found the source code so i decided to implement it myself. 
I've created a simple cli project in go to simulate the server, instead of creating a http server it only read the filepath that you put in the cli args. 
Can someone find a bug or a vulnerability in my code or in the python 3.12 simple http server ?


# Usage

```bash
go run main.go <path>
```

# Example

```bash
└──╼ $ go run main.go .
Entries of `/home/levrone/Documents/github/golang/normpath`: 

d--------- Dec 17 22:11 .git
---------- Dec 17 22:12 README.md
---------- Dec 17 16:18 go.mod
---------- Dec 17 16:18 go.sum
---------- Dec 17 21:57 main.go
d--------- Dec 17 16:27 pathlib
```

