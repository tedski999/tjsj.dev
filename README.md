# tjsj.dev
### Yet another CS students personal web server

currently under development\
its very cool

## Setting up & Running

Clone the repository
```
$ git clone https://github.com/tedski999/tjsj.dev.git
$ cd tjsj.dev
```

Link the CA certificates
```
$ ln -s /wherever/the/live/certs/are/ ./web/certs
```

Compiling the Go program
```
$ make
```

Allowing program to bind to protected ports
```
# setcap CAP_NET_BIND_SERVICE=+eip ./bin/tjsj.dev
```

Starting the server
```
$ ./bin/tjsj.dev
```

The website should now be hosted on your machine accessible via HTTPS:
https://localhost/

