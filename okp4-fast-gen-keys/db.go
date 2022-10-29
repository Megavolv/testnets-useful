package main

import (
	"encoding/binary"
	//"fmt"
	"strconv"

	//"github.com/davecgh/go-spew/spew"

	//"fmt"
	"os"
	"strings"
)

//"os"

const TRIGGER_THRESHOLD = 5

type Db struct {
	KeysName     string
	IndexName    string
	buffer       []string
	counter, num int
	offset       uint32
	stop         chan bool
	fkeys        *os.File

	findex *os.File
	index  []uint32
}

func NewDb(name string, num int) *Db {
	db := &Db{
		KeysName:  name + ".json." + strconv.Itoa(num),
		IndexName: name + ".idx." + strconv.Itoa(num),
		num:       num,
	}

	var err error

	db.fkeys, err = os.OpenFile(db.KeysName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		panic(err)
	}

	db.findex, err = os.OpenFile(db.IndexName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		panic(err)
	}

	db.buffer = make([]string, 16)
	db.index = make([]uint32, 16)

	db.stop = make(chan bool)
	return db
}

func (db *Db) Flush() {

	data := strings.Join(db.buffer[0:db.counter], "")
	_, err := db.fkeys.Write([]byte(data))
	if err != nil {
		panic(err)
	}

	a := make([]byte, db.counter*4)
	for i := 0; i < db.counter; i++ {
		binary.LittleEndian.PutUint32(a[i*4:i*4+4], db.index[i])
	}

	_, err = db.findex.Write(a)
	if err != nil {
		panic(err)
	}

	db.counter = 0
}

func (db *Db) Add(key string) {
	key = key + "\n"

	db.buffer[db.counter] = key
	db.index[db.counter] = db.offset

	//fmt.Println(db.offset, strconv.FormatUint(uint64(db.offset), 16))

	db.offset += uint32(len(key))
	db.counter += 1

	if db.counter > TRIGGER_THRESHOLD {
		//fmt.Println("Flush по триггеру")
		db.Flush()
	}

	if int(db.counter) == len(db.buffer) {
		//fmt.Println("Flush по заполнению буфера")
		//	spew.Dump(db.index)
		db.Flush()
	}

}
