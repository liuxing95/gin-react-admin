package system

import (
	"errors"
	"fmt"

	"gorm.io/gorm"

	"github.com/liuxing95/gin-react-admin/global"
	"github.com/liuxing95/gin-react-admin/model/common/request"
	"github.com/liuxing95/gin-react-admin/model/system"
)

type ApiService struct{}

var ApiServiceApp = new(ApiService)

func (apiService *ApiService) CreateApi(api system.SysApi) (err error) {
	if !errors.Is(global.GRA_DB.Where("path = ? AND method = ?", api.Path, api.Method).First(&system.SysApi{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在相同api")
	}
	return global.GRA_DB.Create(&api).Error
}

// @function: DeleteApi
// @description: 删除基础api
// @param: api model.SysApi
// @return: err error
func (apiService *ApiService) DeleteApi(api system.SysApi) (err error) {
	var entity system.SysApi
	err = global.GRA_DB.Where("id = ?", api.ID).First(&entity).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	err = global.GRA_DB.Delete(&entity).Error
	if err != nil {
		return err
	}
	return nil
}

// @function: GetAPIInfoList
// @description: 分页获取数据,
// @param: api model.SysApi, info request.PageInfo, order string, desc bool
// @return: list interface{}, total int64, err error
func (apiService *ApiService) GetAPIInfoList(api system.SysApi, info request.PageInfo, order string, desc bool) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GRA_DB.Model(&system.SysApi{})
	var apiList []system.SysApi

	if api.Path != "" {
		db = db.Where("path LIKE ?", "%"+api.Path+"%")
	}

	if api.Description != "" {
		db = db.Where("description LIKE ?", "%"+api.Description+"%")
	}

	if api.Method != "" {
		db = db.Where("method = ?", api.Method)
	}

	if api.ApiGroup != "" {
		db = db.Where("api_group = ?", api.ApiGroup)
	}

	err = db.Count(&total).Error

	if err != nil {
		return apiList, total, err
	} else {
		db = db.Limit(limit).Offset(offset)
		if order != "" {
			var orderStr string
			// 设置有效排序key 防止sql注入
			orderMap := make(map[string]bool, 5)
			orderMap["id"] = true
			orderMap["path"] = true
			orderMap["method"] = true
			orderMap["api_group"] = true
			orderMap["description"] = true
			if orderMap[order] {
				if desc {
					orderStr = order + " desc"
				} else {
					orderStr = order
				}
			} else {
				err = fmt.Errorf("非法的排序字段: %v", order)
				return apiList, total, err
			}

			err = db.Order(orderStr).Find(&apiList).Error
		} else {
			err = db.Order("api_group").Find(&apiList).Error
		}

		return apiList, total, err
	}
}

//@function: GetAllApis
//@description: 获取所有的api
//@return:  apis []model.SysApi, err error

func (apiService *ApiService) GetAllApis() (apis []system.SysApi, err error) {
	err = global.GRA_DB.Find(&apis).Error
	return
}

//@function: GetApiById
//@description: 根据id获取api
//@param: id float64
//@return: api model.SysApi, err error

func (apiService *ApiService) GetApiById(id int) (api system.SysApi, err error) {
	err = global.GRA_DB.Where("id = ?", id).First(&api).Error
	return
}

// @function: UpdateApi
// @description: 根据id更新api
// @param: api model.SysApi
// @return: err error
func (apiService *ApiService) UpdateApi(api system.SysApi) (err error) {
	var oldA system.SysApi
	err = global.GRA_DB.Where("id =?", api.ID).First(&oldA).Error
	if oldA.Path != api.Path || oldA.Method != api.Method {
		if !errors.Is(global.GRA_DB.Where("path =? AND method =?", api.Path, api.Method).First(&system.SysApi{}).Error, gorm.ErrRecordNotFound) {
			return errors.New("存在相同api")
		}
	}
	if err != nil {
		return err
	} else {
		err = global.GRA_DB.Save(&api).Error
	}
	return err
}

// @function: DeleteApis
// @description: 删除选中API
// @param: apis []model.SysApi
// @return: err error
func (apiService *ApiService) DeleteApisByIds(ids request.IdsReq) (err error) {
	var apis []system.SysApi
	err = global.GRA_DB.Find(&apis, "id in ?", ids.Ids).Delete(&apis).Error
	if err != nil {
		return err
	}
	return err
}
