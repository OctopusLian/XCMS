/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-01-02 09:14:06
 * @LastEditors: neozhang
 * @LastEditTime: 2022-01-02 09:14:27
 */
package sysinit

import (
	"time"
	"xcms/utils"

	cache "github.com/patrickmn/go-cache"
)

func init() {
	//init cache
	utils.Cache = cache.New(60*time.Minute, 120*time.Minute)

	//init db
	initDB()
}
