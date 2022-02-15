package commoncfg

// DBConfig contains parameters to connect to an SQL DB
type SqlDBConfig struct {
	Host           string `json:"host,omitempty"`
	Port           int    `json:"port,omitempty"`
	User           string `json:"user,omitempty"`
	Password       string `json:"password,omitempty"`
	Database       string `json:"database,omitempty"`
	MaxConnections int    `json:"max_connections,omitempty"`
}
