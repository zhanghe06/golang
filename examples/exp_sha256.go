package main

import (
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"strings"
)

const (
	FirstFactor = 20
	SecondFactor = 40
	ThirdFactor = 80
)

func CountHash(docId string) string  {
	sha := sha256.New()
	sha.Write([]byte(docId))
	bytes := sha.Sum(nil)
	println(binary.LittleEndian.Uint32(bytes[:]))

	FirstLevel := binary.LittleEndian.Uint32(bytes[8:]) % FirstFactor
	SecondLevel := binary.LittleEndian.Uint32(bytes[8:]) % SecondFactor
	ThirdLevel := binary.LittleEndian.Uint32(bytes[8:]) % ThirdFactor

	//println(binary.LittleEndian.Uint32(bytes[0:]))
	//println(binary.LittleEndian.Uint32(bytes[8:]))
	//println(binary.LittleEndian.Uint32(bytes[16:]))
	//println(FirstLevel)
	//println(SecondLevel)
	//println(ThirdLevel)
	res := fmt.Sprintf("%02d/%02d/%02d", FirstLevel, SecondLevel, ThirdLevel)
	println(res)
	return res
}


func CountHashNew(docId string) string  {
	sha := sha256.New()
	sha.Write([]byte(strings.ToUpper(docId)))
	bytes := sha.Sum(nil)

	fmt.Println(bytes[0:8])
	fmt.Println(bytes[8:16])
	fmt.Println(bytes[16:24])
	fmt.Println(bytes[24:32])

	//FirstLevel := binary.LittleEndian.Uint32(bytes[0:]) % 256
	SecondLevel := binary.LittleEndian.Uint32(bytes[8:]) % 256
	ThirdLevel := binary.LittleEndian.Uint32(bytes[16:]) % 256
	FourthLevel := binary.LittleEndian.Uint32(bytes[24:]) % 256

	//res := fmt.Sprintf("%02X/%02X/%02X/%02X", FirstLevel, SecondLevel, ThirdLevel, FourthLevel)
	res := fmt.Sprintf("%02X/%02X/%02X", SecondLevel, ThirdLevel, FourthLevel)
	println(res)
	println(strings.ToUpper(docId))
	return res
}


func main() {
	//CountHash("000C292555BF1EECB189BEC4DA91FB79")
	//CountHashNew("000C292555BF1EECB189BEC4DA91FB79")
	//CountHashNew("123222556778906")
	//CountHashNew("xxx")
	//CountHashNew("000C292555BF1EDDAE82D7835984E73F")
	//CountHashNew("000C292555BF1EDDAE834EFF7833073F")
	//CountHashNew("000C292555BF1EDDBOFA54F2F54B073F")
	//CountHashNew("000C292555BF1EDDB0FDB8ACC4DDC73F")
	//CountHashNew("000C292555BF1EDDB18C07360F17073F")
	//CountHashNew("000C292555BF1EDDB190D6D55808073E")
	CountHashNew("000C292555BF1EDDB192581B05B4E73F")
}