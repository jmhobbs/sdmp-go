package journal

import "fmt"

type JournalEntry struct {
	identifier            string
	publisher_fingerprint string
	resource_identifier   string
}

func (je JournalEntry) String() string {
	return fmt.Sprintf("%s@%s/%s", je.identifier, je.publisher_fingerprint, je.resource_identifier)
}
