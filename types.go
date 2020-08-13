package main

// User struct
// This object represents a Telegram user or bot.
type User struct {
	ID                      int    `json:"id"`
	IsBot                   bool   `json:"is_bot"`
	FirstName               string `json:"first_name"`
	LastName                string `json:"last_name"`
	Username                string `json:"username"`
	LanguageCode            string `json:"language_code"`
	CanJoinGroups           bool   `json:"can_join_groups"`
	CanReadAllGroupMessages bool   `json:"can_read_all_group_messages"`
	SupportsInlineQueries   bool   `json:"supports_inline_queries"`
}

// Chat struct
// This object represents a chat.
type Chat struct {
	ID               int              `json:"id"`
	Type             string           `json:"type"`
	Title            string           `json:"title"`
	Username         string           `json:"username"`
	FirstName        string           `json:"first_name"`
	LastName         string           `json:"last_name"`
	Photo            *ChatPhoto       `json:"photo"`
	Description      string           `json:"description"`
	InviteLink       string           `json:"invite_link"`
	PinnedMessage    *Message         `json:"pinned_message"`
	Permissions      *ChatPermissions `json:"permissions"`
	ShowModeDelay    int              `json:"show_mode_delay"`
	StickerSetName   int              `json:"sticker_set_name"`
	CanSetStickerSet int              `json:"can_set_sticker_set"`
}

type Message struct {
	MessageID             int                   `json:"message_id"`
	From                  *User                 `json:"from"`
	Date                  int                   `json:"date"`
	Chat                  *Chat                 `json:"chat"`
	ForwardFrom           *User                 `json:"forward_from"`
	ForwardFromChat       *Chat                 `json:"forward_from_chat"`
	ForwardFromMessageID  int                   `json:"forward_from_message_id"`
	ForwardSenderName     int                   `json:"forward_sender_name"`
	ForwardSigature       int                   `json:"forward_signature"`
	ForwardDate           int                   `json:"forward_date"`
	ReplyToMessage        *Message              `json:"reply_to_message"`
	EditDate              *Message              `json:"edit_date"`
	MediaGroupID          *Message              `json:"media_group_id"`
	AuthorSignature       *Message              `json:"author_signature"`
	Text                  string                `json:"text"`
	Entities              []MessageEntity       `json:"entities"`
	CaptionEntities       []MessageEntity       `json:"caption_entiies"`
	Audio                 *Audio                `json:"audio"`
	Document              *Document             `json:"document"`
	Animation             *Animation            `json:"animation"`
	Game                  *Game                 `json:"game"`
	Photo                 []PhotoSize           `json:"photo"`
	Sticker               *Sticker              `json:"sticker"`
	Video                 *Video                `json:"video"`
	Voice                 *Voice                `json:"voice"`
	VideoNote             *VideoNote            `json:"vidoe_note"`
	Caption               string                `json:"caption"`
	Contact               *Contact              `json:"contact"`
	Loacation             *Location             `json:"location"`
	Venue                 *Venue                `json:"venue"`
	Poll                  *Poll                 `json:"poll"`
	Dice                  *Dice                 `json:"dice"`
	NewChatMemebers       []User                `json:"new_chat_member"`
	LeftChatMember        User                  `json:"left_chat_member"`
	NewChatTitle          string                `json:"new_chat_title"`
	NewChatPhoto          []PhotoSize           `json:"new_chat_photo"`
	DeleteChatPhoto       bool                  `json:"delete_chat_photo"`
	GroupChatCreated      bool                  `json:"group_chat_created"`
	SuperGroupChatCreated bool                  `json:"supergroup_chat_created"`
	ChannelChatCreated    bool                  `json:"channel_chat_created"`
	MigrateToChatID       int                   `json:"migrate_to_chat_id"`
	MigrageFromChatID     int                   `json:"migrate_from_chat_id"`
	PinnedMessage         *Message              `json:"pinned_message"`
	Invoice               *Invoice              `json:"invoice"`
	SuccessfulPayment     *SuccessfulPayment    `json:"successful_payment"`
	ConnectedWebsite      string                `json:"connected_website"`
	PassportData          *PassportData         `json:"passport_data"`
	ReplyMarkup           *InlineKeyboardMarkup `json:"reply_markup"`
}

