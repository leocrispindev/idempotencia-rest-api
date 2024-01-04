package model

import "time"

type Message struct {
	Info            MessageInfo
	IdempotenciaKey string `bson:"idempotenciaKey"`
}

type MessageInfo struct {
	From    string    `bson:"from"`           // Remetente da mensagem
	To      string    `bson:"to"`             // Destinatário da mensagem
	Content string    `bson:"content"`        // Conteúdo da mensagem (pode ser texto, imagem, vídeo, etc.)
	Time    time.Time `bson:"time,omitempty"` // Hora do envio da mensagem
}
