package handler

type Error struct {
	Message string `json:"message"`
}

type Title struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Title   string `json:"title"`
}

type Response struct {
	Result []Title `json:"result"`
}

func NewResponse() Response {
	return Response{[]Title{}}
}
