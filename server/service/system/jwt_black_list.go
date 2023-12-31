package system

import (
	"context"

	"github.com/liuxing95/gin-react-admin/global"
	"github.com/liuxing95/gin-react-admin/model/system"
	"github.com/liuxing95/gin-react-admin/utils"
	"go.uber.org/zap"
)

type JwtService struct{}

//@function: JsonInBlacklist
//@description: 拉黑jwt
//@param: jwtList model.JwtBlacklist
//@return: err error
func (jwtService *JwtService) JsonInBlacklist(jwtList system.JwtBlacklist) (err error) {
	err = global.GRA_DB.Create(&jwtList).Error
	if err != nil {
		return
	}
	global.BlackCache.SetDefault(jwtList.Jwt, struct{}{})
	return
}

//@function: IsBlacklist
//@description: 判断JWT是否在黑名单内部
//@param: jwt string
//@return: bool
func (jwtService *JwtService) IsBlacklist(jwt string) bool {
	_, ok := global.BlackCache.Get(jwt)
	return ok
}

//@function: GetRedisJWT
//@description: 从redis取jwt
//@param: userName string
//@return: redisJWT string, err error

func (jwtService *JwtService) GetRedisJWT(userName string) (redisJWT string, err error) {
	redisJWT, err = global.GRA_REDIS.Get(context.Background(), userName).Result()
	return redisJWT, err
}

//@function: SetRedisJWT
//@description: jwt存入redis并设置过期时间
//@param: jwt string, userName string
//@return: err error
func (jwtService *JwtService) SetRedisJWT(jwt string, userName string) (err error) {
	// 此处过期时间等于jwt过期时间
	dr, err := utils.ParseDuration(global.GRA_CONFIG.JWT.ExpiresTime)
	if err != nil {
		return err
	}
	timer := dr
	err = global.GRA_REDIS.Set(context.Background(), userName, jwt, timer).Err()
	return err
}

func LoadAll() {
	var data []string
	err := global.GRA_DB.Model(&system.JwtBlacklist{}).Select("jwt").Find(&data).Error
	if err != nil {
		global.GRA_LOG.Error("加载数据库jwt黑名单失败!", zap.Error(err))
	}
	for i := 0; i < len(data); i++ {
		global.BlackCache.SetDefault(data[i], struct{}{})
	} // jwt黑名单 加入 BlackCache 中
}
