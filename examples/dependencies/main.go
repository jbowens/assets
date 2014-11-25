package main

//
// Run from assets/examples. i.e.
//
// go run dependencies/main.go
//

import (
	"fmt"

	"github.com/jbowens/assets"
)

var coreJS assets.AssetBundle
var homeJS assets.AssetBundle

func onStartUp() {

	// On start up, compile your assets.
	pipeline := []assets.Filter{
		assets.Concat(),
		assets.UglifyJS(),
		assets.Fingerprint(),
		assets.WriteToDir("generated"),
	}

	coreJS = assets.Dir("assets/js/core").MustAllFiles().MustFilter(pipeline...)
	homeJS = assets.Dir("assets/js/home").MustAllFiles().MustFilter(pipeline...).Add(coreJS)
}

func whenRenderingHome() {
	// When including your javascript, include all files listed by
	// homeJS.FileNames().
	for _, file := range homeJS.FileNames() {
		fmt.Println(file)
	}
}

func main() {
	onStartUp()
	whenRenderingHome()
}
