package vtag

import (
	"fmt"
	"github.com/vela-ssoc/vela-kit/audit"
	"github.com/vela-ssoc/vela-kit/opcode"
)

type tag struct {
	Del []string `json:"del"`
	Add []string `json:"add"`
}

func newTag() *tag {
	return &tag{}
}

func contain(dst []string, tv string) bool {
	n := len(dst)
	if n == 0 {
		return false
	}

	for i := 0; i < n; i++ {
		if dst[i] == tv {
			return true
		}
	}
	return false
}

func (t *tag) addTag(tv string) {
	if len(tv) == 0 {
		return
	}

	if contain(t.Add, tv) {
		return
	}

	t.Add = append(t.Add, tv)
}

func (t *tag) delTag(tv string) {
	if len(tv) == 0 {
		return
	}

	if contain(t.Del, tv) {
		return
	}

	t.Del = append(t.Del, tv)
}

func (t *tag) send() error {
	if len(t.Del)+len(t.Add) == 0 {
		return fmt.Errorf("tag empty")
	}
	err := xEnv.TnlSend(opcode.OpTag, t)
	if err != nil {
		audit.Errorf("send tag fail %v", t).From("vela-tag").Put()
	}
	return err
}
