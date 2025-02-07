package logic

import (
	"slgserver/server/slgserver/global"
	"slgserver/server/slgserver/logic/mgr"
	"slgserver/util"
)

var ViewWidth = 5
var ViewHeight = 5

//是否在视野范围内
func armyIsInView(rid, x, y int) bool {
	unionId := getUnionId(rid)
	for i := util.MaxInt(x-ViewWidth, 0); i < util.MinInt(x+ViewWidth, global.MapWith) ; i++ {
		for j := util.MaxInt(y-ViewHeight, 0); j < util.MinInt(y+ViewHeight, global.MapHeight) ; j++ {
			build, ok := mgr.RBMgr.PositionBuild(i, j)
			if ok {
				tUnionId := getUnionId(build.RId)
				if (tUnionId != 0 && unionId == tUnionId) || build.RId == rid{
					return true
				}
			}

			city, ok := mgr.RCMgr.PositionCity(i, j)
			if ok {
				tUnionId := getUnionId(city.RId)
				if (tUnionId != 0 && unionId == tUnionId) || city.RId == rid{
					return true
				}
			}
		}
	}

	return false
}