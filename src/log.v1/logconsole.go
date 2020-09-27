package log

import (
	"fmt"
	"os"
)

type Consolelog struct {
	format      string
	loglevel    string
	logdatachan chan *string
}

func Newconsolelog() *Consolelog {
	return &Consolelog{logdatachan: make(chan *string, 50000)}
}

func (c *Consolelog) Debug(args ...interface{}) error {
	c.loglevel = "Debug"
	c.format = Logformat(c.loglevel)
	Writefilelog(c.logdatachan, args...)
	for data := range c.logdatachan {
		_, err := fmt.Fprintf(os.Stdout, c.format, *data)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Consolelog) Info(args ...interface{}) error {
	c.loglevel = "Info"
	c.format = Logformat(c.loglevel)
	Writefilelog(c.logdatachan, args...)
	for data := range c.logdatachan {
		_, err := fmt.Fprintf(os.Stdout, c.format, *data)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Consolelog) Warn(args ...interface{}) error {
	c.loglevel = "Warn"
	c.format = Logformat(c.loglevel)
	Writefilelog(c.logdatachan, args...)
	for data := range c.logdatachan {
		_, err := fmt.Fprintf(os.Stdout, c.format, *data)
		if err != nil {
			return err
		}

	}
	return nil
}

func (c *Consolelog) Error(args ...interface{}) error {
	c.loglevel = "Error"
	c.format = Logformat(c.loglevel)
	Writefilelog(c.logdatachan, args...)
	for data := range c.logdatachan {
		_, err := fmt.Fprintf(os.Stdout, c.format, *data)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Consolelog) Fatal(args ...interface{}) error {
	c.loglevel = "Fatal"
	c.format = Logformat(c.loglevel)
	Writefilelog(c.logdatachan, args...)
	for data := range c.logdatachan {
		_, err := fmt.Fprintf(os.Stdout, c.format, *data)
		if err != nil {
			return err
		}
	}
	return nil
}
