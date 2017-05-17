package dal

import (
	"database/sql"
	"strings"

	"github.com/damoye/gopu/config"
	_ "github.com/go-sql-driver/mysql" // MySQL
)

var db *sql.DB

func init() {
	var err error
	if db, err = sql.Open("mysql", config.Conf.DatabaseString); err != nil {
		panic(err)
	}
}

// InsertTask inserts a new task
func InsertTask(message string) (int64, error) {
	res, err := db.Exec("INSERT INTO task (data) VALUES (?)", message)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

// InsertSubtask inserts a new subtask
func InsertSubtask(taskID int64, tokens []string) error {
	if len(tokens) == 0 {
		return nil
	}
	sql := "INSERT INTO subtask (task_id, token) VALUES (?, ?)" +
		strings.Repeat(" ,(?, ?)", len(tokens)-1)
	args := make([]interface{}, 0, 2*len(tokens))
	for _, token := range tokens {
		args = append(args, taskID, token)
	}
	_, err := db.Exec(sql, args...)
	return err
}

// Deliver updates deliverTime of subtask
func Deliver(taskID int64, token string) error {
	_, err := db.Exec("UPDATE subtask SET is_delivered=TRUE WHERE task_id=(?) AND token=(?)", taskID, token)
	return err
}
