package config

// Shard describes the shard which holds the unique keys
type Shard struct {
	Name    string
	Idx     int
	Address string
}

// Config describes the sharding config
type Config struct {
	Shards []Shard
}
