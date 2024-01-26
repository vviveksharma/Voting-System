package models


type Block struct {
	Pos       int
	Data      interface{} // interface means any value can be entered into it
	Timestamp string
	Hash      string
	PrevHash  string
}

