package app

import (
	"github.com/velocitykode/velocity/pkg/cache"
	"github.com/velocitykode/velocity/pkg/events"
	"github.com/velocitykode/velocity/pkg/log"
	"github.com/velocitykode/velocity/pkg/orm"
	"github.com/velocitykode/velocity/pkg/router"
)

// listenerFunc adapts a plain function to the events.Listener interface.
type listenerFunc struct {
	fn func(event interface{}) error
}

func (l *listenerFunc) Handle(event interface{}) error { return l.fn(event) }
func (l *listenerFunc) ShouldQueue() bool              { return false }

// initEvents registers event listeners for framework observability.
// Customize these listeners to add your own logging, metrics, or tracing.
func initEvents(logger log.Logger, dispatcher events.Dispatcher) {
	// Request lifecycle events
	dispatcher.Listen("request.started", &listenerFunc{fn: func(e interface{}) error {
		if req, ok := e.(*router.RequestStarted); ok {
			logger.Debug("Request started",
				"request_id", req.RequestID,
				"method", req.Method,
				"path", req.Path,
			)
		}
		return nil
	}})

	dispatcher.Listen("request.handled", &listenerFunc{fn: func(e interface{}) error {
		if req, ok := e.(*router.RequestHandled); ok {
			logger.Info("Request completed",
				"request_id", req.RequestID,
				"method", req.Method,
				"path", req.Path,
				"status", req.StatusCode,
				"duration", req.Duration,
			)
		}
		return nil
	}})

	dispatcher.Listen("request.failed", &listenerFunc{fn: func(e interface{}) error {
		if req, ok := e.(*router.RequestFailed); ok {
			logger.Error("Request failed",
				"request_id", req.RequestID,
				"error", req.Error,
				"recovered", req.Recovered,
			)
		}
		return nil
	}})

	// Database query events
	dispatcher.Listen("query.executed", &listenerFunc{fn: func(e interface{}) error {
		if q, ok := e.(*orm.QueryExecuted); ok {
			logger.Debug("Query executed",
				"sql", q.SQL,
				"duration", q.Duration,
				"rows", q.RowsAffected,
			)
		}
		return nil
	}})

	// Cache events
	dispatcher.Listen("cache.hit", &listenerFunc{fn: func(e interface{}) error {
		if c, ok := e.(*cache.CacheHit); ok {
			logger.Debug("Cache hit", "key", c.Key)
		}
		return nil
	}})

	dispatcher.Listen("cache.miss", &listenerFunc{fn: func(e interface{}) error {
		if c, ok := e.(*cache.CacheMiss); ok {
			logger.Debug("Cache miss", "key", c.Key)
		}
		return nil
	}})
}
