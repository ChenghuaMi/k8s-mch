package main

import (
	"k8s-mch/global"
	"k8s-mch/initiallize"
)

func main() {
	r := initiallize.Routers()
	initiallize.Viper()
	initiallize.K8s()
	err := r.Run(global.Conf.System.Address)
	if err != nil {
		panic(err)
	}

}
