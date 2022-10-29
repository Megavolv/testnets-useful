package main

import (
	"encoding/binary"

	"os"
	"strconv"
	"strings"
)

func GetKeysByOneFile(fkey, findex *os.File, start, end int64) (string, error) {
	a, err := GetKeyPosition(findex, start)
	if err != nil {
		return "", err
	}

	b, err := GetKeyPosition(findex, end)
	if err != nil {
		info, _ := fkey.Stat()
		b = uint32(info.Size())
	}

	_, err = fkey.Seek(int64(a), 0)
	if err != nil {
		panic(err)
	}

	kbuf := make([]byte, b-a-1)

	_, err = fkey.Read(kbuf)
	if err != nil {
		panic(err)
	}

	return string(kbuf), nil
}

func GetKeyPosition(findex *os.File, position int64) (key_offset uint32, err error) {
	_, err = findex.Seek(position*4, 0)
	if err != nil {
		return 0, err
	}

	buf := make([]byte, 4)

	_, err = findex.Read(buf)
	if err != nil {
		return 0, err
	}

	key_offset = binary.LittleEndian.Uint32(buf[0:4])
	return
}

func LoadFile(name string) *File {
	data := strings.Split(name, ".")

	start_end := strings.Split(data[1], "-")
	start, err := strconv.ParseInt(start_end[0], 10, 32)

	if err != nil {
		panic(err)
	}

	end, err := strconv.ParseInt(start_end[1], 10, 32)
	if err != nil {
		panic(err)
	}

	f := &File{
		Name:  name,
		Start: start,
		End:   end,
	}

	f.f, err = os.OpenFile(name, os.O_RDONLY, 0755)
	if err != nil {
		panic(err)
	}

	return f
}
