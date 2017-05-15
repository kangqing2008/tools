package cfg

import (
	"github.com/gpmgo/gopm/modules/goconfig"
	"kangqing2008/tools/errors"
)

const(
	DEFAULT_CFG_FILE	= "app_config.ini"
	DEFAULT_DB_SECTION	= "db.default"

	DB_URL				= "url"
	DB_USER				= "user"
	DB_PASSWORD			= "password"
)

type DBConfig struct {
	url			string
	user		string
	password	string
}

func GetDBConfig(filename,section string) DBConfig{
	var dbconfig DBConfig = DBConfig{}
	cfg,err := goconfig.LoadConfigFile(filename)
	errors.PanicIfError(err)
	var err1,err2,err3 error
	dbconfig.url,err1 = cfg.GetValue(section,DB_URL)
	dbconfig.user , err2 = cfg.GetValue(section,DB_USER)
	dbconfig.password,err3 = cfg.GetValue(section,DB_PASSWORD)
	errors.PanicIfErrors(err1,err2,err3)
	return dbconfig
}


func GetDefaultDBConfig()DBConfig{
	return GetDBConfig(DEFAULT_CFG_FILE,DEFAULT_DB_SECTION)
}