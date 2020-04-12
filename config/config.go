package config

import (
	"errors"
	"fmt"
	"github.com/lw000/gocommon/db/mysql"
	"github.com/lw000/gocommon/db/rdsex"
	"strconv"

	"github.com/Unknwon/goconfig"
)

// IniConfig ini配置
type IniConfig struct {
	RdsCfg   *tyrdsex.JsonConfig
	MysqlCfg *tymysql.JsonConfig
	TLS      struct {
		Enable   bool
		CertFile string
		KeyFile  string
	}
	Port     int64
	Debug    int64
	SplitLog int
}

// NewIniConfig ...
func NewIniConfig() *IniConfig {
	return &IniConfig{
		RdsCfg:   &tyrdsex.JsonConfig{},
		MysqlCfg: &tymysql.JsonConfig{},
	}
}

// LoadIniConfig ...
func LoadIniConfig(file string) (*IniConfig, error) {
	cfg := NewIniConfig()
	er := cfg.Load(file)
	return cfg, er
}

// Load ...
func (c *IniConfig) Load(file string) error {
	var (
		er error
		f  *goconfig.ConfigFile
	)

	f, er = goconfig.LoadConfigFile(file)
	if er != nil {
		return fmt.Errorf("读取配置文件失败[%s]", file)
	}

	er = c.readCfg(f)
	if er != nil {
		return er
	}

	er = c.readTlsCfg(f)
	if er != nil {
		return er
	}

	// er = c.readMysqlCfg(f)
	// if er != nil {
	// 	return er
	// }

	// er = c.readRdsCfg(f)
	// if er != nil {
	// 	return er
	// }

	return nil
}

func (c *IniConfig) readCfg(f *goconfig.ConfigFile) error {
	var (
		er       error
		port     string
		debug    string
		splitlog string
	)

	section := "main"

	port, er = f.GetValue(section, "port")
	if er != nil {
		return fmt.Errorf("无法获取键值(%s):%s", "port", er.Error())
	}

	c.Port, er = strconv.ParseInt(port, 10, 64)
	if er != nil {
		return errors.New(er.Error())
	}

	debug, er = f.GetValue(section, "debug")
	if er != nil {
		return fmt.Errorf("无法获取键值(%s):%s", "debug", er.Error())
	}

	c.Debug, er = strconv.ParseInt(debug, 10, 64)
	if er != nil {
		return errors.New(er.Error())
	}

	splitlog, er = f.GetValue(section, "splitlog")
	if er != nil {
		return fmt.Errorf("无法获取键值(%s):%s", "splitlog", er.Error())
	}

	c.SplitLog, er = strconv.Atoi(splitlog)
	if er != nil {
		return errors.New(er.Error())
	}

	return nil
}

func (c *IniConfig) readMysqlCfg(f *goconfig.ConfigFile) error {
	var (
		er           error
		maxOdleConns string
		maxOpenConns string
	)

	section := "mysql"
	c.MysqlCfg.Username, er = f.GetValue(section, "username")
	if er != nil {
		return fmt.Errorf("无法获取键值(%s):%s", "username", er.Error())
	}

	c.MysqlCfg.Password, er = f.GetValue(section, "password")
	if er != nil {
		return fmt.Errorf("无法获取键值(%s):%s", "password", er.Error())
	}

	c.MysqlCfg.Host, er = f.GetValue(section, "host")
	if er != nil {
		return fmt.Errorf("无法获取键值(%s):%s", "host", er.Error())
	}

	c.MysqlCfg.Database, er = f.GetValue(section, "database")
	if er != nil {
		return fmt.Errorf("无法获取键值(%s):%s", "database", er.Error())
	}

	maxOdleConns, er = f.GetValue(section, "MaxOdleConns")
	if er != nil {
		return fmt.Errorf("无法获取键值(%s):%s", "MaxOdleConns", er.Error())
	}
	c.MysqlCfg.MaxOdleConns, er = strconv.ParseInt(maxOdleConns, 10, 64)
	if er != nil {
		return fmt.Errorf("无法获取键值(%s):%s", "MaxOdleConns", er.Error())
	}

	maxOpenConns, er = f.GetValue(section, "MaxOpenConns")
	if er != nil {
		return er
	}
	c.MysqlCfg.MaxOpenConns, er = strconv.ParseInt(maxOpenConns, 10, 64)
	if er != nil {
		return er
	}

	return nil
}

func (c *IniConfig) readRdsCfg(f *goconfig.ConfigFile) error {
	var (
		er           error
		Db           string
		PoolSize     string
		MinIdleConns string
	)

	section := "redis"
	c.RdsCfg.Host, er = f.GetValue(section, "host")
	if er != nil {
		return fmt.Errorf("无法获取键值(%s):%s", "host", er.Error())
	}

	c.RdsCfg.Psd, er = f.GetValue(section, "psd")
	if er != nil {
		return fmt.Errorf("无法获取键值(%s):%s", "psd", er.Error())
	}

	Db, er = f.GetValue(section, "db")
	if er != nil {
		return fmt.Errorf("无法获取键值(%s):%s", "db", er.Error())
	}
	c.RdsCfg.Db, er = strconv.ParseInt(Db, 10, 64)
	if er != nil {
		return er
	}

	PoolSize, er = f.GetValue(section, "poolSize")
	if er != nil {
		return fmt.Errorf("无法获取键值(%s):%s", "poolSize", er.Error())
	}
	c.RdsCfg.PoolSize, er = strconv.ParseInt(PoolSize, 10, 64)
	if er != nil {
		return er
	}

	MinIdleConns, er = f.GetValue(section, "minIdleConns")
	if er != nil {
		return fmt.Errorf("无法获取键值(%s):%s", "minIdleConns", er.Error())
	}
	c.RdsCfg.MinIdleConns, er = strconv.ParseInt(MinIdleConns, 10, 64)
	if er != nil {
		return er
	}
	return nil
}

func (c *IniConfig) readTlsCfg(f *goconfig.ConfigFile) error {
	var (
		er     error
		enable string
	)

	section := "tls"

	enable, er = f.GetValue(section, "enable")
	if er != nil {
		return fmt.Errorf("无法获取键值(%s):%s", "port", er.Error())
	}
	c.TLS.Enable, er = strconv.ParseBool(enable)
	if er != nil {
		return errors.New(er.Error())
	}

	if c.TLS.Enable {
		c.TLS.CertFile, er = f.GetValue(section, "certFile")
		if er != nil {
			return fmt.Errorf("无法获取键值(%s):%s", "certFile", er.Error())
		}
		if c.TLS.CertFile == "" {
			return errors.New("cretFile is empty")
		}

		c.TLS.KeyFile, er = f.GetValue(section, "keyFile")
		if er != nil {
			return fmt.Errorf("无法获取键值(%s):%s", "keyFile", er.Error())
		}
		if c.TLS.KeyFile == "" {
			return errors.New("keyFile is empty")
		}
	}

	return nil
}

func (c IniConfig) String() string {
	return fmt.Sprintf("{%v, %v}", c.MysqlCfg, c.RdsCfg)
}
