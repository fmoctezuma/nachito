package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const (
	matchesapi = "https://api.fifa.com/api/v1/calendar/matches?idCompetition=17&idSeason=254645&idStage=275073&language=en-GB"
)

type MatchesData struct {
	Results []struct {
		IDCompetition string      `json:"IdCompetition"`
		IDSeason      string      `json:"IdSeason"`
		IDStage       string      `json:"IdStage"`
		IDGroup       string      `json:"IdGroup"`
		Weather       interface{} `json:"Weather"`
		Attendance    interface{} `json:"Attendance"`
		IDMatch       string      `json:"IdMatch"`
		MatchDay      interface{} `json:"MatchDay"`
		StageName     []struct {
			Locale      string `json:"Locale"`
			Description string `json:"Description"`
		} `json:"StageName"`
		GroupName []struct {
			Locale      string `json:"Locale"`
			Description string `json:"Description"`
		} `json:"GroupName"`
		CompetitionName []struct {
			Locale      string `json:"Locale"`
			Description string `json:"Description"`
		} `json:"CompetitionName"`
		Date      time.Time `json:"Date"`
		LocalDate time.Time `json:"LocalDate"`
		Home      struct {
			Score      interface{} `json:"Score"`
			Side       interface{} `json:"Side"`
			IDTeam     string      `json:"IdTeam"`
			PictureURL string      `json:"PictureUrl"`
			IDCountry  string      `json:"IdCountry"`
			Tactics    interface{} `json:"Tactics"`
			TeamType   int         `json:"TeamType"`
			AgeType    int         `json:"AgeType"`
			TeamName   []struct {
				Locale      string `json:"Locale"`
				Description string `json:"Description"`
			} `json:"TeamName"`
			FootballType int `json:"FootballType"`
			Gender       int `json:"Gender"`
		} `json:"Home"`
		Away struct {
			Score      interface{} `json:"Score"`
			Side       interface{} `json:"Side"`
			IDTeam     string      `json:"IdTeam"`
			PictureURL string      `json:"PictureUrl"`
			IDCountry  string      `json:"IdCountry"`
			Tactics    interface{} `json:"Tactics"`
			TeamType   int         `json:"TeamType"`
			AgeType    int         `json:"AgeType"`
			TeamName   []struct {
				Locale      string `json:"Locale"`
				Description string `json:"Description"`
			} `json:"TeamName"`
			FootballType int `json:"FootballType"`
			Gender       int `json:"Gender"`
		} `json:"Away"`
		HomeTeamScore          interface{} `json:"HomeTeamScore"`
		AwayTeamScore          interface{} `json:"AwayTeamScore"`
		AggregateHomeTeamScore interface{} `json:"AggregateHomeTeamScore"`
		AggregateAwayTeamScore interface{} `json:"AggregateAwayTeamScore"`
		HomeTeamPenaltyScore   interface{} `json:"HomeTeamPenaltyScore"`
		AwayTeamPenaltyScore   interface{} `json:"AwayTeamPenaltyScore"`
		LastPeriodUpdate       interface{} `json:"LastPeriodUpdate"`
		Leg                    interface{} `json:"Leg"`
		IsHomeMatch            interface{} `json:"IsHomeMatch"`
		Stadium                struct {
			IDStadium string `json:"IdStadium"`
			Name      []struct {
				Locale      string `json:"Locale"`
				Description string `json:"Description"`
			} `json:"Name"`
			IDCity   string `json:"IdCity"`
			CityName []struct {
				Locale      string `json:"Locale"`
				Description string `json:"Description"`
			} `json:"CityName"`
			IDCountry  string `json:"IdCountry"`
			Properties struct {
				IDCBS  string `json:"IdCBS"`
				IDIFES string `json:"IdIFES"`
			} `json:"Properties"`
			IsUpdateable interface{} `json:"IsUpdateable"`
		} `json:"Stadium"`
		MatchTime           interface{} `json:"MatchTime"`
		SecondHalfTime      interface{} `json:"SecondHalfTime"`
		FirstHalfTime       interface{} `json:"FirstHalfTime"`
		FirstHalfExtraTime  interface{} `json:"FirstHalfExtraTime"`
		SecondHalfExtraTime interface{} `json:"SecondHalfExtraTime"`
		Winner              interface{} `json:"Winner"`
		MatchReportURL      interface{} `json:"MatchReportUrl"`
		PlaceHolderA        string      `json:"PlaceHolderA"`
		PlaceHolderB        string      `json:"PlaceHolderB"`
		BallPossession      interface{} `json:"BallPossession"`
		Officials           []struct {
			IDCountry  string `json:"IdCountry"`
			OfficialID string `json:"OfficialId"`
			NameShort  []struct {
				Locale      string `json:"Locale"`
				Description string `json:"Description"`
			} `json:"NameShort"`
			Name []struct {
				Locale      string `json:"Locale"`
				Description string `json:"Description"`
			} `json:"Name"`
			OfficialType  int `json:"OfficialType"`
			TypeLocalized []struct {
				Locale      string `json:"Locale"`
				Description string `json:"Description"`
			} `json:"TypeLocalized"`
		} `json:"Officials"`
		MatchStatus       int  `json:"MatchStatus"`
		ResultType        int  `json:"ResultType"`
		MatchNumber       int  `json:"MatchNumber"`
		TimeDefined       bool `json:"TimeDefined"`
		OfficialityStatus int  `json:"OfficialityStatus"`
		Properties        struct {
			IDCBS  string `json:"IdCBS"`
			IDIFES string `json:"IdIFES"`
		} `json:"Properties"`
		IsUpdateable interface{} `json:"IsUpdateable"`
	} `json:"Results"`
}

func main() {
	// req client
	fifaapiclient := http.Client{
		Timeout: time.Second * 4,
	}
	req, err := http.NewRequest("GET", matchesapi, nil)
	if err != nil {
		log.Fatal(err)
	}
	// Basic headers
	req.Header.Set("User-Agent", "FIFA WorldCup Slack bot")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	// response
	res, getErr := fifaapiclient.Do(req)
	if getErr != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}
	var matches MatchesData
	//matches := MatchesData{}
	json.Unmarshal([]byte(body), &matches)

	for i, match := range matches.Results[0].GroupName {
		fmt.Println(i, match)
	}
	fmt.Println("Group Nme: ", matches.Results[1].GroupName[0].Description)
	fmt.Println("IDMatch: ", matches.Results[4].IDMatch)

}
