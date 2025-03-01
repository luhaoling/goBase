package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var psCertificate sync.Mutex
	var propertyCertificate sync.Mutex

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()

		psCertificate.Lock()
		defer psCertificate.Unlock()

		time.Sleep(5 * time.Millisecond)

		propertyCertificate.Lock()
		propertyCertificate.Unlock()
	}()

	go func() {
		defer wg.Done()

		propertyCertificate.Lock()
		defer propertyCertificate.Unlock()

		time.Sleep(5 * time.Millisecond)

		psCertificate.Lock()
		psCertificate.Unlock()
	}()

	wg.Wait()
	fmt.Println("Done")
}

//解决办法
//
//func main() {
//	var ps sync.Mutex
//	var psCertificate sync.Mutex
//	var propertyCertificate sync.Mutex
//
//	// 引入第三方锁
//	var globalLock sync.Mutex
//
//	var wg sync.WaitGroup
//	wg.Add(2)
//
//	go func() {
//		defer wg.Done()
//
//		// 先获取全局锁，保证获取资源锁时的顺序一致
//		globalLock.Lock()
//		psCertificate.Lock()
//		propertyCertificate.Lock()
//		globalLock.Unlock()
//
//		defer psCertificate.Unlock()
//		defer propertyCertificate.Unlock()
//
//		time.Sleep(5 * time.Millisecond)
//	}()
//
//	go func() {
//		defer wg.Done()
//
//		// 先获取全局锁，保证获取资源锁时的顺序一致
//		globalLock.Lock()
//		propertyCertificate.Lock()
//		psCertificate.Lock()
//		globalLock.Unlock()
//
//		defer propertyCertificate.Unlock()
//		defer psCertificate.Unlock()
//
//		time.Sleep(5 * time.Millisecond)
//	}()
//
//	wg.Wait()
//	fmt.Println("Done")
//}
