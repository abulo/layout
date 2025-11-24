package user

import (
	"cloud/api/common"
	"cloud/code"
	"cloud/dao"
	"cloud/internal/response"
	"cloud/service/sys/menu"
	"context"

	"github.com/abulo/ratel/v3/util"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/gogo/protobuf/proto"
)

// 获取登录用户的菜单
func SysUserMenu(ctx context.Context, newCtx *app.RequestContext) {
	tenantId := newCtx.GetInt64("tenantId") // 租户 Id
	sysUserId := newCtx.GetInt64("userId")  // 用户 Id
	list, currentMenuIds, err := common.SysUserTree(ctx, sysUserId, tenantId)
	if err != nil {
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.TokenError,
			"msg":  err.Error(),
		})
		return
	}
	var curList []dao.SysMenuTree
	var homeId int64
	homeId = 0
	for _, item := range list {
		if item.Type == 2 {
			continue
		}
		if util.InArray(item.Id, currentMenuIds) {
			curList = append(curList, *item)
			if item.Id > homeId {
				homeId = item.Id
			}
		}
	}
	homeId = homeId + 1
	homeMenu := common.SysUserMenuDao(BuildVMenu(homeId))
	// 将homeMenu插入到curList的最前面
	curList = append([]dao.SysMenuTree{*homeMenu}, curList...)
	newList := SysUserMenuTree(curList)
	RefactorMenuRedirect(newList)
	RefactorMenuAffix(newList)
	response.JSON(newCtx, consts.StatusOK, utils.H{
		"code": code.Success,
		"msg":  code.StatusText(code.Success),
		"data": newList,
	})
}

func BuildVMenu(homeId int64) *menu.SysMenuObject {
	item := &menu.SysMenuObject{}
	item.Id = proto.Int64(homeId)
	item.Name = proto.String("首页")
	item.Code = proto.String("home")
	item.Type = proto.Int32(1)
	item.Sort = proto.Int32(0)
	item.ParentId = proto.Int64(0)
	item.Path = proto.String("/home/index")
	item.Icon = proto.String("ep:home-filled")
	item.Component = proto.String("/home/index")
	item.ComponentName = proto.String("Home")
	item.Hide = proto.Int32(0)
	item.Cache = proto.Int32(0)
	item.Full = proto.Int32(0)
	item.Status = proto.Int32(0)
	return item
}

func SysUserMenuTree(menus []dao.SysMenuTree) []*dao.SysMenuTree {
	deptMap := make(map[int64]*dao.SysMenuTree)
	parentMap := make(map[int64]bool)
	var roots []*dao.SysMenuTree
	for i := range menus {
		deptMap[menus[i].Id] = &menus[i]
		parentMap[menus[i].ParentId] = true
	}
	for i := range menus {
		department := &menus[i]
		if _, ok := deptMap[department.ParentId]; !ok {
			roots = append(roots, department)
		} else {
			if parent, ok := deptMap[department.ParentId]; ok {
				parent.Children = append(parent.Children, department)
			}
		}
	}

	return roots
}

func RefactorMenuRedirect(menuList []*dao.SysMenuTree) {
	for _, item := range menuList {
		if item.Component == "" {
			for _, child := range item.Children {
				if !child.Meta.IsHide {
					item.Redirect = child.Path
					break
				}
			}
		}
		if item.Children != nil {
			RefactorMenuRedirect(item.Children)
		}
	}
}

func RefactorMenuAffix(menuList []*dao.SysMenuTree) {
	// 只是需要设置一个元素的值affix 为true
	for _, item := range menuList {
		if item.Component != "" && !item.Meta.IsHide {
			item.Meta.IsAffix = true
			break
		}
		if item.Children != nil {
			RefactorMenuAffix(item.Children)
		}
	}
}
