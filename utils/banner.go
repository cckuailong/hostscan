package utils

import "github.com/fatih/color"

func Banner() {
	banner := `  
/ )( \ /  \ / ___)(_  _)/ ___) / __) / _\ (  ( \
) __ ((  O )\___ \  )(  \___ \( (__ /    \/    /
\_)(_/ \__/ (____/ (__) (____/ \___)\_/\_/\_)__)	

`
	color.HiBlue(banner)
}