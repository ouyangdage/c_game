package controllers

import (
	//	"code.google.com/p/goprotobuf/proto"
	//	"crypto/md5"
	//	"encoding/hex"
	//	"encoding/json"
	"fmt"
	//	"io/ioutil"
	//	"libs/log"
	//	"libs/token"
	"github.com/fhbzyc/c_game/models"
	//	"net/http"
	//	"protocol"
	//	"rpc"
	//	"strconv"
	//	"strings"
	//	"time"
	//	"models/table"
)

func (this *Controller) Signin() error {

	request := this.Request

	username, ok := request.Params[0].(string)
	if !ok {
		return this.returnError(lineNum(), fmt.Errorf("Invalid method parameters"))
	}

	_, ok = request.Params[1].(string)
	if !ok {
		return this.returnError(lineNum(), fmt.Errorf("Invalid method parameters"))
	}

	user, err := models.User.FindOneByPlatformId(0, username)
	if err != nil {
		return this.returnError(lineNum(), err)
	} else if user.PlatformUuid == "" {
		user.PlatformUuid = username
		if err := models.User.Insert(user); err != nil {
			return this.returnError(lineNum(), err)
		}
	}

	this.Connect.Uid = user.Uid

	return this.returnSuccess(user)
}

//	request := new(protocol.LoginRequest)
//	err = proto.Unmarshal(req.Req, request)
//	if err != nil {
//		return err
//	}
//
//	platformId := request.GetPlatformId()
//	platformUuid := request.GetPlatformUuid()
//	//	otherData := request.GetOtherData()
//	//session := request.GetOtherSession()
//	//sign := request.GetOtherSign()
//
//	user, err := models.User.GetUserByPlatformId(platformId, platformUuid)
//	if err != nil {
//		return ReturnError(lineNum(), err)
//	}
//	if user.Uid == 0 {
//
//		user = new(models.UserTable)
//		user.PlatformId = platformId
//		user.PlatformUuid = platformUuid
//		user.Ip = request.GetIp()
//		user.Imei = request.GetImei()
//		err = models.User.Insert(user)
//		if err != nil {
//			return ReturnError(lineNum(), err)
//		}
//	}
//
//	tokenStr, err := token.GameToken.AddToken(user.Uid)
//	if err != nil {
//		return ReturnError(lineNum(), err)
//	}
//
//	areaList, err := models.GameArea.GetAll()
//	if err != nil {
//		return ReturnError(lineNum(), err)
//	}
//
//	var areaProtoList []*protocol.Area
//	for _, area := range areaList {
//		temp := new(protocol.Area)
//		temp.AreaId = proto.Int32(area.AreaId)
//		temp.AreaName = proto.String(area.AreaName)
//		areaProtoList = append(areaProtoList, temp)
//	}
//
//	return ReturnSuccess(r, &protocol.LoginResponse{
//		Token:    proto.String(tokenStr),
//		AreaList: areaProtoList})
//}
//
//func ppLogin(ppToken string) string {
//
//	if len(ppToken) != 32 {
//		return ""
//	}
//
//	pp_id := 4335
//	app_key := "8dbbcdf221234073ccd75b1a277f7255"
//	url := "http://passport_i.25pp.com:8080/account?tunnel-command=2852126760"
//
//	m := md5.New()
//	m.Write([]byte("sid=" + ppToken + app_key))
//	sign := hex.EncodeToString(m.Sum(nil))
//
//	postData := `{"data":{"sid":"%s"},"encrypt":"MD5","game":{"gameId":%d},"id":%d,"service":"account.verifySession","sign":"%s"}`
//	postData = fmt.Sprintf(postData, ppToken, pp_id, time.Now().Unix(), sign)
//
//	client := new(http.Client)
//	resp, err := client.Post(url, "application/x-www-form-urlencoded", strings.NewReader(postData))
//	if err != nil {
//		log.Logger.Error("PPLOGIN ERROR , line: ", lineNum())
//		return ""
//	}
//
//	body, _ := ioutil.ReadAll(resp.Body)
//
//	resp.Body.Close()
//
//	var jsonData map[string]interface{}
//	if json.Unmarshal(body, &jsonData) != nil {
//		log.Logger.Error("PPLOGIN ERROR , line: ", lineNum())
//		return ""
//	}
//
//	code := jsonData["state"].(map[string]interface{})
//	if int(code["code"].(float64)) != 1 {
//		log.Logger.Error("PPLOGIN ERROR , line: ", lineNum())
//		return ""
//	}
//
//	data := jsonData["data"].(map[string]interface{})
//	return data["accountId"].(string)
//}
//
//func login91(uin string, session string) error {
//
//	appId := 100
//	appKey := ""
//
//	sign := strconv.Itoa(appId) + "4" + uin + session + appKey
//	m := md5.New()
//	m.Write([]byte(sign))
//	sign = hex.EncodeToString(m.Sum(nil))
//
//	url := "http://service.sj.91.com/usercenter/ap.aspx?AppId=%d&Act=4&Uin=%s&SessionId=%s&Sign=" + sign
//	url = fmt.Sprintf(url, appId, uin, session)
//
//	client := new(http.Client)
//	resp, err := client.Get(url)
//	if err != nil {
//		log.Logger.Errorf("91LOGIN ERROR , line: %d , %v", lineNum(), err)
//		return err
//	}
//
//	body, _ := ioutil.ReadAll(resp.Body)
//	resp.Body.Close()
//
//	var jsonData map[string]string
//	if err = json.Unmarshal(body, &jsonData); err != nil {
//		log.Logger.Errorf("91LOGIN ERROR , line: %d , %v", lineNum(), err)
//		return err
//	}
//
//	if jsonData["ErrorCode"] != "1" {
//		err = fmt.Errorf("ErrorCode != 1")
//		log.Logger.Errorf("91LOGIN ERROR , line: %d , %v", lineNum(), err)
//		return err
//	}
//
//	return nil
//}
