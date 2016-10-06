.PHONY: demo
demo: *.go deps
	go build -o demo .


.PHONY:deps
deps:
	go get github.com/kr/pty
	go get golang.org/x/net/websocket
