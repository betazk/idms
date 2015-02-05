package imds

import (
	"bufio"
	"io"
)

type Deserializer struct {
	br *bufio.Reader
}

func NewDeserializer(rd io.Reader) *Deserializer {
	b := bufio.NewReader(rd)
	return &Deserializer{br: b}
}

func (d *Deserializer) ReadByte() byte {
	c, _ := d.br.ReadByte()
	return c
}

func (d *Deserializer) ReadBytes(delim byte) []byte {
	bs, _ := d.br.ReadBytes(delim)
	return bs
}

func (d *Deserializer) ReadString(delim byte) string {
	str, _ := d.br.ReadString(delim)
	return str
}

func (d *Deserializer) ReadInt32() int32 {
	return int32(d.ReadByte())<<24 | int32(d.ReadByte())<<16 |
		int32(d.ReadByte())<<8 | int32(d.ReadByte())
}

func (d *Deserializer) ReadUInt32() uint32 {
	return uint32(d.ReadInt32())
}

func (d *Deserializer) ReadInt64() int64 {
	return int64(d.ReadByte())<<56 | int64(d.ReadByte())<<48 |
		int64(d.ReadByte())<<40 | int64(d.ReadByte())<<32 |
		int64(d.ReadByte())<<16 | int64(d.ReadByte())<<8 |
		int64(d.ReadByte())
}

func (d *Deserializer) ReadUInt64() uint64 {
	return uint64(d.ReadInt64())
}

func (d *Deserializer) ReadBool() bool {
	return d.ReadByte() == 1
}
