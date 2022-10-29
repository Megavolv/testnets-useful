package main

import (
	"encoding/binary"
	"fmt"

	//"strconv"

	"os"
	"strings"
)

const TRIGGER_THRESHOLD = 4096 // TODO

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

//Генерация имени для файла с ключами и индексом
func GenFileNames(prefix string, start, end int) (keys_name string, index_name string) {
	return fmt.Sprintf("%s.%d-%d.json", prefix, start, end), fmt.Sprintf("%s.%d-%d.idx", prefix, start, end)
}

func NewDb(keysname, indexname string, num int) *Db {
	db := &Db{
		KeysName:  keysname,
		IndexName: indexname,
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

	db.offset += uint32(len(key))
	db.counter += 1

	if db.counter > TRIGGER_THRESHOLD {
		//fmt.Println("Flush по триггеру")
		db.Flush()
	}

	if int(db.counter) == len(db.buffer) {
		//fmt.Println("Flush по заполнению буфера")
		db.Flush()
	}

}
