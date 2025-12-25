package container

import (
	"example/simple-gin/config"
	"example/simple-gin/database"
	"example/simple-gin/handlers"
	"example/simple-gin/service"
	"log"
)

// Container 依赖注入容器
// 管理应用中所有的依赖和服务实例
type Container struct {
	// Config
	Config *config.Config

	// Database
	DB service.Database

	// Services
	UserService    service.UserService
	ProductService service.ProductService

	// Handlers
	UserHandler    *handlers.UserHandler
	ProductHandler *handlers.ProductHandler

	// Middleware (如果需要注入)
	// 可以在这里添加中间件、日志系统等
}

// NewContainer 创建并初始化容器
func NewContainer(cfg *config.Config) (*Container, error) {
	c := &Container{
		Config: cfg,
	}

	// 初始化数据库
	if err := c.initDatabase(); err != nil {
		return nil, err
	}

	// 初始化服务
	c.initServices()

	// 初始化处理器
	c.initHandlers()

	log.Println("Container initialized successfully")
	return c, nil
}

// initDatabase 初始化数据库层
func (c *Container) initDatabase() error {
	db, err := database.Init(c.Config)
	if err != nil {
		return err
	}
	c.DB = db
	log.Println("Database layer initialized")
	return nil
}

// initServices 初始化服务层
func (c *Container) initServices() {
	c.UserService = service.NewUserService(c.DB)
	c.ProductService = service.NewProductService(c.DB)
	log.Println("Service layer initialized")
}

// initHandlers 初始化处理器层
func (c *Container) initHandlers() {
	c.UserHandler = handlers.NewUserHandler(c.UserService)
	c.ProductHandler = handlers.NewProductHandler(c.ProductService)
	log.Println("Handler layer initialized")
}

// GetMiddlewares 返回所有中间件
// 可以在这里集中管理中间件，便于后续扩展
func (c *Container) GetMiddlewares() []func(*interface{}) {
	// 这里可以添加需要注入的中间件
	// 暂时不需要，但架构上已经预留了位置
	return nil
}

// 如果需要关闭资源（如数据库连接），可以添加 Close 方法
// func (c *Container) Close() error {
//     // 关闭数据库连接等
//     return nil
// }
