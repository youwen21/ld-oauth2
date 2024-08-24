package flag_vars

import (
	"flag"
	"net"
)

var (
	host = flag.String("host", "0.0.0.0", "listen host.")
	port = flag.String("port", "8181", "Port to run the server.")

	root = flag.String("root", "", "specify the root.")

	runmode = flag.String("runmode", "", "runmode default is empty")

	logfile  = flag.String("logfile", "", "logfile default os stdout")
	loglevel = flag.String("loglevel", "error", "log level")

	ldClientId     = flag.String("ld_client_id", "", "")
	ldClientSecret = flag.String("ld_client_secret", "", "")
)

func Address() string {
	return net.JoinHostPort(Host(), Port())
}

func Host() string {
	return *host
}

func Port() string {
	return *port
}

func GetRoot() string {
	return *root
}

func RunMode() string {
	return *runmode
}

func Logfile() string {
	return *logfile
}

func Loglevel() string {
	return *loglevel
}

func LDClientId() string {
	return *ldClientId
}

func LDClientSecret() string {
	return *ldClientSecret
}

func init() {
	flag.Parse()
}
