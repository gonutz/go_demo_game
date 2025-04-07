package ds

func SPEAKER_COMBINED(c, g int) uint32 {
	return uint32(byte(c)) | uint32(byte(g))<<16
}

func SPEAKER_CONFIG(a int) byte {
	return byte(a)
}

func SPEAKER_GEOMETRY(a int) byte {
	return byte((uint32(a) >> 16) & 0x00FF)
}
