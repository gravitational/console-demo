# Web-Based Console Demo

This demo was built for the Gitlab team.

It opens a web-based console connected to a `/bin/bash`
running in a virtual terminal on the server side.

### Build

Simply type `make`.

### Run

Simply run `./demo` and then open http://localhost:5000 
You will get a web-based terminal.


### TODO

See `main.go` and replace the path of `/bin/bash` with
the given `kubectl` command.

That's the simplest implementation of this.


### Possible "nice to have" improvements

Detect "window resize" in a browser and pass the new window 
size into the server via AJAx.

On the server side, call SetSize() on the virtual TTY.
