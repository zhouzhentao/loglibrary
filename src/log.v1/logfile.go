package log

import (
	"fmt"
	"os"
)

type Filelog struct {
	file        *os.File
	filepath    string
	format      string
	loglevel    string
	logdatachan chan *string
}

func Newfilelog(file *os.File, filepath string) *Filelog {
	return &Filelog{file: file, filepath: filepath, logdatachan: make(chan *string, 50000)}
}

func (f *Filelog) Debug(args ...interface{}) error {
	f.loglevel = "Debug"
	f.format = Logformat(f.loglevel)
	Writefilelog(f.logdatachan, args...)
	defer f.file.Close()
	for data := range f.logdatachan {
		_, err := fmt.Fprintf(f.file, f.format, *data)
		if err != nil {
			return err
		}
	}
	return nil
}

func (f *Filelog) Info(args ...interface{}) error {
	f.loglevel = "Info"
	f.format = Logformat(f.loglevel)
	Writefilelog(f.logdatachan, args...)
	defer f.file.Close()
	for data := range f.logdatachan {
		_, err := fmt.Fprintf(f.file, f.format, *data)
		if err != nil {
			return err
		}
	}
	return nil
}

func (f *Filelog) Warn(args ...interface{}) error {
	f.loglevel = "Warn"
	f.format = Logformat(f.loglevel)
	Writefilelog(f.logdatachan, args...)
	defer f.file.Close()
	for data := range f.logdatachan {
		_, err := fmt.Fprintf(f.file, f.format, *data)
		if err != nil {
			return err
		}
	}
	return nil
}

func (f *Filelog) Error(args ...interface{}) error {
	f.loglevel = "Error"
	f.format = Logformat(f.loglevel)
	Writefilelog(f.logdatachan, args...)
	defer f.file.Close()
	for data := range f.logdatachan {
		_, err := fmt.Fprintf(f.file, f.format, *data)
		if err != nil {
			return err
		}
	}
	return nil
}

func (f *Filelog) Fatal(args ...interface{}) error {
	f.loglevel = "Fatal"
	f.format = Logformat(f.loglevel)
	Writefilelog(f.logdatachan, args...)
	defer f.file.Close()
	for data := range f.logdatachan {
		_, err := fmt.Fprintf(f.file, f.format, *data)
		if err != nil {
			return err
		}
	}
	return nil
}
