package log

type Consolelog struct {
	format   string
	loglevel string
}

func Newconsolelog() *Consolelog {
	return &Consolelog{}
}

func (c *Consolelog) Debug(args ...interface{}) error {
	c.loglevel = "Debug"
	c.format = Logformat(c.loglevel)
	//err := Writefilelog(os.Stdout, c.format, args...)
	//if err != nil {
	//	return err
	//}
	return nil
}

func (c *Consolelog) Info(args ...interface{}) error {
	c.loglevel = "Info"
	c.format = Logformat(c.loglevel)
	//err := Writefilelog(os.Stdout, c.format, args...)
	//if err != nil {
	//	return err
	//}
	return nil
}

func (c *Consolelog) Warn(args ...interface{}) error {
	c.loglevel = "Warn"
	c.format = Logformat(c.loglevel)
	//err := Writefilelog(os.Stdout, c.format, args...)
	//if err != nil {
	//	return err
	//}
	return nil
}

func (c *Consolelog) Error(args ...interface{}) error {
	c.loglevel = "Error"
	c.format = Logformat(c.loglevel)
	//err := Writefilelog(os.Stdout, c.format, args...)
	//if err != nil {
	//	return err
	//}
	return nil
}

func (c *Consolelog) Fatal(args ...interface{}) error {
	c.loglevel = "Fatal"
	c.format = Logformat(c.loglevel)
	//err := Writefilelog(os.Stdout, c.format, args...)
	//if err != nil {
	//	return err
	//}
	return nil
}
