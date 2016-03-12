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
	var opt_id = flag.Int("id", 0, "device id")
	flag.Parse()
	offset := 0x10

	id := byte(*opt_id)

	s, err := serial.OpenPort(&serial.Config{Name: *opt_port, Baud: 1500000})
	if err != nil {
		log.Fatal(err)
	}

	servo := b3m.New(s).GetServo(id)

	whoami, err := servo.ReadMem(0x0f + offset, 1)
	if err != nil {
		log.Fatal(err)
	}

	axd := 0
	ayd := 0
	azd := 0

	if whoami[0] == 0x33 {
		// LIS3DH Accelometer
		err = servo.WriteMem(0x23 + offset, []byte{0x88})
		if err != nil {
			log.Fatal(err)
		}

		err = servo.WriteMem(0x20 + offset, []byte{0x77})
		if err != nil {
			log.Fatal(err)
		}
		axd = -37
		ayd = 18
		azd = 0
	} else if whoami[0] == 0xd4{
		// L3GD20 Gyro sensor(experimental)
		err = servo.WriteMem(0x23 + offset, []byte{0x10}) // 00h:250dps 10h:500dps 20h:200dps
		if err != nil {
			log.Fatal(err)
		}

		err = servo.WriteMem(0x20 + offset, []byte{0x0f})
		if err != nil {
			log.Fatal(err)
		}
		axd = 16
		ayd = -15
		azd = 3
	} else {
		// Unknown device
		log.Fatal(whoami)
	}

	for {
		acc, err := servo.ReadMem(0x28 + offset, 6)
		if err != nil {
			log.Printf("error %v", err)
		} else {
			ax := int(acc[0] >> 4)  + int(int8(acc[1])) << 4 + axd
			ay := int(acc[2] >> 4)  + int(int8(acc[3])) << 4 + ayd
			az := int(acc[4] >> 4)  + int(int8(acc[5])) << 4 + azd
			log.Printf("Acc: %v\t %v\t %v", ax, ay, az)
		}
		time.Sleep(20 * time.Millisecond)
	}
	log.Printf("ok")
}
