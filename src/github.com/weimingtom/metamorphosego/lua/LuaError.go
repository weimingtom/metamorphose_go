package metamorphosego

type LuaError struct {
	_errorStatus int
	
	message string //FIXME:
}

func NewLuaError(errorStatus int) *LuaError {
	self := &LuaError{}
	self.message = "Lua error" //super("Lua error");
	self._errorStatus = errorStatus
	return self 
}

func (self *LuaError) GetErrorStatus() int {
	return self._errorStatus;
}
