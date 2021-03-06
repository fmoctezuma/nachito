type LivescoreData struct {
	Stadium struct {
		IDStadium string `json:"IdStadium"`
		Name      []struct {
			Locale      string `json:"Locale"`
			Description string `json:"Description"`
		} `json:"Name"`
		Capacity   interface{} `json:"Capacity"`
		WebAddress interface{} `json:"WebAddress"`
		Built      interface{} `json:"Built"`
		Roof       bool        `json:"Roof"`
		Turf       interface{} `json:"Turf"`
		IDCity     string      `json:"IdCity"`
		CityName   []struct {
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
	ResultType             int    `json:"ResultType"`
	MatchDay               string `json:"MatchDay"`
	IDMatch                string `json:"IdMatch"`
	IDStage                string `json:"IdStage"`
	IDGroup                string `json:"IdGroup"`
	IDSeason               string `json:"IdSeason"`
	IDCompetition          string `json:"IdCompetition"`
	HomeTeamPenaltyScore   int    `json:"HomeTeamPenaltyScore"`
	AwayTeamPenaltyScore   int    `json:"AwayTeamPenaltyScore"`
	AggregateHomeTeamScore int    `json:"AggregateHomeTeamScore"`
	AggregateAwayTeamScore int    `json:"AggregateAwayTeamScore"`
	Weather                struct {
		Humidity      interface{}   `json:"Humidity"`
		Temperature   interface{}   `json:"Temperature"`
		WindSpeed     interface{}   `json:"WindSpeed"`
		Type          interface{}   `json:"Type"`
		TypeLocalized []interface{} `json:"TypeLocalized"`
	} `json:"Weather"`
	Attendance          interface{} `json:"Attendance"`
	Date                time.Time   `json:"Date"`
	LocalDate           time.Time   `json:"LocalDate"`
	MatchTime           string      `json:"MatchTime"`
	SecondHalfTime      interface{} `json:"SecondHalfTime"`
	FirstHalfTime       interface{} `json:"FirstHalfTime"`
	FirstHalfExtraTime  interface{} `json:"FirstHalfExtraTime"`
	SecondHalfExtraTime interface{} `json:"SecondHalfExtraTime"`
	Winner              interface{} `json:"Winner"`
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
		Coaches       []interface{} `json:"Coaches"`
		Players       []interface{} `json:"Players"`
		Bookings      []interface{} `json:"Bookings"`
		Goals         []interface{} `json:"Goals"`
		Substitutions []interface{} `json:"Substitutions"`
		FootballType  int           `json:"FootballType"`
		Gender        int           `json:"Gender"`
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
		Coaches       []interface{} `json:"Coaches"`
		Players       []interface{} `json:"Players"`
		Bookings      []interface{} `json:"Bookings"`
		Goals         []interface{} `json:"Goals"`
		Substitutions []interface{} `json:"Substitutions"`
		FootballType  int           `json:"FootballType"`
		Gender        int           `json:"Gender"`
	} `json:"AwayTeam"`
	BallPossession struct {
		Intervals   []interface{} `json:"Intervals"`
		LastX       []interface{} `json:"LastX"`
		OverallHome interface{}   `json:"OverallHome"`
		OverallAway interface{}   `json:"OverallAway"`
	} `json:"BallPossession"`
	TerritorialPossesion      interface{}   `json:"TerritorialPossesion"`
	TerritorialThirdPossesion interface{}   `json:"TerritorialThirdPossesion"`
	Officials                 []interface{} `json:"Officials"`
	MatchStatus               int           `json:"MatchStatus"`
	GroupName                 []struct {
		Locale      string `json:"Locale"`
		Description string `json:"Description"`
	} `json:"GroupName"`
	StageName []struct {
		Locale      string `json:"Locale"`
		Description string `json:"Description"`
	} `json:"StageName"`
	OfficialityStatus int  `json:"OfficialityStatus"`
	TimeDefined       bool `json:"TimeDefined"`
	Properties        struct {
		IDCBS  string `json:"IdCBS"`
		IDIFES string `json:"IdIFES"`
	} `json:"Properties"`
	IsUpdateable interface{} `json:"IsUpdateable"`
}
