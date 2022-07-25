package app

import "CricRadio/controllers/matches"

func mapUrls() {
	router.GET("/matches/list", matches.ListMatches)

}
