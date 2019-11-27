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

func (b LogBook) Record (log []byte) (error) {
	logTime := jodaTime.Format ("MMM dd, yyyy [hh:mm:ss a Z UTC]: ", time.Now ())

	_, errX := b.logBook.Write ([]byte (logTime))
	if errX != nil {
		return err.New ("Log not recorded or log recorded partially.", nil, nil, errX)
	}

	_, errY := b.logBook.Write (log)
	if errY != nil {
		return err.New ("Log recorded partially.", nil, nil, errY)
	}

	_, errZ := b.logBook.Write ([]byte ("\n"))
	if errZ != nil {
		return err.New ("Log recorded partially.", nil, nil, errZ)
	}

	return nil
}
