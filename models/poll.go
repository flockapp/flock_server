package models

type Poll struct {
	Id         int64 `json:"id"`
	ActivityId int64 `json:"-"`
	UserId     int64 `json:"-"`
	PosVotes   int64 `json:"posVotes"`
	NegVotes   int64 `json:"negVotes"`
}

type PollResponse struct {
	Id      int64  `json:"id"`
	PollId  int64  `json:"-"`
	UserId  int64  `json:"-"`
	Body    string `json:"body"`
	Verdict bool   `json:"verdict"`
}