type MessageEntity struct {
	Type     string `json:"type"`
	Offset   int    `json:"offset"`
	Length   int    `json:"length"`
	URL      string `json:"url"`
	User     User   `json:"user"`
	Language string `json:"language"`
}

// PhotoSize This object represents one size of a photo or a file / sticker thumbnail.
type PhotoSize struct {
	FileID       string `json:"file_id"`
	FileUniqueID string `json:"file_unique_id"`
	Width        int    `json:"width"`
	Height       int    `json:"height"`
	FileSize     int    `json:"file_size"`
}

// Audio This object represents an audio file to be treated as music by the Telegram clients.
type Audio struct {
	FileID       string     `json:"file_id"`
	FileUniqueID string     `json:"file_unique_id"`
	Duration     int        `json:"duration"`
	Performer    int        `json:"performer"`
	Title        string     `json:"title"`
	MimeType     string     `json:"mime_type"`
	FileSize     int        `json:"file_size"`
	Thumb        *PhotoSize `json:"thumb"`
}

// Document This object represents a general file (as opposed to photos, voice messages and audio files).
type Document struct {
	FileID       string     `json:"file_id"`
	FileUniqueID string     `json:"file_unique_id"`
	Thumb        *PhotoSize `json:"thumb"`
	FileName     string     `json:"file_name"`
	MimeType     string     `json:"mime_type"`
	FileSize     int        `json:"file_size"`
}

// Video This object represents a video file.
type Video struct {
	FileID       string     `json:"file_id"`
	FileUniqueID string     `json:"file_unique_id"`
	Width        int        `json:"width"`
	Height       int        `json:"height"`
	Duration     int        `json:"duration"`
	Thumb        *PhotoSize `json:"thumb"`
	MimeType     string     `json:"mime_type"`
	FileSize     int        `json:"file_size"`
}

// Animation This object represents an animation file (GIF or H.264/MPEG-4 AVC video without sound)
type Animation struct {
	FileID       string     `json:"file_id"`
	FileUniqueID string     `json:"file_unique_id"`
	Width        int        `json:"width"`
	Height       int        `json:"height"`
	Duration     int        `json:"duration"`
	Thumb        *PhotoSize `json:"thumb"`
	FileName     string     `json:"file_name"`
	MimeType     string     `json:"mime_type"`
	FileSize     int        `json:"file_size"`
}

// Voice This object represents a voice note.
type Voice struct {
	FileID       string `json:"file_id"`
	FileUniqueID string `json:"file_unique_id"`
	Duration     int    `json:"duration"`
	MimeType     string `json:"mime_type"`
	FileSize     int    `json:"file_size"`
}

// VideoNote This object represents a video message (available in Telegram apps as of v.4.0).
type VideoNote struct {
	FileID       string     `json:"file_id"`
	FileUniqueID string     `json:"file_unique_id"`
	Length       int        `json:"length"`
	Duration     int        `json:"duration"`
	Thumb        *PhotoSize `json:"thumb"`
	FileSize     int        `json:"file_size"`
}

