package app

import (
	"github.com/velocitykode/velocity/pkg/cache"
	"github.com/velocitykode/velocity/pkg/events"
	"github.com/velocitykode/velocity/pkg/log"
	"github.com/velocitykode/velocity/pkg/orm"
	"github.com/velocitykode/velocity/pkg/router"
)

// initEvents registers event listeners for framework observability.
// Customize these listeners to add your own logging, metrics, or tracing.
func initEvents() {
	// Request lifecycle events
	events.On("request.started", func(e interface{}) error {
		if req, ok := e.(*router.RequestStarted); ok {
			log.Debug("Request started",
				"request_id", req.RequestID,
				"method", req.Method,
				"path", req.Path,
			)
		}
		return nil
	})

	events.On("request.handled", func(e interface{}) error {
		if req, ok := e.(*router.RequestHandled); ok {
			log.Info("Request completed",
				"request_id", req.RequestID,
				"method", req.Method,
				"path", req.Path,
				"status", req.StatusCode,
				"duration", req.Duration,
			)
		}
		return nil
	})

	events.On("request.failed", func(e interface{}) error {
		if req, ok := e.(*router.RequestFailed); ok {
			log.Error("Request failed",
				"request_id", req.RequestID,
				"error", req.Error,
				"recovered", req.Recovered,
			)
		}
		return nil
	})

	// Database query events
	events.On("query.executed", func(e interface{}) error {
		if q, ok := e.(*orm.QueryExecuted); ok {
			log.Debug("Query executed",
				"sql", q.SQL,
				"duration", q.Duration,
				"rows", q.RowsAffected,
			)
		}
		return nil
	})

	// Cache events
	events.On("cache.hit", func(e interface{}) error {
		if c, ok := e.(*cache.CacheHit); ok {
			log.Debug("Cache hit", "key", c.Key)
		}
		return nil
	})

	events.On("cache.miss", func(e interface{}) error {
		if c, ok := e.(*cache.CacheMiss); ok {
			log.Debug("Cache miss", "key", c.Key)
		}
		return nil
	})
}
