package wamd

import (
	"database/sql"
	"time"
)

func AddSendMessageText(jid string, content string, priority int) {
	var db, _ = sql.Open("sqlite", "./dbdata/deliveries.db")
	defer db.Close()
	now := time.Now()
	_, err := db.Exec("INSERT INTO deliveries VALUES (NULL,?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		jid, "text", "", content, "", "", "pending", priority, 0, 0, 0, "", "", "", now, now)
	if err != nil {
		return
	}
}
