package config

var (
	host     = "localhost"
	port     = "5432"
	dbuser   = "root"
	password = "root"
	dbname   = "postgres"
)

func Postgres() map[string]string {
	return map[string]string{
		"host":     host,
		"port":     port,
		"dbuser":   dbuser,
		"password": password,
		"dbname":   dbname,
	}
}
