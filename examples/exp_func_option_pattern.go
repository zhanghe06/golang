package main

import "fmt"

// Go 函数选项模式（Functional Options Pattern）

type Message struct {
	// 标题、内容、信息类型
	title, message, messageType string

	// 账号
	account     string
	accountList []string

	// token
	token     string
	tokenList []string
}

type MessageOption func(*Message)

func NewMessage(title, message, messageType string, opts ...MessageOption) *Message {
	msg := &Message{
		title:       title,
		message:     message,
		messageType: messageType,
	}

	for _, opt := range opts {
		opt(msg)
	}

	fmt.Println(msg)
	return msg
}

func WithAccount(account string) MessageOption {
	return func(message *Message) {
		message.account = account
	}
}

func WithAccountList(accountList []string) MessageOption {
	return func(message *Message) {
		message.accountList = accountList
	}
}

func WithToken(token string) MessageOption {
	return func(message *Message) {
		message.token = token
	}
}

func WithTokenList(tokenList []string) MessageOption {
	return func(message *Message) {
		message.tokenList = tokenList
	}
}

func main() {
	// 单账号推送
	_ = NewMessage(
		"来自Golang的信息",
		"金榜题名",
		"单账号推送",
		WithAccount("123456"),
	)

	// 多账号推送
	_ = NewMessage(
		"来自Golang的信息",
		"金榜题名",
		"多账号推送",
		WithAccountList([]string{"123456", "654321"}),
	)
}
