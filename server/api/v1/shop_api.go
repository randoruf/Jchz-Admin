package v1

import (
	"github.com/gin-gonic/gin"
	"jchz-admin/core/system"
	"jchz-admin/model/request"
	"jchz-admin/model/response"
	"net/http"
	"strconv"
)

type ShopApi struct{}

var ShopApiGroup = new(ShopApi)

func (s *ShopApi) GetShopList(c *gin.Context) {
	var QueryParam request.ShopQueryRequest
	var err error
	QueryParam.Query = c.Query("query")
	pageNumStr, ok := c.GetQuery("pagenum")
	if !ok {
		response.FailWithDetailed(http.StatusBadRequest, "参数错误", c)
		return
	}
	QueryParam.Pagenum, err = strconv.Atoi(pageNumStr)
	if err != nil {
		response.FailWithDetailed(http.StatusBadRequest, "参数错误", c)
		return
	}
	pageSizeStr, ok := c.GetQuery("pagesize")
	if !ok {
		response.FailWithDetailed(http.StatusBadRequest, "参数错误", c)
		return
	}
	QueryParam.Pagesize, err = strconv.Atoi(pageSizeStr)
	if err != nil {
		response.FailWithDetailed(http.StatusBadRequest, "参数错误", c)
		return
	}

	ShopList, total, err := shopService.QueryShopsList(QueryParam)
	if err != nil {
		response.FailWithDetailed(http.StatusInternalServerError, "获取商店列表失败", c)
		return
	}

	c.JSON(http.StatusOK, response.ShopPageResponse{
		Data: &response.ShopPageData{
			Total:   total,
			PageNum: int64(QueryParam.Pagenum),
			Shops:   ShopList,
		},
		Meta: &response.Meta{
			Msg:    "获取商店列表成功",
			Status: http.StatusOK,
		},
	})
}

func (s *ShopApi) AddShop(c *gin.Context) {
	var addShopReq request.AddShopRequest
	err := c.ShouldBindJSON(&addShopReq)
	if err != nil {
		response.FailWithDetailed(http.StatusBadRequest, "参数错误", c)
		return
	}
	newShop, err := shopService.AddShop(addShopReq)
	if err != nil {
		response.FailWithDetailed(http.StatusInternalServerError, "添加用户失败", c)
		return
	}
	c.JSON(http.StatusOK, response.AddShopResponse{
		ShopData: newShop,
		Meta: &response.Meta{
			Msg:    "添加商店成功",
			Status: http.StatusOK,
		},
	})
}

func (s *ShopApi) UpdateShop(c *gin.Context) {
	var updateShopReq request.UpdateShopRequest
	uid := c.Param("id")
	if uid == "" {
		response.FailWithDetailed(http.StatusBadRequest, "参数错误", c)
		return
	}
	err := c.ShouldBindJSON(&updateShopReq)
	if err != nil {
		response.FailWithDetailed(http.StatusBadRequest, "参数错误", c)
		return
	}
	newShop, err := shopService.UpdateShop(uid, updateShopReq)
	if err != nil {
		system.PrintError(err)
		response.FailWithDetailed(http.StatusInternalServerError, "更新商店失败", c)
		return
	}
	c.JSON(http.StatusOK, response.UpdateShopResponse{
		ShopData: newShop,
		Meta: &response.Meta{
			Msg:    "更新商店成功",
			Status: http.StatusOK,
		},
	})
}

func (s *ShopApi) DeleteShop(c *gin.Context) {
	uid := c.Param("id")
	if uid == "" {
		response.FailWithDetailed(http.StatusBadRequest, "参数错误", c)
		return
	}
	flag, err := shopService.DeleteShop(uid)
	if err != nil {
		response.FailWithDetailed(http.StatusBadRequest, "参数错误", c)
		return
	}
	if flag == false {
		system.PrintError(err)
		response.FailWithDetailed(http.StatusInternalServerError, "删除商店失败", c)
		return
	}
	response.SuccessWithNullData(http.StatusOK, "删除商店成功", c)
	return
}
