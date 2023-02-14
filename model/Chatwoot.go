package model

type ContactChatwoot struct {
	Payload struct {
		Contact struct {
			AvailabilityStatus string      `json:"availability_status,omitempty"`
			Email              string      `json:"email,omitempty"`
			ID                 int         `json:"id,omitempty"`
			Name               string      `json:"name,omitempty"`
			PhoneNumber        string      `json:"phone_number,omitempty"`
			Identifier         interface{} `json:"identifier,omitempty"`
			Thumbnail          string      `json:"thumbnail,omitempty"`
			CustomAttributes   struct {
			} `json:"custom_attributes,omitempty"`
		} `json:"contact,omitempty"`
	} `json:"payload,omitempty"`
}

type CreateContact struct {
	InboxID           string           `json:"inbox_id"`
	Name              string           `json:"name"`
	Email             string           `json:"email"`
	PhoneNumber       string           `json:"phone_number"`
	ContactAttributes CustomAttributes `json:"custom_attributes"`
}

type SearchContactChatwoot struct {
	Meta struct {
		Count       int `json:"count"`
		CurrentPage int `json:"current_page"`
	} `json:"meta"`
	Payload []struct {
		AdditionalAttributes struct {
		} `json:"additional_attributes"`
		Email       interface{} `json:"email"`
		ID          int         `json:"id"`
		Name        string      `json:"name"`
		PhoneNumber string      `json:"phone_number"`
	} `json:"payload"`
}

type CreateConversation struct {
	InboxID   string `json:"inbox_id"`
	ContactID int    `json:"contact_id"`
	Status    string `json:"status"`
}

type ConversationChatwoot struct {
	ID      int `json:"id"`
	InboxID int `json:"inbox_id"`
}

type SearchConversationChatwoot struct {
	Payload []struct {
		ID      int `json:"id"`
		InboxID int `json:"inbox_id"`
	} `json:"payload"`
}

type SendMessage struct {
	Content     string `json:"content"`
	MessageType string `json:"message_type"`
	Private     bool   `json:"private"`
}

type SendMessageResult struct {
	ID             int         `json:"id"`
	Content        string      `json:"content"`
	InboxID        int         `json:"inbox_id"`
	ConversationID int         `json:"conversation_id"`
	MessageType    int         `json:"message_type"`
	ContentType    string      `json:"content_type"`
	Status         string      `json:"status"`
	CreatedAt      int         `json:"created_at"`
	Private        bool        `json:"private"`
	SourceID       interface{} `json:"source_id"`
	Sender         struct {
		ID                 int    `json:"id"`
		Name               string `json:"name"`
		AvailableName      string `json:"available_name"`
		AvatarURL          string `json:"avatar_url"`
		Type               string `json:"type"`
		AvailabilityStatus string `json:"availability_status"`
		Thumbnail          string `json:"thumbnail"`
		ContentAttributes  struct {
		} `json:"content_attributes"`
	} `json:"sender"`
}

type ChatwootWebhook struct {
	Account              Account              `json:"account"`
	AdditionalAttributes AdditionalAttributes `json:"additional_attributes"`
	ContentAttributes    AdditionalAttributes `json:"content_attributes"`
	ContentType          string               `json:"content_type"`
	Content              string               `json:"content"`
	Conversation         Conversation         `json:"conversation"`
	CreatedAt            string               `json:"created_at"`
	ID                   int64                `json:"id"`
	Inbox                Account              `json:"inbox"`
	MessageType          string               `json:"message_type"`
	Private              bool                 `json:"private"`
	Sender               Sender               `json:"sender"`
	SourceID             interface{}          `json:"source_id"`
	Event                string               `json:"event"`
}

type AdditionalAttributes struct {
}

