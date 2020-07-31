package lib

import (
	"fmt"
	"os"
	"sgame/servers/comm"
	"time"
)

type FileConfig struct {
	//	ProcName string `json:"proc_name"`
	ConnServ      int      `json:"conn_serv"`
	MasterDb      int      `json:"master_db"`
	SlaveDb       int      `json:"slave_db"`
	LogFile       string   `json:"log_file"`
	ManageAddr    []string `json:"manage_addr"`
	MaxOnline     int      `json:"max_online"`
	ClientTimeout int      `json:"client_timeout"`
	MonitorInv    int      `json:"monitor_inv"` //monitor interval seconds
}

type Config struct {
	NameSpace  string
	ProcId     int
	ProcName   string
	ConfigFile string
	Daemon     bool
	FileConfig *FileConfig
	Comm       *comm.CommConfig
	TableMap   comm.TableMap
	ReportServ *comm.ReportServ //report to manger
	Users      *OnLineList
}

//Comm Config Setting
func CommSet(pconfig *Config) bool {
	var _func_ = "<CommSet>"
	//daemonize
	if pconfig.Daemon {
		comm.Daemonize()
	}

	//file config
	pconfig.FileConfig = new(FileConfig)
	if pconfig.FileConfig == nil {
		fmt.Printf("new FileConfig failed!\n")
		return false
	}

	//load file config
	if comm.LoadJsonFile(pconfig.ConfigFile, pconfig.FileConfig, nil) != true {
		fmt.Printf("%s failed!\n", _func_)
		return false
	}

	//comm config
	pconfig.Comm = comm.InitCommConfig(pconfig.FileConfig.LogFile, pconfig.NameSpace, pconfig.ProcId)
	if pconfig.Comm == nil {
		fmt.Printf("%s init comm config failed!", _func_)
		return false
	}
	pconfig.Comm.ServerCfg = pconfig

	var log = pconfig.Comm.Log
	//lock uniq
	if comm.LockUniqFile(pconfig.Comm, pconfig.NameSpace, pconfig.ProcId, pconfig.ProcName) == false {
		log.Err("%s lock uniq file failed!", _func_)
		return false
	}

	return true
}

//Local Proc Setting
func LocalSet(pconfig *Config) bool {
	var _func_ = "<LocalSet>"
	var log = pconfig.Comm.Log

	//users
	pconfig.Users = new(OnLineList)
	if pconfig.Users == nil {
		log.Err("%s new users failed!", _func_)
		return false
	}
	pconfig.Users.user_map = make(map[int64]*UserOnLine)

	//reg table-map
	if RegistTableMap(pconfig) == false {
		log.Err("%s regist table map failed!", _func_)
		return false
	}

	//load table-map
	if comm.LoadTableFiles(pconfig.TableMap, pconfig.Comm) == false {
		log.Err("%s load table files failed!", _func_)
		return false
	}

	//start report serv
	pconfig.ReportServ = comm.StartReport(pconfig.Comm, pconfig.ProcId, pconfig.ProcName, pconfig.FileConfig.ManageAddr, comm.REPORT_METHOD_ALL,
		pconfig.FileConfig.MonitorInv)
	if pconfig.ReportServ == nil {
		log.Err("%s fail! start report failed!", _func_)
		return false
	}
	pconfig.ReportServ.Report(comm.REPORT_PROTO_SERVER_START, time.Now().Unix(), "", nil)

	//add ticker
	pconfig.Comm.TickPool.AddTicker("heart_beat", comm.TICKER_TYPE_CIRCLE, 0, comm.PERIOD_HEART_BEAT_DEFAULT, SendHeartBeatMsg, pconfig)
	pconfig.Comm.TickPool.AddTicker("report_sync", comm.TICKER_TYPE_CIRCLE, 0, comm.PERIOD_REPORT_SYNC_DEFAULT, ReportSyncServer, pconfig)
	pconfig.Comm.TickPool.AddTicker("check_client_heart", comm.TICKER_TYPE_CIRCLE, 0, int64(pconfig.FileConfig.ClientTimeout*1000),
		CheckClientTimeout, pconfig)
	pconfig.Comm.TickPool.AddTicker("recv_cmd", comm.TICKER_TYPE_CIRCLE, 0, comm.PERIOD_RECV_REPORT_CMD_DEFAULT, RecvReportCmd, pconfig)

	return true
}

func ServerExit(pconfig *Config) {
	//close proc
	if pconfig.Comm.Proc != nil {
		pconfig.Comm.Proc.Close()
	}

	//close report_serv
	if pconfig.ReportServ != nil {
		pconfig.ReportServ.Report(comm.REPORT_PROTO_SERVER_STOP, time.Now().Unix(), "", nil)
		time.Sleep(time.Second)
		pconfig.ReportServ.Close()
	}

	//unlock uniq
	comm.UnlockUniqFile(pconfig.Comm, pconfig.NameSpace, pconfig.ProcId, pconfig.ProcName)

	//close log
	if pconfig.Comm.Log != nil {
		pconfig.Comm.Log.Info("%s", "server exit...")
		pconfig.Comm.Log.Close()
	}
	time.Sleep(1 * time.Second)
	os.Exit(0)
}

//Main Proc
func ServerStart(pconfig *Config) {
	var log = pconfig.Comm.Log
	var default_sleep = time.Duration(comm.DEFAULT_SERVER_SLEEP_IDLE)
	log.Info("%s starts---%v", pconfig.ProcName, os.Args)

	//each support routine
	go comm.HandleSignal(pconfig.Comm)

	//main loop
	for {
		//handle info
		handle_info(pconfig)

		//recv pkg
		RecvMsg(pconfig)

		//tick
		handle_tick(pconfig)

		//sleep
		time.Sleep(time.Millisecond * default_sleep)
	}
}

//After ReLoad Config If Need Handle
func AfterReLoadConfig(pconfig *Config, old_config *FileConfig, new_config *FileConfig) {
	var _func_ = "<AfterReLoadConfig>"
	log := pconfig.Comm.Log

	log.Info("%s finish!", _func_)
	return
}

/*----------------Static Func--------------------*/
func handle_info(pconfig *Config) {
	var _func_ = "<handle_info>"
	log := pconfig.Comm.Log
	select {
	case m := <-pconfig.Comm.ChInfo:
		switch m {
		case comm.INFO_EXIT:
			ServerExit(pconfig)
		case comm.INFO_USR1:
			log.Info(">>reload config!")
			var new_config FileConfig
			ret := comm.LoadJsonFile(pconfig.ConfigFile, &new_config, pconfig.Comm)
			if !ret {
				log.Err("%s reload config failed!", _func_)
			} else {
				AfterReLoadConfig(pconfig, pconfig.FileConfig, &new_config)
				*(pconfig.FileConfig) = new_config
			}
		case comm.INFO_USR2:
			log.Info(">>reload tables")
			comm.ReLoadTableFiles(pconfig.TableMap, pconfig.Comm)
		case comm.INFO_PPROF:
			log.Info(">>profiling")
			comm.DefaultHandleProfile(pconfig.Comm)
		default:
			pconfig.Comm.Log.Info("unknown msg")
		}
	default:
		break
	}
}

//each ticker
func handle_tick(pconfig *Config) {
	pconfig.Comm.TickPool.Tick(0)
}
