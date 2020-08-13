package main

import (
	"errors"
	// "io/ioutil"
	"net/url"
	"reflect"
	"strconv"

	// "strconv"
	"strings"

	// "bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

// Bot struct
// This represents a Bot object that'll make api calls to the telegram api
type Bot struct {
	Name     string
	Token    string
	Client   *http.Client
	APIURL   string
	endpoint string
}

var bot *Bot

func init() {
	minit()
	fmt.Println("in init bot.go")
	bot, err = NewBot(token)
	if err != nil {
		log.Fatal("Error creating bot ", err)
	}
	fmt.Println("Bot created:", bot.Name)
}

// WebhookResp this struct holds the data from the HandleRequest api endpoint
type WebhookResp struct {
	OK          bool        `json:"ok"`
	Result      interface{} `json:"result"`
	ErrorCode   int         `json:"error_code"`
	Description string      `json:"description"`
}

// NewBot creates a new bot instance
func NewBot(token string) (*Bot, error) {
	bot := &Bot{
		Name:     "yugolee",
		Token:    token,
		Client:   &http.Client{},
		APIURL:   "https://api.telegram.org/",
		endpoint: "bot%s/",
	}
	bot.endpoint = fmt.Sprintf(bot.endpoint, bot.Token)
	// fmt.Println("bot.Token", bot.Token)
	// fmt.Println("bot.endpoint", bot.endpoint)
	if err := bot.Ready(); err != nil {
		return bot, err
	}

	return bot, nil
}

// Ready method sets the webhook on telegram
func (bot *Bot) Ready() error {
	URL := os.Getenv("WEBHOOKURL")
	if URL == "" {
		return errors.New("environment variable WEBHOOKURL not set")
	}
	data := url.Values{}
	data.Set("url", URL)
	// fmt.Println("func Ready: data:", data)

	urlStr, err := bot.makeURL(SetWebhook)
	if err != nil {
		return err
	}

	resp, err := bot.Request(http.MethodPost, urlStr, data)
	if err != nil {
		log.Println()
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return errors.New("StatusCode not Ok")
	}

	var hookResp WebhookResp
	if err := json.NewDecoder(resp.Body).Decode(&hookResp); err != nil {
		return err
	}

	if hookResp.Result == false {
		return errors.New("Webhook result not Ok")
	}

	return nil
}

// makeURL
func (bot *Bot) makeURL(telemethod string) (string, error) {
	method := bot.endpoint + telemethod

	u, err := url.ParseRequestURI(bot.APIURL)
	if err != nil {
		return "", err
	}
	u.Path = method
	urlStr := u.String()
	return urlStr, nil
}

func (bot *Bot) Request(method, url string, data url.Values) (*http.Response, error) {
	r, _ := http.NewRequest(method, url, strings.NewReader(data.Encode()))     // URL-encoded payload
	r.Header.Add("Authorization", fmt.Sprintf("auth_token=\"%s\"", bot.Token)) // htis line not really needed
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")          // remove this line to delete webhook
	// r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	resp, err := bot.Client.Do(r)
	if err != nil {
		log.Println()
		return nil, err
	}
	return resp, nil
}

func (bot *Bot) SendMessage(chatID int, text string) (*Message, error) {
	msgToSend := &SendMessageBody{
		ChatID:      chatID,
		Text:        text,
		ReplyMarkup: nil,
	}
	data := structToMap(msgToSend)

	// fmt.Println(data)

	urlStr, err := bot.makeURL(SendMessage)
	if err != nil {
		return nil, err
	}

	resp, err := bot.Request(http.MethodPost, urlStr, data)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("StatusCode not Ok")
	}

	msgReceived := &Message{}
	err = json.NewDecoder(resp.Body).Decode(msgReceived)
	if err != nil {
		return nil, err
	}
	return msgReceived, nil
}

func structToMap(inter interface{}) (values url.Values) {
	values = url.Values{}
	iVal := reflect.ValueOf(inter).Elem()
	typ := iVal.Type()
	for i := 0; i < iVal.NumField(); i++ {

		f := iVal.Field(i)

		var v string
		switch f.Interface().(type) {
		case int, int8, int16, int32, int64:
			v = strconv.FormatInt(f.Int(), 10)
			if v == "" {
				continue
			}
		case uint, uint8, uint16, uint32, uint64:
			v = strconv.FormatUint(f.Uint(), 10)
			if v == "" {
				continue
			}
		case float32:
			v = strconv.FormatFloat(f.Float(), 'f', 4, 32)
			if v == "" {
				continue
			}
		case float64:
			v = strconv.FormatFloat(f.Float(), 'f', 4, 64)
			if v == "" {
				continue
			}
		case []byte:
			v = string(f.Bytes())
			if v == "" {
				continue
			}
		case string:
			v = f.String()
			if v == "" {
				continue
			}
		default:
			continue
		}
		values.Set(typ.Field(i).Tag.Get("json"), v)

	}
	return
}

