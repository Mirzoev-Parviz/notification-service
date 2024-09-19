package models

type Settings struct {
	AppParams  AppParams
	DBSettings DBSettings
}

type AppParams struct {
	PortRun string
}

type DBSettings struct {
	User     string
	Password string
	Host     string
	Port     string
	Protocol string
	Database string
	TLS      string
}
