package vars

import "flag"

var (
	Version          = flag.Bool("v", false, "Show hostscan version")
	Host             = flag.String("d", "", "Host to test")
	Ip               = flag.String("i", "", "Nginx IP")
	Timeout          = flag.Int("t", 5, "Timeout for Http connection.")
	Thread			 = flag.Int("T", 3, "Thread for Http connection.")
	HostFile         = flag.String("D", "", "Hosts in file to test")
	IpFile     	     = flag.String("I", "", "Nginx Ip in file to test")
	OutFile			 = flag.String("O", "result.txt", "Output File")
	IsRandUA	     = flag.Bool("U", false, "Open to send random UserAgent to avoid bot detection.")
)