package logBook

import (
	"github.com/vjeantet/jodaTime"
	"github.com/qamarian-dtp/err"
	"io"
	"time"
)

func New (logBook io.Writer) (LogBook) {
	return LogBook {logBook}
}

type LogBook struct {
	logBook io.Writer
}

// Method Record () records the log on a new line, in the log book.
//
// Example:
//
// 	Nov 27, 2019 [UTC+0100 21:52:53]: hello world! An example.
//
func (b LogBook) Record (log []byte) (error) {
	logTime := jodaTime.Format ("\nMMM dd, yyyy ['UTC'Z HH:mm:ss]: ", time.Now ())

	_, errX := b.logBook.Write ([]byte (logTime))
	if errX != nil {
		return err.New ("Log not recorded or log recorded partially.", nil, nil,
			errX)
	}

	_, errY := b.logBook.Write (log)
	if errY != nil {
		return err.New ("Log recorded partially.", nil, nil, errY)
	}

	return nil
}
