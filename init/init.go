package init

import (
	"fmt"
	global "github.com/AmbitionLover/go-snake/pkg"
	"github.com/AmbitionLover/go-snake/pkg/core"
)

/*
@Date:2022/2/7
@Author Ambi 赵帅
System init
*/

func init() {
	showLogo()
	baseInit()
	RunServe()
}

// Version 版本号
const Version = "0.0.1"

func showLogo() {
	fmt.Println("go-snake 当前版本：", Version)
	fmt.Println("go-snake 作者：Ambi")
	fmt.Println("go-snake 作者联系方式：QQ：1277518148")
}

func baseInit() {
	global.VP = core.Viper() // 初始化Viper
	global.LOG = core.Zap()  // 初始化zap日志库
	//global.DB = initialize.Gorm() // gorm连接数据库
	//if global.GVA_DB != nil {
	//	initialize.RegisterTables(global.GVA_DB) // 初始化表
	//	// 程序结束前关闭数据库链接
	//	db, _ := global.GVA_DB.DB()
	//	defer db.Close()
	//}

}
