package main

import "io"

type ascii struct {
	w io.Writer
	r io.Reader
}

func (a *ascii) Write(p []byte) (int, error) {
	buf := make([]byte, 0, len(p))
	var lastB byte

	for _, b := range p {
		if b == '\n' && lastB != '\r' {
			buf = append(buf, '\r')
		}
		buf = append(buf, b)
		lastB = b
	}

	n, err := a.w.Write(buf)
	if n > len(p) {
		return len(p), err
	}
	return n, err
}

func (a *ascii) Read(p []byte) (int, error) {
	buf := make([]byte, len(p))

	n, err := a.r.Read(buf)
	next := 0

	for i := 0; i < n; i++ {
		if buf[i] == '\r' {
			continue
		}
		p[next] = buf[i]
		next++
	}
	return next, err
}
