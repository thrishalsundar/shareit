package models

type Resp struct {
	Msg  string
	Err  error
	Data any
	Stat int
}

func PResMaker(msg string, data any, stat int, err error) *Resp {
	return &Resp{
		Msg: msg, Data: data, Stat: stat, Err: err,
	}
}

type Custom struct {
	Data any
	Msg  string
}

// PASSWORDMANAGEMENT
