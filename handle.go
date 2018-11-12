package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// EventInst event data
type EventInst struct {
	ID          int64       `json:"event_id,omitempty"`
	TxnID       string      `json:"txn_id"`
	EventType   string      `json:"event_type"`
	Action      string      `json:"action"`
	ActionTime  time.Time   `json:"action_time"`
	ObjType     string      `json:"obj_type"`
	Data        []EventData `json:"data"`
	OwnerID     string      `json:"bk_supplier_account"`
	RequestID   string      `json:"request_id"`
	RequestTime time.Time   `json:"request_time"`
}

// EventData event instance detail
type EventData struct {
	CurData interface{} `json:"cur_data"`
	PreData interface{} `json:"pre_data"`
}

// EventAction
const (
	EventActionCreate = "create"
	EventActionUpdate = "update"
	EventActionDelete = "delete"
)

// ObjectEvent accept event
func ObjectEvent(w http.ResponseWriter, r *http.Request) {

	input := new(EventInst)
	if err := json.NewDecoder(r.Body).Decode(input); err != nil {
		log.Fatalf("objectEvent fail to decode event info request body. err: %v\n", err)
		rsp := GetResponseErrorBody(ERRHTTPJSONUnmarshalFailure, err.Error())
		WriteJSON(http.StatusBadRequest, rsp, w)
		return
	}

	err := handleObjectEventInfo(input)
	if err != nil {
		rsp := GetResponseErrorBody(ERREventHandlelFailure, err.Error())
		WriteJSON(http.StatusInternalServerError, rsp, w)
		return
	}

	rsp := GetResponseSuccBody(nil)
	WriteJSON(http.StatusOK, rsp, w)

}

func handleObjectEventInfo(eventData *EventInst) error {

	for _, data := range eventData.Data {
		var mapData map[string]interface{}
		var ok bool
		if eventData.Action == EventActionDelete {
			mapData, ok = data.PreData.(map[string]interface{})
		} else {
			mapData, ok = data.CurData.(map[string]interface{})
		}
		if !ok {
			err := fmt.Errorf("event: model id:%s, action:%s, instance info not map", eventData.ObjType, eventData.Action)
			log.Fatalln(err.Error())
			return err
		}
		instName, ok := mapData["bk_inst_name"].(string)
		if !ok {
			err := fmt.Errorf("event: model id:%s, action:%s, not found instance name", eventData.ObjType, eventData.Action)
			log.Fatalln(err.Error())
			return err
		}
		msg := fmt.Sprintf("[event change] model id:%s, action:%s, instance name:%s", eventData.ObjType, eventData.Action, instName)
		fmt.Println(msg)
	}
	return nil

}
