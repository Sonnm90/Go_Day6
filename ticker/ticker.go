package ticker

import (
	"fmt"
	"time"
)

func Demo() {
	ticker := time.NewTicker(1 * time.Second) // Tạo Ticker với chu kỳ 1 giây
	defer ticker.Stop()                       // Đảm bảo dừng Ticker khi chương trình kết thúc

	done := make(chan bool) // Kênh để đồng bộ kết thúc chương trình
	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Ticker fired at", t)
			}
		}
	}()

	time.Sleep(5 * time.Second) // Chờ 5 giây
	done <- true                // Gửi tín hiệu để dừng goroutine

	fmt.Println("Ticker stopped")
}