type Conversation struct {
	AdditionalAttributes AdditionalAttributes `json:"additional_attributes"`
	CanReply             bool                 `json:"can_reply"`
	Channel              string               `json:"channel"`
	ContactInbox         ContactInbox         `json:"contact_inbox"`
	ID                   int64                `json:"id"`
	InboxID              int64                `json:"inbox_id"`
	Messages             []Message            `json:"messages"`
	Labels               []interface{}        `json:"labels"`
	Meta                 Meta                 `json:"meta"`
	Status               string               `json:"status"`
	CustomAttributes     AdditionalAttributes `json:"custom_attributes"`
	SnoozedUntil         interface{}          `json:"snoozed_until"`
	UnreadCount          int64                `json:"unread_count"`
	FirstReplyCreatedAt  interface{}          `json:"first_reply_created_at"`
	AgentLastSeenAt      int64                `json:"agent_last_seen_at"`
	ContactLastSeenAt    int64                `json:"contact_last_seen_at"`
	Timestamp            int64                `json:"timestamp"`
}

type CustomAttributes struct {
	GroupID string `json:"group_id"`
	IsGroup bool   `json:"is_group"`
}

type Sender struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Type  string `json:"type"`
}

type Account struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type Attachment struct {
	ID        int64       `json:"id"`
	MessageID int64       `json:"message_id"`
	FileType  string      `json:"file_type"`
	AccountID int64       `json:"account_id"`
	Extension interface{} `json:"extension"`
	DataURL   string      `json:"data_url"`
	ThumbURL  string      `json:"thumb_url"`
	FileSize  int64       `json:"file_size"`
}

type ContactInbox struct {
	ID           int64  `json:"id"`
	ContactID    int64  `json:"contact_id"`
	InboxID      int64  `json:"inbox_id"`
	SourceID     string `json:"source_id"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
	HmacVerified bool   `json:"hmac_verified"`
	PubsubToken  string `json:"pubsub_token"`
}

type Message struct {
	ID             int64               `json:"id"`
	Content        string              `json:"content"`
	AccountID      int64               `json:"account_id"`
	InboxID        int64               `json:"inbox_id"`
	ConversationID int64               `json:"conversation_id"`
	MessageType    int64               `json:"message_type"`
	CreatedAt      int64               `json:"created_at"`
	UpdatedAt      string              `json:"updated_at"`
	Private        bool                `json:"private"`
	Status         string              `json:"status"`
	SourceID       interface{}         `json:"source_id"`
	ContentType    string              `json:"content_type"`
	SenderType     string              `json:"sender_type"`
	SenderID       int64               `json:"sender_id"`
	LabelList      interface{}         `json:"label_list"`
	Conversation   MessageConversation `json:"conversation"`
	Attachments    []Attachment        `json:"attachments"`
	Sender         Assignee            `json:"sender"`
}

type MessageConversation struct {
	AssigneeID  int64 `json:"assignee_id"`
	UnreadCount int64 `json:"unread_count"`
}

type Assignee struct {
	ID                 int64       `json:"id"`
	Name               string      `json:"name"`
	AvailableName      string      `json:"available_name"`
	AvatarURL          string      `json:"avatar_url"`
	Type               string      `json:"type"`
	AvailabilityStatus interface{} `json:"availability_status"`
	Thumbnail          string      `json:"thumbnail"`
}

type Meta struct {
	Sender       MetaSender  `json:"sender"`
	Assignee     Assignee    `json:"assignee"`
	Team         interface{} `json:"team"`
	HmacVerified bool        `json:"hmac_verified"`
}

type MetaSender struct {
	AdditionalAttributes AdditionalAttributes `json:"additional_attributes"`
	CustomAttributes     CustomAttributes     `json:"custom_attributes"`
	Email                string               `json:"email"`
	ID                   int64                `json:"id"`
	Identifier           interface{}          `json:"identifier"`
	Name                 string               `json:"name"`
	PhoneNumber          string               `json:"phone_number"`
	Thumbnail            string               `json:"thumbnail"`
	Type                 string               `json:"type"`
}
