package main

import (
	"encoding/binary"
	"fmt"
	"path/filepath"
	"strings"

	//"time"
	"os"

	flag "github.com/spf13/pflag"
	//"github.com/davecgh/go-spew/spew"
)

var keysname string

func init() {
	flag.StringVar(&keysname, "keysname", "", "Keys name")
	flag.Parse()
}

type Index struct {
	data     []uint64
	position uint64
}

func NewIndex() *Index {
	index := &Index{data: make([]uint64, 62500001)}
	index.Add(0)
	return index
}

func (i *Index) Add(offset uint64) {
	i.data[i.position] = offset
	i.position++
}

func (i *Index) Last() uint64 {
	return i.data[i.position-1]
}

func read() {}

func main() {
	fkeys, err := os.Open(keysname)
	if err != nil {
		panic(err)
	}
	defer fkeys.Close()

	name := strings.TrimSuffix(keysname, filepath.Ext(keysname))
	findex, err := os.OpenFile(name+".idx", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		panic(err)
	}
	defer findex.Close()

	index := NewIndex()

	var count uint64 = 128000000

	var offset int64

	for {
		buf := make([]byte, count)
		got, err := fkeys.Read(buf)
		if err != nil {
			fmt.Errorf(err.Error())
			break
		}

		for i := 0; i < got; i++ {
			if buf[i] == 10 {
				index.Add(uint64(offset) + uint64(i) + 1)
				if index.position%100000 == 0 {
					fmt.Println(index.position)
				}
			}
		}

		sub := int64(got) - (int64(index.Last()) - int64(offset))

		//fmt.Println("sub ", sub)
		//fmt.Println("got ", got)
		//return
		//fmt.Println(string(buf[0 : int64(got)-sub]))

		offset, err = fkeys.Seek(0-sub, 1)
		if err != nil {
			panic(err)
		}

		//time.Sleep(1 * time.Second)
	}

	bytedata := make([]byte, index.position*8)
	var i uint64
	for i = 0; i < index.position; i++ {
		binary.LittleEndian.PutUint64(bytedata[i*8:i*8+8], index.data[i])
	}

	_, err = findex.Write(bytedata)
	if err != nil {
		fmt.Errorf(err.Error())
	}

}
