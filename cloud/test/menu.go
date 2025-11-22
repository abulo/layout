package main

import (
	"cloud/dao"
	"encoding/json"
	"fmt"

	"github.com/abulo/ratel/v3/util"
)

func main() {
	jsonString := `[
        {
            "path": "/sys",
            "name": "",
            "meta": {
                "icon": "ep:setting",
                "title": "系统管理",
                "isHide": false,
                "isFull": false,
                "isAffix": false,
                "isKeepAlive": false
            },
            "children": [
                {
                    "path": "/sys/menu",
                    "name": "SysMenu",
                    "component": "/sys/menu/index",
                    "meta": {
                        "icon": "ep:menu",
                        "title": "菜单管理",
                        "isHide": false,
                        "isFull": false,
                        "isAffix": false,
                        "isKeepAlive": false
                    }
                },
				{
                    "path": "/sys/menu",
                    "name": "SysMenu",
                    "component": "/sys/menu/index",
                    "meta": {
                        "icon": "ep:menu",
                        "title": "菜单管理",
                        "isHide": false,
                        "isFull": false,
                        "isAffix": false,
                        "isKeepAlive": false
                    }
                },
                {
                    "path": "/sys/dict",
                    "name": "SysDictType",
                    "component": "/sys/dict/index",
                    "meta": {
                        "icon": "ep:collection",
                        "title": "字典管理",
                        "isHide": false,
                        "isFull": false,
                        "isAffix": false,
                        "isKeepAlive": false
                    },
                    "children": [
                        {
                            "path": "/sys/dict/:dictTypeId",
                            "name": "SysDict",
                            "component": "/sys/dict/item",
                            "meta": {
                                "icon": "ep:data-board",
                                "title": "字典列表",
                                "isHide": true,
                                "isFull": false,
                                "isAffix": false,
                                "isKeepAlive": false,
                                "activeMenu": "/sys/dict"
                            }
                        },
						{
                            "path": "/sys/dict/:dictTypeId",
                            "name": "SysDict1",
                            "component": "/sys/dict/item1",
                            "meta": {
                                "icon": "ep:data-board",
                                "title": "字典列表",
                                "isHide": false,
                                "isFull": false,
                                "isAffix": false,
                                "isKeepAlive": false,
                                "activeMenu": "/sys/dict"
                            }
                        }
                    ]
                }
            ]
        }
    ]`
	var menuList []*dao.SysMenuTree
	json.Unmarshal([]byte(jsonString), &menuList)
	refactorMenuRedirect(menuList)
	refactorMenuAffix(menuList)
	fmt.Println(util.JSONString(menuList))
}

// 当 menuList中的元素component是空值时需要设置 redirect 的值为  children中有值中 isHide 是否为false的children 的 path
// 希望是递归函数来完成, 返回新的menuList
func refactorMenuRedirect(menuList []*dao.SysMenuTree) {
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
			refactorMenuRedirect(item.Children)
		}
	}
}

func refactorMenuAffix(menuList []*dao.SysMenuTree) {
	// 只是需要设置一个元素的值affix 为true
	for _, item := range menuList {
		if item.Component != "" && !item.Meta.IsHide {
			item.Meta.IsAffix = true
			break
		}
		if item.Children != nil {
			refactorMenuAffix(item.Children)
		}
	}
}
