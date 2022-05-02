package handler

import (
	"context"
	"cronitor-server/db"
	"cronitor-server/lib"
	"cronitor-server/util"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorhill/cronexpr"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetAllMonitors(c *gin.Context) {
	// /api/monitors

	page := c.DefaultQuery("page", "1")

	monitorsArray := make([]lib.Monitor, 0)

	monitorsArray, err := db.GetAllMonitors(page, monitorsArray)

	if Error(c, err) {
		return //exit
	}

	count, err := db.GetMonitorsCount()

	if Error(c, err) {
		fmt.Println("not found in monitor")
		return //exit
	}
	em := lib.ExpectedResponse{}

	em.PageSize = 50
	em.Page = page
	em.TotalMonitorCount = count
	em.Version = "2020-10-01"
	em.Monitors = monitorsArray

	c.JSON(http.StatusOK, em)

}

func GetMonitorActivityV3(c *gin.Context) {

}

func GetMonitorInfo(c *gin.Context) {

	var withInvocations bool
	var withEvents bool

	if c.Query("withInvocations") == "true" {
		withInvocations = true
	}

	if c.Query("withEvents") == "true" {
		withEvents = true
	}

	monitor_code := c.Param("monitor_code")

	client := db.Conn()

	MonitorCollection := client.Database("cronitor").Collection("Monitor")

	InvocationCollection := client.Database("cronitor").Collection("Invocation")

	monitor := new(lib.Monitor)

	err := MonitorCollection.FindOne(context.TODO(), bson.M{"code": monitor_code}).Decode(&monitor)

	if Error(c, err) {
		fmt.Println("not found in monitor")
		return //exit
	}

	inv := new(lib.Invocation)

	opts := options.FindOne()

	opts.SetSort(bson.D{{"series", -1}})

	filter := bson.M{"code": monitor_code, "state": "complete"}

	err = InvocationCollection.FindOne(context.TODO(), filter, opts).Decode(&inv)

	latest_event := new(lib.Event)

	if err != mongo.ErrNoDocuments {

		latest_event.Stamp = inv.Events[1].Stamp
		latest_event.Msg = inv.Events[1].Msg
		latest_event.State = inv.Events[1].State
		latest_event.Duration = inv.Events[1].Duration

	}

	nextTime := cronexpr.MustParse(monitor.Rules[0].Value).Next(time.Now())

	next_expected_at := nextTime.Unix()

	md := lib.MonitorDetail{
		Name:             monitor.Name,
		Key:              monitor.Code,
		Latest_event:     *latest_event,
		Note:             monitor.Note,
		Type:             monitor.Type,
		Timezone:         monitor.Timezone,
		Schedule:         monitor.Rules[0].Value,
		Next_expected_at: next_expected_at,
		Platform:         "Linux",
		Created:          monitor.Created.String(),
	}

	if withInvocations {

	}

	if withEvents {
		// complete_event := new(lib.Event)
		// fail_event := new(lib.Event)
		// run_event := new(lib.Event)

		// latest_events := map[string]lib.Event{
		// 	"complete": complete_event,
		// 	"fail":     fail_event,
		// 	"run":      run_event,
		// }

		// md.Latest_events = latest_events
		// Limit by 10 documents only
		// options.SetLimit(10)

	}

	c.JSON(http.StatusOK, md)

}

func GetMonitorsV3(c *gin.Context) {

	// auth := c.Request.Header.Get("Authorization")

	// if auth == "" {
	// 	c.String(http.StatusForbidden, "No Authorization header provided")
	// 	c.Abort()
	// 	return
	// }

	// token := strings.TrimPrefix(auth, "Basic ")

	// if token == auth {
	// 	c.String(http.StatusForbidden, "Could not find bearer token in Authorization header")
	// 	c.Abort()
	// 	return
	// }

	page := c.DefaultQuery("page", "1")

	monitorsArray := make([]lib.Monitor, 0)

	monitorsArray, err := db.GetAllMonitors(page, monitorsArray)

	if Error(c, err) {
		return //exit
	}

	count, err := db.GetMonitorsCount()

	em := lib.ExpectedResponse{}

	em.PageSize = 50
	em.Page = page
	em.TotalMonitorCount = count
	em.Version = "v3"
	em.Monitors = monitorsArray

	c.JSON(http.StatusOK, em)

}

