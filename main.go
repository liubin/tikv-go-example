package main

import (
	"log"
	"flag"
	"fmt"
)

var (
	pdAddr = flag.String("pd", "localhost:2379", "pd address:localhost:2379")
)


func main() {

	flag.Parse()

	tikv,err := NewTiKVServer(*pdAddr)
	if err != nil{
		log.Fatalf("init tikv error %s", err.Error())
	}

	tikv.Set("aaa", "abc")
	s, _ := tikv.Get("aaa")
	fmt.Println(string(s))

	tikv.Set("aaa", "123")
	s, _ = tikv.Get("aaa")
	fmt.Println(string(s))
}
