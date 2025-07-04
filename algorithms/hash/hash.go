package hash

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
)

/* NOTE:
These algorithms are all written by AI for me to study and slightly adjusted (mainly the conversion to a byte array)
and some comments from my side.
It's hard for me to confirm if they are correct at the moment,
since actual implementations are often way more complicated
because of interface contract reasons, security and other unkown reasons.
And also my current lack of fully understanding the papers for these hashing algorithms.

Online converters are giving different results for reasons I do not fully understand. Perhaps because they are taking strings
as input, whereas I am converting any input into a []byte to be able to use any data type as input,
or also because of different undisclosed prime numbers and seeds. 

For this reason, I have not written any tests.


However, the results provided are pretty good and definetly usable to proceed with implementing a hash-map.
Especially the Murmur3 implementation seems to have very good distribution when I tested it, so I use it for my hash-map implementation.

I will revisit hashing algorithms at a later point.
*/

// DJB2 This implementation looks to be correct to my knowledge, but online converters yield different results, probably due to my input to []byte conversion
func DJB2(input any) (uint32, error) {
	in, err := convertToByteArray(input)
	if err != nil {
		return 0, err
	}

	var hash uint32 = 5381
	for i := 0; i < len(in); i++ {
		// ((hash << 5) + hash) is simply the same as hash * 33. But with the bit shift, it is faster on many cpus.
		hash = ((hash << 5) + hash) + uint32(in[i])
	}

	return hash, nil
}

func XxHash(input any) (uint64, error) {
	in, err := convertToByteArray(input)
	if err != nil {
		return 0, err
	}

	prime := uint64(11400714785074694791)
	var hash uint64 = 0

	for i := 0; i < len(in); i++ {
		hash = hash*prime + uint64(in[i])
	}

	return hash, nil
}

// Murmur3 should correctly calculate a hash for almost all types. One exception is type chan (any) which is not supported.
func Murmur3(input any, seed uint32) (uint32, error) {
	in, err := convertToByteArray(input)
	if err != nil {
		return 0, err
	}

	const (
		c1 = 0xcc9e2d51
		c2 = 0x1b873593
		r1 = 15
		r2 = 13
		m  = 5
		n  = 0xe6546b64
	)

	hash := seed
	nblocks := len(in) / 4

	for i := 0; i < nblocks; i++ {
		k := binary.LittleEndian.Uint32(in[i*4 : (i+1)*4])
		k *= c1
		k = (k << r1) | (k >> (32 - r1))
		k *= c2

		hash ^= k
		hash = ((hash<<r2)|(hash>>(32-r2)))*m + n
	}

	tail := in[nblocks*4:]
	k1 := uint32(0)

	switch len(in) & 3 {
	case 3:
		k1 ^= uint32(tail[2]) << 16
		fallthrough
	case 2:
		k1 ^= uint32(tail[1]) << 8
		fallthrough
	case 1:
		k1 ^= uint32(tail[0])
		k1 *= c1
		k1 = (k1 << r1) | (k1 >> (32 - r1))
		k1 *= c2
		hash ^= k1
	}

	hash ^= uint32(len(in))
	hash ^= hash >> 16
	hash *= 0x85ebca6b
	hash ^= hash >> 13
	hash *= 0xc2b2ae35
	hash ^= hash >> 16

	return hash, nil
}

func convertToByteArray(data any) ([]byte, error) {
	var buf bytes.Buffer
	encoder := gob.NewEncoder(&buf)
	err := encoder.Encode(data)
	if err != nil {
		return make([]byte, 0), err
	}
	return buf.Bytes(), nil
}
