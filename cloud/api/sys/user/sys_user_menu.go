package user

import (
	"cloud/api/common"
	"cloud/code"
	"cloud/dao"
	"cloud/internal/response"
	"context"

	"github.com/abulo/ratel/v3/util"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
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
	for _, item := range list {
		if item.Type == 2 {
			continue
		}
		if util.InArray(item.Id, currentMenuIds) {
			curList = append(curList, *item)
		}
	}
	newList := SysUserMenuTree(curList)
	RefactorMenuRedirect(newList)
	RefactorMenuAffix(newList)
	response.JSON(newCtx, consts.StatusOK, utils.H{
		"code": code.Success,
		"msg":  code.StatusText(code.Success),
		"data": newList,
	})
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
