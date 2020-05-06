package models

import (
	"fmt"
	"reportGameErr/global"
	"strconv"

	log "github.com/sirupsen/logrus"
)

type ErrorItem struct {
	Source string `json:"source"`
	Name   string `json:"name"`
	ErrMsg string `json:"err_msg"`
	Line   int    `json:"line"`
	Column int    `json:"column"`
}

func (p ErrorItem) String() string {
	return fmt.Sprintf("Source: %s, Name: %s, Line:%d, Column:%d, ErrMsg:%s", p.Source, p.Name, p.Line, p.Column, p.ErrMsg)
}

type LogRecordConfig struct {
	module string
	line   string
	column string
	info   string
}

func (cfg *LogRecordConfig) SetColumn(column string) {
	cfg.column = column
}

func (cfg *LogRecordConfig) SetInfo(info string) {
	cfg.info = info
}

func (cfg *LogRecordConfig) SetModule(module string) {
	cfg.module = module
}

func (cfg *LogRecordConfig) SetLine(line string) {
	cfg.line = line
}

func NewLogConfig() *LogRecordConfig {
	return &LogRecordConfig{}
}

func (cfg *LogRecordConfig) Save(clientIP string) error {
	var (
		err error
		row int
		col int
	)

	row, err = strconv.Atoi(cfg.line)
	if err != nil {
		log.Error(err)
		return err
	}

	col, err = strconv.Atoi(cfg.column)
	if err != nil {
		log.Error(err)
		return err
	}

	go cfg.writeLog(row, col)

	return nil
}

func (cfg *LogRecordConfig) writeLog(row, col int) {
	defer func() {
		if x := recover(); x != nil {
			log.Error(x)
		}
	}()

	source, name, line, column, ok := global.SourceMapServer.Get(row, col)
	if ok {
		p := &ErrorItem{Source: source, Name: name, Line: line, Column: column, ErrMsg: cfg.info}
		log.Error(p)
	} else {
		log.WithFields(log.Fields{"row": row, "col": col}).Error("unknown error")
	}
}
