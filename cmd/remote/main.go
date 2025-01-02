package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	"net/http"
	_ "net/http/pprof"

	"github.com/DeBankDeFi/nodex/pkg/db"
	"github.com/DeBankDeFi/nodex/pkg/reader"
	"github.com/DeBankDeFi/nodex/pkg/utils"
)

func main() {
	config := &utils.Config{}
	flag.StringVar(&config.RemoteListenAddr, "listen_addr", "0.0.0.0:7654", "remote server listen address")
	flag.StringVar(&config.S3ProxyAddr, "s3proxy_addr", "127.0.0.1:8765", "s3 address")
	flag.StringVar(&config.KafkaAddr, "kafka_addr", "172.31.25.142:9092", "kafka address")
	flag.StringVar(&config.MetricEndpoint, "metric_address", ":10086", "metric address")
	flag.StringVar(&config.Env, "env", "test", "env")
	flag.StringVar(&config.ChainId, "chain_id", "eth", "chain id")
	flag.StringVar(&config.Role, "role", "master", "master or backup")
	flag.StringVar(&config.DBInfoPath, "db_info_path", "dbinfo.json", "db info path")
	flag.IntVar(&config.ReorgDeep, "reorg_deep", 128, "chain reorg deep")
	flag.IntVar(&config.DBCacheSize, "db_cache_size", 2048, "db cache size in MB")
	flag.StringVar(&config.NdrcAddr, "ndrc_addrs", "127.0.0.1:8089", "ndrc addrs")
	flag.Parse()
	stopChan := make(chan os.Signal, 1)

	signal.Notify(stopChan, syscall.SIGTERM, syscall.SIGINT)

	pool := db.NewDBPool()

	reader, err := reader.NewReader(config, pool)
	if err != nil {
		panic(err)
	}

        err = reader.Start()
        if err != nil {
                panic(err)
        }

        defer func() {
                reader.Stop()
        }()

	go func() {
		http.ListenAndServe("localhost:6060", nil)
	}()

	go func() {
		http.Handle("/metrics", promhttp.Handler())
		http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("pong"))
		})
		http.ListenAndServe(config.MetricEndpoint, nil)
	}()

	err = reader.Start()
	if err != nil {
		panic(err)
	}

	defer func() {
		reader.Stop()
	}()
	<-stopChan
}
