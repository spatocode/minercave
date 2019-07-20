package net

type Config struct {
	Cryptocurrency		string		`json:"cryptocurrency"`
	Threads				int			`json:"threads"`
	Log					bool		`json:"log"`
	Solo				bool		`json:"solo"`
	Address				string		`json:"address"`
	Pools				[]Pools		`json:"pools"`
}

type Pools struct {
	Url					string		`json:"url"`
	User				string		`json:"user"`
	Password			string		`json:"password"`
	KeepAlive			bool		`json:"keepalive"`
	RigID				string		`json:"rig-id"`
}
