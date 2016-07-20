package api

import (
	"github.com/leancloud/lean-cli/lean/apps"
)

// GetAppListResult is GetAppList function's result type
type GetAppListResult struct {
	AppID     string `json:"app_id"`
	AppKey    string `json:"app_key"`
	AppName   string `json:"app_name"`
	MasterKey string `json:"master_key"`
	AppDomain string `json:"app_domain"`
}

// GetAppList returns the current user's all LeanCloud application
func GetAppList() ([]*GetAppListResult, error) {
	client := NewClient()

	resp, err := client.getX("/1/clients/self/apps", nil)
	if err != nil {
		return nil, err
	}

	var result []*GetAppListResult
	err = resp.JSON(&result)
	return result, err
}

// DeployAppFromGit will deploy applications with user's git repo
// returns the event token for polling deploy log
func DeployAppFromGit(projectPath string, groupName string) (string, error) {
	client := NewClient()

	appID, err := apps.GetCurrentAppID(projectPath)
	if err != nil {
		return "", err
	}

	opts, err := client.options()
	if err != nil {
		return "", err
	}
	opts.Headers["X-LC-Id"] = appID

	resp, err := client.post("/1.1/functions/_ops/groups/"+groupName+"/buildAndDeploy", map[string]interface{}{
		"comment":             "",
		"noDependenciesCache": false,
		"async":               true,
	}, opts)

	if err != nil {
		return "", err
	}

	return resp.Get("eventToken").MustString(), nil
}

// DeployAppFromFile will deploy applications with specific file
// returns the event token for polling deploy log
func DeployAppFromFile(projectPath string, groupName string, fileURL string) (string, error) {
	client := NewClient()

	appID, err := apps.GetCurrentAppID(projectPath)
	if err != nil {
		return "", err
	}

	opts, err := client.options()
	if err != nil {
		return "", err
	}
	opts.Headers["X-LC-Id"] = appID

	resp, err := client.post("/1.1/functions/_ops/groups/"+groupName+"/buildAndDeploy", map[string]interface{}{
		"zipUrl":              fileURL,
		"comment":             "",
		"noDependenciesCache": false,
		"async":               true,
	}, opts)

	if err != nil {
		return "", err
	}

	return resp.Get("eventToken").MustString(), nil

}

// GetAppInfoResult is GetAppInfo function's result type
type GetAppInfoResult struct {
	AppID          string `json:"app_id"`
	AppKey         string `json:"app_key"`
	AppName        string `json:"app_name"`
	MasterKey      string `json:"master_key"`
	AppDomain      string `json:"app_domain"`
	LeanEngineMode string `json:"leanengine_mode"`
}

// GetAppInfo returns the application's detail info
func GetAppInfo(appID string) (*GetAppInfoResult, error) {
	client := NewClient()
	resp, err := client.getX("/1.1/clients/self/apps/"+appID, nil)
	if err != nil {
		return nil, err
	}
	result := new(GetAppInfoResult)
	err = resp.JSON(result)
	return result, err
}

// GetGroupsResult is GetGroups's result struct
type GetGroupsResult struct {
	GroupName string `json:"groupName"`
	Prod      int    `json:"prod"`
	Quota     int    `json:"quota"`
}

// GetGroups returns the application's engine groups
func GetGroups(appID string) ([]*GetGroupsResult, error) {
	client := NewClient()
	opts, err := client.options()
	if err != nil {
		return nil, err
	}
	opts.Headers["X-LC-Id"] = appID

	resp, err := client.getX("/1.1/functions/_ops/groups", opts)

	var result []*GetGroupsResult
	err = resp.JSON(&result)

	return result, err
}