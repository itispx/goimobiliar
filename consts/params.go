package consts

type RunMultiOutput[T any] []*RunMultiOutputEntry[T]

type RunMultiOutputEntry[T any] struct {
	Success bool   `json:"success,omitempty"`
	ImobId  string `json:"imobId,omitempty"`
	Data    T      `json:"data,omitempty"`
	Error   Error  `json:"error,omitempty"`
}

type Error struct {
	Message string `json:"message,omitempty"`
}

type RunMultiInput[T any] struct {
	Parallel bool
	Entries  []*RunMultiInputEntry[T]
}

type RunMultiInputEntry[T any] struct {
	Endpoint string
	ImobId   string
	UserId   string
	UserPass string
	Input    T
}
