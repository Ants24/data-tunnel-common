package datatunnelcommon

type DBType string

const (
	OracleType       DBType = oracleString
	PostgresType     DBType = postgresString
	MysqlType        DBType = mySQLString
	PieCloudDBType   DBType = pieCloudDBString
	PieCloudDBTPType DBType = pieCloudDBTPString
	PieCloudDBXPType DBType = pieCloudDBXPString
	DamengType       DBType = damengString
)

type OracleConnectType string

const (
	SID          OracleConnectType = sid
	SERVICE_NAME OracleConnectType = service_name
)

type DBConfig struct {
	DBType            DBType            `json:"dbType"`
	OracleConnectType OracleConnectType `json:"oracleConnectType"`
	Version           string            `json:"version"`
	Host              string            `json:"host"`
	Port              int               `json:"port"`
	Username          string            `json:"username"`
	Password          string            `json:"password"`
	Database          string            `json:"database"`
	ConnectionParams  string            `json:"connectionParams"`
}
