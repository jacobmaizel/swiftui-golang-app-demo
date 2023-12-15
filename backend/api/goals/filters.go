package goals

import (
	"github.com/gin-gonic/gin"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/competition"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/goal"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/profile"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/squad"
	u "github.com/jacobmaizel/swiftui-golang-app-demo/utils"
)

func handleGoalFilters(c *gin.Context, q *ent.GoalQuery) *ent.GoalQuery {

	if profileId, err := u.UuidFromQuery(c, "profile"); err == nil {
		q = q.Where(
			goal.HasProfileWith(
				profile.ID(profileId),
			),
		)
	}

	if competitionId, err := u.UuidFromQuery(c, "competition"); err == nil {
		q = q.Where(
			goal.HasCompetitionWith(
				competition.ID(competitionId),
			),
		)

	}

	if squadId, err := u.UuidFromQuery(c, "squad"); err == nil {
		q = q.Where(
			goal.HasSquadWith(
				squad.ID(squadId),
			),
		)
	}

	return q
}
