package param

type WebAppParam struct {
	Flag1       bool
	Version     string
	ServiceName string

	SqliteDsn string
	MysqlDsn  string
	PgDsn     string
}
