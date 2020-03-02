package fdfs_client

import cfgWs "state-server/conf"

type config struct {
	trackerAddr []string
	maxConns    int
}

func newConfig() (*config, error) {
	config := &config{}
	config.trackerAddr = append(config.trackerAddr, cfgWs.TrackerServerAddr)
	config.maxConns = cfgWs.MaxConn

	return config, nil
}
