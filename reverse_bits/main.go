package main

// num = 964176192 (00111001011110000010100101000000)
func reverseBits(num uint32) uint32 {

	// 01010101010101010101010101010101 10101010101010101010101010101010
	num = (num&0x55555555)<<1 | (num&0xAAAAAAAA)>>1

	// num = 00110110101101000001011010000000

	// 00110011001100110011001100110011 11001100110011001100110011001100
	num = (num&0x33333333)<<2 | (num&0xCCCCCCCC)>>2

	// num = 11001001111000010100100100100000

	// 00001111000011110000111100001111 11110000111100001111000011110000
	num = (num&0x0F0F0F0F)<<4 | (num&0xF0F0F0F0)>>4

	// num = 10011100000111101001010000000010

	// 00000000111111110000000011111111 11111111000000001111111100000000
	num = (num&0x00FF00FF)<<8 | (num&0xFF00FF00)>>8

	// num = 00011110100111000000001010010100

	// 00000000000000001111111111111111 11111111111111110000000000000000
	num = (num&0x0000FFFF)<<16 | (num&0xFFFF0000)>>16

	// num = 00000010100101000001111010011100

	return num
}
