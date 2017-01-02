package enumdroptypedeclaration

import "testing"

func TestEnumWorks(t *testing.T) {
	var original, decoded *Message
	original = &Message{EnumeratedField: B}
	decoded = &Message{}
	encoded, err := original.Marshal()
	if err != nil {
		t.Errorf("Could not marshal Message: %v\n", err)
	}

	err = decoded.Unmarshal(encoded)
	if err != nil {
		t.Errorf("Could not unmarshal the encoded Message: %v\n", err)
	}

	if decoded.EnumeratedField != B {
		t.Errorf("Decoded enumerated field is not the original value")
	}
}
