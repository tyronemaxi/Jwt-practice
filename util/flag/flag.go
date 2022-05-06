package flag

import "flag"

const (
	MysqlDriver = "mysql"
)

var (
	// server flags
	HealthPort int
	ListenPort int
	RunMode    string
	LocalDebug bool

	// db flags
	DBHost                    string
	DBPort                    int
	DBUser                    string
	DBPassword                string
	DBDatabase                string
	MysqlDriverParams         string
	DBDriver                  string
	DBConnectTimeoutInSeconds int
	PGIdleTransactionTimeout  int
	PGLockTimeout             int
	DBConnectionMaxLifetime   int
	DBMaxOpenConn             int
	DBMaxIdleConn             int
	DBDebugMode               bool
)

func init() {
	// db flags
	flag.StringVar(&DBHost, "db-host", "", "db host")
	flag.IntVar(&DBPort, "db-port", 3306, "db port")
	flag.StringVar(&DBUser, "db-user", "", "db user")
	flag.StringVar(&DBPassword, "db-password", "", "db password")
	flag.StringVar(&DBDatabase, "db-database", "", "db database")
	flag.StringVar(&DBDriver, "db-driver", MysqlDriver, "db driver")
	flag.IntVar(&DBConnectTimeoutInSeconds, "db-conn-timeout", 30, "db connect time out in seconds")
	flag.IntVar(&PGIdleTransactionTimeout, "pg-idle-in-transaction-session-timeout", 30000, "pg idle transaction session time out in millisecond")
	flag.IntVar(&PGLockTimeout, "pg-lock-timeout", 30000, "pg lock time out in millisecond")
	flag.StringVar(&MysqlDriverParams, "mysql-driver-params", "charset=utf8&parseTime=true&loc=Local",
		"mysql connection parameters for go-sql-driver")
	flag.IntVar(&DBConnectionMaxLifetime, "db-connection-max-lifetime-in-seconds", 60,
		"the maximum amount of second a connection may be reused, "+
			"this flag is used to avoid dbproxy disconnect unexpectedly")
	flag.IntVar(&DBMaxOpenConn, "db-max-open-conn", 20,
		"db max open connection")
	flag.IntVar(&DBMaxIdleConn, "db-max-idle-conn", 5,
		"db max idle connection")
	flag.BoolVar(&DBDebugMode, "gorm-debug-mode", false,
		"debug sql switch")
	flag.BoolVar(&LocalDebug, "local-debug-mode", false,
		"local debug switch")

}