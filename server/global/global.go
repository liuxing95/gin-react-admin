package global

import (
	"internal/singleflight"
	"sync"

	"github.com/liuxing95/gin-react-admin/config"
	"github.com/liuxing95/gin-react-admin/utils/timer"
	"github.com/redis/go-redis/v9"
	"github.com/songzhibin97/gkit/cache/local_cache"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	GRA_DB     *gorm.DB
	GRA_DBList map[string]*gorm.DB
	GRA_REDIS  *redis.Client
	GRA_CONFIG config.Server
	// GRA_VP     *viper.Viper

	GRA_LOG                 *zap.Logger
	GRA_Timer               timer.Timer = timer.NewTimerTask()
	GRA_Concurrency_Control             = &singleflight.Group{}
	BlackCache              local_cache.Cache
	lock                    sync.RWMutex
)

// GetGlobalDBByDBName 通过名称获取db list中的db
func GetGlobalDBByDBName(dbname string) *gorm.DB {
	lock.RLock()
	defer lock.RUnlock()
	return GRA_DBList[dbname]
}

// MustGetGlobalDBByDBName 通过名称获取db 如果不存在则panic
func MustGetGlobalDBByDBName(dbname string) *gorm.DB {
	lock.RLock()
	defer lock.RUnlock()
	db, ok := GRA_DBList[dbname]
	if !ok || db == nil {
		panic("db no init")
	}
	return db
}
