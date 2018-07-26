package irc

import (
	"bufio"
	"errors"
	"io"
	"strings"
)

// Represents message commands.
const (
	PasswordMessage     Command = "PASS"
	NicknameMessage     Command = "NICK"
	UserMessage         Command = "USER"
	OperatorMessage     Command = "OPER"
	ModeMessage         Command = "MODE"
	ServiceMessage      Command = "SERVICE"
	QuitMessage         Command = "QUIT"
	SquitMessage        Command = "SQUIT"
	JoinMessage         Command = "JOIN"
	PartMessage         Command = "PART"
	TopicMessage        Command = "TOPIC"
	NamesMessage        Command = "NAMES"
	ListMessage         Command = "LIST"
	InviteMessage       Command = "INVITE"
	KickMessage         Command = "KICK"
	PrivateMessage      Command = "PRIVMSG"
	NoticeMessage       Command = "NOTICE"
	MotdMessage         Command = "MOTD"
	ListUsersMessage    Command = "LUSERS"
	VersionMessage      Command = "VERSION"
	StatsMessage        Command = "STATS"
	LinksMessage        Command = "LINKS"
	TimeMessage         Command = "TIME"
	ConnectMessage      Command = "CONNECT"
	TraceMessage        Command = "TRACE"
	AdminMessage        Command = "ADMIN"
	InfoMessage         Command = "INFO"
	ServiceListMessage  Command = "SERVLIST"
	ServiceQueryMessage Command = "SQUERY"
	WhoMessage          Command = "WHO"
	WhoisMessage        Command = "WHOIS"
	WhowasMessage       Command = "WHOWAS"
	KillMessage         Command = "KILL"
	PingMessage         Command = "PING"
	PongMessage         Command = "PONG"
	ErrorMessage        Command = "ERROR"
	AwayMessage         Command = "AWAY"
	RehashMessage       Command = "REHASH"
	DieMessage          Command = "DIE"
	RestartMessage      Command = "RESTART"
	SummonMessage       Command = "SUMMON"
	UsersMessage        Command = "USERS"
	WallopsMessage      Command = "WALLOPS"
	UserHostMessage     Command = "USERHOST"
	IsonMessage         Command = "ISON"
)

// Represents reply message commands.
const (
	WelcomeReply           Command = "001"
	YourHostReply          Command = "002"
	CreatedReply           Command = "003"
	MyInfoReply            Command = "004"
	BounceReply            Command = "005"
	UserHostReply          Command = "302"
	IsonReply              Command = "303"
	AwayReply              Command = "301"
	UnawayReply            Command = "305"
	NowAwayReply           Command = "306"
	WhoisUserReply         Command = "311"
	WhoisServerReply       Command = "312"
	WhoisOperatorReply     Command = "313"
	WhoisIdleReply         Command = "317"
	EndOfWhoisReply        Command = "318"
	WhoisChannelsReply     Command = "319"
	WhowasUserReply        Command = "314"
	EndOfWhowasReply       Command = "369"
	ListReply              Command = "322"
	EndOfListReply         Command = "323"
	UniqueOperatorisReply  Command = "325"
	ChannelModeisReply     Command = "324"
	NoTopicReply           Command = "331"
	TopicReply             Command = "332"
	InvitingReply          Command = "341"
	SummoningReply         Command = "342"
	InviteListReply        Command = "346"
	EndOfInviteListReply   Command = "347"
	ExceptListReply        Command = "348"
	EndOfExceptListReply   Command = "349"
	VersionReply           Command = "351"
	WhoReply               Command = "352"
	EndOfWhoReply          Command = "315"
	NamesReply             Command = "353"
	EndOfNamesReply        Command = "366"
	LinksReply             Command = "364"
	EndOfLinksReply        Command = "365"
	BanListReply           Command = "367"
	EndOfBanListReply      Command = "368"
	InfoReply              Command = "371"
	EndOfInfoReply         Command = "374"
	MotdStartReply         Command = "375"
	MotdReply              Command = "372"
	EndOfMotdReply         Command = "376"
	YoureOperatorReply     Command = "381"
	RehashingReply         Command = "382"
	YoureServiceReply      Command = "383"
	TimeReply              Command = "391"
	UsersStartReply        Command = "392"
	UsersReply             Command = "393"
	EndOfUsersReply        Command = "394"
	NoUsersReply           Command = "395"
	TraceLinkReply         Command = "200"
	TraceConnectingReply   Command = "201"
	TraceHandshakeReply    Command = "202"
	TraceUnknownReply      Command = "203"
	TraceOperatorReply     Command = "204"
	TraceUserReply         Command = "205"
	TraceServerReply       Command = "206"
	TraceServiceReply      Command = "207"
	TraceNewTypeReply      Command = "208"
	TraceClassReply        Command = "209"
	TraceLogReply          Command = "261"
	EndOfTraceReply        Command = "262"
	StatsLinkInfoReply     Command = "211"
	StatsCommandsReply     Command = "212"
	EndOfStatsReply        Command = "219"
	StatsOlineReply        Command = "243"
	UserModeisReply        Command = "221"
	ServiceListReply       Command = "234"
	EndOfServiceListReply  Command = "235"
	ListUsersClientReply   Command = "251"
	ListUsersOperatorReply Command = "252"
	ListUsersUnknownReply  Command = "253"
	ListUsersChannelsReply Command = "254"
	ListUsersMeReply       Command = "255"
	AdminMeReply           Command = "256"
	AdminLoc1Reply         Command = "257"
	AdminLoc2Reply         Command = "258"
	AdminEmailReply        Command = "259"
	TryAgainReply          Command = "263"
)

