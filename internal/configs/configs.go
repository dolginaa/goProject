package configs

import "fmt"

const (
	user     = "schedule_user"
	dbname   = "schedule"
	password = "user"
	sslmode  = "disable"
	host     = "localhost"
	port     = "5432"
)

var ConnString = generateDsn()

func generateDsn() string {
	return fmt.Sprintf("user=%s dbname=%s password=%s sslmode=%s host=%s port=%s", user, dbname, password, sslmode, host, port)
}
