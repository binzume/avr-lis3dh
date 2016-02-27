package main

import (
	"github.com/binzume/gob3m/b3m"
	"github.com/tarm/serial"
	"flag"
	"time"
	"log"
)

func main() {
	var opt_port = flag.String("port", "COM1", "Serial port")
	var opt_id = flag.Int("id", 0, "servo id")
	flag.Parse()

	id := byte(*opt_id)

	s, err := serial.OpenPort(&serial.Config{Name: *opt_port, Baud: 1500000})
	if err != nil {
		log.Fatal(err)
	}

	servo := b3m.New(s).GetServo(id)

	err = servo.WriteMem(0x23, []byte{0x88})
	if err != nil {
		log.Fatal(err)
	}

	err = servo.WriteMem(0x20, []byte{0x77})
	if err != nil {
		log.Fatal(err)
	}

	axd := -37
	ayd := 18
	azd := 0

	for {
		acc, err := servo.ReadMem(0x28, 6)
		if err != nil {
			log.Printf("error %v", err)
		} else {
			ax := int(acc[0] >> 4)  + int(int8(acc[1])) << 4 + axd
			ay := int(acc[2] >> 4)  + int(int8(acc[3])) << 4 + ayd
			az := int(acc[4] >> 4)  + int(int8(acc[5])) << 4 + azd
			log.Printf("Acc: \t %v\t %v\t %v", ax, ay, az)
		}
		time.Sleep(20 * time.Millisecond)
	}
	log.Printf("ok")
}
