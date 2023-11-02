/**
 * @author mch
 */

package initiallize

import (
	"fmt"
	"github.com/spf13/viper"
	"k8s-mch/global"
)

func Viper() {
	v := viper.New()
	v.SetConfigFile("config.yaml")
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(err.Error())
	}
	err = v.Unmarshal(&global.Conf)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("config:",global.Conf)
}
