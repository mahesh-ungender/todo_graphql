package runner

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"
	"todo_graphql/api/router"
	"todo_graphql/api/service"
	"todo_graphql/logger"
	// For doc generation
	// _ "auth/api/swagger"
)

type API interface{
	Go(ctx context.Context, wg *sync.WaitGroup)
}

type api struct {
}

func (runner *api) Go(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	services := service.Init()

	//i think i have to add logger file
	logger.Log.Infof("Starting API server on %v...", "3000")

	routerV1 := router.Init(services)
	s := &http.Server{
		Addr:         fmt.Sprintf(":%s","3000"),
		Handler:      routerV1,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}
	s.SetKeepAlivesEnabled(true)
	s.ListenAndServe()
}

// NewAPI returns an instance of the REST API runner
func NewAPI() API {
	return &api{}
}