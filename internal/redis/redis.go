package redis

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gomodule/redigo/redis"
	"google.golang.org/appengine"
)

var redisPool *redis.Pool

func initializeRedis() (*redis.Pool, error) {
	redisAddr := os.Getenv("REDISADDR")
	redisPassword := os.Getenv("REDIS_PASSWORD")
	const maxConnections = os.Getenv("REDIS_MAX_CONN")

	redisPool = &redis.Pool{
		maxIdle: maxConnections,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", redisAddr)
			if redisPassword == "" {
				return c, err
			}
			if err != nil {
				return nil, err
			}
			if _, err := c.Do("AUTH", redisPassword); err != nil {
				c.Close()
				return nil, err
			}
			return c, nil
		},
	}

	http.HandleFunc("/", handle)
	appengine.Main()
}

func handle(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	redisConn := redisPool.Get()
	defer redisConn.Close()

	count, err := redisConn.Do("INCR", "count")
	if err != nil {
		msg := fmt.Sprintf("cloud not increment count : %v", err)
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Count: %d", count)
}