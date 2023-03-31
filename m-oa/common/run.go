package common

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Run(r *gin.Engine, srvName string, Addr string) {
	srv := &http.Server{
		Addr:    Addr,
		Handler: r,
	}

	go func() {
		log.Printf("%s server runing in %s\n", srvName, srv.Addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalln(err)
		}
	}()

	quit := make(chan os.Signal)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGALRM)
	<-quit
	log.Printf("shutting Down project %s...\n", srvName)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalln("web server shutdown,cause by", err)
	}

	select {
	case <-ctx.Done():
		log.Println("关闭超时")
	}
	log.Printf("%s stop success\n", srvName)
}
