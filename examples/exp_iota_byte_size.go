package main

type ByteSize float64

const (
	_           = iota             // ignore first value by assigning to blank identifier
	KB ByteSize = 1 << (10 * iota) // 1 << (10*1) 2的10次方
	MB                             // 1 << (10*2)
	GB                             // 1 << (10*3)
	TB                             // 1 << (10*4)
	PB                             // 1 << (10*5)
	EB                             // 1 << (10*6)
	ZB                             // 1 << (10*7)
	YB                             // 1 << (10*8)
)

func main() {
	println(KB) // +1.024000e+003
	println(MB) // +1.048576e+006
	println(GB) // +1.073742e+009
	println(TB) // +1.099512e+012
	println(PB) // +1.125900e+015
	println(EB) // +1.152922e+018
	println(ZB) // +1.180592e+021
	println(YB) // +1.208926e+024
}
