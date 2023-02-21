package wamd

import (
	"database/sql"
	"go.mau.fi/whatsmeow"
	"time"
)

func AddSendMessage(jid string, content string, priority int) {
	var db, _ = sql.Open("sqlite", "./dbdata/deliveries.db")
	defer db.Close()
	now := time.Now().Format("2006-01-02 15:04:05")
	msgid := whatsmeow.GenerateMessageID()
	_, err := db.Exec("INSERT INTO deliveries VALUES (NULL,?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		jid, "text", msgid, content, "", "", "pending", priority, 0, 0, 0, "", "", "", now, now)
	if err != nil {
		return
	}
}
