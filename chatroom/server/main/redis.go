package main

import (
	"github.com/garyburd/redigo/redis"
	"time"
)

// 定义一个全局的Pool
var pool *redis.Pool

func initPool(address string, maxIdle, maxActive int, idleTimeout time.Duration) {
	pool = &redis.Pool{
		MaxIdle:     maxIdle,     // 最大空闲连接数
		MaxActive:   maxActive,   // 和数据库最大连接数，0表示无限制
		IdleTimeout: idleTimeout, //最大空闲时间
		Dial: func() (conn redis.Conn, err error) {
			return redis.Dial("tcp", address)
		},
	}
}
