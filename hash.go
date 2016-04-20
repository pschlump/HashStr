package HashStr

//
// Hash string into number or into short name.
//
// (C) Philip Schlump, 2013-2015.
// MIT Licensed
//

// Hashes a string into a number
func HashStr(s []byte) (n int) {
	h := len(s)
	for _, c := range s {
		h += int(c)
		h += (h << 3)
		h = h ^ (h >> 11)
		h += (h << 15)
		n = h
	}
	return
}

// Hashes a string into a set of 5 bytes - a name for the string.
func HashStrToName(s string) (rv string) {
	n := HashStr([]byte(s))
	rv = string(((n & 0xff) + 1) & 0xff)
	n >>= 8
	rv += string(((n & 0xff) + 1) & 0xff)
	n >>= 8
	rv += string(((n & 0xff) + 1) & 0xff)
	n >>= 8
	rv += string(((n & 0xff) + 1) & 0xff)
	n >>= 8
	rv += string(((n & 0xff) + 1) & 0xff)
	return
}
