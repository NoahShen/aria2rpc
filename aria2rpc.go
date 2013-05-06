package aria2rpc

import (
	"bufio"
	"encoding/base64"
	"github.com/kdar/httprpc"
	"io/ioutil"
	"os"
)

var RpcUrl = "http://127.0.0.1:6800/jsonrpc"
var RpcVersion = "2.0"

func AddUri(uri string, params map[string]interface{}) (string, error) {
	method := "aria2.addUri"
	paramArr := make([]interface{}, 1)
	uris := make([]string, 1)
	uris[0] = uri
	paramArr[0] = uris
	if params != nil {
		paramArr = append(paramArr, params)
	}

	var replyGID string
	err := httprpc.CallJson(RpcVersion, RpcUrl, method, paramArr, &replyGID)
	if err != nil {
		return "", err
	}
	return replyGID, nil
}

func AddTorrent(path string) (string, error) {
	file, openFileErr := os.Open(path)
	if openFileErr != nil {
		return "", nil
	}

	fileBytes, readError := ioutil.ReadAll(bufio.NewReader(file))
	if readError != nil {
		return "", readError
	}
	coder := base64.StdEncoding
	encoded := coder.EncodeToString(fileBytes)

	method := "aria2.addTorrent"
	paramArr := make([]interface{}, 1)
	paramArr[0] = encoded
	var replyGID string
	err := httprpc.CallJson(RpcVersion, RpcUrl, method, paramArr, &replyGID)
	if err != nil {
		return "", err
	}
	return replyGID, nil
}

func Remove(gid string, force bool) (string, error) {
	var method string
	if force {
		method = "aria2.forceRemove"
	} else {
		method = "aria2.remove"
	}
	paramArr := make([]interface{}, 1)
	paramArr[0] = gid
	var replyGID string
	err := httprpc.CallJson(RpcVersion, RpcUrl, method, paramArr, &replyGID)
	if err != nil {
		return "", err
	}
	return replyGID, nil
}

func Pause(gid string, force bool) (string, error) {
	var method string
	if force {
		method = "aria2.forcePause"
	} else {
		method = "aria2.pause"
	}
	paramArr := make([]interface{}, 1)
	paramArr[0] = gid
	var replyGID string
	err := httprpc.CallJson(RpcVersion, RpcUrl, method, paramArr, &replyGID)
	if err != nil {
		return "", err
	}
	return replyGID, nil
}

func PauseAll() (string, error) {
	method := "aria2.pauseAll"
	paramArr := make([]string, 0)
	var reply string
	err := httprpc.CallJson(RpcVersion, RpcUrl, method, paramArr, &reply)
	if err != nil {
		return "", err
	}
	return reply, nil
}

func Unpause(gid string) (string, error) {
	method := "aria2.unpause"
	paramArr := make([]interface{}, 1)
	paramArr[0] = gid
	var replyGID string
	err := httprpc.CallJson(RpcVersion, RpcUrl, method, paramArr, &replyGID)
	if err != nil {
		return "", err
	}
	return replyGID, nil
}

func UnpauseAll() (string, error) {
	method := "aria2.unpauseAll"
	paramArr := make([]string, 0)
	var reply string
	err := httprpc.CallJson(RpcVersion, RpcUrl, method, paramArr, &reply)
	if err != nil {
		return "", err
	}
	return reply, nil
}

func GetStatus(gid string, keys []string) (map[string]interface{}, error) {
	method := "aria2.tellStatus"
	paramArr := make([]interface{}, 1)
	paramArr[0] = gid
	if keys != nil && len(keys) > 0 {
		paramArr = append(paramArr, keys)
	}
	var reply = make(map[string]interface{})
	err := httprpc.CallJson(RpcVersion, RpcUrl, method, paramArr, &reply)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func GetActive(keys []string) ([]map[string]interface{}, error) {
	method := "aria2.tellActive"
	paramArr := make([]interface{}, 0)
	if keys != nil && len(keys) > 0 {
		paramArr = append(paramArr, keys)
	}
	var reply = make([]map[string]interface{}, 10)
	err := httprpc.CallJson(RpcVersion, RpcUrl, method, paramArr, &reply)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func GetWaiting(offset, num int, keys []string) ([]map[string]interface{}, error) {
	method := "aria2.tellWaiting"
	paramArr := make([]interface{}, 2)
	paramArr[0] = offset
	paramArr[1] = num
	if keys != nil && len(keys) > 0 {
		paramArr = append(paramArr, keys)
	}
	var reply = make([]map[string]interface{}, 10)
	err := httprpc.CallJson(RpcVersion, RpcUrl, method, paramArr, &reply)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func GetStopped(offset, num int, keys []string) ([]map[string]interface{}, error) {
	method := "aria2.tellStopped"
	paramArr := make([]interface{}, 2)
	paramArr[0] = offset
	paramArr[1] = num
	if keys != nil && len(keys) > 0 {
		paramArr = append(paramArr, keys)
	}
	var reply = make([]map[string]interface{}, 10)
	err := httprpc.CallJson(RpcVersion, RpcUrl, method, paramArr, &reply)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func GetOption(gid string) (map[string]interface{}, error) {
	method := "aria2.getOption"
	paramArr := make([]interface{}, 1)
	paramArr[0] = gid
	var reply = make(map[string]interface{})
	err := httprpc.CallJson(RpcVersion, RpcUrl, method, paramArr, &reply)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func ChangeOption(gid string, params map[string]string) (string, error) {
	method := "aria2.changeOption"
	paramArr := make([]interface{}, 1)
	paramArr[0] = gid
	paramArr = append(paramArr, params)

	var reply string
	err := httprpc.CallJson(RpcVersion, RpcUrl, method, paramArr, &reply)
	if err != nil {
		return "", err
	}
	return reply, nil
}

func GetGlobalOption() (map[string]interface{}, error) {
	method := "aria2.getGlobalOption"
	paramArr := make([]interface{}, 0)
	var reply = make(map[string]interface{})
	err := httprpc.CallJson(RpcVersion, RpcUrl, method, paramArr, &reply)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func ChangeGlobalOption(params map[string]string) (string, error) {
	method := "aria2.changeGlobalOption"
	paramArr := make([]interface{}, 0)
	paramArr = append(paramArr, params)

	var reply string
	err := httprpc.CallJson(RpcVersion, RpcUrl, method, paramArr, &reply)
	if err != nil {
		return "", err
	}
	return reply, nil
}

func GetGlobalStat() (map[string]interface{}, error) {
	method := "aria2.getGlobalStat"
	paramArr := make([]interface{}, 0)
	var reply = make(map[string]interface{})
	err := httprpc.CallJson(RpcVersion, RpcUrl, method, paramArr, &reply)
	if err != nil {
		return nil, err
	}
	return reply, nil
}
