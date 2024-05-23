package client3xui

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
)

type PanelSettings struct {
	WebListen        string `json:"webListen"`
	WebDomain        string `json:"webDomain"`
	WebPort          int    `json:"webPort"`
	WebCertFile      string `json:"webCertFile"`
	WebKeyFile       string `json:"webKeyFile"`
	WebBasePath      string `json:"webBasePath"`
	SessionMaxAge    int    `json:"sessionMaxAge"`
	PageSize         int    `json:"pageSize"`
	ExpireDiff       int    `json:"expireDiff"`
	TrafficDiff      int    `json:"trafficDiff"`
	RemarkModel      string `json:"remarkModel"`
	TgBotEnable      bool   `json:"tgBotEnable"`
	TgBotToken       string `json:"tgBotToken"`
	TgBotProxy       string `json:"tgBotProxy"`
	TgBotChatId      string `json:"tgBotChatId"`
	TgRunTime        string `json:"tgRunTime"`
	TgBotBackup      bool   `json:"tgBotBackup"`
	TgBotLoginNotify bool   `json:"tgBotLoginNotify"`
	TgCpu            int    `json:"tgCpu"`
	TgLang           string `json:"tgLang"`
	TimeLocation     string `json:"timeLocation"`
	SecretEnable     bool   `json:"secretEnable"`
	SubEnable        bool   `json:"subEnable"`
	SubListen        string `json:"subListen"`
	SubPort          int    `json:"subPort"`
	SubPath          string `json:"subPath"`
	SubDomain        string `json:"subDomain"`
	SubCertFile      string `json:"subCertFile"`
	SubKeyFile       string `json:"subKeyFile"`
	SubUpdates       int    `json:"subUpdates"`
	SubEncrypt       bool   `json:"subEncrypt"`
	SubShowInfo      bool   `json:"subShowInfo"`
	SubURI           string `json:"subURI"`
	SubJsonPath      string `json:"subJsonPath"`
	SubJsonURI       string `json:"subJsonURI"`
	SubJsonFragment  string `json:"subJsonFragment"`
	SubJsonMux       string `json:"subJsonMux"`
	SubJsonRules     string `json:"subJsonRules"`
	Datepicker       string `json:"datepicker"`
}

type PanelSettingsResponse struct {
	Success bool           `json:"success"`
	Msg     string         `json:"msg"`
	Obj     *PanelSettings `json:"obj"`
}

func (c *Client) GetPanelSettings(ctx context.Context) (*PanelSettings, error) {
	resp := &PanelSettingsResponse{}
	err := c.Do(ctx, http.MethodPost, "/panel/setting/all", nil, resp)
	return resp.Obj, err
}

func panelSettingsToMap(obj PanelSettings) map[string]string {
	result := make(map[string]string)
	v := reflect.ValueOf(obj)
	t := reflect.TypeOf(obj)

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		key := t.Field(i).Tag.Get("json")

		var value string
		switch field.Kind() {
		case reflect.String:
			value = field.String()
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			value = strconv.FormatInt(field.Int(), 10)
		case reflect.Bool:
			value = strconv.FormatBool(field.Bool())
		case reflect.Float32, reflect.Float64:
			value = strconv.FormatFloat(field.Float(), 'f', -1, 64)
		default:
			value = fmt.Sprintf("%v", field.Interface())
		}

		result[key] = value
	}

	return result
}

func (c *Client) EditPanelSettings(ctx context.Context, settings PanelSettings) error {
	settingsMap := panelSettingsToMap(settings)
	formValues := url.Values{}
	for k, v := range settingsMap {
		formValues.Add(k, v)
	}
	resp, err := c.DoRaw(ctx, http.MethodPost, c.url, "/panel/setting/update",
		"application/x-www-form-urlencoded", []byte(formValues.Encode()))
	if err != nil {
		return err
	}
	var genericResp ApiResponse
	if err := json.Unmarshal(resp, &genericResp); err != nil {
		return err
	}
	if !genericResp.Success {
		return fmt.Errorf("Failed to edit panel settings")
	}
	return nil
}

func (c *Client) RestartPanel(ctx context.Context) (*ApiResponse, error) {
	resp := &ApiResponse{}
	err := c.Do(ctx, http.MethodPost, "/panel/setting/restartPanel", nil, resp)
	return resp, err
}
