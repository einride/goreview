package a

import "sync"

func Clean() {
	var mu sync.Mutex
	mu.Lock()
	defer mu.Unlock()
	mu.Unlock()
}

type muxStruct struct {
	mu sync.Mutex
}

type muxStruct2 struct {
	s muxStruct
}

func Clean2() {
	var s muxStruct
	s.mu.Lock()
	defer s.mu.Unlock()
	s.mu.Unlock()

	var s2 muxStruct2
	s2.s.mu.Lock()
	if true {
		s2.s.mu.Unlock()
	}
}
