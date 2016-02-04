package log

import (
	"log"

	seelog "github.com/cihub/seelog"
)

func init() {
	conf := `
<seelog type="sync">
	<outputs formatid="main">
		<console />
	</outputs>
	<formats>
        <format id="main" format="%Date %Time [%Level] %FullPath %RelFile %Func %Line - %Msg%n" />
    </formats>
</seelog>
`
	logger, err := seelog.LoggerFromConfigAsBytes([]byte(conf))
	if err != nil {
		log.Panic(err)
	} else {
		seelog.ReplaceLogger(logger)
	}
}
