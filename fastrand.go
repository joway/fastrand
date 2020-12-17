package fastrand

import (
	"math/rand"
	"sync"
	"time"
)

type poolSource struct {
	p *sync.Pool
}

func (s *poolSource) Int63() int64 {
	v := s.p.Get()
	defer s.p.Put(v)
	return v.(rand.Source).Int63()
}

func (s *poolSource) Seed(seed int64) {
	v := s.p.Get()
	defer s.p.Put(v)
	v.(rand.Source).Seed(seed)
}

func (s *poolSource) Uint64() uint64 {
	v := s.p.Get()
	defer s.p.Put(v)
	return v.(rand.Source64).Uint64()
}

func newPoolSource() *poolSource {
	s := &poolSource{}
	p := &sync.Pool{New: func() interface{} {
		return rand.NewSource(time.Now().Unix())
	}}
	s.p = p
	return s
}

//New create a thread-safe rand.Rand object
func New() *rand.Rand {
	return rand.New(newPoolSource())
}

//NewUnsafety create a non-thread-safe rand.Rand object with current time seed
func NewUnsafety() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().Unix()))
}
