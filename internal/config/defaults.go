package config

func (cfg *Config) LoadDefaults() {
	if cfg.LogDirectory == nil {
		cfg.LogDirectory = new(string)

		*cfg.LogDirectory = "./logs"
	}

	cfg.Api.Host = "http://localhost"
	cfg.Api.Authorization = ""
	cfg.Api.Timeout = 5

	cfg.WebApi.Host = "http://localhost"
	cfg.WebApi.Endpoint = "/api/spy/get"
	cfg.WebApi.Timeout = 5
	cfg.WebApi.Interval = 120

	cfg.Vms.Enabled = false
	cfg.Vms.MinWait = 60
	cfg.Vms.MaxWait = 180
	cfg.Vms.Timeout = 5
	cfg.Vms.Limit = 100
	cfg.Vms.ExcludeEmpty = true
	cfg.Vms.SubBots = true
	cfg.Vms.AddOnly = true
	cfg.Vms.SetOffline = true

	cfg.RemoveInactive.Enabled = false
	cfg.RemoveInactive.InactiveTime = 2592000
	cfg.RemoveInactive.Interval = 86400
	cfg.RemoveInactive.Timeout = 5

	cfg.RemoveDups.Enabled = false
	cfg.RemoveDups.Interval = 120
	cfg.RemoveDups.Limit = 100
	cfg.RemoveDups.MaxServers = 100
	cfg.RemoveDups.Timeout = 5
}
