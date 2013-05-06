package aria2rpc

import (
	"log"
	"testing"
)

func NoTestAddUri(t *testing.T) {
	//uri := "https://www.kernel.org/pub/linux/kernel/v3.x/linux-3.8.4.tar.xz"
	uri := "http://mirrors.ustc.edu.cn/ubuntu-releases//quantal/ubuntu-12.10-desktop-i386.iso"
	//uri := "http://bt.ktxp.com/torrents/2013/03/24/3a091e44394e2ec6345cc263accf31058eda504e.torrent"

	params := make(map[string]string)
	params["max-download-limit"] = "1K"
	gid, err := AddUri(uri, params)
	if err != nil {
		log.Fatal("add Uri error:", err)
	}
	log.Println(gid)
}

func NoTestAddTorrent(t *testing.T) {
	gid, err := AddTorrent("/home/noah/DueWest.torrent")
	if err != nil {
		log.Fatal("add Torrent error:", err)
	}
	log.Println(gid)
}

func NoTestRemove(t *testing.T) {
	gid, err := Remove("12", true)
	if err != nil {
		log.Fatal("Remove error:", err)
	}
	log.Println(gid)
}

func NoTestPause(t *testing.T) {
	gid, err := Pause("16", false)
	if err != nil {
		log.Fatal("Pause error:", err)
	}
	log.Println(gid)
}

func NoTestPauseAll(t *testing.T) {
	reply, err := PauseAll()
	if err != nil {
		log.Fatal("Pause All error:", err)
	}
	log.Println(reply)
}

func NoTestUnpause(t *testing.T) {
	gid, err := Unpause("16")
	if err != nil {
		log.Fatal("Unpause error:", err)
	}
	log.Println(gid)
}

func NoTestUnPauseAll(t *testing.T) {
	reply, err := UnpauseAll()
	if err != nil {
		log.Fatal("Unpause All error:", err)
	}
	log.Println(reply)
}

func TestGetStatus(t *testing.T) {
	//keys := []string{"gid", "status"}
	reply, err := GetStatus("27", nil)
	if err != nil {
		log.Fatal("GetStatus", err)
	}
	log.Println(reply)
}

func NoTestGetActive(t *testing.T) {
	//keys := []string{"gid", "status"}
	//reply, err := GetActive(keys)
	//if err != nil {
	//	log.Fatal("GetActive", err)
	//}
	//log.Println(reply)

	reply2, err2 := GetActive(nil)
	if err2 != nil {
		log.Fatal("GetActive", err2)
	}
	log.Println(reply2)
}

func NoTestGetWaiting(t *testing.T) {
	keys := []string{"gid", "status"}
	reply, err := GetWaiting(0, 10, keys)
	if err != nil {
		log.Fatal("Waiting error:", err)
	}
	log.Println(reply)
}

func NoTestGetStopped(t *testing.T) {
	reply, err := GetStopped(0, 10, nil)
	if err != nil {
		log.Fatal("Waiting error:", err)
	}
	log.Println(reply)
}

func NoTestGetOption(t *testing.T) {
	reply, err := GetOption("16")
	if err != nil {
		log.Fatal("GetOption error:", err)
	}
	log.Println(reply)
}

func NoTestGetGlobalOption(t *testing.T) {
	reply, err := GetGlobalOption()
	if err != nil {
		log.Fatal("GetGlobalOption error:", err)
	}
	log.Println(reply)
}

func NoTestChangeOption(t *testing.T) {
	params := make(map[string]string)
	params["max-download-limit"] = "12K"
	reply, err := ChangeOption("2", params)
	if err != nil {
		log.Fatal("ChangeOption error:", err)
	}
	log.Println(reply)
}

func NoTestChangeGlobalOption(t *testing.T) {
	params := make(map[string]string)
	params["max-overall-download-limit"] = "1K"
	reply, err := ChangeGlobalOption(params)
	if err != nil {
		log.Fatal("ChangeGlobalOption error:", err)
	}
	log.Println(reply)
}

func NoTestGetGlobalStat(t *testing.T) {
	reply, err := GetGlobalStat()
	if err != nil {
		log.Fatal("GetGlobalStat error:", err)
	}
	log.Println(reply)
}
