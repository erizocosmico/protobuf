package droptypedeclaration

import "testing"

func TestDroppedMessage(t *testing.T) {
	var original, decoded *Dropped
	original = &Dropped{Name: "Sergio", Age: 29}
	decoded = &Dropped{}
	encoded, err := original.Marshal()
	if err != nil {
		t.Errorf("Could not marshal Dropped: %v\n", err)
	}

	err = decoded.Unmarshal(encoded)
	if err != nil {
		t.Errorf("Could not unmarshal the encoded Dropped: %v\n", err)
	}

	if original.Name != decoded.Name {
		t.Errorf("Name in Dropped was not transmitted correctly: %v\n", err)
	}

	if original.Age != decoded.Age {
		t.Errorf("Age in Dropped was not transmitted correctly: %v\n", err)
	}
}

func TestKeptMessage(t *testing.T) {
	var original, decoded *Kept
	original = &Kept{Name: "Sergio", Age: 29}
	decoded = &Kept{}
	encoded, err := original.Marshal()
	if err != nil {
		t.Errorf("Could not marshal Kept: %v\n", err)
	}

	err = decoded.Unmarshal(encoded)
	if err != nil {
		t.Errorf("Could not unmarshal the encoded Kept: %v\n", err)
	}

	if original.Name != decoded.Name {
		t.Errorf("Name in Kept was not transmitted correctly: %v\n", err)
	}

	if original.Age != decoded.Age {
		t.Errorf("Age in Kept was not transmitted correctly: %v\n", err)
	}
}
