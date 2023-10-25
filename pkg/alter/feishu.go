package alter

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func NewFeiShu(hook, msg, at string) Alerter {
	return &FeiShu{
		Hook: hook,
		Msg:  msg,
		At:   at,
	}
}

type FeiShu struct {
	Hook string `json:"hook"`
	Msg  string `json:"msg"`
	At   string `json:"at"`
}

func (f *FeiShu) Send() (err error) {
	f.Msg = fmt.Sprintf(`{"msg_type":"text","content":{"text":"<at user_id=\"%s\">%s</at>%s"}}`, f.At, f.At, f.Msg)
	body := strings.NewReader(f.Msg)
	res, err := http.Post(f.Hook, `application/json`, body)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	resp, _ := io.ReadAll(res.Body)
	fmt.Println(string(resp))
	
	return
}
