package main

import (
	"fornever.org/app"
	"github.com/urfave/cli"
)

var commandEntry = cli.Command{
	Name:   "entry",
	Usage:  "program entry",
	Action: entry,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:   "addr",
			EnvVar: "LISTEN_ADDR",
			Value:  "0.0.0.0:8080",
		},

		cli.StringFlag{
			Name:   "sqlite-dsn",
			EnvVar: "SQLITE_DSN",
		},

		cli.StringFlag{
			Name:   "mysql-dsn",
			EnvVar: "MYSQL_DSN",
		},

		cli.StringFlag{
			Name:   "pg-dsn",
			EnvVar: "PG_DSN",
		},
	},
}

func entry(c *cli.Context) error {

	inst := app.CreateApp(&app.WebAppParam{
		ServiceName: AppName,
		Version:     Version,
		Flag1:       false,
		SqliteDsn:   c.String("sqlite-dsn"),
		MysqlDsn:    c.String("mysql-dsn"),
		PgDsn:       c.String("pg-dsn"),
	})

	return inst.Run(c.String("addr"))

}
