package mjson

var emptyArray = []byte("[]")

// ToArray assumes the list of byte slices are raw JSON messages, and joins
// them into an array. Note: if you know the total length of the items
// beforehand, use ToArrayWithLength.
func ToArray(items [][]byte) []byte {
	length := 0
	for _, item := range items {
		length += len(item)
	}

	return ToArrayWithLength(items, length)
}

// ToArrayWithLength joins the items into a JSON array.
func ToArrayWithLength(items [][]byte, totalSize int) []byte {
	if len(items) == 0 {
		return emptyArray
	}

	output := make([]byte, totalSize+len(items)+1)
	output[0] = '['
	n := 1
	for i, item := range items {
		if i > 0 {
			output[n] = ','
			n++
		}

		n += copy(output[n:], item)
	}
	output[n] = ']'

	return output
}
