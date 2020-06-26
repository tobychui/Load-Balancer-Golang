package main

import (
    "net/http"
    "strconv"
    "flag"
    "log"
)

var configName = flag.String("conf", "config.yml", "The filepath for the config file")
var useTLS =  flag.Bool("tls", false, "Enable TLS mode on the load balancing server. Require paramter -cert and -key")
var tls_cert = flag.String("cert", "localhost.crt", "TLS certificate file (.crt)")
var tls_key = flag.String("key", "localhost.key", "TLS key file (.key)")

func main() {
    //Prase the flags 
    flag.Parse()
    log.Println("Spinning up load balancer...")
    log.Println("Reading Config.yml...")

    //Read configuration from file
    proxy, err := ReadConfig(*configName)
    if err != nil {
      LogErr("An error occurred while trying to parse config.yml")
      LogErrAndCrash(err.Error())
    }

    http.HandleFunc("/", proxy.handler)
    log.Println("Listening to requests on port: " + strconv.Itoa(proxy.Port))

    //Serve the web, using the user launch config
    if (*useTLS){
        err := http.ListenAndServeTLS(":" + strconv.Itoa(proxy.Port), *tls_cert, *tls_key, nil)
        if err != nil {
            LogErr("Failed to bind to port " + strconv.Itoa(proxy.Port))
            panic(err)
        }
    }else{
        err = http.ListenAndServe(":" + strconv.Itoa(proxy.Port), nil)
        if err != nil {
            LogErr("Failed to bind to port " + strconv.Itoa(proxy.Port))
            panic(err)
        }
    }
    
    
}
