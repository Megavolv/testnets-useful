package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"os"
)

func main() {

	var err error
	k, err := os.OpenFile("prefix.json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		panic(err)
	}
	defer k.Close()

	i, err := os.OpenFile("prefix.idx", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		panic(err)
	}
	defer i.Close()

	k0, err := os.OpenFile("prefix.json.0", os.O_RDONLY, 0755)
	if err != nil {
		panic(err)
	}
	defer k0.Close()

	k1, err := os.OpenFile("prefix.json.1", os.O_RDONLY, 0755)
	if err != nil {
		panic(err)
	}
	defer k1.Close()

	i0, err := os.OpenFile("prefix.idx.0", os.O_RDONLY, 0755)
	if err != nil {
		panic(err)
	}
	defer i0.Close()

	i1, err := os.OpenFile("prefix.idx.1", os.O_RDONLY, 0755)
	if err != nil {
		panic(err)
	}
	defer i1.Close()

	// копируем первую пачку ключей
	keyBytes, err := io.Copy(k, k0)
	if err != nil {
		panic(err)
	}

	//keyBytes - объем первой пачки
	fmt.Println("keyBytes", keyBytes)
	// копируем первую часть индекса без изменений
	_, err = io.Copy(i, i0)
	if err != nil {
		panic(err)
	}

	// создаем слайс для чтения второго индекса целиком
	ibuf := make([]byte, 16000) //TODO

	l1, err := i1.Read(ibuf)
	if err != nil {
		panic(err)
	}
	// l1 - длина прочитанного индекса

	fixbuf := make([]byte, l1)

	for i := 0; i < l1/4; i++ {
		b := binary.LittleEndian.Uint32(ibuf[i*4 : i*4+4])
		binary.LittleEndian.PutUint32(fixbuf[i*4:i*4+4], b+uint32(keyBytes))
	}

	_, err = i.Write(fixbuf)
	if err != nil {
		panic(err)
	}

	// копируем первую пачку ключей
	_, err = io.Copy(k, k1)
	if err != nil {
		panic(err)
	}

}
