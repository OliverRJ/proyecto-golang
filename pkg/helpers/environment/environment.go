package environment

import "os"

type Environment string

const (
	LOCAL Environment = "local"
	DEV   Environment = "dev"
	TEST  Environment = "test"
	PROD  Environment = "prod"
)

var environments = map[string]Environment{
	"local": LOCAL,
	"dev":   DEV,
	"test":  TEST,
	"prod":  PROD,
}

func of(str string) Environment {
	return environments[str]
}

func GetEnv() Environment {
	return environments[os.Getenv("ENV")]
}
