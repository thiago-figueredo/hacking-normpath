Hey guys, i have tried to bypass python 3.12 builtin http file server (`python3 -m http.server <port> `).
I was intrigued that we cannot read a file that is outside the base path of the server. 
If the server has a base path like `/home/foo/bar`, we can only access files or directories inside it, and /home/foo/bar/../../../etc/passwd it's redirect to `/home/foo/bar`.
I've read the python 3.12 source code and it seems it's using a function called normpath from a posix module. 
I've not found the source code so i decided to implement it myself. 
I've created a simple cli project in go to simulate the server, instead of creating a http server it only read the filepath that you put in the cli args. 
Can someone find a bug or a vulnerability in my code or in the python 3.12 simple http server ?

