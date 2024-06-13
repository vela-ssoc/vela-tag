package vtag

import (
	"github.com/vela-ssoc/vela-kit/lua"
	"github.com/vela-ssoc/vela-kit/vela"
)

var xEnv vela.Environment

func newLuaTagL(L *lua.LState) int {
	L.CheckCodeVM("vela-tag")
	L.Push(NewTag())
	return 1
}

func newLuaSetTagL(L *lua.LState) int {
	L.CheckCodeVM("vela-tag")
	t := NewTag()
	t.Range(L, t.AddTag)
	t.Send()
	return 0
}

func newLuaDelTagL(L *lua.LState) int {
	L.CheckCodeVM("vela-tag")
	t := NewTag()
	t.Range(L, t.delTag)
	t.Send()
	return 0
}

func WithEnv(env vela.Environment) {
	xEnv = env
	xEnv.Set("tag", lua.NewFunction(newLuaTagL))
	xEnv.Set("set_tag", lua.NewFunction(newLuaSetTagL))
	xEnv.Set("del_tag", lua.NewFunction(newLuaDelTagL))
}
