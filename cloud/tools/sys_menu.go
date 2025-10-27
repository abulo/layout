package tools

import "cloud/dao"

func SysMenuTree(menus []dao.SysMenu) []*dao.SysMenu {
	deptMap := make(map[int64]*dao.SysMenu)
	parentMap := make(map[int64]bool)
	var roots []*dao.SysMenu
	for i := range menus {
		deptMap[*menus[i].Id] = &menus[i]
		parentMap[*menus[i].ParentId] = true
	}
	for i := range menus {
		department := &menus[i]
		if _, ok := deptMap[*department.ParentId]; !ok {
			roots = append(roots, department)
		} else {
			if parent, ok := deptMap[*department.ParentId]; ok {
				parent.Children = append(parent.Children, department)
			}
		}
	}

	return roots
}
