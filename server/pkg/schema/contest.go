package schema

type ContestItem struct {
	Contest_ID  int
	Title       string
	Start_time  string
	End_time    string
	Password    string
	Administrator string
	Description string
}

type ContestDBItem struct {
	Contest_ID  int
	Title       string
	Start_time  string
	End_time    string
	Password    string
	Administrator string
	Description string
}
