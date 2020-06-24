package main

import (
    ll "sgame/servers/manage_serv/lib"
    "flag"
    "fmt"
)


var config ll.Config;
var pconfig *ll.Config = &config;

//option
var name_space = flag.String("N", "", "name space ");
//var proc_id = flag.Int("p", 0, "proc id in proc_bridge sys");
var config_file = flag.String("f", "", "config file");
var proc_name = flag.String("P", "" , "proc name ");

func init() {
}

func parse_flag() bool {
    //check flag
	flag.Parse();
	if len(*name_space) <=0 || len(*config_file)<=0 || len(*proc_name)<=0 {
		flag.PrintDefaults();
		return false;
	}
	pconfig.ProcId = 0;
	pconfig.NameSpace = *name_space;
    pconfig.ConfigFile = *config_file;
    pconfig.ProcName = *proc_name;
    return true;	
}

func main() {		
	//parse flag
	if parse_flag() == false {
		return;
	}
	
    //comm set
	if ll.CommSet(pconfig) == false {
		fmt.Printf("comm set failed!\n");
		return;
	}
	
	//self set
	if ll.SelfSet(pconfig) == false {
		fmt.Printf("self set failed!\n");
		ll.ServerExit(pconfig);
		return;
	}
    
        
    //start server
    ll.ServerStart(pconfig);
}