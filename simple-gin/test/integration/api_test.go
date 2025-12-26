// Package integration 集成测试
// 运行方式: go test -v ./test/integration/...
package integration

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"example/simple-gin/internal/config"
	"example/simple-gin/internal/container"
	"example/simple-gin/internal/router"

	"github.com/gin-gonic/gin"
)

// setupTestRouter 创建测试用的路由
func setupTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)

	// 使用测试配置
	cfg := &config.Config{
		Server: config.ServerConfig{
			Port: 8080,
			Mode: "debug",
		},
		DB: config.DatabaseConfig{
			Host: "localhost",
			Port: 5432,
		},
	}

	c, _ := container.NewContainer(cfg)

	r := gin.New()
	routerCfg := &router.RouterConfig{
		EnableSwagger: true, // 测试环境启用 Swagger
	}
	router.SetupRoutes(r, c.UserService, c.ProductService, routerCfg)

	return r
}

// TestPingEndpoint 测试健康检查接口
func TestPingEndpoint(t *testing.T) {
	r := setupTestRouter()

	req, _ := http.NewRequest("GET", "/ping", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)

	// 使用 pkg/response 后，响应格式为 {"code": 0, "msg": "success", "data": {...}}
	if response["code"].(float64) != 0 {
		t.Errorf("Expected code 0, got %v", response["code"])
	}

	if response["msg"] != "success" {
		t.Errorf("Expected msg 'success', got %v", response["msg"])
	}

	data := response["data"].(map[string]interface{})
	if data["message"] != "pong" {
		t.Errorf("Expected data.message 'pong', got %v", data["message"])
	}
}

// TestGetUsers 测试获取用户列表
func TestGetUsers(t *testing.T) {
	r := setupTestRouter()

	req, _ := http.NewRequest("GET", "/api/v1/users", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)

	if response["code"].(float64) != 0 {
		t.Errorf("Expected code 0, got %v", response["code"])
	}

	// 验证返回的是数组
	data, ok := response["data"].([]interface{})
	if !ok {
		t.Error("Expected data to be an array")
	}

	// 应该有预置的用户数据
	if len(data) < 1 {
		t.Error("Expected at least 1 user in seed data")
	}
}

// TestCreateUser 测试创建用户
func TestCreateUser(t *testing.T) {
	r := setupTestRouter()

	user := map[string]string{
		"name":  "测试用户",
		"email": "test@example.com",
		"phone": "13900139000",
	}
	body, _ := json.Marshal(user)

	req, _ := http.NewRequest("POST", "/api/v1/users", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("Expected status 201, got %d", w.Code)
	}

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)

	if response["code"].(float64) != 0 {
		t.Errorf("Expected code 0, got %v", response["code"])
	}

	data := response["data"].(map[string]interface{})
	if data["name"] != "测试用户" {
		t.Errorf("Expected name '测试用户', got %v", data["name"])
	}
}

// TestCreateUserInvalidData 测试创建用户 - 无效数据
func TestCreateUserInvalidData(t *testing.T) {
	r := setupTestRouter()

	// 缺少必填字段
	user := map[string]string{
		"name": "测试用户",
		// 缺少 email 和 phone
	}
	body, _ := json.Marshal(user)

	req, _ := http.NewRequest("POST", "/api/v1/users", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status 400, got %d", w.Code)
	}
}

// TestGetUserNotFound 测试获取不存在的用户
func TestGetUserNotFound(t *testing.T) {
	r := setupTestRouter()

	req, _ := http.NewRequest("GET", "/api/v1/users/9999", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("Expected status 404, got %d", w.Code)
	}
}

// TestGetProducts 测试获取产品列表
func TestGetProducts(t *testing.T) {
	r := setupTestRouter()

	req, _ := http.NewRequest("GET", "/api/v1/products", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)

	data, ok := response["data"].([]interface{})
	if !ok {
		t.Error("Expected data to be an array")
	}

	if len(data) < 1 {
		t.Error("Expected at least 1 product in seed data")
	}
}

// TestCreateUserWithInvalidEmail 测试创建用户 - 无效邮箱格式
func TestCreateUserWithInvalidEmail(t *testing.T) {
	r := setupTestRouter()

	user := map[string]string{
		"name":  "测试用户",
		"email": "invalid-email", // 无效邮箱
		"phone": "13900139000",
	}
	body, _ := json.Marshal(user)

	req, _ := http.NewRequest("POST", "/api/v1/users", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status 400, got %d", w.Code)
	}

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)

	// 验证使用了 pkg/validator 验证
	if response["msg"] != "invalid email format" {
		t.Errorf("Expected msg 'invalid email format', got %v", response["msg"])
	}
}

// TestCreateUserWithInvalidPhone 测试创建用户 - 无效手机号
func TestCreateUserWithInvalidPhone(t *testing.T) {
	r := setupTestRouter()

	user := map[string]string{
		"name":  "测试用户",
		"email": "test@example.com",
		"phone": "123", // 无效手机号
	}
	body, _ := json.Marshal(user)

	req, _ := http.NewRequest("POST", "/api/v1/users", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status 400, got %d", w.Code)
	}

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)

	// 验证使用了 pkg/validator 验证
	if response["msg"] != "invalid phone format" {
		t.Errorf("Expected msg 'invalid phone format', got %v", response["msg"])
	}
}
