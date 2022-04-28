package global

var (
	ServerInfo    *Server
	WebSocketInfo *WebSocket
	MysqlInfo     *Mysql
	RedisInfo     *Redis
	GRPCInfo      *GRPC
)

type Server struct {
	Host string
	Port string
}

type WebSocket struct {
	Host string
	Port string
}

type GRPC struct {
	Host string
	Port string
}

type Mysql struct {
	Host     string
	Port     string
	User     string
	PassWord string
	DBName   string
}

type Redis struct {
	Host     string
	Port     string
	PassWord string
	DB       int
}
