package redis

import (
	"context"
	"encoding/json"
	"log"

	"github.com/google/uuid"
	"github.com/jacobmaizel/swiftui-golang-app-demo/api/competitions"
	"github.com/jacobmaizel/swiftui-golang-app-demo/api/squads"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent"
	"github.com/jacobmaizel/swiftui-golang-app-demo/server"
	"github.com/jacobmaizel/swiftui-golang-app-demo/shared"
)

func CreateRedisIndices(c context.Context) {
	// initialize the indexes, if they were created new, we need to push existing data to them

	prof_idx_create_cmd := server.Redis.
		B().
		FtCreate().
		Index("profileIdx").OnJson().
		Prefix(1).Prefix("profile:").
		Schema().
		FieldName("$.first_name").As("first_name").Text().
		FieldName("$.last_name").As("last_name").Text().
		FieldName("$.public").As("public").Tag().
		Build()

	if prof_idx_res := server.Redis.Do(c, prof_idx_create_cmd); prof_idx_res.Error() != nil {

		// log.Println("hey there", reflect.TypeOf(prof_idx_res.Error()))
		log.Println("error creating profile index:::", prof_idx_res.Error())
	}

	comp_idx_create_cmd := server.Redis.
		B().
		FtCreate().
		Index("competitionIdx").OnJson().
		Prefix(1).Prefix("competition:").
		Schema().
		FieldName("$.title").As("title").Text().
		FieldName("$.description").As("description").Text().
		FieldName("$.participant_types").As("participant_types").Text().
		FieldName("$.workout_types").As("workout_types").Text().
		FieldName("$.type").As("type").Text().
		FieldName("$.public").As("public").Tag().
		Build()

	if comp_idx_res := server.Redis.Do(c, comp_idx_create_cmd); comp_idx_res.Error() != nil {
		// log.Println("hey there copm", reflect.TypeOf(comp_idx_res.Error()))
		log.Println("error creating competition index:::", comp_idx_res.Error())
	}
	squad_idx_create_cmd := server.Redis.
		B().
		FtCreate().
		Index("squadIdx").OnJson().
		Prefix(1).Prefix("squad:").
		Schema().
		FieldName("$.title").As("title").Text().
		FieldName("$.public").As("public").Tag().
		Build()

	if squad_idx_res := server.Redis.Do(c, squad_idx_create_cmd); squad_idx_res.Error() != nil {
		log.Println("error creating squad index:::", squad_idx_res.Error())
	}

}

func IndexProfileDoc(c context.Context, prof *ent.Profile) {

	profToIdx := shared.PartialProfile{
		ID:        prof.ID,
		FirstName: prof.FirstName,
		LastName:  prof.LastName,
		Picture:   prof.Picture,
		Public:    &prof.Public,
	}

	prof_string, err := json.Marshal(profToIdx)

	if err != nil {
		log.Println("\n\n FAILED TO MARSHAL PROFILE WHILE INSERTING REDIS DOCUMENTS", err)
		return
	}

	cmd := server.Redis.B().
		JsonSet().
		Key("profile:" + prof.ID.String()).
		Path(".").
		Value(string(prof_string)).
		Build()

	if res := server.Redis.Do(c, cmd); res.Error() != nil {
		log.Println("\n\nFAILED TO INSERT PROFILE INTO REDIS" + prof.ID.String() + res.Error().Error())
		return
	}

}

func IndexCompetitionDoc(c context.Context, comp *ent.Competition) {

	// log.Println("indexing competition doc iD: ", comp.ID.String())

	var err error
	if comp.Edges.Owner == nil {
		// log.Println("INDEXCOMPDOC::: Pulling owner to index it.")
		comp.Edges.Owner, err = comp.QueryOwner().First(c)

		if err != nil {
			comp.Edges.Owner = &ent.Profile{
				ID:        uuid.New(),
				FirstName: "Missing",
				LastName:  "Owner Info",
				Picture:   "https://lh3.googleusercontent.com/a/AAcHTtdX9WVfVj5dWY5vVLyJozVxIdjY76d7qLJjgI7u2a4TK1xU=s20-c",
			}
		}
	}

	var compToIndex competitions.ResponseCompetition = competitions.ResponseCompetition{
		ID:                comp.ID,
		Title:             comp.Title,
		Description:       comp.Description,
		Scheduled:         &comp.Scheduled,
		Participant_types: comp.ParticipantTypes,
		Type:              comp.Type,
		Status:            comp.Status,
		Start:             comp.Start,
		End:               comp.End,
		Public:            &comp.Public,
		Workout_types:     comp.WorkoutTypes,
		Owner: shared.PartialProfile{
			ID:        comp.Edges.Owner.ID,
			FirstName: comp.Edges.Owner.FirstName,
			LastName:  comp.Edges.Owner.LastName,
			Picture:   comp.Edges.Owner.Picture,
			Public:    &comp.Edges.Owner.Public,
		},
		CreatedAt: comp.CreatedAt,
		UpdatedAt: comp.UpdatedAt,
	}

	comp_string, err := json.Marshal(compToIndex)

	if err != nil {
		log.Println("\n\n FAILED TO MARSHAL competition WHILE INSERTING REDIS DOCUMENTS", err)
		return
	}

	cmd := server.Redis.B().
		JsonSet().
		Key("competition:" + comp.ID.String()).
		Path(".").
		Value(string(comp_string)).
		Build()

	if res := server.Redis.Do(c, cmd); res.Error() != nil {
		log.Println("\n\nFAILED TO INSERT competition INTO REDIS" + comp.ID.String() + res.Error().Error())
		return
	}

}

