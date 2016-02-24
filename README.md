

３軸加速度センサ LIS3DH にAVRからSPIで接続するサンプル．

http://akizukidenshi.com/catalog/g/gK-06791/


シリアルからコマンドを受け付けてSPIに渡すだけです．

```
PC <--(シリアル)--> AVR <--(SPI)--> LIS3DH
```

近藤科学のB3Mシリーズのサーボモータと同一フォーマットのコマンドに返答します．(1.5Mbpsで通信するためにはAVRを12MHzで動作させる必要があります)


Golangから使う場合：

https://github.com/binzume/gob3m


```go
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

	s, err := serial.OpenPort(&serial.Config{Name: *opt_port, Baud: 1500000, ReadTimeout: 100 * time.Millisecond})
	if err != nil {
		log.Fatal(err)
	}

	servo := b3m.GetServo(s, id)

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
			log.Fatal(err)
		}
		ax := int(acc[0] >> 4)  + int(int8(acc[1])) << 4 + axd
		ay := int(acc[2] >> 4)  + int(int8(acc[3])) << 4 + ayd
		az := int(acc[4] >> 4)  + int(int8(acc[5])) << 4 + azd
		log.Printf("Acc: \t %v\t %v\t %v", ax, ay, az)

		time.Sleep(10 * time.Millisecond)
	}
	log.Printf("ok")
}
```

GetServo()とかしてますがサーボモータではないので加速度の値が取れます．


