package vars

import "flag"

var (
	Version          = flag.Bool("v", false, "Show hostscan version")
	Host             = flag.String("d", "", "Host to test")
	Ip               = flag.String("i", "", "Nginx IP. \nExample: 1.1.1.1 or 1.2.3.4/24")
	Timeout          = flag.Int("t", 5, "Timeout for Http connection.")
	Thread			 = flag.Int("T", 3, "Thread for Http connection.")
	HostFile         = flag.String("D", "", "Hosts in file to test")
	IpFile     	     = flag.String("I", "", "Nginx Ip in file to test")
	Iports	         = flag.String("p", "", "Port List of Nginx IP. If the flag is set, hostscan will ignore the port in origin IP input. \nExample: 80,8080,8000-8009")
	OutFile			 = flag.String("O", "result.txt", "Output File")
	IsRandUA	     = flag.Bool("U", false, "Open to send random UserAgent to avoid bot detection.")
	Verbose	     	 = flag.Bool("V", false, "Output All scan Info. \nDefault is false, only output the result with title.")
	FilterRespStatusCodes = flag.String("F", "", "Filter result with List of Response Status Code. \nExample: 200,201,302")
)