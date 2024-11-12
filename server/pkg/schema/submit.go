package schema

type Submission struct {
	ID        int    `json:"id"`
    ProblemID int    `json:"problem_id"`
    UserID    string `json:"user_id"`
    Input     string `json:"input"`
	Output   string `json:"output"`
	Status   string `json:"status"`
	ErrorMsg string `json:"errorMsg"`
}
