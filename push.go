package main

import (
	"database/sql"
	"errors"
	"strconv"
	"time"

	"gopkg.in/redis.v5"
)

var redisClient = redis.NewClient(&redis.Options{Addr: config.RedisAddress})

type pushRequest struct {
	Message    string
	ClientKeys []string
	TTL        time.Time
}

func push(req pushRequest) error {
	if len(req.ClientKeys) == 0 {
		return errors.New("ClientKeys is empty")
	}
	// database
	db, err := sql.Open("mysql", config.DatabaseString)
	if err != nil {
		return err
	}
	defer db.Close()
	res, err := db.Exec(`INSERT INTO Task (Message) VALUES (?)`, req.Message)
	if err != nil {
		return err
	}
	taskID, err := res.LastInsertId()
	if err != nil {
		return err
	}
	sql := "INSERT INTO Subtask (Client, TaskID, TTL) VALUES"
	args := make([]interface{}, 0, 3*len(req.ClientKeys))
	for _, item := range req.ClientKeys {
		sql += " (?, ?, ?),"
		args = append(args, item, taskID, req.TTL)
	}
	sql = sql[0 : len(sql)-1]
	if _, err := db.Exec(sql, args...); err != nil {
		return err
	}
	// redis
	taskKey := strconv.FormatInt(taskID, 36)
	pipe := redisClient.Pipeline()
	pipe.Set(getRedisMessageKey(taskKey), req.Message, 30*time.Second)
	for _, item := range req.ClientKeys {
		pipe.Publish(getRedisClientKey(item), taskKey)
	}
	results, err := pipe.Exec()
	if err != nil {
		return err
	}
	for _, item := range results {
		if item.Err() != nil {
			return err
		}
	}
	return nil
}
