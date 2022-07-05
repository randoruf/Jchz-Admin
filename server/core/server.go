package core

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"jchz-admin/Initialize"
	"jchz-admin/config"
	"jchz-admin/core/system"
	"jchz-admin/global"
)

func RunServer() {
	gin.SetMode(gin.ReleaseMode)
	fmt.Println("--------------------JCHZ-ADMIN--------------------")
	fmt.Println(system.SysInfo, "JCHZ-ADMIN system Initializing...")
	global.JA_CONFIG = config.InitConfig()
	global.JA_DB = Initialize.InitGorm()
	if global.JA_DB == nil {
		system.PrintCrash(errors.New("DB initialization failed, please check config.yaml whether the database information you entered is correct"))
		system.PrintCrash(errors.New("Exiting..."))
		return
	}
	router := Initialize.InitRouters()
	system.PrintSysInfo("JCHZ-ADMIN system initialization completed!")
	fmt.Printf(`
	::::::'##::'######::'##::::'##:'########:
	:::::: ##:'##... ##: ##:::: ##:..... ##::
	:::::: ##: ##:::..:: ##:::: ##::::: ##:::
	:::::: ##: ##::::::: #########:::: ##::::
	'##::: ##: ##::::::: ##.... ##::: ##:::::
	 ##::: ##: ##::: ##: ##:::: ##:: ##::::::
	. ######::. ######:: ##:::: ##: ########:
	:......::::......:::..:::::..::........::
	:::'###::::'########::'##::::'##:'####:'##::: ##:
	::'## ##::: ##.... ##: ###::'###:. ##:: ###:: ##:
	:'##:. ##:: ##:::: ##: ####'####:: ##:: ####: ##:
	'##:::. ##: ##:::: ##: ## ### ##:: ##:: ## ## ##:
	 #########: ##:::: ##: ##. #: ##:: ##:: ##. ####:
	 ##.... ##: ##:::: ##: ##:.:: ##:: ##:: ##:. ###:
	 ##:::: ##: ########:: ##:::: ##:'####: ##::. ##:
	..:::::..::........:::..:::::..::....::..::::..::
	` + "\n")
	system.PrintSysInfo("JCHZ-ADMIN system is listening on port: " + global.JA_CONFIG.SystemConfig.Port)
	err := router.Run(":" + global.JA_CONFIG.SystemConfig.Port)
	if err != nil {
		system.PrintCrash(err)
	}
}
