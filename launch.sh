# !/bin/bash

set -e

trap 'killall distribsilly' SIGINT

cd $(dirname $0)

killall distribsilly || true
sleep 0.1

go build -o distribsilly
./distribsilly -db-location=shard0.db -http-addr=127.0.0.1:8080 -config-file=sharding.toml -shard=shard0 &
./distribsilly -db-location=shard1.db -http-addr=127.0.0.1:8081 -config-file=sharding.toml -shard=shard1 &
./distribsilly -db-location=shard2.db -http-addr=127.0.0.1:8082 -config-file=sharding.toml -shard=shard2 &

wait
