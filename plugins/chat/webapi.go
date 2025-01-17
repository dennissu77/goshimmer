package chat

import (
	"net/http"

	"github.com/labstack/echo"

	"github.com/iotaledger/goshimmer/packages/jsonmodels"
	"github.com/iotaledger/goshimmer/plugins/messagelayer"
	"github.com/iotaledger/goshimmer/plugins/webapi"
)

func configureWebAPI() {
	webapi.Server().POST("chat", SendChatMessage)
}

// SendChatMessage sends a chat message.
func SendChatMessage(c echo.Context) error {
	req := &Request{}
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, jsonmodels.NewErrorResponse(err))
	}
	chatPayload := NewPayload(req.From, req.To, req.Message)

	msg, err := messagelayer.Tangle().IssuePayload(chatPayload)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{Error: err.Error()})
	}

	return c.JSON(http.StatusOK, Response{MessageID: msg.ID().Base58()})
}

// Request defines the chat message to send
type Request struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Message string `json:"message"`
}

// Response contains the ID of the message sent.
type Response struct {
	MessageID string `json:"messageID,omitempty"`
	Error     string `json:"error,omitempty"`
}