func PutMonitorsV3(c *gin.Context) {

	monitors := []lib.Monitor{}

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.AbortWithError(400, err)
		return
	}

	err = json.Unmarshal(body, &monitors)
	if err != nil {
		c.AbortWithError(400, err)
		return
	}

	if len(monitors) > 0 {

		client := db.Conn()

		MonitorCollection := client.Database("cronitor").Collection("Monitor")

		for _, monitor := range monitors {

			result := new([]lib.Monitor)

			// count, _ := MonitorCollection.CountDocuments(context.TODO(), bson.M{"key": monitor.Key})
			err := MonitorCollection.FindOne(context.TODO(), bson.M{"key": monitor.Key}).Decode(&result)

			if err == mongo.ErrNoDocuments {

				monitor.Code = util.GenCode()
				monitor.Created = time.Now()
				monitor.Updated = time.Now()
				monitor.Passing = false
				monitor.Platform = "Linux cron"

				if monitor.Name == "" {
					monitor.Name = monitor.DefaultName
					monitor.DefaultName = ""
				}

				monitor.Note = monitor.DefaultNote
				monitor.DefaultNote = ""

				_, err := MonitorCollection.InsertOne(context.TODO(), monitor)

				if Error(c, err) {
					return //exit
				}

			} else {
				monitor.Updated = time.Now()
				updatedMonitor := bson.M{
					"note":         monitor.Note,
					"name":         monitor.Name,
					"rules":        monitor.Rules,
					"host":         monitor.Host,
					"commandtorun": monitor.CommandToRun,
					"runas":        monitor.RunAs,
					"updatedat":    time.Now(),
				}
				_, err = MonitorCollection.UpdateOne(context.TODO(), bson.M{"key": monitor.Key}, bson.M{"$set": updatedMonitor})

				if Error(c, err) {
					return //exit
				}

			}

		}
	}

	// response
	monitorsArray := make([]lib.Monitor, 0)

	page := c.DefaultQuery("page", "1")

	monitorsArray, err = db.GetAllMonitors(page, monitorsArray)

	monitorsArray = make([]lib.Monitor, 0)

	if Error(c, err) {
		return //exit
	}

	c.JSON(200, monitorsArray)
}

func Ping(c *gin.Context) {

	q := c.Request.URL.Query()

	api_key := c.Param("apiKey")
	code := c.Param("code")

	series := q.Get("series")
	state := q.Get("state")
	host := q.Get("host")
	ip := c.ClientIP()
	client := c.GetHeader("User-Agent")

	event := lib.Event{
		State: state,
		Msg:   q.Get("msg"),
	}

	event.Try, _ = strconv.Atoi(q.Get("try"))

	event.Status, _ = strconv.Atoi(q.Get("status_code"))

	event.Stamp, _ = strconv.ParseFloat(q.Get("stamp"), 64)

	if state != "run" {

		event.Duration, _ = strconv.ParseFloat(q.Get("duration"), 64)

	}

	// _, err := db.AddEvent(event)

	// if Error(c, err) {
	// 	return //exit
	// }

	if api_key != "" || code != "" {

		c.String(http.StatusOK, "{\"d\": null}")
	}

	if state == "run" {

		events := make([]lib.Event, 0)
		events = append(events, event)

		invocation := lib.Invocation{
			ApiKey:         api_key,
			Code:           code,
			Type:           "execution",
			Start:          event.Stamp,
			Localized_time: util.ConvStampToDateStr(event.Stamp),
			State:          "running",
			Series:         series,
			Ip:             ip,
			Host:           host,
			Client:         client,
			Events:         events,
		}

		_, err := db.AddInvocation(invocation)

		if Error(c, err) {
			return //exit
		}

		updateMonitorStatus := bson.M{
			"running": true,
		}
		_, err = db.UpdateMonitorStatus(code, updateMonitorStatus)

		if Error(c, err) {
			return //exit
		}

	}

	if state != "run" {

		var status string
		var passing bool

		if state == "complete" {
			state = "complete"
			status = "0"
			passing = true
		}

		if state == "fail" {
			state = "failed"
			status = "1"
			passing = false
		}

		runningInvocation := new(lib.Invocation)

		err := db.GetRunningInvocation(api_key, code, series, runningInvocation)

		if Error(c, err) {
			return //exit
		}

		events := runningInvocation.Events
		events = append(events, event)

		updateInvocation := bson.M{
			"end":      event.Stamp,
			"status":   status,
			"duration": event.Duration,
			"state":    state,
			"events":   events,
		}
		_, err = db.UpdateRunningInvocation(api_key, code, series, updateInvocation)

		if Error(c, err) {
			return //exit
		}

		// update monitor passing
		updateMonitorStatus := bson.M{
			"passing": passing,
		}
		_, err = db.UpdateMonitorStatus(code, updateMonitorStatus)

		if Error(c, err) {
			return //exit
		}
	}

}

func GetMonitorAggregates(c *gin.Context) {

	// q := c.Request.URL.Query()

	// fmt.Println(q)

	// fmt.Println(q["field"])

	// for _, v := range q["field"] {
	// 	fmt.Println(v)
	// }

}