func (bot *Bot) SendLocation(chatID int, latitude, longitude float32) (interface{}, error) {
	msgToSend := &SendLocationBody{
		Latitude:  latitude,
		Longitude: longitude,
		ChatID:    chatID,
	}
	data := structToMap(msgToSend)

	urlStr, err := bot.makeURL(SendLocation)
	if err != nil {
		return nil, err
	}

	resp, err := bot.Request(http.MethodPost, urlStr, data)
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("StatusCode not Ok")
	}

	var hookResp WebhookResp
	if err := json.NewDecoder(resp.Body).Decode(&hookResp); err != nil {
		return nil, err
	}
	fmt.Println(hookResp.Result)
	if hookResp.Result == false {
		return nil, errors.New("Webhook result not Ok")
	}
	body, err := json.Marshal(hookResp.Result)
	if err != nil {
		return nil, errors.New("i told you ")
	}

	var msgReceived interface{}
	// body, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, msgReceived)
	// err = json.NewDecoder(body).Decode(msgReceived)
	if err != nil {
		return nil, errors.New("dfd")
	}
	return nil, nil
}

func (bot *Bot) ForwardMessages(chatID, fromChatID, MessageID int) (*Message, error) {
	msgToSend := &ForwardMessageBody{
		ChatID:     chatID,
		FromChatID: fromChatID,
		MessageID:  MessageID,
	}
	data := structToMap(msgToSend)

	urlStr, err := bot.makeURL(ForwardMessage)
	if err != nil {
		return nil, err
	}

	resp, err := bot.Request(http.MethodPost, urlStr, data)
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("StatusCode not Ok")
	}

	var hookResp WebhookResp
	if err := json.NewDecoder(resp.Body).Decode(&hookResp); err != nil {
		return nil, err
	}

	if hookResp.Result == false {
		return nil, errors.New("Webhook result not Ok")
	}

	msgReceived := &Message{}
	err = json.NewDecoder(resp.Body).Decode(msgReceived)
	if err != nil {
		return nil, err
	}
	return msgReceived, nil
}

type getCommandResults struct {
	OK     bool         `json:"ok"`
	Result []BotCommand `json:"result"`
}

// GetMyCommands this was used to see the setcommands for the bot
func (bot *Bot) GetMyCommands() (interface{}, error) {
	urlStr, err := bot.makeURL(GetMyCommands)
	if err != nil {
		return nil, err
	}

	resp, err := bot.Request(http.MethodPost, urlStr, nil)
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("StatusCode not Ok")
	}
	fmt.Println(resp.Body)

	// var hookResp WebhookResp
	// if err := json.NewDecoder(resp.Body).Decode(&hookResp); err != nil {
	// 	return nil, err

	var arr interface{}
	err = json.NewDecoder(resp.Body).Decode(&arr)
	if err != nil {
		return nil, err
	}
	return arr, nil

}

// SetMyCommands this was used to setup the commands for the bot
func (bot *Bot) SetMyCommands(commands []BotCommand) error {
	var msgToSend []BotCommand
	msgToSend = append(msgToSend, commands...)

	reqbytes, err := json.Marshal(msgToSend)
	if err != nil {
		return err
	}

	data := url.Values{}
	data.Set("commands", string(reqbytes))

	urlStr, err := bot.makeURL(SetMyCommands)
	if err != nil {
		return err
	}

	resp, err := bot.Request(http.MethodPost, urlStr, data)
	if resp.StatusCode != http.StatusOK {
		return errors.New("StatusCode not Ok")
	}

	var hookResp WebhookResp
	if err := json.NewDecoder(resp.Body).Decode(&hookResp); err != nil {
		return err
	}

	if hookResp.Result == false || hookResp.OK == false {
		return errors.New("Webhook result not Ok")
	}

	return nil

}

// Run this is hopefully to make the bot run
func (bot *Bot) Run() {
	http.HandleFunc("/", HandleRequests)
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal("port already taken")
	}
	fmt.Println("Running")
}

// setupCommands this was used to setup the
func setupCommands(bot Bot) error {
	var commands []BotCommand

	startCommand := BotCommand{
		Command:     "start",
		Description: "Initiates conversation with yugoyomi",
	}
	sayHelloCommand := BotCommand{
		Command:     "sayhello",
		Description: "Greets you Hello",
	}
	sendLocationCommand := BotCommand{
		Command:     "sendlocation",
		Description: "Send my location to you",
	}
	playGameCommand := BotCommand{
		Command:     "playgame",
		Description: "We play a game to see if you are smarter than me.",
	}

	commands = append(commands, []BotCommand{startCommand, sayHelloCommand, sendLocationCommand, playGameCommand}...)
	err := bot.SetMyCommands(commands)
	if err != nil {
		return err
	}
	bot.Ready()
	return nil
}

// TODO: make it work asynchronously
// TODO: make the existing commands work better
// TODO: add more commands 
// TODO: refactor
// TODO: probably add a game
// TODO: implement a small db to store users session (not sure how this would work but it is most probably for the game)
