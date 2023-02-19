package main

import (
	"database/sql"
	"fmt"
	"github.com/spf13/viper"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"
	"wuzapi/model"
	"wuzapi/service/Chatwoot"
)

type deliveries struct {
	id              int
	sid             string
	rel_id          int
	holder_id       int
	tag_id          int
	contact_id      int
	parent_id       int
	jid             string
	message_type    string
	message_id      string
	content         string
	attachment      string
	attachment_name string
	token           string
	delivery_status string
	read_status     int
	priority        int
	attempt         int
	is_schedule     int
	is_group        int
	commit_tag      int
	payload         string
	sent_after_unix int
	send_after      string
	read_at         string
	created_at      string
	updated_at      string
}

func worker() {
	db, err := sql.Open("sqlite", "./dbdata/deliveries.db")
	if err != nil {
		log.Fatal().Err(err).Msg("Could not open/create ./dbdata/deliveries.db")
		os.Exit(1)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Error().Msg(err.Error())
		}
	}(db)

	for {
		min := 1
		max := 5
		randSend := rand.Intn(max-min) + min
		fmt.Println()
		rows, err := db.Query("SELECT id, jid, message_type, content, attachment, attachment_name, is_group FROM deliveries WHERE delivery_status = 'pending' ORDER BY priority")
		if err != nil {
			return
		}
		defer rows.Close()

		var result []deliveries

		for rows.Next() {
			var each = deliveries{}
			var err = rows.Scan(&each.id, &each.jid, &each.message_type, &each.content, &each.attachment, &each.attachment_name, &each.is_group)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			result = append(result, each)
		}

		if err = rows.Err(); err != nil {
			fmt.Println(err.Error())
			return
		}
		now := time.Now()

		if len(result) == 0 {
			fmt.Println("No pending message " + now.Format("2006-01-02 15:04:05"))
		} else {
			fmt.Println("Pending message: ", len(result), now.Format("2006-01-02 15:04:05"))
			i := 0
			switch result[i].message_type {
			case "text":
				time.Sleep(time.Duration(randSend) * time.Second)
				kode, respond, msgid := SendMessageProcess(result[i].jid, result[i].content, 1, nil, nil)
				if viper.GetString("chatwoot.baseUrl") != "" && kode == http.StatusOK {
					isGroup := strings.Contains(result[i].jid, "@g.us")
					chat := result[i].jid
					if isGroup {
						chat = strings.Split(chat, "@")[0]
						chat = chat[3:]
					}
					var oneSenderWebhook model.OneSenderWebhook
					oneSenderWebhook.Chat = chat
					oneSenderWebhook.SenderPhone = chat
					oneSenderWebhook.MessageText = result[i].content
					oneSenderWebhook.MessageType = "text"
					oneSenderWebhook.IsFromMe = true
					oneSenderWebhook.IsGroup = isGroup
					oneSenderWebhook.MessageID = msgid
					Chatwoot.IncomingMessageApi(oneSenderWebhook)
				}
				if kode == 200 {
					now := time.Now()
					_, err = db.Exec("UPDATE deliveries SET message_id=?,delivery_status=?, updated_at=? WHERE id=?", msgid, "sent", now, result[i].id)
					if err != nil {
						fmt.Println(err.Error())
						return
					}
					log.Info().Msg("Message sent to " + result[i].jid)
					fmt.Println(respond)
				}
			}
		}
		time.Sleep(1 * time.Second)
	}
}

func receipt(id string, status string, timestamp time.Time) {
	fmt.Println("Receipt: ", id, status, timestamp)
	db, err := sql.Open("sqlite", "./dbdata/deliveries.db")
	if err != nil {
		log.Fatal().Err(err).Msg("Could not open/create ./dbdata/deliveries.db")
		os.Exit(1)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Error().Msg(err.Error())
		}
	}(db)
	_, err = db.Exec("UPDATE deliveries SET delivery_status=? WHERE message_id=?", status, id)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
