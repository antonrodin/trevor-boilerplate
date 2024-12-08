package config

import "github.com/alexedwards/scs/v2"

type AppConfig struct {
	AppTitle     string
	InProduction bool
	Session      *scs.SessionManager
}
