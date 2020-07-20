package analyzer

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/q8s-io/heimdall/pkg/infrastructure/kafka"
)

func Signal() {
	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-sigterm:
		log.Println("signal cancelled")

		// 处理业务逻辑
		close(kafka.Queue)

		// time.Sleep(1 * time.Second)
		log.Println("server stop")
		os.Exit(0)
	}
}
