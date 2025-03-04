package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"strconv"
	"sync"
)

// 实现用户服务 API
// 实现注意细节，方法写在 interface{} 里面，struct 一般是定义一个带有“物理”属性的内容，如存储块、用户属性等

//type User struct {
//	ID   int
//	Name string
//	Age  string
//}
//
//type UserService interface {
//	GetUser(id string) (*User, error)
//}
//
//type MemoryUserService struct {
//	U map[string]*User
//}
//
//type UserS struct {
//	MemoryUserService
//}
//
//var service MemoryUserService
//
//func main() {
//	service = NewMemoryUserService()
//	r := gin.Default()
//	r.GET("/user/:id", GetUserByID)
//	r.Run(":8080")
//}
//
//func NewMemoryUserService() MemoryUserService {
//	b := MemoryUserService{
//		U: make(map[string]*User),
//	}
//	b.AddUser(&User{ID: 1, Name: "Alice"})
//	return b
//}
//
//func (m *MemoryUserService) AddUser(u *User) {
//	m.U[strconv.Itoa(u.ID)] = u
//}
//
//func GetUserByID(ctx *gin.Context) {
//	id := ctx.Param("id")
//	user := &User{}
//	u, err := user.GetUser(id)
//	if err != nil {
//		fmt.Println(err)
//		ctx.JSON(404, gin.H{
//			"User":  nil,
//			"error": err.Error(),
//		})
//		return
//	}
//	ctx.JSON(200, gin.H{
//		"User":  u,
//		"error": nil,
//	})
//
//}
//
//func (u *User) GetUser(id string) (*User, error) {
//	if c, ok := service.U[id]; !ok {
//		return nil, errors.New("user not exist")
//	} else {
//		return c, nil
//	}
//}

// 优化

//// User 结构体定义
//type User struct {
//	ID    string `json:"id"`
//	Name  string `json:"name"`
//	Email string `json:"email"`
//}
//
//// UserService 接口定义
//type UserService interface {
//	CreateUser(user User) (User, error)
//	GetUser(id string) (User, error)
//	UpdateUser(user User) (User, error)
//	DeleteUser(id string) error
//}
//
//// InMemoryUserService 内存实现
//type InMemoryUserService struct {
//	users map[string]User
//	mu    sync.RWMutex
//}
//
//func NewInMemoryUserService() *InMemoryUserService {
//	return &InMemoryUserService{
//		users: make(map[string]User),
//	}
//}
//
//func (s *InMemoryUserService) CreateUser(user User) (User, error) {
//	s.mu.Lock()
//	defer s.mu.Unlock()
//
//	if user.ID == "" {
//		return User{}, fmt.Errorf("user ID is required")
//	}
//
//	if _, exists := s.users[user.ID]; exists {
//		return User{}, fmt.Errorf("user ID %s already exists", user.ID)
//	}
//
//	s.users[user.ID] = user
//	return user, nil
//}
//
//func (s *InMemoryUserService) GetUser(id string) (User, error) {
//	s.mu.RLock()
//	defer s.mu.RUnlock()
//
//	user, exists := s.users[id]
//	if !exists {
//		return User{}, fmt.Errorf("user not found")
//	}
//
//	return user, nil
//}
//
//func (s *InMemoryUserService) UpdateUser(user User) (User, error) {
//	s.mu.Lock()
//	defer s.mu.Unlock()
//
//	if _, exists := s.users[user.ID]; !exists {
//		return User{}, fmt.Errorf("user not found")
//	}
//
//	s.users[user.ID] = user
//	return user, nil
//}
//
//func (s *InMemoryUserService) DeleteUser(id string) error {
//	s.mu.Lock()
//	defer s.mu.Unlock()
//
//	if _, exists := s.users[id]; !exists {
//		return fmt.Errorf("user not found")
//	}
//
//	delete(s.users, id)
//	return nil
//}
//
//// Gin 路由处理函数
//func createUserHandler(service UserService) gin.HandlerFunc {
//	return func(c *gin.Context) {
//		var user User
//		if err := c.ShouldBindJSON(&user); err != nil {
//			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//			return
//		}
//
//		createdUser, err := service.CreateUser(user)
//		if err != nil {
//			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
//			return
//		}
//
//		c.JSON(http.StatusCreated, createdUser)
//	}
//}
//
//func getUserHandler(service UserService) gin.HandlerFunc {
//	return func(c *gin.Context) {
//		id := c.Param("id")
//
//		user, err := service.GetUser(id)
//		if err != nil {
//			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
//			return
//		}
//
//		c.JSON(http.StatusOK, user)
//	}
//}
//
//func updateUserHandler(service UserService) gin.HandlerFunc {
//	return func(c *gin.Context) {
//		id := c.Param("id")
//		var user User
//
//		if err := c.ShouldBindJSON(&user); err != nil {
//			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//			return
//		}
//
//		if user.ID != id {
//			c.JSON(http.StatusBadRequest, gin.H{"error": "ID mismatch"})
//			return
//		}
//
//		updatedUser, err := service.UpdateUser(user)
//		if err != nil {
//			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
//			return
//		}
//
//		c.JSON(http.StatusOK, updatedUser)
//	}
//}
//
//func deleteUserHandler(service UserService) gin.HandlerFunc {
//	return func(c *gin.Context) {
//		id := c.Param("id")
//
//		if err := service.DeleteUser(id); err != nil {
//			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
//			return
//		}
//
//		c.Status(http.StatusNoContent)
//	}
//}
//
//func main() {
//	// 初始化服务
//	userService := NewInMemoryUserService()
//
//	// 配置 Gin 路由
//	router := gin.Default()
//
//	// 注册路由
//	router.POST("/users", createUserHandler(userService))
//	router.GET("/users/:id", getUserHandler(userService))
//	router.PUT("/users/:id", updateUserHandler(userService))
//	router.DELETE("/users/:id", deleteUserHandler(userService))
//
//	// 启动服务
//	router.Run(":8080")
//}

// 优化2：

// Controller
type Controller interface {
	GetUserByID()
}

// Service
type MemoryService interface {
	GetUser(int) (*User, error)
}

type User struct {
	ID   int
	Name string
	Age  int
}

type Memory struct {
	m map[int]*User
	sync.Mutex
}

type Global struct {
	Memory
	MemoryService
}

func NewGlobal() *Global {
	g := &Global{
		Memory: Memory{
			m:     make(map[int]*User),
			Mutex: sync.Mutex{},
		},
		MemoryService: &Memory{},
	}
	g.Memory.m[1] = &User{ID: 1, Name: "Alice"}
	return g
}

func main() {
	r := gin.Default()
	g := NewGlobal()
	r.GET("/user/:id", g.GetUserByID)
	r.Run(":8080")
}

func (g *Global) GetUserByID(c *gin.Context) {
	p := c.Param("id")
	id, err := strconv.Atoi(p)
	if err != nil {
		c.JSON(500, gin.H{
			"User": nil,
			"Err":  err,
		})
		return
	}
	user, err := g.Memory.GetUser(id)
	if err != nil {
		c.JSON(404, err)
		return
	}
	c.JSON(200, gin.H{
		"User": user,
		"Err":  nil,
	})
	return
}

func (m *Memory) GetUser(id int) (*User, error) {
	if user, ok := m.m[id]; !ok {
		return nil, errors.New("user not exist")
	} else {
		return user, nil
	}
}
