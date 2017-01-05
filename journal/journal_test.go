package journal

import (
	"bufio"
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/kdar/stringio"
)

func TestConstructor(t *testing.T) {
	journal := New("abcdef")
	if journal == nil || journal.node_fingerprint != "abcdef" {
		t.Error("Journal did not initialize")
	}
}

func TestAppend(t *testing.T) {
	journal := New("abcdef")
	journal.Append(JournalEntry{publisher_fingerprint: "hijklmn", resource_identifier: "opqrstuv"})
	e := journal.journal.Back()

	if e == nil {
		t.Error("Did not insert")
	}

	if e != nil && e.Value.(JournalEntry).publisher_fingerprint != "hijklmn" {
		t.Error("Back doesn't match.")
	}

	journal.Append(JournalEntry{publisher_fingerprint: "omgwtfbbq", resource_identifier: "opqrstuv"})
	e = journal.journal.Back()

	if e != nil && e.Value.(JournalEntry).publisher_fingerprint != "omgwtfbbq" {
		t.Error("Back doesn't match.")
	}
}

func TestFromExample(t *testing.T) {
	journal := New("GlvAreTo0lCSyum7Wzh8pzhxYOOu-gMIgO2N95AAwAGP6-nR8xCvWvIW0t9rF_ZZfpCY_fDV38JDFKaOU91A8Q")
	journal.Append(JournalEntry{publisher_fingerprint: "To0lCSGlvAreyum7WzYOOu-gh8pzhxMIgO2N95AAwAGP6-nR8xCvWvIW0t9rF_ZZfpCY_fDV38JDFKaOU91A8Q", resource_identifier: "lvAlCSreTo0Gyzh8pureTo0Gm7WzhxYu-O2N9AGP65AAw-nR8vIW0xCvWt9rFpCY__ZZffDVJOU91DFKaA8Q38"})
	journal.Append(JournalEntry{publisher_fingerprint: "GlvAreTo0lCSyum7Wzh8pzhxYOOu-gMIgO2N95AAwAGP6-nR8xCvWvIW0t9rF_ZZfpCY_fDV38JDFKaOU91A8Q", resource_identifier: "reTo0GlvAlCSyzh8pum7WzhxYgMIgOOu-O2N9AGP65AAw-nR8vIW0xCvWt9rFpCY__ZZffDVJOU91DFKaA8Q38"})
	journal.Append(JournalEntry{publisher_fingerprint: "To0lCSGlvAreyum7WzYOOu-gh8pzhxMIgO2N95AAwAGP6-nR8xCvWvIW0t9rF_ZZfpCY_fDV38JDFKaOU91A8Q", resource_identifier: "0GlvreToAlCSpum7yzh8WzhxYzh8gOOu-O2N9AGP65AAw-nR8vIW0xCvWt9rFpCY__ZZffDVJOU91DFKaA8Q38"})
	journal.Append(JournalEntry{publisher_fingerprint: "To0lCSGlvAreyum7WzYOOu-gh8pzhxMIgO2N95AAwAGP6-nR8xCvWvIW0t9rF_ZZfpCY_fDV38JDFKaOU91A8Q", resource_identifier: "7WzhxYgMIgOreTo0GlvAlCSyzh-O2N9AGP65AAw-nR8vIW0xCvWt9rFpCY__ZZffD0xCVJOU91DFKad3cA8Q38"})

	sio := stringio.New()

	journal.Serialize(bufio.NewWriter(sio))
	sio.Seek(0, 0)
	result := sio.GetValueString()

	expected := `GlvAreTo0lCSyum7Wzh8pzhxYOOu-gMIgO2N95AAwAGP6-nR8xCvWvIW0t9rF_ZZfpCY_fDV38JDFKaOU91A8Q
dvgOj3boYBzvEtLGz2DOaGW6SKKTmu-jtgi38nn7t40TBW4ObYTQSmUIJk4xMRJaH-ePzvptJgyaR8J0feJ5jw@To0lCSGlvAreyum7WzYOOu-gh8pzhxMIgO2N95AAwAGP6-nR8xCvWvIW0t9rF_ZZfpCY_fDV38JDFKaOU91A8Q/lvAlCSreTo0Gyzh8pureTo0Gm7WzhxYu-O2N9AGP65AAw-nR8vIW0xCvWt9rFpCY__ZZffDVJOU91DFKaA8Q38
d0vXOkAbZf3SnefTJJfWtx7E_dH9gJt2LVcX4wduKFn0gvCuCcDxOSzW8Qt-0v1PZVoCRYLkTRe1ZEEHBRJVMA@GlvAreTo0lCSyum7Wzh8pzhxYOOu-gMIgO2N95AAwAGP6-nR8xCvWvIW0t9rF_ZZfpCY_fDV38JDFKaOU91A8Q/reTo0GlvAlCSyzh8pum7WzhxYgMIgOOu-O2N9AGP65AAw-nR8vIW0xCvWt9rFpCY__ZZffDVJOU91DFKaA8Q38
jIHdWukvQ-O__5gkPGeJ0DxZ1Ae1Ri9mIVXieXebXGKo2xCTnE5DwXFm97_q5FjQDyDPIQyJcqvATLmE3zP6MQ@To0lCSGlvAreyum7WzYOOu-gh8pzhxMIgO2N95AAwAGP6-nR8xCvWvIW0t9rF_ZZfpCY_fDV38JDFKaOU91A8Q/0GlvreToAlCSpum7yzh8WzhxYzh8gOOu-O2N9AGP65AAw-nR8vIW0xCvWt9rFpCY__ZZffDVJOU91DFKaA8Q38
afPovW-Z-Rb3gzbESOrKRukhLjdcoEfphbhq6DLW4YWdGL_ht-x9BbioGWlJnnQQ-I-KwygscarFHydSxoK6Pw@To0lCSGlvAreyum7WzYOOu-gh8pzhxMIgO2N95AAwAGP6-nR8xCvWvIW0t9rF_ZZfpCY_fDV38JDFKaOU91A8Q/7WzhxYgMIgOreTo0GlvAlCSyzh-O2N9AGP65AAw-nR8vIW0xCvWt9rFpCY__ZZffD0xCVJOU91DFKad3cA8Q38`

	if 0 != strings.Compare(result, expected) {
		t.Errorf("Journal does not match expected.\n\nResult:\n\n%s\n\nExpected:\n\n%s", result, expected)
	}
}

