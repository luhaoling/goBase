package errorAndPanic

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"testing"
	"time"
)

type Config struct{}

// Error 显式传递特性
// 强制开发者通过返回值显式处理潜在问题
func ReadConfig(path string) (Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return Config{}, fmt.Errorf("打开文件配置失败:%w", err)
	}
	defer file.Close()
	return Config{}, nil
}

// Panic 的异常传播机制
// 当程序遇到无法继续执行的严重错误时，pannic 会终止当前 goroutine 的正常执行流程，并开始栈展开过程
func MustConnectDB(connStr string) *sql.DB {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(fmt.Sprintf("数据库连接失败:%v", err))
	}
	return db
}

// recover 机制捕获 panic
// recover 使用注意事项
// 1. recover 必须在 defer 函数中调用
// 2. 只能捕获同一 goroutine 的 panic
// 3. 恢复后程序继续执行而不是回滚
func SafeExecute(fn func()) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("捕获到 panic：%v", r)
		}
	}()
	fn()
}

var ErrOrderNotFound error = errors.New("订单处理失败")

type Order struct{}

func FetchOrder(orderID string) (Order, error) {
	return Order{}, nil
}

// 典型应用场景
// 1.可预期的业务错误
func ProcessOrder(orderID string) error {
	order, err := FetchOrder(orderID)
	if errors.Is(err, ErrOrderNotFound) {
		return fmt.Errorf("订单处理失败:%w", err)
	}
	fmt.Println(order)
	return nil
}

type Result struct{}

// 2.外部依赖的暂时故障
func RetryAPIcall() (Result, error) {
	for i := 0; i < 3; i++ {
		res, err := CallAPI()
		if err == nil {
			return res, nil
		}
		time.Sleep(time.Second)
	}
	return Result{}, fmt.Errorf("Api 调用失败")
}

func CallAPI() (Result, error) {
	return Result{}, nil
}

type User struct {
	Name     string
	Password string
}

// 3. 用户输入校验
func ValidateUser(u User) error {
	var errs []error
	if u.Name == "" {
		errs = append(errs, errors.New("用户名不能为空"))
	}
	if len(u.Password) < 8 {
		errs = append(errs, errors.New("密码长度不足"))
	}
	return errors.Join(errs...)
}

// 适用于 Panic 的场景
// 1.程序启动时依赖缺失
func Start() {
	if err := loadConfig(); err != nil {
		panic("关键配置加载失败: " + err.Error())
	}
}
func loadConfig() error {
	return nil
}

type Cache struct {
	closed bool
	store  map[string]string
}

// 2. 不可恢复的异常状态
func (c *Cache) Get(key string) interface{} {
	if c.closed {
		panic("访问已关闭的缓存")
	}
	return c.store[key]
}

// 3. 测试中的断言失败
func TestDivision(t *testing.T) {
	asserEqual := func(a, b int) {
		if a != b {
			panic(fmt.Sprintf("%d!=%d", a, b))
		}
	}
	asserEqual(Divide(10, 2), 5)
}
func Divide(a, b int) int {
	return a / b
}

// 错误包装最佳实践
// 最佳实践1：使用 fmt.Errorf()的 %w 谓词创建错误链
func TestFormatW() {
	fmt.Println(fmt.Errorf("%w", errors.New("error")))
}

// 最佳实践2：使用 errors.Is/As 进行错误判别
func TestErrorsIsAs(err error) {
	if errors.Is(err, ErrOrderNotFound) {

	}
}

// Panic 恢复模式
// 实践1：在 goroutine 入口处设置恢复
func SafeGo(fn func()) {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("goroutine panic: %v", r)
			}
		}()
		fn()
	}()
}

// 实践2：对于 HTTP 服务
func RevoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				w.WriteHeader(http.StatusInternalServerError)
				log.Printf("请求处理 panic")
			}
		}()
		next.ServeHTTP(w, r)
	})
}