// Contact This object represents a phone contact.
type Contact struct {
	PhoneNumber string `json:"phone_number"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	UserID      int    `json:"user_id"`
	VCard       string `json:"vcard"`
}

// Location his object represents a point on the map.
type Location struct {
	Longitude float32 `json:"longitude"`
	Latitude  float32 `json:"latitude"`
}

// Venue This object represents a venue.
type Venue struct {
	Location       *Location `json:"location"`
	Title          string    `json:"title"`
	Address        string    `json:"address"`
	FourSquareID   string    `json:"foursquare_id"`
	FourSquareType string    `json:"foursquare_type"`
}

// This object contains information about one answer option in a poll.
type PollOption struct {
	Text       string `json:"text"`
	VoterCount int    `json:"voter_count"`
}

// PollAnswer This object represents an answer of a user in a non-anonymous poll.
type PollAnswer struct {
	PollID    string `json:"poll_id"`
	User      *User  `json:"user"`
	OptionIDs []int  `json:"option_ids"`
}

// Poll This object contains information about a poll.
type Poll struct {
	ID                    string          `json:"id"`
	Question              string          `json:"question"`
	Options               []PollOption    `json:"options"`
	TotalVoterCount       int             `json:"total_voter_count"`
	IsClosed              bool            `json:"is_closed"`
	IsAnonymous           bool            `json:"is_anonymous"`
	Type                  string          `json:"type"`
	AllowsMultipleAnswers bool            `json:"allows_multiple_answers"`
	CorrectOPtionID       int             `json:"correct_option_id"`
	Explanation           string          `json:"explanation"`
	ExplanationEntities   []MessageEntity `json:"explanation_entities"`
	OpenPeriod            int             `json:"open_period"`
	CloseDate             int             `json:"close_date"`
}

// Dice This object represents a dice with a random value from 1 to 6 for currently supported base emoji. (Yes, we're aware of the “proper” singular of die. But it's awkward, and we decided to help it change. One dice at a time!)
type Dice struct {
	Emoji string `json:"emoji"`
	Value int    `json:"value"`
}

// UserProfilePhotos This object represent a user's profile pictures.
type UserProfilePhotos struct {
	TotalCount int         `json:"total_count"`
	Photos     []PhotoSize `json:"photos"`
}

// File This object represents a file ready to be downloaded. The file can be downloaded via the link https://api.telegram.org/file/bot<token>/<file_path>. It is guaranteed that the link will be valid for at least 1 hour. When the link expires, a new one can be requested by calling getFile.
// Maximum file size to download is 20 MB
type File struct {
	FileID       string `json:"file_id"`
	FileUniqueID string `json:"file_unique_id"`
	FileSize     int    `json:"file_size"`
	FilePath     string `json:"file_path"`
}

// ReplyKeyboardMarkup o
// This object represents a custom keyboard with reply options (see Introduction to bots for details and examples).
type ReplyKeyboardMarkup struct {
	Keyboard        []KeyboardButton `json:"keyboard"`
	ResizeKeyboard  bool             `json:"resize_keyboard"`
	OneTimeKeyboard bool             `json:"one_time_keyboard"`
	Selective       bool             `json:"selective"`
}

// KeyboardButton struct
// This object represents one button of the reply keyboard. For simple text buttons String can be used instead of this object to specify text of the button. Optional fields request_contact, request_location, and request_poll are mutually exclusive.
type KeyboardButton struct {
	Text            bool                    `json:"selective"`
	RequestContact  bool                    `json:"request_contact"`
	RequestLocation bool                    `json:"request_location"`
	RequestPoll     *KeyboardButtonPollType `json:"request_poll"`
}

// KeyboardButtonPollType struct
// This object represents type of a poll, which is allowed to be created and sent when the corresponding button is pressed.
type KeyboardButtonPollType struct {
	Type string `json:"type"`
}

// ReplyKeyboardRemove struct
// Upon receiving a message with this object, Telegram clients will remove the current custom keyboard and display the default letter-keyboard. By default, custom keyboards are displayed until a new keyboard is sent by a bot. An exception is made for one-time keyboards that are hidden immediately after the user presses a button (see ReplyKeyboardMarkup).
type ReplyKeyboardRemove struct {
	RemoveKeyboard bool `json:"remove_keyboard"`
	Selective      bool `json:"selective"`
}

// InlineKeyboardMarkup struct
// This object represents an inline keyboard that appears right next to the message it belongs to.
type InlineKeyboardMarkup struct {
	InlineKeyboard []InlineKeyboardButton `json:"inline_keyboard"`
}

// InlineKeyboardButton struct
// This object represents one button of an inline keyboard. You must use exactly one of the optional fields.
type InlineKeyboardButton struct {
	Text                         string        `json:"text"`
	URL                          string        `json:"url"`
	LoginURL                     *LoginURL     `json:"login_url"`
	CallbackData                 string        `json:"callback_data"`
	SwitchInlineQuery            string        `json:"switch_inline_query"`
	SwitchInlineQueryCurrentChat string        `json:"switch_inline_query_current_chat"`
	CallbackGame                 *CallbackGame `json:"callback_game"`
	Pay                          bool          `json:"pay"`
}

// LoginUrl struct
// This object represents a parameter of the inline keyboard button used to automatically authorize a user. Serves as a great replacement for the Telegram Login Widget when the user is coming from Telegram. All the user needs to do is tap/click a button and confirm that they want to log in:
type LoginURL struct {
	URL                string `json:"url"`
	ForwardText        string `json:"forward_text"`
	BotUsername        string `json:"bot_username"`
	RequestWriteAccess bool   `json:"request_write_access"`
}

// CallbackQuery struct
// This object represents an incoming callback query from a callback button in an inline keyboard. If the button that originated the query was attached to a message sent by the bot, the field message will be present. If the button was attached to a message sent via the bot (in inline mode), the field inline_message_id will be present. Exactly one of the fields data or game_short_name will be present.
type CallbackQuery struct {
	ID              string   `json:"url"`
	From            *User    `json:"from"`
	Message         *Message `json:"message"`
	InlineMessageID string   `json:"inline_message_id"`
	ChatInstance    string   `json:"chat_instance"`
	Data            string   `json:"data"`
	GameShortName   string   `json:"game_short_name"`
}

// ForceReply struct
// Upon receiving a message with this object, Telegram clients will display a reply interface to the user (act as if the user has selected the bot‘s message and tapped ’Reply'). This can be extremely useful if you want to create user-friendly step-by-step interfaces without having to sacrifice privacy mode.
type ForceReply struct {
	ForceReply bool `json:"force_reply"`
	Selective  bool `json:"selective"`
}

// ChatPhoto struct
// This object represents a chat photo.
type ChatPhoto struct {
	SmallFileID       string `json:"small_file_id"`
	SmallFileUniqueID string `json:"small_file_unique_id"`
	BigFileID         string `json:"big_file_id"`
	BigFileUniqueID   string `json:"big_file_unique_id"`
}

// ChatMember struct
// This object contains information about one member of a chat.
type ChatMember struct {
	User                  *User  `json:"user"`
	Status                string `json:"status"`
	CustomTitle           string `json:"custom_title"`
	UntilDate             int    `json:"until_date"`
	CanBeEdited           bool   `json:"can_be_edited"`
	CanPostMessages       bool   `json:"can_post_messages"`
	CanEditMessages       bool   `json:"can_edit_messages"`
	CanDeleteMessages     bool   `json:"can_delete_messages"`
	CanRestrictMembers    bool   `json:"can_restrict_members"`
	CanChangeInfo         bool   `json:"can_change_info"`
	CanInviteUsers        bool   `json:"can_invite_users"`
	CanPinMessages        bool   `json:"can_pin_messages"`
	IsMember              bool   `json:"is_member"`
	CanSendMessages       bool   `json:"can_send_messages"`
	CanSendMediaMessages  bool   `json:"can_send_media_messages"`
	CanSendPolls          bool   `json:"can_send_polls"`
	CanSendOtherMessages  bool   `json:"can_send_other_messages"`
	CanAddWebPagePreviews bool   `json:"can_send_web_page_previews"`
}

// ChatPermissions struct
// Describes actions that a non-administrator user is allowed to take in a chat
type ChatPermissions struct {
	CanSendMessages       bool `json:"can_send_messages"`
	CanSendMediaMessages  bool `json:"can_send_media_messages"`
	CanSendPolls          bool `json:"can_send_polls"`
	CanSendOtherMessages  bool `json:"can_send_other_messages"`
	CanAddWebPagePreviews bool `json:"can_send_web_page_previews"`
	CanChangeInfo         bool `json:"can_change_info"`
	CanInviteUsers        bool `json:"can_invite_users"`
	CanPinMessages        bool `json:"can_pin_messages"`
}

// BotCommand struct
// This object represents a bot command.
type BotCommand struct {
	Command     string `json:"command"`
	Description string `json:"description"`
}

// ResponseParameters struct
// Contains information about why a request was unsuccessful.
type ResponseParameters struct {
	MigrateToChatID int `json:"migrate_to_chat_id"`
	RetryAfter      int `json:"retry_after"`
}

// InputMedia interface
// This object represents the content of a media message to be sent. It should be one of
// InputMediaAnimation
// InputMediaDocument
// InputMediaAudio
// InputMediaPhoto
// InputMediaVideo
type InputMedia interface {
}

// InputMediaPhoto struct
// Represents a photo to be sent.
type InputMediaPhoto struct {
	Type      string `json:"type"`
	Media     string `json:"media"`
	Caption   string `json:"caption"`
	ParseMode string `json:"parse_mode"`
}

// InputMediaVideo struct
// Represents a video to be sent.
type InputMediaVideo struct {
	Type              string     `json:"type"`
	Media             string     `json:"media"`
	Thumb             *InputFile `json:"thumb"`
	Caption           string     `json:"caption"`
	ParseMode         string     `json:"parse_mode"`
	Width             int        `json:"width"`
	Height            int        `json:"height"`
	Duration          int        `json:"duration"`
	SupportsStreaming bool       `json:"supports_streaming"`
}

// InputMediaAnimation struct
// Represents an animation file (GIF or H.264/MPEG-4 AVC video without sound) to be sent.
type InputMediaAnimation struct {
	Type      string     `json:"type"`
	Media     string     `json:"media"`
	Thumb     *InputFile `json:"thumb"`
	Caption   string     `json:"caption"`
	ParseMode string     `json:"parse_mode"`
	Width     int        `json:"width"`
	Height    int        `json:"height"`
	Duration  int        `json:"duration"`
}

// InputMediaAudio struct
// Represents an audio file to be treated as music to be sent.
type InputMediaAudio struct {
	Type      string     `json:"type"`
	Media     string     `json:"media"`
	Thumb     *InputFile `json:"thumb"`
	Caption   string     `json:"caption"`
	ParseMode string     `json:"parse_mode"`
	Width     int        `json:"width"`
	Height    int        `json:"height"`
	Duration  int        `json:"duration"`
	Performer string     `json:"performer"`
	Title     string     `json:"title"`
}

// InputMediaDocument struct
// Represents a general file to be sent.
type InputMediaDocument struct {
	Type      string     `json:"type"`
	Media     string     `json:"media"`
	Thumb     *InputFile `json:"thumb"`
	Caption   string     `json:"caption"`
	ParseMode string     `json:"parse_mode"`
}

// InputFile struct
// This object represents the contents of a file to be uploaded. Must be posted using multipart/form-data in the usual way that files are uploaded via the browser.
type InputFile struct {
}

// Sticker struct
// This object represents a sticker.
type Sticker struct {
	FileID       string        `json:"file_id"`
	FileUniqueID string        `json:"file_unique_id"`
	Width        string        `json:"width"`
	Height       string        `json:"height"`
	IsAnimated   bool          `json:"is_animated"`
	Thumb        *PhotoSize    `json:"thumb"`
	Emoji        string        `json:"emoji"`
	SetName      string        `json:"set_name"`
	MaskPosition *MaskPosition `json:"mask_position"`
	FileSize     int           `json:"file_size"`
}

// StickerSet struct
// This object represents a sticker set.
type StickerSet struct {
	Name          string     `json:"name"`
	Title         string     `json:"title"`
	IsAnimated    bool       `json:"is_animated"`
	ContainsMasks bool       `json:"contains_masks"`
	Stickers      []Sticker  `json:"stickers"`
	Thumb         *PhotoSize `json:"thumb"`
}

// MaskPosition struct
// This object describes the position on faces where a mask should be placed by default
type MaskPosition struct {
	Point  string  `json:"point"`
	XShift float32 `json:"x_shift"`
	YShift float32 `json:"y_shift"`
	Scale  float32 `json:"scale"`
}

// InlineQueryResult interface
// This object represents one result of an inline query. Telegram clients currently support results of the following 20 types:
type InlineQueryResult interface{}

// Game struct
// This object represents a game. Use BotFather to create and edit games, their short names will act as unique identifiers.
type Game struct {
	Title        string          `json:"title"`
	Description  string          `json:"description"`
	Photo        []PhotoSize     `json:"photo"`
	Text         string          `json:"text"`
	TextEntities []MessageEntity `json:"text_entities"`
	Animation    *Animation      `json:"animation"`
}

// CallbackGame interface
// A placeholder, currently holds no information. Use BotFather to set up your game.
type CallbackGame interface {
	// SetGameScore(userID int, score, int, force, disabelEditMessage bool, chatID, messageID int, inllineMessageID string) (*Message, bool)
	// GetGameHighScores(userID, chatID, messageID int, InlineMessageID string) []GetGameHighScore
}

// GameHighScore struct
// This object represents one row of the high scores table for a game.
type GameHighScore struct {
	Position int   `json:"position"`
	User     *User `json:"user"`
	Score    int   `json:"score"`
}

// InlineQuery struct
// This object represents an incoming inline query. When the user sends an empty query, your bot could return some default or trending results.
type InlineQuery struct {
	ID       string   `json:"id"`
	From     *User    `json:"from"`
	Location Location `json:"location"`
	Query    string   `json:"query"`
	Offset   string   `json:"offset"`
}

// Update struct
// This object represents an incoming update.
// At most one of the optional parameters can be present in any given update.
type Update struct {
	UpdateId           int                `json:"update_id"`
	Message            *Message           `json:"message"`
	EditedMessage      *Message           `json:"edited_message"`
	ChannelPost        *Message           `json:"channel_post"`
	EditedChannelPost  *Message           `json:"edited_channel_post"`
	InlineQuery        *InlineQuery       `json:"inline_query"`
	ChosenInlineResult *ChosenInlineQuery `json:"chosen_inline_query"`
	CallbackQuery      *CallbackQuery     `json:"callback_query"`
	ShippingQuery      *ShippingQuery     `json:"shipping_query"`
	PreCheckoutQuery   *PreCheckoutQuery  `json:"pre_checkout_query"`
	Poll               Poll               `json:"poll"`
	PollAnswer         *PollAnswer        `json:"poll_answer"`
}

type ChosenInlineQuery struct{}
type ShippingQuery struct{}
type PreCheckoutQuery struct{}
type PassportData struct{}
type Invoice struct{}
type SuccessfulPayment struct{}

// SendMessageBody struct helper for sending messages
type SendMessageBody struct {
	ChatID                int         `json:"chat_id"`
	Text                  string      `json:"text"`
	ParseMode             string      `json:"parse_mode"`
	DisableWebPagePreview bool        `json:"disable_web_page_preview"`
	DisableNotification   bool        `json:"disable_notification"`
	ReplyToMessageID      int         `json:"reply_to_message_id"`
	ReplyMarkup           interface{} `json:"reply_markup"`
}

// ForwardMessageBody struct helper for sending messages
type ForwardMessageBody struct {
	ChatID              int  `json:"chat_id"`
	FromChatID          int  `json:"from_chat_id"`
	DisableNotification bool `json:"disable_notification"`
	MessageID           int  `json:"message_id"`
}

// SendPhotoBody struct helper for sending messages
type SendPhotoBody struct {
	ChatID              int         `json:"chat_id"`
	Photo               string      `json:"photo"`
	Caption             string      `json:"caption"`
	ParseMode           string      `json:"parse_mode"`
	DisableNotification bool        `json:"disable_notification"`
	ReplyToMessageID    int         `json:"reply_to_message_id"`
	ReplyMarkup         interface{} `json:"reply_markup"`
}

// SendLocationBody struct helper for sending messages
type SendLocationBody struct {
	ChatID              int         `json:"chat_id"`
	Latitude            float32     `json:"latitude"`
	Longitude           float32     `json:"longitude"`
	LivePeriod          int         `json:"live_period"`
	DisableNotification bool        `json:"disable_notification"`
	ReplyToMessageID    int         `json:"reply_to_message_id"`
	ReplyMarkup         interface{} `json:"reply_markup"`
}

// SendContactBody struct helper for sending messages
type SendContactBody struct {
	ChatID              int         `json:"chat_id"`
	Latitude            float32     `json:"latitude"`
	Longitude           float32     `json:"longitude"`
	LivePeriod          int         `json:"live_period"`
	DisableNotification bool        `json:"disable_notification"`
	ReplyToMessageID    int         `json:"reply_to_message_id"`
	ReplyMarkup         interface{} `json:"reply_markup"`
}

// SendContactBody struct helper for sending messages
type SetMycommandsBody struct {
	Commands []BotCommand `json:"commands"`
}


type MsgRcv struct {
	OK     bool     `json:"ok,omitempty"`
	Result *Message `json:"result,omitempty"`
}