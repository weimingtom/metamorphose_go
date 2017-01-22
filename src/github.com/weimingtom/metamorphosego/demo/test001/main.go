package main

import (
	"fmt"
	. "github.com/weimingtom/metamorphosego/lua"
)

type HookImpl struct {
	
}

func (self *HookImpl) LuaHook(L *Lua, ar *Debug) int {
	fmt.Print("HookImpl.LuaHook is called\n")
	return 0
}

func main() {
	fmt.Print("metamorphosego start...\n")
	NewBlockCnt()
	NewConsControl(nil)
	NewDebug()
	NewExpdesc()
	NewLua()
	NewLuaError(0)
	
	var hook Hook
	hook = &HookImpl{}
	hook.LuaHook(nil, nil)
	
	NewLHSAssign()
	NewLocVar()
}
