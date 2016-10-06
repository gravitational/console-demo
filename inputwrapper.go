/*
Copyright 2015 Gravitational, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package main

import (
	"bytes"
	"io"

	"golang.org/x/net/websocket"
)

// InputWrapper implements io.Reader interface and it wraps the web
// socket. This is why:
//
// something inside term.js causes to inject a strange sequence of fake
// keyboard inputs into the websocket. to reproduce, launch vim inside
// a web browser when you're on OS X.
//
// I never figured out where this sequence comes from, but this wrapper
// simply intercepts & drops these sequences
type InputWrapper struct {
	ws *websocket.Conn
}

const patternLen = 5

// ignoredInputs are the strange input bytes that we look for and drop
var ignoredInputs = [][patternLen]byte{
	{27, 91, 62, 48, 59},
	{27, 80, 48, 43, 114},
}

func (this *InputWrapper) Read(out []byte) (n int, err error) {
	var data []byte
	err = websocket.Message.Receive(this.ws, &data)
	if err != nil {
		return 0, io.EOF
	}

	if len(data) >= patternLen {
		for i := range ignoredInputs {
			pattern := ignoredInputs[i]
			if bytes.Equal(pattern[:], data[:patternLen]) {
				return 0, nil
			}
		}
	}
	return copy(out, data), nil
}
