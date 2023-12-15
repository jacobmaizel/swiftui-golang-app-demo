package services

import (
	"context"
	"encoding/json"
	"errors"
	"log"

	"github.com/jacobmaizel/swiftui-golang-app-demo/api/competitions"
	"github.com/jacobmaizel/swiftui-golang-app-demo/api/profiles"
	"github.com/jacobmaizel/swiftui-golang-app-demo/api/squads"
	"github.com/jacobmaizel/swiftui-golang-app-demo/server"
	"github.com/redis/rueidis"
)

func generateSearchResults(c context.Context, for_type, term string) (int64, any, error) {
	// return type (query_string, index_name, error)

	var query_string, index_name string

	switch for_type {
	case "profiles":
		index_name = "profileIdx"
		query_string = "@first_name|last_name: *" + term + "*" + "@public: {true}"
		// var prof_res_list []profiles.ResponseProfile = []profiles.ResponseProfile{}

		count, search_res, err := ftSearch(c, query_string, index_name)

		if err != nil {
			return 0, "", err
		}

		var res []profiles.ResponseProfile = []profiles.ResponseProfile{}

		for _, redis_prof_doc := range search_res {

			doc := redis_prof_doc.Doc["$"]

			var p profiles.ResponseProfile

			if err := json.Unmarshal([]byte(doc), &p); err != nil {
				return 0, "", err
			}
			res = append(res, p)
		}

		return count, res, nil

	case "competitions":
		index_name = "competitionIdx"
		query_string = "@title|description|participant_types|workout_types|type: *" + term + "*" + "@public: {true}"
		// var comp_res_list []competitions.ResponseCompetition = []competitions.ResponseCompetition{}

		count, search_res, err := ftSearch(c, query_string, index_name)

		if err != nil {
			return 0, "", err
		}

		var res []competitions.ResponseCompetition = []competitions.ResponseCompetition{}

		for _, redis_comp_doc := range search_res {

			doc := redis_comp_doc.Doc["$"]

			var comp_res competitions.ResponseCompetition

			err := json.Unmarshal([]byte(doc), &comp_res)

			if err != nil {
				// return 0, "", err
				log.Println("error unmarshalling comp_res", err)
			}

			res = append(res, comp_res)
			// res = append(comp_res_list, comp_res)

		}

		return count, res, nil
	case "squads":
		index_name = "squadIdx"
		query_string = "@title: *" + term + "*" + "@public: {true}"

		// var squad_res_list []squads.ResponseSquad = []squads.ResponseSquad{}

		count, search_res, err := ftSearch(c, query_string, index_name)

		if err != nil {
			return 0, "", err
		}

		var res []squads.ResponseSquad = []squads.ResponseSquad{}

		for _, redis_squad_doc := range search_res {

			doc := redis_squad_doc.Doc["$"]

			var squad_res squads.ResponseSquad

			if err := json.Unmarshal([]byte(doc), &squad_res); err != nil {
				return 0, "", err
			}
			res = append(res, squad_res)

		}

		return count, res, nil
	default:
		return 0, "", errors.New("invalid type")
	}

}

func ftSearch(c context.Context, query_string, index_name string) (int64, []rueidis.FtSearchDoc, error) {

	// query_string, index_name, err := generateSearchParams(c, for_type, search_term)

	cmd := server.Redis.B().
		FtSearch().
		Index(index_name).
		Query(query_string).
		Build()

	count, search_res, err := server.Redis.Do(c, cmd).AsFtSearch()

	if err != nil {
		log.Println("error doing ft search", err)
		return 0, nil, err

	}

	// log.Println("ftSearch:::", search_res)

	return count, search_res, nil
}
