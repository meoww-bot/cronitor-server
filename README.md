# cronitor-server

It's a fake version of cronitor server, compatible to mod version cronitor-cli https://github.com/meoww-bot/cronitor-cli/ .

### Requirements

- DB: mongodb

### Develop progress


#### For web ui

- ✅ GET "/ping/:apiKey/:code", for `cronitor exec`
- ✅ GET "/api/monitors", get all monitors
- ✅ GET "/api/monitors/:monitor_code", get specific monitor
    - (WIP) withEvents=true, `latest_events` field
    - (WIP) withInvocations=true, `latest_invocations` field
- ❌ GET "/api/aggregates", aggregates for events stats 


#### For cronitor-cli

- ✅ PUT "/v3/monitors", for `cronitor discover`
- ✅ GET "/v3/monitors", for `cronitor status`
    - (WIP) STATUS field: "Completed x hours ago"
- ❌ GET "/v3/monitors/:monitor_code/activity", for `cronitor activity`. It's low priority, due to it output raw json data when running `cronitor activity` .

#### For user

No plan for now, cause we don't need it, we just need to monitor all crontab job.

- ❌ user module