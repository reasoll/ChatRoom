package config

import _ "gopkg.in/yaml.v2"

type Mysql struct {
	Addr         string `yaml:"Addr"`
	Port         string `yaml:"Port"`
	UserNmae     string `yaml:"UserName"`
	Password     string `yaml:"Password"`
	DataBaseName string `yaml:"DataBaseName"`
}

func (mysql *Mysql) DefaultMySqlConfig() {
	//mysql.Addr = "23.234.252.205"
	mysql.Addr = "172.16.1.227"
	mysql.Port = "3306"
	mysql.DataBaseName = "chatroom"
	//mysql.DataBaseName = "ChatRoom"
	mysql.UserNmae = "root"
	mysql.Password = "123456"
	//mysql.Password = "!QAZxsw2"
}

func (mysql *Mysql) InitMysqlConfig() {
	mysql.DefaultMySqlConfig()
}
