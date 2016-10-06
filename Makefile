.PHONY: demo
demo: *.go deps
	go build -o demo .


.PHONY:deps
deps:
	go get github.com/kr/pty
	go get golang.org/x/net/websocket
	go get golang.org/x/net/websocket
	go get golang.org/x/text/encoding
	go get golang.org/x/text/encoding/unicode

