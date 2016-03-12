

STマイクロエレクトロニクスの3軸加速度センサ LIS3DH にAVRからSPIで接続するサンプル．
同じくSTマイクロのジャイロセンサ L3GD20 も使えます．


- LIS3DH http://akizukidenshi.com/catalog/g/gK-06791/
- L3GD20 http://akizukidenshi.com/catalog/g/gK-06779/

## 利用方法

```
PC等 <--(シリアル)--> AVR <--(SPI)--> LIS3DH
```

- TODO: 回路図
- SPIで通信するのでSCK,MISO,MOSIをつないでください
- PB1, PB2 をCSにつなぐことで2デバイス接続できます
- シリアルからコマンドを受け付けます

近藤科学のB3Mシリーズのサーボモータと同一フォーマットのコマンドに返答します．(1.5Mbpsで通信するためにはAVRを12MHzで動作させる必要があります)

### メモリマップ

- 0x00 device id. (R/W)
- 0x10～0x4F sensor1 (センサのデータシート参照)
- 0x50～0x8F sensor2 (センサのデータシート参照)
- 0xA2～0xB2 device description (R)


### Golangから使う場合

- [sample.go](sample.go)
- https://github.com/binzume/gob3m を使っています

GetServo()とかしてますがサーボモータではないので加速度や角速度の値が取れます．


## ライセンス

Copyright 2016 Kousuke Kawahira

Released under the MIT license

