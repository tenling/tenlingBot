package Action

import "fmt"

func replyAirplan(name string) string {
	reply := airplanMap[name]
	if reply == nil {
		return replyAirplanFail
	}
	return fmt.Sprint(reply)
}

func addAirplan(name string, context string) string {
	reply := airplanMap[name]
	if reply == nil {
		newAirplan := []string{
			context,
		}
		airplanMap[name] = newAirplan
		return name + "的航班，我知道了喵～"
	}
	reply = append(reply, context)
	return name + "的航班，我知道了喵～"
}

func addKeyword(keyword string, reply string, user string) string {
	setRely := autoRelyMap[keyword]
	if setRely == "" {
		autoRelyMap[keyword] = reply
		return addKeywordSuccess
	}

	if user == tenling {
		autoRelyMap[keyword] = reply
		return addKeywordSuccess
	}
	return addKeywordFail
}

func autoRely(keyword string) string {
	reply := autoRelyMap[keyword]
	if reply == "" {
		return keyword
	}
	return reply
}
