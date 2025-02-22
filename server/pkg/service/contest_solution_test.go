package service

import (
	"context"
	"testing"

	"github.com/ForwardGlimpses/OJ/server/pkg/schema"
	gomonkey "github.com/agiledragon/gomonkey/v2"

	"github.com/onsi/gomega"
)

func TestGetContestRanking(t *testing.T) {
	var mockresult schema.ContestSolutionItems
	patches := gomonkey.ApplyMethod(ContestSolutionSvc, "GetContestSolutions", func(this *ContestSolutionService, ctx context.Context, contestID int) (schema.ContestSolutionItems, error) {
		return mockresult, nil
	})
	defer patches.Reset()

	tests := []struct {
		name   string
		input  schema.ContestSolutionItems
		output []ContestRankingItem
	}{
		{
			input:  schema.ContestSolutionItems{},
			output: []ContestRankingItem{},
		},
	}

	for _, tt := range tests {
		mockresult = tt.input
		res, _, _ := ContestSolutionSvc.Query(context.Background(), schema.ContestSolutionParams{})
		gomega.Expect(tt.output, gomega.Equal(res))
	}
}
