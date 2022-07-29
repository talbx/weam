package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/talbx/weam/internal/pkg/service"
	"github.com/talbx/weam/internal/pkg/util"
)

func main() {

	var AppConf util.AppConfig
	util.ReadConfig(&AppConf)

	r := gin.Default()
	r.POST("/notify", func(c *gin.Context) {
		jsonData, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			panic(err)
		}
		var payload util.Payload
		c.Bind(&payload)
		json.Unmarshal(jsonData, &payload)
		var notifier = service.WebexNotifier{}

		mapper := service.PayloadMapper{}
		mds := mapper.Apply(payload)

		forEach(mds, func(md util.Markdown) {

			webex := util.Webex{
				RoomId:   AppConf.Receiver,
				Markdown: md,
			}

			notifier.Notify(webex, &AppConf)
		})

		c.JSON(http.StatusAccepted, gin.H{
			"msg": "success",
		})
	})
	r.Run(":2000")
}

func forEach(mds []util.Markdown, fun func(md util.Markdown)) {
	for _, md := range mds {
		if md.Alertname != "" {
			fun(md)
		}
	}
}
