package v1

import (
	"github.com/gin-gonic/gin"
	"jchz-admin/core/system"
	"jchz-admin/model/response"
	"net/http"
)

type MenuApi struct{}

var MenuApiGroup = new(MenuApi)

func (m *MenuApi) MenuList(c *gin.Context) {
	list, err := menuService.GetMenuList()
	if err != nil {
		system.PrintError(err)
		response.FailWithDetailed(http.StatusInternalServerError, "获取菜单失败", c)
	}

	var data []*response.MenusData

	for i := 0; i < len(list); i++ {
		if list[i].Order == 0 {
			continue
		}
		var children []*response.MenusData
		for j := 0; j < len(list); j++ {
			if list[j].ParentID == list[i].Id {
				children = append(children, &response.MenusData{
					Id:       list[j].Id,
					AuthName: list[j].AuthName,
					Path:     list[j].Path,
					Children: []*response.MenusData{},
					Order:    nil,
				})
			}
		}
		data = append(data, &response.MenusData{
			Id:       list[i].Id,
			AuthName: list[i].AuthName,
			Path:     list[i].Path,
			Children: children,
			Order:    list[i].Order,
		})
	}
	res := &response.MenusResponse{Data: data, Meta: &response.Meta{
		Msg:    "获取菜单成功",
		Status: http.StatusOK,
	}}
	c.JSON(http.StatusOK, res)
}
