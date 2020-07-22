package protocol

// NeteaseSongDetailResponseData is a response struct of NCM API: song detail
type NeteaseSongDetailResponseData struct {
	Songs []struct {
		Name string `json:"name"`
		ID   int    `json:"id"`
		Pst  int    `json:"pst"`
		T    int    `json:"t"`
		Ar   []struct {
			ID    int           `json:"id"`
			Name  string        `json:"name"`
			Tns   []interface{} `json:"tns"`
			Alias []interface{} `json:"alias"`
		} `json:"ar"`
		Alia []interface{} `json:"alia"`
		Pop  int           `json:"pop"`
		St   int           `json:"st"`
		Rt   string        `json:"rt"`
		Fee  int           `json:"fee"`
		V    int           `json:"v"`
		Crbt interface{}   `json:"crbt"`
		Cf   string        `json:"cf"`
		Al   struct {
			ID     int           `json:"id"`
			Name   string        `json:"name"`
			PicURL string        `json:"picUrl"`
			Tns    []interface{} `json:"tns"`
			Pic    int64         `json:"pic"`
		} `json:"al"`
		Dt int `json:"dt"`
		H  struct {
			Br   int `json:"br"`
			Fid  int `json:"fid"`
			Size int `json:"size"`
			Vd   int `json:"vd"`
		} `json:"h"`
		M struct {
			Br   int `json:"br"`
			Fid  int `json:"fid"`
			Size int `json:"size"`
			Vd   int `json:"vd"`
		} `json:"m"`
		L struct {
			Br   int `json:"br"`
			Fid  int `json:"fid"`
			Size int `json:"size"`
			Vd   int `json:"vd"`
		} `json:"l"`
		A               interface{}   `json:"a"`
		Cd              string        `json:"cd"`
		No              int           `json:"no"`
		RtURL           interface{}   `json:"rtUrl"`
		Ftype           int           `json:"ftype"`
		RtUrls          []interface{} `json:"rtUrls"`
		DjID            int           `json:"djId"`
		Copyright       int           `json:"copyright"`
		SID             int           `json:"s_id"`
		Mark            int           `json:"mark"`
		OriginCoverType int           `json:"originCoverType"`
		Single          int           `json:"single"`
		NoCopyrightRcmd interface{}   `json:"noCopyrightRcmd"`
		Mv              int           `json:"mv"`
		Mst             int           `json:"mst"`
		Cp              int           `json:"cp"`
		Rtype           int           `json:"rtype"`
		Rurl            interface{}   `json:"rurl"`
		PublishTime     int64         `json:"publishTime"`
	} `json:"songs"`
	Privileges []struct {
		ID             int  `json:"id"`
		Fee            int  `json:"fee"`
		Payed          int  `json:"payed"`
		St             int  `json:"st"`
		Pl             int  `json:"pl"`
		Dl             int  `json:"dl"`
		Sp             int  `json:"sp"`
		Cp             int  `json:"cp"`
		Subp           int  `json:"subp"`
		Cs             bool `json:"cs"`
		Maxbr          int  `json:"maxbr"`
		Fl             int  `json:"fl"`
		Toast          bool `json:"toast"`
		Flag           int  `json:"flag"`
		PreSell        bool `json:"preSell"`
		PlayMaxbr      int  `json:"playMaxbr"`
		DownloadMaxbr  int  `json:"downloadMaxbr"`
		ChargeInfoList []struct {
			Rate          int         `json:"rate"`
			ChargeURL     interface{} `json:"chargeUrl"`
			ChargeMessage interface{} `json:"chargeMessage"`
			ChargeType    int         `json:"chargeType"`
		} `json:"chargeInfoList"`
	} `json:"privileges"`
	Code int `json:"code"`
}

// NeteaseSongURLResponseData is a response struct of NCM API: song URLS
type NeteaseSongURLResponseData struct {
	Data []struct {
		ID            int         `json:"id"`
		URL           string      `json:"url"`
		Br            int         `json:"br"`
		Size          int         `json:"size"`
		Md5           string      `json:"md5"`
		Code          int         `json:"code"`
		Expi          int         `json:"expi"`
		Type          string      `json:"type"`
		Gain          int         `json:"gain"`
		Fee           int         `json:"fee"`
		Uf            interface{} `json:"uf"`
		Payed         int         `json:"payed"`
		Flag          int         `json:"flag"`
		CanExtend     bool        `json:"canExtend"`
		FreeTrialInfo interface{} `json:"freeTrialInfo"`
		Level         string      `json:"level"`
		EncodeType    string      `json:"encodeType"`
	} `json:"data"`
	Code int `json:"code"`
}