func IndexSquadDoc(c context.Context, squad *ent.Squad) {

	if squad.Edges.Owner == nil {
		// log.Println("INDEXSQUADDOC::: Pulling owner to index it.")
		var err error
		squad.Edges.Owner, err = squad.QueryOwner().First(c)

		if err != nil {
			log.Println("\n\nFAILED TO PULL OWNER, using fake profile as owner FOR SQUAD WHILE INSERTING REDIS DOCUMENTS for squad: ", squad.ID.String())

			// set owner to fake profile with random info
			squad.Edges.Owner = &ent.Profile{
				ID:        uuid.New(),
				FirstName: "Missing",
				LastName:  "Owner Info",
				Picture:   "https://lh3.googleusercontent.com/a/AAcHTtdX9WVfVj5dWY5vVLyJozVxIdjY76d7qLJjgI7u2a4TK1xU=s20-c",
			}
		}
	}

	if squad.Edges.Members == nil {
		// log.Println("INDEXSQUADDOC::: Pulling members to index it.")
		squad.Edges.Members = squad.QueryMembers().AllX(c)
	}

	member_list := []shared.PartialProfile{}

	for _, member := range squad.Edges.Members {
		member_list = append(member_list, shared.PartialProfile{
			ID:        member.ID,
			FirstName: member.FirstName,
			LastName:  member.LastName,
			Picture:   member.Picture,
			Public:    &member.Public,
		})
	}

	squadToIndex := squads.ResponseSquad{

		ID:      squad.ID,
		Title:   squad.Title,
		Public:  squad.Public,
		Members: member_list,
		Owner: shared.PartialProfile{
			ID:        squad.Edges.Owner.ID,
			FirstName: squad.Edges.Owner.FirstName,
			LastName:  squad.Edges.Owner.LastName,
			Picture:   squad.Edges.Owner.Picture,
			Public:    &squad.Edges.Owner.Public,
		},
		CreatedAt: squad.CreatedAt,
		UpdatedAt: squad.UpdatedAt,
	}

	squad_string, err := json.Marshal(squadToIndex)

	if err != nil {
		log.Println("\n\n FAILED TO MARSHAL squad WHILE INSERTING REDIS DOCUMENTS", err)
		return
	}

	cmd := server.Redis.B().
		JsonSet().
		Key("squad:" + squad.ID.String()).
		Path(".").
		Value(string(squad_string)).
		Build()

	if res := server.Redis.Do(c, cmd); res.Error() != nil {
		log.Println("\n\nFAILED TO INSERT squad INTO REDIS" + squad.ID.String() + res.Error().Error())
		return
	}

}

// utilities to marshal data into indexes to put into redisJSON

func InsertProfileDocuments(c context.Context) {

	all_profiles, err := server.Db.Profile.Query().All(c)

	if err != nil {
		log.Println("\n\nFAILED TO PULL PROFILES WHILE INSERTING REDIS DOCUMENTS")
		return
	}

	for _, prof := range all_profiles {

		IndexProfileDoc(c, prof)
	}

}

func InsertCompetitionDocuments(c context.Context) {

	all_competitions, err := server.Db.Competition.Query().All(c)

	if err != nil {
		log.Println("\n\nFAILED TO PULL COMPETITIONS WHILE INSERTING REDIS DOCUMENTS")
		return
	}

	for _, comp := range all_competitions {

		IndexCompetitionDoc(c, comp)
	}

}

func InsertSquadDocuments(c context.Context) {

	all_squads, err := server.Db.Squad.Query().All(c)

	if err != nil {
		log.Println("\n\nFAILED TO PULL squads WHILE INSERTING REDIS DOCUMENTS")
		return
	}

	for _, squad := range all_squads {

		IndexSquadDoc(c, squad)
	}

}
