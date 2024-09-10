package entity

type TwitchUserResponse struct {
	Data []*TwitchUserInfo `json:"data"`
}

type TwitchUserInfo struct {
	Id                string `json:"id"`
	Login             string `json:"login"`
	Display_name      string `json:"display_name"`
	Type              string `json:"type"`
	Broadcaster_type  string `json:"broadcaster_type"`
	Description       string `json:"description"`
	Profile_image_url string `json:"profile_image_url"`
	Offline_image_url string `json:"offline_image_url"`
	View_count        int    `json:"view_count"`
	Email             string `json:"email"`
	Created_at        string `json:"created_at"`
}

type TwitchStreamScheduleResponse struct {
	Data *TwitchStreamScheduleInfo `json:"data"`
}

type TwitchStreamScheduleInfo struct {
	Segments          []*TwitchStreamScheduleSegment `json:"segments"`
	Broadcaster_id    string                         `json:"broadcaster_id"`
	Broadcaster_name  string                         `json:"broadcaster_name"`
	Broadcaster_login string                         `json:"broadcaster_login"`
	Vacation          *TwitchStreamerVacation        `json:"vacation"`
	Pagination        *TwitchPagination              `json:"pagination"`
}

type TwitchStreamScheduleSegment struct {
	Id             string                        `json:"id"`
	Start_time     string                        `json:"start_time"`
	End_time       string                        `json:"end_time"`
	Title          string                        `json:"title"`
	Canceled_until string                        `json:"canceled_until"`
	Category       *TwitchStreamScheduleCategory `json:"category"`
	Is_recurring   bool                          `json:"is_recurring"`
}

type TwitchStreamScheduleCategory struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type TwitchStreamerVacation struct {
	Start_time string `json:"start_time"`
	End_time   string `json:"end_time"`
}

type TwitchPagination struct {
	Cursor string `json:"cursor"`
}
