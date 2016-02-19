package lib

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
)

type CustomFormatter struct {
}

func (f *CustomFormatter) Format(entry *log.Entry) (result []byte, err error) {
	data := map[string]interface{}{
		"time":  entry.Time.Format(time.RFC3339),
		"msg":   entry.Message,
		"level": entry.Level.String(),
	}
	for k, v := range entry.Data {
		switch v := v.(type) {
		case error:
			data[k] = v.Error()
		default:
			data[k] = v
		}
	}

	buffer := new(bytes.Buffer)
	if err = gob.NewEncoder(buffer).Encode(data); err != nil {
		log.Fatal(err)
	}
	result = buffer.Bytes()

	return result, nil
}

type Logger struct {
	*log.Logger
	*Config
}

func NewLogger(config *Config) (l *Logger) {
	l = &Logger{
		Logger: log.New(),
		Config: config,
	}
	log.SetFormatter(&log.TextFormatter{})
	l.setOutput()
	return
}

func (self *Logger) setOutput() {
	output := fmt.Sprintf("log/%v", self.Default("log.output"))
	file, err := os.OpenFile(output, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fmt.Errorf("| Error | Opening file: %v \n", err)
		return
	}
	log.SetOutput(file)
}

func (self *Logger) Printf(format string, args ...interface{}) {
	fields := log.Fields{"Data": args}
	log.WithFields(fields).Info(format)
}
