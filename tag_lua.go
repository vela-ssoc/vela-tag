package vtag

import (
	"github.com/bytedance/sonic"
	"github.com/vela-ssoc/vela-kit/auxlib"
	"github.com/vela-ssoc/vela-kit/lua"
)

func (t *tag) String() string                         { return auxlib.B2S(t.Byte()) }
func (t *tag) Type() lua.LValueType                   { return lua.LTObject }
func (t *tag) AssertFloat64() (float64, bool)         { return 0, false }
func (t *tag) AssertString() (string, bool)           { return "", false }
func (t *tag) AssertFunction() (*lua.LFunction, bool) { return nil, false }
func (t *tag) Peek() lua.LValue                       { return t }

func (t *tag) Byte() []byte {
	chunk, err := sonic.Marshal(t)
	if err != nil {
		return []byte("[]")
	}
	return chunk
}

func (t *tag) Range(L *lua.LState, handle func(string)) int {
	n := L.GetTop()
	if n == 0 {
		return 0
	}

	for i := 1; i <= n; i++ {
		handle(L.CheckString(i))
	}
	return 0
}

func (t *tag) addL(L *lua.LState) int {
	return t.Range(L, t.addTag)
}

func (t *tag) delL(L *lua.LState) int {
	return t.Range(L, t.delTag)
}

func (t *tag) sendL(L *lua.LState) int {
	t.send()
	return 0
}

func (t *tag) Index(L *lua.LState, key string) lua.LValue {
	switch key {
	case "add":
		return L.NewFunction(t.addL)
	case "del":
		return L.NewFunction(t.delL)
	case "send":
		return L.NewFunction(t.sendL)

	}
	return lua.LNil
}
