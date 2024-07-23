package logic

import (
	"weikang/services/user_svc/proto/user"
)

type Server struct {
	user.UnimplementedUserServer
}
