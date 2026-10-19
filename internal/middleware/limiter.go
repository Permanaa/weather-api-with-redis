package middleware

import (
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/permanaa/weather-api-with-redis/internal/helper"
	"golang.org/x/time/rate"
)

type ClientLimiter struct {
	limiter  *rate.Limiter
	lastSeen time.Time
}

type IPRateLimiter struct {
	mu      sync.Mutex
	clients map[string]*ClientLimiter
}

func NewIPRateLimiter() *IPRateLimiter {
	return &IPRateLimiter{
		clients: make(map[string]*ClientLimiter),
	}
}

func (i *IPRateLimiter) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip := getClientIp(r)
		limiter := i.GetLimiter(ip)

		if !limiter.Allow() {
			w.Header().Set("Retry-After", "12")
			helper.JsonError(w, http.StatusTooManyRequests, "rate limit exceeded")
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (i *IPRateLimiter) GetLimiter(ip string) *rate.Limiter {
	i.mu.Lock()
	defer i.mu.Unlock()

	client, exists := i.clients[ip]
	// add client to IPRateLimiter.clients list if client does not exist
	if !exists {
		client = &ClientLimiter{
			limiter: rate.NewLimiter(rate.Every(12*time.Second), 5),
		}

		i.clients[ip] = client
	}

	client.lastSeen = time.Now()

	return client.limiter
}

func getClientIp(r *http.Request) string {
	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr
	}
	return host
}
