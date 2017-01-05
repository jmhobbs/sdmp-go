package util

import (
	"strings"
	"testing"
)

func TestSHA512(t *testing.T) {
	if 0 != strings.Compare(SHA512("hello world"), "MJ7MSJwS1utMxA9QyQLytNDtd-5RGnx6m808qG1M2G-YndNbxf9JlnDaNCVbRbDP2DDoH2Bdz33FVC6TrpzXbw") {
		t.Error("SHA512 Does Not Match")
	}
}
