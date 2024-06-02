package logic

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"net/url"
)

func CreatedListAll(c *gin.Context) {
	upMid := c.Query("up_mid")
	targetUrl := "https://api.bilibili.com/x/v3/fav/folder/created/list-all"
	u, _ := url.ParseRequestURI(targetUrl)

	data := url.Values{}
	data.Set("up_mid", upMid)

	u.RawQuery = data.Encode()

	resp, err := http.Get(u.String())
	if err != nil {
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	dataBackup := make(map[string]any)
	err = json.Unmarshal(body, &dataBackup)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": dataBackup["data"],
		"msg":  "success",
	})
}

func ResourceList(c *gin.Context) {
	targetUrl := "https://api.bilibili.com/x/v3/fav/resource/list"
	u, _ := url.ParseRequestURI(targetUrl)

	data := url.Values{}
	data.Set("media_id", c.Query("media_id"))
	data.Set("pn", c.Query("pn"))
	data.Set("ps", c.Query("ps"))
	data.Set("keyword", c.Query("keyword"))
	data.Set("order", c.Query("order"))
	data.Set("type", c.Query("type"))
	data.Set("tid", c.Query("tid"))
	data.Set("platform", c.Query("platform"))

	u.RawQuery = data.Encode()

	resp, err := http.Get(u.String())
	if err != nil {
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	dataBackup := make(map[string]any)
	err = json.Unmarshal(body, &dataBackup)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": dataBackup["data"],
		"msg":  "success",
	})
}

func CollectedList(c *gin.Context) {
	targetUrl := "https://api.bilibili.com/x/v3/fav/folder/collected/list"
	u, _ := url.ParseRequestURI(targetUrl)

	data := url.Values{}
	data.Set("pn", c.Query("pn"))
	data.Set("ps", c.Query("ps"))
	data.Set("up_mid", c.Query("up_mid"))
	data.Set("platform", c.Query("platform"))

	u.RawQuery = data.Encode()

	resp, err := http.Get(u.String())
	if err != nil {
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	dataBackup := make(map[string]any)
	err = json.Unmarshal(body, &dataBackup)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": dataBackup["data"],
		"msg":  "success",
	})
}