// Represents error message commands.
const (
	NoSuchNicknameError        Command = "401"
	NoSuchServerError          Command = "402"
	NoSuchChannelError         Command = "403"
	CannotSendToChannelError   Command = "404"
	TooManyChannelsError       Command = "405"
	WasNoSuchNicknameError     Command = "406"
	TooManyTargetsError        Command = "407"
	NoSuchServiceError         Command = "408"
	NoOriginError              Command = "409"
	NoRecipientError           Command = "411"
	NoTextToSendError          Command = "412"
	NoTopLevelError            Command = "413"
	WildcardTopLevelError      Command = "414"
	BadMaskError               Command = "415"
	UnknownCommandError        Command = "421"
	NoMotdError                Command = "422"
	NoAdminInfoError           Command = "423"
	FileError                  Command = "424"
	NoNicknameGivenError       Command = "431"
	ErroneusNicknameError      Command = "432"
	NicknameInUseError         Command = "433"
	NicknameCollisionError     Command = "436"
	UnavailableResourceError   Command = "437"
	UserNotInChannelError      Command = "441"
	NotOnChannelError          Command = "442"
	UserOnChannelError         Command = "443"
	NoLoginError               Command = "444"
	SummonDisabledError        Command = "445"
	UserDisabledError          Command = "446"
	NotRegisteredError         Command = "451"
	NeedMoreParamsError        Command = "461"
	AlreadyRegisteredError     Command = "462"
	NoPermissionForHostError   Command = "463"
	PasswordMismatchError      Command = "464"
	YoureBannedError           Command = "465"
	YouWillBeBannedError       Command = "466"
	KeySetError                Command = "467"
	ChannelIsFullError         Command = "471"
	UnknownModeError           Command = "472"
	InviteOnlyChannelError     Command = "473"
	BannedFromChannelError     Command = "474"
	BadChannelKeyError         Command = "475"
	BadChannelMaskError        Command = "476"
	NoChannelModesError        Command = "477"
	BanListFullError           Command = "478"
	NoPrivilegesError          Command = "481"
	ChannelOperatorNeededError Command = "482"
	CannotKillServerError      Command = "483"
	RestrictedError            Command = "484"
	UniqueOperatorNeededError  Command = "485"
	NoOperatorHostError        Command = "491"
	UserModeUnknownFlagError   Command = "501"
	UsersDontMatchError        Command = "502"
)

// ErrInvalidMessage raised when input message contains invalid token or format.
var ErrInvalidMessage = errors.New("irc: invalid input message")

// Sender represents structured sender information. Nickname can also acts as
// the server name if originated from relay chat server.
type Sender struct {
	Nickname string
	User     string
	Host     string
}

// Command represents a valid relay chat command.
type Command string

// builder represents reusable buffer/string builder.
type builder interface {
	io.Writer
	WriteByte(byte) error
	WriteRune(rune) (int, error)
	WriteString(string) (int, error)
}

// Message represents parsed raw relay chat message.
type Message struct {
	Sender  Sender
	Command Command
	Params  []string
}

