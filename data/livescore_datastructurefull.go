type livescore struct {
	ContinuationToken string `json:"ContinuationToken"`
	ContinuationHash  string `json:"ContinuationHash"`
	Results           []struct {
		Stadium struct {
			IDStadium interface{} `json:"IdStadium"`
			Name      []struct {
				Locale      string `json:"Locale"`
				Description string `json:"Description"`
			} `json:"Name"`
			Capacity           interface{}   `json:"Capacity"`
			WebAddress         interface{}   `json:"WebAddress"`
			Built              interface{}   `json:"Built"`
			Roof               bool          `json:"Roof"`
			Turf               interface{}   `json:"Turf"`
			IDCity             interface{}   `json:"IdCity"`
			CityName           []interface{} `json:"CityName"`
			IDCountry          string        `json:"IdCountry"`
			PostalCode         interface{}   `json:"PostalCode"`
			Street             interface{}   `json:"Street"`
			Email              interface{}   `json:"Email"`
			Fax                interface{}   `json:"Fax"`
			Phone              interface{}   `json:"Phone"`
			AffiliationCountry interface{}   `json:"AffiliationCountry"`
			AffiliationRegion  interface{}   `json:"AffiliationRegion"`
			Latitude           interface{}   `json:"Latitude"`
			Longitude          interface{}   `json:"Longitude"`
			Length             interface{}   `json:"Length"`
			Width              interface{}   `json:"Width"`
			Properties         interface{}   `json:"Properties"`
			IsUpdateable       interface{}   `json:"IsUpdateable"`
		} `json:"Stadium"`
		ResultType             int         `json:"ResultType"`
		MatchDay               string      `json:"MatchDay"`
		IDMatch                string      `json:"IdMatch"`
		IDStage                string      `json:"IdStage"`
		IDGroup                interface{} `json:"IdGroup"`
		IDSeason               string      `json:"IdSeason"`
		IDCompetition          string      `json:"IdCompetition"`
		HomeTeamPenaltyScore   interface{} `json:"HomeTeamPenaltyScore"`
		AwayTeamPenaltyScore   interface{} `json:"AwayTeamPenaltyScore"`
		AggregateHomeTeamScore interface{} `json:"AggregateHomeTeamScore"`
		AggregateAwayTeamScore interface{} `json:"AggregateAwayTeamScore"`
		Weather                struct {
			Humidity      interface{} `json:"Humidity"`
			Temperature   interface{} `json:"Temperature"`
			WindSpeed     interface{} `json:"WindSpeed"`
			Type          int         `json:"Type"`
			TypeLocalized []struct {
				Locale      string `json:"Locale"`
				Description string `json:"Description"`
			} `json:"TypeLocalized"`
		} `json:"Weather"`
		Attendance          string      `json:"Attendance"`
		Date                time.Time   `json:"Date"`
		LocalDate           time.Time   `json:"LocalDate"`
		MatchTime           interface{} `json:"MatchTime"`
		SecondHalfTime      interface{} `json:"SecondHalfTime"`
		FirstHalfTime       interface{} `json:"FirstHalfTime"`
		FirstHalfExtraTime  int         `json:"FirstHalfExtraTime"`
		SecondHalfExtraTime int         `json:"SecondHalfExtraTime"`
		Winner              string      `json:"Winner"`
		Period              int         `json:"Period"`
		HomeTeam            struct {
			Score      int         `json:"Score"`
			Side       interface{} `json:"Side"`
			IDTeam     string      `json:"IdTeam"`
			PictureURL string      `json:"PictureUrl"`
			IDCountry  string      `json:"IdCountry"`
			TeamType   int         `json:"TeamType"`
			AgeType    int         `json:"AgeType"`
			Tactics    string      `json:"Tactics"`
			TeamName   []struct {
				Locale      string `json:"Locale"`
				Description string `json:"Description"`
			} `json:"TeamName"`
			Coaches []interface{} `json:"Coaches"`
			Players []struct {
				IDPlayer      string      `json:"IdPlayer"`
				IDTeam        string      `json:"IdTeam"`
				ShirtNumber   int         `json:"ShirtNumber"`
				Status        int         `json:"Status"`
				SpecialStatus interface{} `json:"SpecialStatus"`
				Captain       bool        `json:"Captain"`
				PlayerName    []struct {
					Locale      string `json:"Locale"`
					Description string `json:"Description"`
				} `json:"PlayerName"`
				ShortName []struct {
					Locale      string `json:"Locale"`
					Description string `json:"Description"`
				} `json:"ShortName"`
				Position    int         `json:"Position"`
				FieldStatus int         `json:"FieldStatus"`
				LineupX     interface{} `json:"LineupX"`
				LineupY     interface{} `json:"LineupY"`
			} `json:"Players"`
			Bookings []struct {
				Card        int         `json:"Card"`
				Period      int         `json:"Period"`
				IDEvent     interface{} `json:"IdEvent"`
				EventNumber interface{} `json:"EventNumber"`
				IDPlayer    string      `json:"IdPlayer"`
				IDTeam      string      `json:"IdTeam"`
				Minute      string      `json:"Minute"`
				Reason      string      `json:"Reason"`
			} `json:"Bookings"`
			Goals []struct {
				Type           int         `json:"Type"`
				IDPlayer       string      `json:"IdPlayer"`
				Minute         string      `json:"Minute"`
				IDAssistPlayer string      `json:"IdAssistPlayer"`
				Period         int         `json:"Period"`
				IDGoal         interface{} `json:"IdGoal"`
				IDTeam         string      `json:"IdTeam"`
			} `json:"Goals"`
			Substitutions []struct {
				IDEvent            interface{} `json:"IdEvent"`
				Period             int         `json:"Period"`
				Reason             int         `json:"Reason"`
				SubstitutePosition int         `json:"SubstitutePosition"`
				IDPlayerOff        string      `json:"IdPlayerOff"`
				IDPlayerOn         string      `json:"IdPlayerOn"`
				PlayerOffName      []struct {
					Locale      string `json:"Locale"`
					Description string `json:"Description"`
				} `json:"PlayerOffName"`
				PlayerOnName []struct {
					Locale      string `json:"Locale"`
					Description string `json:"Description"`
				} `json:"PlayerOnName"`
				Minute string `json:"Minute"`
				IDTeam string `json:"IdTeam"`
			} `json:"Substitutions"`
			FootballType int `json:"FootballType"`
			Gender       int `json:"Gender"`
		} `json:"HomeTeam"`
		AwayTeam struct {
			Score      int         `json:"Score"`
			Side       interface{} `json:"Side"`
			IDTeam     string      `json:"IdTeam"`
			PictureURL string      `json:"PictureUrl"`
			IDCountry  string      `json:"IdCountry"`
			TeamType   int         `json:"TeamType"`
			AgeType    int         `json:"AgeType"`
			Tactics    string      `json:"Tactics"`
			TeamName   []struct {
				Locale      string `json:"Locale"`
				Description string `json:"Description"`
			} `json:"TeamName"`
			Coaches []struct {
				IDCoach   string `json:"IdCoach"`
				IDCountry string `json:"IdCountry"`
				Name      []struct {
					Locale      string `json:"Locale"`
					Description string `json:"Description"`
				} `json:"Name"`
				Alias []struct {
					Locale      string `json:"Locale"`
					Description string `json:"Description"`
				} `json:"Alias"`
				Role int `json:"Role"`
			} `json:"Coaches"`
			Players []struct {
				IDPlayer      string      `json:"IdPlayer"`
				IDTeam        string      `json:"IdTeam"`
				ShirtNumber   int         `json:"ShirtNumber"`
				Status        int         `json:"Status"`
				SpecialStatus interface{} `json:"SpecialStatus"`
				Captain       bool        `json:"Captain"`
				PlayerName    []struct {
					Locale      string `json:"Locale"`
					Description string `json:"Description"`
				} `json:"PlayerName"`
				ShortName []struct {
					Locale      string `json:"Locale"`
					Description string `json:"Description"`
				} `json:"ShortName"`
				Position    int         `json:"Position"`
				FieldStatus int         `json:"FieldStatus"`
				LineupX     interface{} `json:"LineupX"`
				LineupY     interface{} `json:"LineupY"`
			} `json:"Players"`
			Bookings []struct {
				Card        int         `json:"Card"`
				Period      int         `json:"Period"`
				IDEvent     interface{} `json:"IdEvent"`
				EventNumber interface{} `json:"EventNumber"`
				IDPlayer    string      `json:"IdPlayer"`
				IDTeam      string      `json:"IdTeam"`
				Minute      string      `json:"Minute"`
				Reason      string      `json:"Reason"`
			} `json:"Bookings"`
			Goals         []interface{} `json:"Goals"`
			Substitutions []struct {
				IDEvent            interface{} `json:"IdEvent"`
				Period             int         `json:"Period"`
				Reason             int         `json:"Reason"`
				SubstitutePosition int         `json:"SubstitutePosition"`
				IDPlayerOff        string      `json:"IdPlayerOff"`
				IDPlayerOn         string      `json:"IdPlayerOn"`
				PlayerOffName      []struct {
					Locale      string `json:"Locale"`
					Description string `json:"Description"`
				} `json:"PlayerOffName"`
				PlayerOnName []struct {
					Locale      string `json:"Locale"`
					Description string `json:"Description"`
				} `json:"PlayerOnName"`
				Minute string `json:"Minute"`
				IDTeam string `json:"IdTeam"`
			} `json:"Substitutions"`
			FootballType int `json:"FootballType"`
			Gender       int `json:"Gender"`
		} `json:"AwayTeam"`
		BallPossession struct {
			Intervals   []interface{} `json:"Intervals"`
			LastX       []interface{} `json:"LastX"`
			OverallHome int           `json:"OverallHome"`
			OverallAway int           `json:"OverallAway"`
		} `json:"BallPossession"`
		TerritorialPossesion      interface{} `json:"TerritorialPossesion"`
		TerritorialThirdPossesion interface{} `json:"TerritorialThirdPossesion"`
		Officials                 []struct {
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
		MatchStatus int           `json:"MatchStatus"`
		GroupName   []interface{} `json:"GroupName"`
		StageName   []struct {
			Locale      string `json:"Locale"`
			Description string `json:"Description"`
		} `json:"StageName"`
		OfficialityStatus int  `json:"OfficialityStatus"`
		TimeDefined       bool `json:"TimeDefined"`
		Properties        struct {
			IDCBS        string `json:"IdCBS"`
			IDInfostrada string `json:"IdInfostrada"`
		} `json:"Properties"`
		IsUpdateable interface{} `json:"IsUpdateable"`
	} `json:"Results"`
}
