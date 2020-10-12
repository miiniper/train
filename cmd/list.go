package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/modood/table"
	"github.com/spf13/cobra"
)

var list = &cobra.Command{
	Use:   "list",
	Short: "get train list",
	Run: func(cmd *cobra.Command, args []string) {
		from, err := cmd.Flags().GetString("from")
		if err != nil {
			fmt.Errorf("get from flag err : %s ", err)
		}
		to, err := cmd.Flags().GetString("to")
		if err != nil {
			fmt.Errorf("get to flag err : %s ", err)
		}
		date, err := cmd.Flags().GetString("date")
		if err != nil {
			fmt.Errorf("get date flag err : %s ", err)
		}
		url := "https://kyfw.12306.cn/otn/leftTicket/query?leftTicketDTO.train_date=" + date + "&leftTicketDTO.from_station=" + from + "&leftTicketDTO.to_station=" + to + "&purpose_codes=ADULT"
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			fmt.Errorf("new request err : %s ", err)
		}
		client := &http.Client{}
		res, err := client.Do(req)
		if err != nil {
			fmt.Errorf("do request err : %s ", err)
		}

		defer res.Body.Close()
		resData := RespStatus{}
		body, _ := ioutil.ReadAll(res.Body)
		err = json.Unmarshal(body, &resData)
		if err != nil {
			fmt.Errorf("json unmarshal err : %s ", err)
		}

		var Trains []TrainInfo

		for _, j := range resData.Data.Result {
			s := strings.Split(j, "|")
			t := TrainInfo{s[3], s[8], s[13], s[32], s[31], s[30], s[23], s[28], s[29]}
			Trains = append(Trains, t)
		}

		table.Output(Trains)

	},
}

func init() {
	list.Flags().StringP("from", "f", "", "start station")
	list.Flags().StringP("to", "t", "", "end station")
	list.Flags().StringP("date", "d", "", "start date,eg :2020-10-11")
}

type TrainInfo struct {
	TrainNo       string `json:"train_no"`
	StartTime     string `json:"start_time"`
	Date          string `json:"date"`
	BusinessClass string `json:"business_class"`
	FirstClass    string `json:"first_class"`
	SecondClass   string `json:"second_class"`
	SoftBerth     string `json:"soft_berth"`
	HardBerth     string `json:"hard_berth"`
	HardSeat      string `json:"hard_seat"`
}

type RespStatus struct {
	Httpstatus int    `json:"httpstatus"`
	Data       Data   `json:"data"`
	Messages   string `json:"messages"`
	Status     bool   `json:"status"`
}
type Map struct {
	XHP string `json:"XHP"`
	VAP string `json:"VAP"`
	QIP string `json:"QIP"`
	BJP string `json:"BJP"`
	BXP string `json:"BXP"`
	VJP string `json:"VJP"`
}
type Data struct {
	Result []string `json:"result"`
	Flag   string   `json:"flag"`
	Map    Map      `json:"map"`
}
