package journal

import "testing"

func TestJournaleEntryToString(t *testing.T) {
	je := &JournalEntry{identifier: "abcdefg", publisher_fingerprint: "hijklmn", resource_identifier: "opqrstuv"}

	if je.String() != "abcdefg@hijklmn/opqrstuv" {
		t.Errorf("%s != 'abcdefg@hijklmn/opqrstuv'", je.String())
	}
}
