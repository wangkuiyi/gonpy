package gonpy

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"log"

	. "github.com/topicai/candy"
	"github.com/wangkuiyi/gonpy/header"
)

func Load(r *bufio.Reader) (matrix *Matrix, err error) {
	defer func() {
		// gonpy/header.Parse calls lexer and parser, which panic for errors.
		if e := recover(); e != nil {
			matrix = nil
			err = fmt.Errorf("%v", e)
		}
	}()

	magic := make([]byte, 6)
	read(r, magic)
	if string(magic) != "\x93NUMPY" {
		return nil, fmt.Errorf("Unknow npy magic code: %s", magic)
	}

	var major, minor uint8
	var headerSize uint16
	binary.Read(r, binary.LittleEndian, &major)
	binary.Read(r, binary.LittleEndian, &minor)
	binary.Read(r, binary.LittleEndian, &headerSize)

	headerBuf := make([]byte, headerSize)
	read(r, headerBuf)

	h := header.Parse(bytes.NewReader(headerBuf))
	dtype := h["descr"].(string)
	shape := h["shape"].(*header.Shape)
	ret := &Matrix{
		Shape: shape,
		Data:  make([]float64, shape.Row*shape.Col),
	}

	if h["fortran_order"].(bool) {
		for col := 0; col < shape.Col; col++ {
			for row := 0; row < shape.Row; row++ {
				ret.Data[row*shape.Col+col] = readElement(r, dtype)
			}
		}
	} else {
		for row := 0; row < shape.Row; row++ {
			for col := 0; col < shape.Col; col++ {
				ret.Data[row*shape.Col+col] = readElement(r, dtype)
			}
		}
	}

	return ret, nil
}

func read(r *bufio.Reader, buf []byte) {
	n, e := r.Read(buf)

	if e != nil {
		log.Panicf("Reading magic error %v", e)
	}
	if n != len(buf) {
		log.Panicf("Read magic with %d bytes", n)
	}
}

func readElement(r *bufio.Reader, dtype string) float64 {
	var bo binary.ByteOrder = binary.LittleEndian
	if dtype[0] == '>' {
		bo = binary.BigEndian
	}

	switch dtype[1:] {
	case "f4":
		var v float32
		Must(binary.Read(r, bo, &v))
		return float64(v)
	case "f8":
		var v float64
		Must(binary.Read(r, bo, &v))
		return v
	case "i8":
		var v int64
		Must(binary.Read(r, bo, &v))
		return float64(v)
	case "i4":
		var v int32
		Must(binary.Read(r, bo, &v))
		return float64(v)
	default:
		log.Panicf("Unknown dtype/descr %v", dtype)
	}

	return 0 // Not really reachable.
}
