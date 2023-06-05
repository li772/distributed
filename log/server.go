package log

import stlog "log"

var log *stlog.Logger

type fileLog string

func (fl fileLog) Write(data []byte) (int, err) {

}
