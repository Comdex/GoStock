package main

import (
	//"code.google.com/p/plotinum/plot"
	//"code.google.com/p/plotinum/plotter"
	"io/ioutil"
	"fmt"
	"strconv"
)

//2013 05 01 00 00 00

type Date struct {
    year uint32
	month uint16
	day uint16
	hour uint16
	minute uint16
	second uint16
}

type Entry struct {
	d Date
	bid float32
	ask float32
}


func main() {

	results, err := ioutil.ReadFile("GBPUSD/GBPUSD1d.txt")

	if err != nil {
		fmt.Println("Error reading GBPUSD1d.txt!")
		return
	}

	parseLine(results[0:50])
}

func parseLine(b []byte) (ret Entry) {

	//s := string(b[0:4])
	//fmt.Println(s)

	var s string
	var n int64
	
	s = string(b[0:4])
	n, _ = strconv.ParseInt(s, 10, 32)
	ret.d.year = uint32(n)

	s = string(b[4:6])
	n, _ = strconv.ParseInt(s, 10, 32)
	ret.d.month = uint16(n)

	s = string(b[6:8])
	n, _ = strconv.ParseInt(s, 10, 32)
	ret.d.day = uint16(n)

	s = string(b[8:10])
	n, _ = strconv.ParseInt(s, 10, 32)
	ret.d.hour = uint16(n)

	s = string(b[10:12])
	n, _ = strconv.ParseInt(s, 10, 32)
	ret.d.minute = uint16(n)

	s = string(b[12:14])
	n, _ = strconv.ParseInt(s, 10, 32)
	ret.d.second = uint16(n)

	s = 
	
	
	return
}
