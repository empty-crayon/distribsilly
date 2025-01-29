package config

import (
	"fmt"
	"hash/fnv"

	"github.com/BurntSushi/toml"
)

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

// Shards represents an easier to use reprensetation of sharding config
type Shards struct {
	Count  int
	CurIdx int
	Addrs  map[int]string
}

// Parse File parses the config and returns it 
func ParseFile (filename string) (Config, error) {
	var c Config
	if _, err := toml.DecodeFile(filename, &c); err != nil {
		return Config{}, err
	}

	return c, nil 
}

// Parseshards converts and verifies list of shards specified in the config into a form that can be used for routing
func ParseShards(shards []Shard, curShardName string) (*Shards, error) {
	shardCount := len(shards)
	shardIdx := -1
	addrs := make(map[int]string)

	for _, s := range shards {
		// checking
		if _, ok := addrs[s.Idx]; ok {
			return nil, fmt.Errorf("duplicate shard index: %d", s.Idx)
		}

		addrs[s.Idx] = s.Address

		if s.Name == curShardName {
			shardIdx = s.Idx
		}
	}

	for i := 0; i < shardCount; i++ {
		if _, ok := addrs[i]; !ok {
			return nil, fmt.Errorf("shard %d not found", i)
		}
	}

	if shardIdx < 0 {
		return nil, fmt.Errorf("shard %q was not found", curShardName)
	}

	return &Shards{
		Count:  shardCount,
		CurIdx: shardIdx,
		Addrs:  addrs,
	}, nil
}

// index returns the shard number for the corresponding key
func (s *Shards) GetShard(key string) int {
	hash := fnv.New64()
	hash.Write([]byte(key))
	shardId := int(hash.Sum64() % uint64(s.Count))
	return shardId
}
