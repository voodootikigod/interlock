package config

// the extension config has all options for all load balancer extensions
// the extension itself will use whichever options needed
type ExtensionConfig struct {
	Name                   string
	ConfigPath             string
	PidPath                string // haproxy, nginx
	BackendOverrideAddress string // haproxy, nginx
	ConnectTimeout         int    // haproxy
	ServerTimeout          int    // haproxy
	ClientTimeout          int    // haproxy
	MaxConn                int    // haproxy, nginx
	Port                   int    // haproxy, nginx
	SyslogAddr             string // haproxy
	AdminUser              string // haproxy
	AdminPass              string // haproxy
	SSLCertPath            string // haproxy
	SSLCert                string // haproxy
	SSLPort                int    // haproxy, nginx
	SSLOpts                string // haproxy
	User                   string // nginx
	WorkerProcesses        int    // nginx
	RLimitNoFile           int    // nginx
	ProxyConnectTimeout    int    // nginx
	ProxySendTimeout       int    // nginx
	ProxyReadTimeout       int    // nginx
	SendTimeout            int    // nginx
	SSLCiphers             string // nginx
	SSLProtocols           string // nginx
}

type Config struct {
	ListenAddr    string
	DockerURL     string
	TLSCACert     string
	TLSCert       string
	TLSKey        string
	AllowInsecure bool
	Extensions    []ExtensionConfig
}
