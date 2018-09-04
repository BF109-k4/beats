package clustersummary

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func getFileds(address string, timeout time.Duration) (map[string]interface{}, error){
	res, err := http.Get("http://"+address+"/api/v1/topology/summary")
	if err != nil  {
		fmt.Print(err)
	}
	//fmt.Println(reflect.TypeOf(res))
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil  {
		fmt.Print(err)
	}
	var dat map[string]interface{}
	json.Unmarshal([]byte(string(body)), &dat)
	var aa = dat["topologies"]
	bb := aa.([]interface{})
	var count = len(bb)
	var clusterInfo = make(map[int]interface{})
	for i := 0; i < count; i=i+1 {
		var topoInfo TopoInfo
		topoInfo.Name = bb[i].(map[string]interface{})["name"].(string)
		topoInfo.Status = bb[i].(map[string]interface{})["status"].(string)
		topoInfo.Uptime = bb[i].(map[string]interface{})["uptime"].(string)
		topoInfo.WorkersTotal = bb[i].(map[string]interface{})["workersTotal"].(float64)
		topoInfo.TasksTotal = bb[i].(map[string]interface{})["tasksTotal"].(float64)
		topoInfo.ExecutorsTotal = bb[i].(map[string]interface{})["executorsTotal"].(float64)
		topoInfo.ReplicationCount = bb[i].(map[string]interface{})["replicationCount"].(float64)
		topoInfo.AssignedMemOnHeap = bb[i].(map[string]interface{})["assignedMemOnHeap"].(float64)
		topoInfo.Spouts, topoInfo.Bolts = getdetailInfo(bb[i].(map[string]interface{})["id"].(string), topoInfo)
		jsTopoInfo, _ := json.Marshal(topoInfo)
		clusterInfo[i] =jsTopoInfo
	}
	var data map[string]interface{}
	json.Unmarshal([]byte(clusterInfo[0].([]byte)), &data)
	//fmt.Println(data)
	return data, err
}

func getdetailInfo(topoId string, topoInfo TopoInfo)([]SpoutInfo, []BoltInfo) {
	var tmpSpoutsInfo []SpoutInfo
	var tmpBoltsInfo []BoltInfo
	res, err := http.Get("http://10.20.13.162:8080/api/v1/topology/" + topoId)
	if err != nil  {
		fmt.Print(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	var dat map[string]interface{}
	json.Unmarshal([]byte(string(body)), &dat)
	spouts := dat["spouts"].([]interface{})
	bolts := dat["bolts"].([]interface{})
	for i:= 0; i < len(spouts); i=i+1 {
		var spoutInfo SpoutInfo
		spoutInfo.SpoutId = spouts[i].(map[string]interface{})["spoutId"].(string)
		spoutInfo.Acked = spouts[i].(map[string]interface{})["acked"].(float64)
		spoutInfo.Emitted = spouts[i].(map[string]interface{})["emitted"].(float64)
		spoutInfo.Tasks = spouts[i].(map[string]interface{})["tasks"].(float64)
		spoutInfo.Failed = spouts[i].(map[string]interface{})["failed"].(float64)
		tmpSpoutsInfo = append(tmpSpoutsInfo, spoutInfo)
	}
	for i:= 0; i < len(bolts); i=i+1 {
		var boltInfo BoltInfo
		boltInfo.BoltId = bolts[i].(map[string]interface{})["boltId"].(string)
		boltInfo.Acked = bolts[i].(map[string]interface{})["acked"].(float64)
		boltInfo.Emitted = bolts[i].(map[string]interface{})["emitted"].(float64)
		boltInfo.Tasks = bolts[i].(map[string]interface{})["tasks"].(float64)
		boltInfo.Failed = bolts[i].(map[string]interface{})["failed"].(float64)
		tmpBoltsInfo = append(tmpBoltsInfo, boltInfo)
	}
	return tmpSpoutsInfo,tmpBoltsInfo
}
