package SourceMap

import (
	"io/ioutil"

	"github.com/go-sourcemap/sourcemap"
)

type SourceMapManager struct {
	ok   bool
	smap *sourcemap.Consumer
}

func NewSourceMapManager() *SourceMapManager {
	return &SourceMapManager{}
}

func (s *SourceMapManager) Parse(filename string) error {
	var (
		er   error
		data []byte
	)

	data, er = ioutil.ReadFile(filename)
	if er != nil {
		return er
	}
	if er != nil {
		s.ok = false
		return er
	}

	s.smap, er = sourcemap.Parse(filename, data)
	if er != nil {
		return er
	}

	//var json = jsoniter.ConfigCompatibleWithStandardLibrary
	//er = json.Unmarshal(data, &s.smap)
	//if er != nil {
	//	return er
	//}

	if er != nil {
		s.ok = false
		return er
	}

	s.ok = true

	return nil
}

func (s *SourceMapManager) Get(r, c int) (source, name string, line, column int, ok bool) {
	return s.smap.Source(r, c)
}
