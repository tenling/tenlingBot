package Action

import (
	"strings"
)

func Action(msg string, user string) string {
	msg = msg[1:]
	reply := ""
	split := strings.Split(msg, " ")
	switch split[0] {
	case "喵":
		if len(split) != 3 {
			return unknowMessage
		}
		reply = addKeyword(split[1], split[2], user)
	case "airplan":
		switch len(split) {
		case 3:
			reply = addAirplan(user, split[2])
		case 2:
			reply = replyAirplan(user)
		default:
			reply = unknowMessage
		}
	default:
		reply = autoRely(msg)
	}
	return reply
}
