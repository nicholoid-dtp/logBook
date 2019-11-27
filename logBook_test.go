package logBook

import (
	"gopkg.in/qamarian-lib/str.v3"
	"os"
	"testing"
)

func TestLogBook (t *testing.T) {
	str.PrintEtr ("LogBook testing has started...\n", "std", "TestLogBook ()")
	b := New (os.Stdout)
	for i := 1; i <= 10; i ++ {
		b.Record ([]byte ("hello world!"))
		b.Record ([]byte {'=', '/', '='})
	}
}
