package lib

import (
	"fmt"
	"os"
	"sgame/servers/comm"
	"time"
)

type FileConfig struct {
	//ProcName string `json:"proc_name"`
	LogicServ  int      `json:"logic_serv"`
	LogFile    string   `json:"log_file"`
	MaxConn    int      `json:"max_conn"`
	ListenAddr string   `json:"listen_addr"`
	ManageAddr []string `json:"manage_addr"`
	ZlibOn     int      `json:"zlib_on"` //json compessed by zlib
	MonitorInv int      `json:"monitor_inv"` //monitor interval seconds
}

type Config struct {
	//comm
	NameSpace  string
	ProcId     int
	ProcName   string
	ConfigFile string
	Daemon bool
	FileConfig *FileConfig
	Comm       *comm.CommConfig
	ReportServ *comm.ReportServ //report to manger
	//local
	TcpServ    *comm.TcpServ
	Ckey2Uid   map[int64]int64  //client key to uid. used for search login user
	Uid2Ckey   map[int64]int64  //uid to client key. used for login user
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

	//lock uniq
	if comm.LockUniqFile(pconfig.Comm, pconfig.NameSpace, pconfig.ProcId, pconfig.ProcName) == false {
		pconfig.Comm.Log.Err("%s lock uniq file failed!", _func_)
		return false
	}
	return true
}

//Local Proc Setting
func LocalSet(pconfig *Config) bool {
	var _func_ = "<LocalSet>"
	log := pconfig.Comm.Log

	//start tcp serv to listen clients
	pserv := comm.StartTcpServ(pconfig.Comm, pconfig.FileConfig.ListenAddr, pconfig.FileConfig.MaxConn)
	if pserv == nil {
		log.Err("%s fail! start_tcp_serv at %s failed!", _func_, pconfig.FileConfig.ListenAddr)
		return false
	}
	pconfig.TcpServ = pserv

	//init some map
	pconfig.Ckey2Uid = make(map[int64]int64)
	pconfig.Uid2Ckey = make(map[int64]int64)

	//start report serv
	pconfig.ReportServ = comm.StartReport(pconfig.Comm, pconfig.ProcId, pconfig.ProcName, pconfig.FileConfig.ManageAddr, comm.REPORT_METHOD_ALL ,
		pconfig.FileConfig.MonitorInv);
	if pconfig.ReportServ == nil {
		log.Err("%s fail! start report failed!" , _func_);
		return false;
	}
	pconfig.ReportServ.Report(comm.REPORT_PROTO_SERVER_START, time.Now().Unix(), "", nil)

	//add ticker
	pconfig.Comm.TickPool.AddTicker("heart_beat", comm.TICKER_TYPE_CIRCLE, 0, comm.PERIOD_HEART_BEAT_DEFAULT, SendHeartBeatMsg, pconfig)
	pconfig.Comm.TickPool.AddTicker("report_sync", comm.TICKER_TYPE_CIRCLE, 0, comm.PERIOD_REPORT_SYNC_DEFAULT, ReportSyncServer, pconfig)
	pconfig.Comm.TickPool.AddTicker("recv_cmd" , comm.TICKER_TYPE_CIRCLE , 0 , comm.PERIOD_RECV_REPORT_CMD_DEFAULT , RecvReportCmd , pconfig);


	return true
}

func ServerExit(pconfig *Config) {
	//close proc
	if pconfig.Comm.Proc != nil {
		pconfig.Comm.Proc.Close()
	}

	//close tcp_serv
	if pconfig.TcpServ != nil {
		pconfig.TcpServ.Close(pconfig.Comm)
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

		//read client
		ReadClients(pconfig)

		//sleep
		time.Sleep(time.Millisecond * default_sleep)
	}
}

//After ReLoad Config If Need Handle
func AfterReLoadConfig(pconfig *Config , old_config *FileConfig , new_config *FileConfig)  {
	var _func_ = "<AfterReLoadConfig>";
	log := pconfig.Comm.Log;

    //Set MaxConn
    if old_config.MaxConn != new_config.MaxConn {
    	log.Info("%s reset max_conn %d-->%d" , _func_ , old_config.MaxConn , new_config.MaxConn);
		pconfig.TcpServ.SetMaxConn(new_config.MaxConn);
	}
	return;
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
			var new_config FileConfig;
			ret := comm.LoadJsonFile(pconfig.ConfigFile , &new_config , pconfig.Comm);
			if !ret {
				log.Err("%s reload config failed!" , _func_)
			} else {
				AfterReLoadConfig(pconfig , pconfig.FileConfig , &new_config);
				*(pconfig.FileConfig) = new_config;
			}
		case comm.INFO_USR2:
			log.Info(">>info usr2")
		case comm.INFO_PPROF:
			log.Info(">>profiling");
			comm.DefaultHandleProfile(pconfig.Comm);
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


