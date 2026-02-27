package config

type Queues struct {
	DeleteOldEvents      Queue `yaml:"delete_old_events"`
	SendNotifiableEvents Queue `yaml:"send_notifiable_events"`
}

type Queue struct {
	Enabled      bool `yaml:"enabled"`
	WorkersCount int  `yaml:"workers_count"`
}
