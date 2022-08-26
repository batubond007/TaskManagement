package Models

type Task struct {
	ID      int    `json:"id" gorm:"primary_key"`
	Type    string `json:"type"`
	Title   string `json:"title"`
	Context string `json:"context"`
}
