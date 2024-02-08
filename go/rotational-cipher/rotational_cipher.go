package rotationalcipher

// RotationalCipher encrypts a string using a shift key, shifts ASCII only
func RotationalCipher(plain string, shiftKey int) string {
	shiftKey = shiftKey % 26
	shifted := make([]byte, len(plain))
	for i := 0; i < len(plain); i++ {
		if plain[i] >= 'a' && plain[i] <= 'z' {
			shifted[i] = shiftByte(plain[i], 'a', shiftKey)
		} else if plain[i] >= 'A' && plain[i] <= 'Z' {
			shifted[i] = shiftByte(plain[i], 'A', shiftKey)
		} else {
			shifted[i] = plain[i]
		}
	}
	return string(shifted)
}

func shiftByte(cur byte, init byte, shift int) byte {
	return (cur-init+byte(shift))%26 + init
}