// Build message with specified builder input.
func (m Message) build(b builder) error {
	// Process prefix if available
	if len(m.Sender.Nickname) > 0 {
		if err := b.WriteByte(':'); err != nil {
			return err
		}
		if _, err := b.WriteString(m.Sender.Nickname); err != nil {
			return err
		}
		if len(m.Sender.User) > 0 {
			if err := b.WriteByte('!'); err != nil {
				return err
			}
			if _, err := b.WriteString(m.Sender.User); err != nil {
				return err
			}
		}
		if len(m.Sender.Host) > 0 {
			if err := b.WriteByte('@'); err != nil {
				return err
			}
			if _, err := b.WriteString(m.Sender.Host); err != nil {
				return err
			}
		}
		if err := b.WriteByte(' '); err != nil {
			return err
		}
	}
	// Write command
	if _, err := b.WriteString(string(m.Command)); err != nil {
		return err
	}
	// Write trailing params, if available
	for i, param := range m.Params {
		if err := b.WriteByte(' '); err != nil {
			return err
		}
		if i == len(m.Params)-1 {
			if err := b.WriteByte(':'); err != nil {
				return err
			}
		}
		if _, err := b.WriteString(param); err != nil {
			return err
		}
	}
	// Write trailing carriage return and newline
	_, err := b.WriteString("\r\n")
	return err
}

// String encodes message into relay chat compatible message protocol.
func (m Message) String() string {
	var b strings.Builder
	m.build(&b)
	return b.String()
}

// IsError determine whether a message contains error reply or not.
// TryAgainReply is classified as error because it does not process given
// command from client.
func (m Message) IsError() bool {
	if len(m.Command) == 3 && (m.Command[0] == '4' || m.Command[0] == '5' || m.Command == TryAgainReply) {
		return true
	}
	return false
}

// Split input with defined separator character.
func split(input string, sep byte) (string, string) {
	if pos := strings.IndexByte(input, sep); pos >= 0 {
		return input[:pos], input[pos+1:]
	}
	return input, ""
}

// Writer represents the relay chat message writer. It is primarily used to wrap
// net.Conn and continuously parse outgoing messages to server.
type Writer struct {
	wr *bufio.Writer
}

// Write message to writer stream.
func (w *Writer) Write(msg Message) error {
	if err := msg.build(w.wr); err != nil {
		return err
	}
	return w.wr.Flush()
}

// NewWriter create new relay chat message writer.
func NewWriter(w io.Writer) *Writer {
	return &Writer{
		wr: bufio.NewWriter(w),
	}
}

// Reader represents the relay chat message reader. It is primarily used to wrap
// net.Conn and continuously parse incoming messages from server.
type Reader struct {
	rd *bufio.Reader
}

// Read stream and parse them into message.
func (r *Reader) Read() (Message, error) {
	// Read input message as string
	msg, err := r.rd.ReadString('\n')
	if err != nil {
		return Message{}, err
	}
	// Parse message
	return NewMessage(msg)
}

// NewReader create new relay chat message reader.
func NewReader(r io.Reader) *Reader {
	return &Reader{
		rd: bufio.NewReader(r),
	}
}

// NewMessage create new parsed raw message from input string.
func NewMessage(input string) (Message, error) {
	// Structured message output
	var msg Message
	// Do not process empty message
	if len(input) == 0 {
		return msg, ErrInvalidMessage
	}
	// Strip carriage return and/or new line token in input
	if strings.HasSuffix(input, "\r\n") {
		input = input[:len(input)-2]
	} else if input[len(input)-1] == '\n' {
		input = input[:len(input)-1]
	}
	// Check if input has sender (prefix) information
	if input[0] == ':' {
		// Trim prefix token
		input = input[1:]
		// Split prefix and input message
		var prefix string
		if prefix, input = split(input, ' '); len(input) == 0 || len(prefix) == 0 {
			return msg, ErrInvalidMessage
		}
		// Get nickname, username, and hostname
		msg.Sender.Nickname, msg.Sender.Host = split(prefix, '@')
		msg.Sender.Nickname, msg.Sender.User = split(msg.Sender.Nickname, '!')
	}
	// Parse message command
	var command string
	command, input = split(input, ' ')
	if len(command) == 0 {
		return msg, ErrInvalidMessage
	}
	msg.Command = Command(command)
	// Parse message params (if available)
	for len(input) > 0 {
		// Check if current params is the last param
		if input[0] == ':' {
			msg.Params = append(msg.Params, input[1:])
			break
		}
		// Separate parameter from the next one
		var param string
		param, input = split(input, ' ')
		// Append param into param list
		msg.Params = append(msg.Params, param)
	}
	// Return parsed message
	return msg, nil
}
