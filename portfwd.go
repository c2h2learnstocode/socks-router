package main

import (
	"fmt"
	"io"
	"net"
	"os"
	//"bufio"
	"bytes"
	//"ioutil"

)

func main() {
	fmt.Println(os.Args)
	if len(os.Args) != 4 {
		fmt.Println("prog listen_port remote_host, remote_port")
		os.Exit(3)

        }
	listen_port := os.Args[1]
	rmt_host := os.Args[2]
	rmt_port := os.Args[3]
	fmt.Println("localhost:" +  listen_port + "  -> " + rmt_host+":"+rmt_port)
	ln, err := net.Listen("tcp", ":"+listen_port)
	if err != nil {
		panic(err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}

		go handleRequest(conn, rmt_host, rmt_port)
	}
}

func handleRequest(conn net.Conn, rmt_host string, rmt_port string) {

	proxy, err := net.Dial("tcp", rmt_host+":"+rmt_port)
	if err != nil {
		panic(err)
	}


//	buffer := make([]byte, 60000)
//	conn.Read(buffer)
//	s := string(buffer)
//	fmt.Println(s)

	go copyIO(conn, proxy) //local to remote
	go copyIO(proxy, conn)
}

func copyIO(src, dest net.Conn) {
	defer src.Close()
	defer dest.Close()

	var b bytes.Buffer
    _ = io.Writer(&b)
	io.Copy(src, io.TeeReader(dest, &b))


	fmt.Println(b.String())
}