func TestSerialize(t *testing.T) {
	journal := New("abcdef")

	sio := stringio.New()

	journal.Serialize(bufio.NewWriter(sio))
	sio.Seek(0, 0)
	out := sio.GetValueString()

	if 0 != strings.Compare(out, "abcdef") {
		t.Error("Journal does not match expected.")
	}
}

func TestNewFromFile(t *testing.T) {
	file, err := os.Open("./testdata/valid.txt")
	if err != nil {
		t.Error("Couldn't open test data.")
	}

	defer file.Close()

	journal := NewFromFile(file)
	sio := stringio.New()

	journal.Serialize(bufio.NewWriter(sio))
	sio.Seek(0, 0)
	result := sio.GetValueString()

	expected, err := ioutil.ReadFile("./testdata/valid.txt")
	if err != nil {
		t.Error("Couldn't read test data.")
	}

	if 0 != strings.Compare(result, string(expected)) {
		t.Errorf("Journal does not match expected.\nGot:\n%s\nEND\n\nExpected:\n%s\nEND", result, string(expected))
	}
}

func TestValidateValid(t *testing.T) {
	file, err := os.Open("./testdata/valid.txt")
	if err != nil {
		t.Error("Couldn't open test data.")
	}

	defer file.Close()

	journal := NewFromFile(file)

	if !journal.Validate() {
		t.Error("Journal does not validate.")
	}
}

func TestValidateInvalid(t *testing.T) {
	file, err := os.Open("./testdata/invalid.txt")
	if err != nil {
		t.Error("Couldn't open test data.")
	}

	defer file.Close()

	journal := NewFromFile(file)

	if journal.Validate() {
		t.Error("Journal should not validate.")
	}
}
