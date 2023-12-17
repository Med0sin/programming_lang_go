package mascot_test

import "testing"

func TestMascot(t *testing.T) {
	if masoct.BestMascot() != "go Gopher" {
		t.Fatal("wrong Mascot :( )")
	}
}
