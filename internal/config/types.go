package config

type VMS struct {
	Enabled  bool   `json:"enabled"`
	Timeout  int    `json:"timeout"`
	ApiToken string `json:"api_token"`
}

type API struct {
	Host          string `json:"host"`
	Authorization string `json:"authorization"`
	Ssl           bool   `json:"ssl"`
	Timeout       uint   `json:"timeout"`
}

type Scanner struct {
	MinWait int `json:"min_wait"`
	MaxWait int `json:"max_wait"`
}

type Config struct {
	Verbose int `json:"verbose"`

	Api     API     `json:"api"`
	Vms     VMS     `json:"vms"`
	Scanner Scanner `json:"scanner"`
}
