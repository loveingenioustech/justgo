package main

import (
	"github.com/cihub/seelog"
	"fmt"
)

func main() {
	demoSeeLog1()

	demoCustomSeeLog()
}

func demoSeeLog1() {
	defer seelog.Flush()
	seelog.Info("Hello from Seelog!")
}

var Logger seelog.LoggerInterface

func loadAppConfig() {
	appConfig := `
<seelog minlevel="warn">
    <outputs formatid="common">
        <rollingfile type="size" filename="D:\tmp\logs\roll.log" maxsize="100000" maxrolls="5"/>
		<filter levels="critical">
            <file path="/data/logs/critical.log" formatid="critical"/>
            <smtp formatid="criticalemail" senderaddress="test@126.com" sendername="ShortUrl API" hostname="smtp.126.com" hostport="25" username="demo" password="sfadasdf">
                <recipient address="ff@gmail.com"/>
            </smtp>
        </filter>
    </outputs>
    <formats>
        <format id="common" format="%Date/%Time [%LEV] %Msg%n" />
	    <format id="critical" format="%File %FullPath %Func %Msg%n" />
	    <format id="criticalemail" format="Critical error on our server!\n    %Time %Date %RelFile %Func %Msg \nSent by Seelog"/>
    </formats>
</seelog>
`
	logger, err := seelog.LoggerFromConfigAsBytes([]byte(appConfig))
	if err != nil {
		fmt.Println(err)
		return
	}
	UseLogger(logger)
}

func init() {
	DisableLog()
	loadAppConfig()
}

// DisableLog disables all library log output
func DisableLog() {
	Logger = seelog.Disabled
}

// UseLogger uses a specified seelog.LoggerInterface to output library log.
// Use this func if you are using Seelog logging system in your app.
func UseLogger(newLogger seelog.LoggerInterface) {
	Logger = newLogger
}

func demoCustomSeeLog() {
	Logger.Critical("test Critical message")
}
