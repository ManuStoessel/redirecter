package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	Handler "github.com/ManueStoessel/redirecter/handler"
	Store "github.com/ManueStoessel/redirecter/store"

	"github.com/gin-gonic/gin"
)

var rh *Handler.RedirecterHandler

func main() {
	rh = &Handler.RedirecterHandler{}
	rh.Store = Store.InitializeStore()

	router := gin.Default()

	router.GET("/", Handler.RootHandler)
	//router.POST("/url/:type", Handler.CreateURL)
	//router.GET("/url", Handler.ListURLs)
	//router.GET("/url/:shorthand", Handler.GetURL)
	router.GET("/to/:shorthand", rh.Redirecter)

	srv := &http.Server{
		Addr:    ":80",
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			log.Printf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}
