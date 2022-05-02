package lib

import "time"

// type RuleValue string

type Rule struct {
	RuleType     string `json:"rule_type"`
	Value        string `json:"value"`
	TimeUnit     string `json:"time_unit,omitempty"`
	GraceSeconds uint   `json:"grace_seconds,omitempty"`
}

type Monitor struct {
	Name             string              `json:"name,omitempty"`
	DefaultName      string              `json:"defaultName,omitempty" bson:"defaultname,omitempty"`
	Host             string              `json:"host"`
	CommandToRun     string              `json:"commandToRun"`
	RunAs            string              `json:"runAs"`
	Key              string              `json:"key"`
	Rules            []Rule              `json:"rules"`
	Tags             []string            `json:"tags"`
	Type             string              `json:"type"`
	Code             string              `json:"code,omitempty"`
	Timezone         string              `json:"timezone,omitempty"`
	Note             string              `json:"note,omitempty" bson:"note,omitempty"`
	Passing          bool                `json:"passing"` // status: Healthy or Failing
	Paused           bool                `json:"paused"`
	Platform         string              `json:"platform"`
	Running          bool                `json:"running"`
	DefaultNote      string              `json:"defaultNote,omitempty" bson:"defaultNote,omitempty"`
	Notifications    map[string][]string `json:"notifications,omitempty"`
	NoStdoutPassthru bool                `json:"-"`
	Created          time.Time           `json:"created,omitempty"`
	Updated          time.Time           `json:"updated,omitempty"`
}

type Latest_events struct {
	Run      Event `json:"run"`
	Complete Event `json:"complete,omitempty"`
	Fail     Event `json:"fail,omitempty"`
}

type MonitorDetail struct {
	Created            string        `json:"created"`
	Disabled           bool          `json:"disabled"`
	Key                string        `json:"key"`
	Latest_event       Event         `json:"latest_event"`  // lastest complete or fail event
	Latest_events      Latest_events `json:"latest_events"` // complete/fail,run
	Latest_incident    Incident      `json:"latest_incident,omitempty"`
	Latest_invocations []Invocation  `json:"latest_invocations"`
	Name               string        `json:"name"`
	Next_expected_at   int64         `json:"next_expected_at"`
	Note               string        `json:"note"`
	Passing            bool          `json:"passing"` // status: Healthy or Failing
	Paused             bool          `json:"paused"`
	Platform           string        `json:"platform"`
	Running            bool          `json:"running"`
	Schedule           string        `json:"schedule"`
	Timezone           string        `json:"timezone"`
	Type               string        `json:"type"` // job
	// Realert_interval   string           `json:"latest_invocations"`

}

type Incident struct {
	Stamp float64 `json:"stamp,omitempty" bson:"stamp,omitempty"`
	State string  `json:"state,omitempty" bson:"state,omitempty"` // open
}

type Event struct {
	ApiKey   string  `json:"api_key,omitempty" bson:"api_key,omitempty"`
	Code     string  `json:"code,omitempty" bson:"code,omitempty"`
	State    string  `json:"state,omitempty" bson:"state,omitempty"`
	Try      int     `json:"try,omitempty" bson:"try,omitempty"`
	Stamp    float64 `json:"stamp,omitempty" bson:"stamp,omitempty"`
	Msg      string  `json:"msg,omitempty" bson:"msg,omitempty"`
	Host     string  `json:"host,omitempty" bson:"host,omitempty"`
	Ip       string  `json:"ip,omitempty" bson:"ip,omitempty"`
	Duration float64 `json:"duration,omitempty" bson:"duration,omitempty"`
	Series   string  `json:"series,omitempty" bson:"series,omitempty"`
	Status   int     `json:"status,omitempty" bson:"status,omitempty"` // status_code
	Client   string  `json:"client,omitempty" bson:"client,omitempty"`
	// Pinged   time.Time `json:"pinged,omitempty" bson:"pinged,omitempty"`
}

type Invocation struct {
	ApiKey         string  `json:"api_key,omitempty" bson:"api_key,omitempty"`
	Code           string  `json:"code,omitempty" bson:"code,omitempty"`
	Type           string  `json:"type,omitempty" bson:"type,omitempty"` // execution
	Start          float64 `json:"start,omitempty" bson:"start,omitempty"`
	End            float64 `json:"end,omitempty" bson:"end,omitempty"`
	Localized_time string  `json:"localized_time,omitempty" bson:"localized_time,omitempty"`
	State          string  `json:"state,omitempty" bson:"state,omitempty"`
	Series         string  `json:"series,omitempty" bson:"series,omitempty"`
	Ip             string  `json:"ip,omitempty" bson:"ip,omitempty"`
	Host           string  `json:"host,omitempty" bson:"host,omitempty"`
	Status         string  `json:"status,omitempty" bson:"status,omitempty"`
	Client         string  `json:"client,omitempty" bson:"client,omitempty"`
	Duration       float64 `json:"duration,omitempty" bson:"duration,omitempty"`
	Events         []Event `json:"events,omitempty" bson:"events,omitempty"` // stamp = start,end

}

// type Metrics struct {
// 	Duration float32 `json:"duration"`
// }

type MonitorSummary struct {
	Name        string `json:"name,omitempty"`
	DefaultName string `json:"defaultName"`
	Key         string `json:"key"`
	Code        string `json:"code,omitempty"`
}

type CronitorApi struct {
	IsDev          bool
	IsAutoDiscover bool
	ApiKey         string
	UserAgent      string
	Logger         func(string)
}

type Line struct {
	Name           string
	FullLine       string
	LineNumber     int
	CronExpression string
	CommandToRun   string
	Code           string
	RunAs          string
	Mon            Monitor
}

type MonitorAggregates struct {
	Duration_mean   float64 `json:"duration_mean"`
	Event_count     int     `json:"event_count"`
	Run_count       int     `json:"run_count"`
	Fail_count      int     `json:"fail_count"`
	Complete_count  int     `json:"complete_count"`
	Uptime          int     `json:"uptime"`
	Alert_count     int     `json:"alert_count"`
	Count_sum       int     `json:"count_sum"`
	Error_count_sum int     `json:"error_count_sum"`
}

type Crontab struct {
	User                    string
	IsUserCrontab           bool
	IsSaved                 bool
	Filename                string
	Lines                   []*Line
	TimezoneLocationName    *TimezoneLocationName
	UsesSixFieldExpressions bool
}

type TimezoneLocationName struct {
	Name string
}

type ExpectedResponse struct {
	TotalMonitorCount int64     `json:"total_monitor_count"`
	PageSize          int       `json:"page_size"`
	Page              string    `json:"page"`
	Version           string    `json:"version"`
	Monitors          []Monitor `json:"monitors"`
}

type StatusMonitor struct {
	Name    string `json:"name"`
	Code    string `json:"code"`
	Passing bool   `json:"passing"`
	Status  string `json:"status"`
}

type StatusMonitors struct {
	Monitors []StatusMonitor `json:"monitors"`
}

type ExistingMonitors struct {
	Monitors    []Monitor
	Names       []string
	CurrentKey  string
	CurrentCode string
}
