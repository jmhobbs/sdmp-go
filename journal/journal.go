package journal

import (
	"bufio"
	"container/list"
	"os"
	"strings"

	"github.com/jmhobbs/sdmp-go/util"
)

type Journal struct {
	node_fingerprint string
	journal          list.List
}

func New(node_fingerprint string) *Journal {
	return &Journal{node_fingerprint: node_fingerprint}
}

func NewFromFile(f *os.File) *Journal {
	journal := &Journal{}

	scanner := bufio.NewScanner(f)

	scanner.Scan()
	journal.node_fingerprint = scanner.Text()

	for scanner.Scan() {
		line := scanner.Text()
		splits := strings.SplitN(line, "@", 2)
		resource := strings.SplitN(splits[1], "/", 2)
		journal.journal.PushBack(JournalEntry{identifier: splits[0], publisher_fingerprint: resource[0], resource_identifier: resource[1]})
	}

	return journal
}

func (j *Journal) Validate() bool {
	last_line_identifier := util.SHA512(j.node_fingerprint)
	for e := j.journal.Front(); e != nil; e = e.Next() {
		je := e.Value.(JournalEntry)
		if 0 != strings.Compare(last_line_identifier, je.identifier) {
			return false
		}
		last_line_identifier = util.SHA512(je.String())
	}
	return true
}

func (j *Journal) Serialize(w *bufio.Writer) {
	w.WriteString(j.node_fingerprint)
	for e := j.journal.Front(); e != nil; e = e.Next() {
		w.Write([]byte("\n"))
		w.WriteString(e.Value.(JournalEntry).String())
	}
	w.Flush()
}

func (j *Journal) Append(je JournalEntry) {
	e := j.journal.Back()
	if e == nil {
		je.identifier = util.SHA512(j.node_fingerprint)
	} else {
		je.identifier = util.SHA512(e.Value.(JournalEntry).String())
	}
	j.journal.PushBack(je)
}
