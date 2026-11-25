package store

import (
	"context"
	"database/sql"
	"time"
)

type ChatStore struct {
	db *sql.DB
}

func NewChatStore(db *sql.DB) *ChatStore { return &ChatStore{db: db} }

type Room struct {
	ID          int64
	Name        string
	Description string
	AgeLimit    int
	MemberLimit int
	IsDM        bool
	CreatedAt   time.Time
}
type Message struct {
	ID        int64
	RoomID    int64
	SenderID  int64
	Body      string
	CreatedAt time.Time
}

func (c *ChatStore) CreateRoom(ctx context.Context, room Room) (int64, error) {
	row := c.db.QueryRowContext(ctx, `insert into rooms (name, description, age_limit, member_limit, is_dm) values($1, $2, $3, $4, $5) RETURNING id`, room.Name, room.Description, room.AgeLimit, room.MemberLimit, room.IsDM)

	var id int64
	err := row.Scan(&id)
	if err != nil {
		return -1, err
	}
	return id, nil
}

func (c *ChatStore) JoinRoom(ctx context.Context, roomID int64, userID int64, isadmin bool) error {
	_, err := c.db.ExecContext(ctx, `insert into room_members (room_id, user_id, is_admin) values($1, $2, $3)`, roomID, userID, isadmin)
	return err
}

func (c *ChatStore) SendMessageToRoom(ctx context.Context, message Message) (int64, error) {
	row := c.db.QueryRowContext(ctx, `insert into messages (room_id, sender_id, body) values($1, $2, $3) RETURNING id`, message.RoomID, message.SenderID, message.Body)

	var id int64
	err := row.Scan(&id)
	if err != nil {
		return -1, err
	}
	return id, nil
}

func (c *ChatStore) GetRoomHistory(ctx context.Context, roomID int64, limit int64) ([]Message, error) {
	rows, err := c.db.QueryContext(ctx, `select id, room_id, sender_id, body, created_at
 from messages where room_id = $1 order by created_at desc, id desc limit $2`, roomID, limit)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var msgs []Message
	for rows.Next() {
		var m Message
		if err := rows.Scan(&m.ID, &m.RoomID, &m.SenderID, &m.Body, &m.CreatedAt); err != nil {
			return nil, err
		}
		msgs = append(msgs, m)
	}
	if err := rows.Err(); err != nil { // errors that ended the loop early
		return nil, err
	}
	return msgs, nil
}
