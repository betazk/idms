package imds

import (
	"bufio"
	"io"
)

type Serializer struct {
	bw *bufio.Writer
}

func NewSerializer(rw io.Writer) *Serializer {
	w := bufio.NewWriter(rw)
	return &Serializer{bw: w}
}

func (s *Serializer) WriteByte(value byte) {
	s.bw.WriteByte(value)
}

func (s *Serializer) WriteBytes(value []byte) {
	s.bw.Write(value)
}

func (s *Serializer) WriteString(value string) {
	s.bw.WriteString(value)
}

func (s *Serializer) WriteInt32(value int32) {
	s.WriteByte(byte(value >> 24 & 0xff))
	s.WriteByte(byte(value >> 16 & 0xff))
	s.WriteByte(byte(value >> 8 & 0xff))
	s.WriteByte(byte(value & 0xff))
}

func (s *Serializer) WriteUInt32(value uint32) {
	s.WriteInt32(int32(value))
}

func (s *Serializer) WriteInt64(value int64) {
	s.WriteByte(byte(value >> 56 & 0xff))
	s.WriteByte(byte(value >> 48 & 0xff))
	s.WriteByte(byte(value >> 40 & 0xff))
	s.WriteByte(byte(value >> 32 & 0xff))
	s.WriteByte(byte(value >> 24 & 0xff))
	s.WriteByte(byte(value >> 16 & 0xff))
	s.WriteByte(byte(value >> 8 & 0xff))
	s.WriteByte(byte(value & 0xff))
}

func (s *Serializer) WriteUInt64(value uint64) {
	s.WriteInt64(int64(value))
}

func (s *Serializer) WriteBool(value bool) {
	if value {
		s.WriteByte(byte(1))
	} else {
		s.WriteByte(byte(0))
	}
}

func (s *Serializer) Flush() {
	s.bw.Flush()
}
