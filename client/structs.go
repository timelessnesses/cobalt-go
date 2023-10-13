package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"

	"github.com/schollz/progressbar/v3"
	"github.com/valyala/fastjson"
)

type Client struct {
	endpoint string
	http     http.Client
}

type Setting struct {
	Url             string
	VCodec          VideoCodec
	VQuality        VideoQuality
	AFormat         AudioFormat
	IsAudioOnly     bool
	IsNoTTWatermark bool
	IsTTFullAudio   bool
	IsAudioMuted    bool
	DubLang         bool
	DisableMetadata bool
}

func NewClient(endpoint string) Client {
	endpoint = strings.TrimSuffix(endpoint, "/")
	return Client{
		endpoint: endpoint,
		http:     http.Client{},
	}
}

type cleansed struct {
	Url             string `json:"url"`
	VCodec          string `json:"vCodec"`
	VQuality        string `json:"vQuality"`
	AFormat         string `json:"aFormat"`
	IsAudioOnly     bool   `json:"isAudioOnly"`
	IsNoTTWatermark bool   `json:"isNoTTWatermark"`
	IsTTFullAudio   bool   `json:"isTTFullAudio"`
	IsAudioMuted    bool   `json:"isAudioMuted"`
	DubLang         bool   `json:"dubLang"`
	DisableMetadata bool   `json:"disableMetadata"`
}

func stripped(options Setting) cleansed {
	return cleansed{
		Url:             options.Url,
		VCodec:          options.VCodec.vCodec,
		VQuality:        options.VQuality.vQuality,
		AFormat:         options.AFormat.format,
		IsAudioOnly:     options.IsAudioOnly,
		IsNoTTWatermark: options.IsNoTTWatermark,
		IsTTFullAudio:   options.IsTTFullAudio,
		IsAudioMuted:    options.IsAudioMuted,
		DubLang:         options.DubLang,
		DisableMetadata: options.DisableMetadata,
	}
}

type GetInfoResult struct {
	Status     string
	Text       string
	Url        string
	PickerType string
	Picker     []*fastjson.Value
	Audio      string
}

func (self *Client) GetInfo(options Setting) (GetInfoResult, error) {
	cleansified := stripped(options)
	jsonified, _ := json.Marshal(cleansified)
	byted := bytes.NewBuffer(jsonified)
	parsed, _ := url.Parse(self.endpoint + "/api/json")
	headers := http.Header{}
	headers.Add("Accept", "application/json")
	headers.Add("Content-Type", "application/json")
	things := http.Request{
		Method: "POST",
		URL:    parsed,
		Header: headers,
		Body:   io.NopCloser(byted),
	}
	resp, err := self.http.Do(
		&things,
	)
	if err != nil {
		return GetInfoResult{}, err
	}
	content, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return GetInfoResult{}, err
	}
	serialized, err := fastjson.ParseBytes(content)
	if err != nil {
		return GetInfoResult{}, err
	}
	return GetInfoResult{
		Status:     string(serialized.GetStringBytes("status")),
		Text:       string(serialized.GetStringBytes("text")),
		Url:        string(serialized.GetStringBytes("url")),
		PickerType: string(serialized.GetStringBytes("pickerType")),
		Picker:     serialized.GetArray("picker"),
		Audio:      string(serialized.GetStringBytes("audio")),
	}, nil
}

func (self *Client) Download(result GetInfoResult, path string) error {
	req, _ := http.NewRequest(
		"GET",
		result.Url,
		nil,
	)
	resp, err := self.http.Do(
		req,
	)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	name, err := get_file_name_from_header(resp)

	if err != nil {
		return err
	}

	f, err := os.OpenFile(name, os.O_CREATE|os.O_WRONLY, fs.ModeDevice)
	defer f.Close()

	bar := progressbar.DefaultBytes(
		resp.ContentLength,
		fmt.Sprintf("Downloading: %s", name),
	)

	io.Copy(io.MultiWriter(f, bar), resp.Body)
	return nil
}

func get_file_name_from_header(r *http.Response) (string, error) {
	name := r.Header.Get("Content-Disposition")
	re := regexp.MustCompile(`filename="([^"]+)"`)
	matched := re.FindStringSubmatch(name)
	if len(matched) == 2 {
		return matched[1], nil
	}
	return "", errors.New("Unreachable?")
}
