package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	appservice "yiwen/go-ddd/internal/application/service"
	domainservice "yiwen/go-ddd/internal/domain/service"
	"yiwen/go-ddd/internal/infrastructure/config"
	"yiwen/go-ddd/internal/infrastructure/persistence/model"
	mysqlrepo "yiwen/go-ddd/internal/infrastructure/persistence/mysql"
	"yiwen/go-ddd/internal/interfaces/api/handler"
	"yiwen/go-ddd/internal/interfaces/api/middleware"
	"yiwen/go-ddd/internal/interfaces/api/router"
)

func main() {
	// 解析命令行参数
	configPath := flag.String("config", "config/config.yaml", "config file path")
	flag.Parse()

	// 加载配置
	cfg, err := config.Load(*configPath)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// 设置Gin模式
	gin.SetMode(cfg.App.Mode)

	// 初始化数据库
	db, err := initDatabase(cfg)
	if err != nil {
		log.Fatalf("Failed to connect database: %v", err)
	}

	// 依赖注入
	// 这是DDD中的重要实践：在应用启动时进行依赖注入
	// 各层之间通过接口解耦，便于测试和维护

	// 1. 初始化仓储层（基础设施层）
	userRepo := mysqlrepo.NewUserRepository(db)

	// 2. 初始化领域服务（领域层）
	userDomainService := domainservice.NewUserDomainService(userRepo)

	// 3. 初始化应用服务（应用层）
	userAppService := appservice.NewUserApplicationService(userRepo, userDomainService)

	// 4. 初始化JWT认证
	jwtAuth := middleware.NewJWTAuth(cfg.JWT.Secret, cfg.JWT.ExpireHour, cfg.JWT.Issuer)

	// 5. 初始化HTTP处理器（接口层）
	userHandler := handler.NewUserHandler(userAppService, jwtAuth)

	// 6. 初始化路由
	r := router.NewRouter(userHandler, jwtAuth)
	engine := r.Setup()

	// 启动服务
	addr := fmt.Sprintf(":%d", cfg.App.Port)
	log.Printf("Server starting on %s", addr)
	log.Printf("Health check: http://localhost%s/health", addr)
	log.Printf("API base: http://localhost%s/api/v1", addr)

	if err := engine.Run(addr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

// initDatabase 初始化数据库连接
func initDatabase(cfg *config.Config) (*gorm.DB, error) {
	// 配置GORM日志
	var logLevel logger.LogLevel
	switch cfg.App.Mode {
	case "debug":
		logLevel = logger.Info
	case "test":
		logLevel = logger.Warn
	default:
		logLevel = logger.Error
	}

	db, err := gorm.Open(mysql.Open(cfg.Database.DSN()), &gorm.Config{
		Logger: logger.Default.LogMode(logLevel),
	})
	if err != nil {
		return nil, err
	}

	// 获取底层 sql.DB
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	// 设置连接池
	sqlDB.SetMaxIdleConns(cfg.Database.MaxIdleConns)
	sqlDB.SetMaxOpenConns(cfg.Database.MaxOpenConns)

	// 自动迁移（开发环境使用，生产环境建议使用SQL脚本）
	if cfg.App.Mode == "debug" {
		if err := db.AutoMigrate(&model.UserModel{}); err != nil {
			return nil, err
		}
	}

	return db, nil
}
