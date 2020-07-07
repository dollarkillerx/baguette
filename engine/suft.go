package engine

import (
	"encoding/json"
	"errors"
	"github.com/dollarkillerx/baguette/define"
	"github.com/gin-gonic/gin"
	"github.com/henrylee2cn/surfer"
	"io/ioutil"
	"log"
	"time"
)

type Suft struct{}

func (s *Suft) Get(ctx *gin.Context) {
	defer ctx.Request.Body.Close()
	all, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.JSON(500, errors.New("params error"))
		return
	}
	params := define.Params{}
	err = json.Unmarshal(all, &params)
	if err != nil {
		ctx.JSON(500, errors.New("params error"))
		return
	}

	resp, err := surfer.Download(&surfer.Request{
		Url: params.Url,
		ConnTimeout: time.Duration(params.TimeOut) * time.Second,
		DialTimeout: time.Duration(params.TimeOut) * time.Second,
		EnableCookie: params.Cookie,
	})
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		body = []byte{}
	}
	header, err := json.Marshal(resp.Header)
	if err != nil {
		log.Fatalln(err)
		header = []byte{}
	}


	ctx.JSON(200, define.Response{Body: string(body), Header: string(header), StatusCode: resp.StatusCode})
}
