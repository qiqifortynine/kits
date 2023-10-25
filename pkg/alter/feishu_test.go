package alter

import (
	"fmt"
	"testing"
)

func TestFeiShuAlert(t *testing.T) {
	addr := `https://open.feishu.cn/open-apis/bot/v2/hook/07493646-b7ec-4871-96ac-61ec98d1829c`
	f := NewFeiShu(addr, `测试消息`, `all`)
	fmt.Println(f.Send())
}
