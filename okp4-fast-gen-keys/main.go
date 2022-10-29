package main

import (
	"runtime"
	"strconv"
	"sync"

	"github.com/Megavolv/okp4lib"
	"github.com/sirupsen/logrus"
	flag "github.com/spf13/pflag"
)

var prefix string
var from int
var to int
var cpus int
var out string

func init() {
	flag.StringVar(&prefix, "prefix", "keys", "Set prefix for generated keys")
	flag.IntVar(&from, "from", 0, "Lower bound of the range")
	flag.IntVar(&to, "to", 100, "Upper limit of the range")
	flag.IntVar(&cpus, "cpus", runtime.NumCPU(), "Number of cores to use. By default - all cores")
	flag.Parse()
}

func main() {
	logger := logrus.New()
	logger.SetLevel(logrus.WarnLevel)

	chunk := (to - from) / cpus

	var wg sync.WaitGroup

	for n := 0; n < cpus; n++ {
		wg.Add(1)
		go func(n int) {
			okp4 := okp4lib.NewOkp4()
			start := from + n*chunk
			end := from + (n * chunk) + chunk

			// Особое условие для последнего потока в связи с погрешностью деления
			if n == cpus-1 {
				end = to
			}

			keys_name, index_name := GenFileNames(prefix, start, end)
			mydb := NewDb(keys_name, index_name, n)

			for i := start; i < end; i++ {
				key, err := okp4.CreateJsonedKey(prefix + strconv.Itoa(i))
				if err != nil {
					logger.Error(err)
					break
				}
				mydb.Add(key)
			}

			mydb.Flush()
			defer mydb.fkeys.Close()
			defer mydb.findex.Close()
			wg.Done()
		}(n)
	}

	wg.Wait()
}
