package schema

type SolutionItem struct {
	Solution_ID int
	Problem_ID  int
	User_ID     string
	Time        int
	Memory      int
	In_date     string
	Language    string
	Code_length string
	Juage_time  string
	Juager      string
	Pass_rate   string
}

type SolutionDBItem struct {
	Solution_ID int
	Problem_ID  int
	User_ID     string
	Time        int
	Memory      int
	In_date     string
	Language    string
	Code_length string
	Juage_time  string
	Juager      string
	Pass_rate   string
}
