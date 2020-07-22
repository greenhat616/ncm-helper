package protocol

// NCMCreatePlaylistResponseData is the json struct of response data, performing a request that create a playlist
type NCMCreatePlaylistResponseData struct {
	ID       int64 `json:"id"`
	Code     int   `json:"code"`
	Playlist struct {
		Subscribers           []interface{} `json:"subscribers"`
		Subscribed            interface{}   `json:"subscribed"`
		Creator               interface{}   `json:"creator"`
		Artists               interface{}   `json:"artists"`
		Tracks                interface{}   `json:"tracks"`
		UpdateFrequency       interface{}   `json:"updateFrequency"`
		BackgroundCoverID     int           `json:"backgroundCoverId"`
		BackgroundCoverURL    interface{}   `json:"backgroundCoverUrl"`
		TitleImage            int           `json:"titleImage"`
		TitleImageURL         interface{}   `json:"titleImageUrl"`
		EnglishTitle          interface{}   `json:"englishTitle"`
		OpRecommend           bool          `json:"opRecommend"`
		RecommendInfo         interface{}   `json:"recommendInfo"`
		AdType                int           `json:"adType"`
		TrackNumberUpdateTime int           `json:"trackNumberUpdateTime"`
		SubscribedCount       int           `json:"subscribedCount"`
		CloudTrackCount       int           `json:"cloudTrackCount"`
		UserID                int           `json:"userId"`
		CreateTime            int64         `json:"createTime"`
		HighQuality           bool          `json:"highQuality"`
		CoverImgID            int64         `json:"coverImgId"`
		NewImported           bool          `json:"newImported"`
		Anonimous             bool          `json:"anonimous"`
		UpdateTime            int64         `json:"updateTime"`
		SpecialType           int           `json:"specialType"`
		CommentThreadID       string        `json:"commentThreadId"`
		CoverImgURL           string        `json:"coverImgUrl"`
		TotalDuration         int           `json:"totalDuration"`
		Privacy               int           `json:"privacy"`
		TrackUpdateTime       int           `json:"trackUpdateTime"`
		TrackCount            int           `json:"trackCount"`
		PlayCount             int           `json:"playCount"`
		Ordered               bool          `json:"ordered"`
		Tags                  []interface{} `json:"tags"`
		Description           interface{}   `json:"description"`
		Status                int           `json:"status"`
		Name                  string        `json:"name"`
		ID                    int64         `json:"id"`
	} `json:"playlist"`
}
