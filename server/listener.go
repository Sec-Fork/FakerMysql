/*
 * @Date: 2021-11-23 10:27:42
 * @LastEditors: recar
 * @LastEditTime: 2021-11-23 11:08:26
 */
// Copyright 2020-2021 Dolthub, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package server

import (
	"net"
)

type Listener struct {
	net.Listener
	h *Handler
}

// NewListener creates a new Listener.
func NewListener(protocol, address string, handler *Handler) (*Listener, error) {
	// l, err := net.Listen(protocol, address)
	// if err != nil {
	// 	return nil, err
	// }
	var l net.Listener
	return &Listener{l, handler}, nil
}

func (l *Listener) Accept() (net.Conn, error) {
	return l.Listener.Accept()
}
