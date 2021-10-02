package vars

import "github.com/schollz/progressbar/v3"


var (
	Schemes = []string{"http", "https"}
	Hosts = []string{}
	Ips = []string{}
)

var ProcessBar *progressbar.ProgressBar

const (
	VersionInfo = "0.0.1"
)