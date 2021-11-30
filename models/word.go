package models

type Word struct {
	ID   int    `json:"id" gorm:"primary_key"`
	English string `json:"english"`
	Shan string `json:"shan"`
}

type Antonym struct {
	WORD_ID   int   `json:"word_id" gorm:"primary_key"`
	English string `json:"english"`
	Shan string `json:"shan"`
}

type Definition struct {
	WORD_ID   int  `json:"word_id" gorm:"primary_key"`
	English string `json:"english"`
	Shan string `json:"shan"`
	Pos string `json:"pos"`
	Uncount string `json:"uncount"`
}

type Translate struct {
	Word
	Antonym Antonym
	Definition Definition
}

// rename to Word table
type Tabler interface {
	TableName() string
}

func (w *Word) TableName() string {
	return "word"
}

func (a *Antonym) TableName() string {
	return "antonym"
}

func (d *Definition) TableName() string {
	return "definition"
}

func (t *Translate) TableName() string {
	return "translate"
}