package metamorphosego

type LuaError struct {
	_errorStatus int
	
	message string //FIXME:
}

func NewLuaError(errorStatus int) *LuaError {
	e := &LuaError{}
	e.message = "Lua error"; //super("Lua error");
	e._errorStatus = errorStatus;
	return e 
}

func (self *LuaError) GetErrorStatus() int {
	return self._errorStatus;
}
