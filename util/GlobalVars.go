package util

import "fmt"

var PacketDelimiter = "\r\n"
var PayloadDelimiter = ":"
var Version = "1.0"

var Caps = []string{"ATCINFO=1", "ICAOEQ=1", "FASTPOS=1"}

var SoftwareTitle = fmt.Sprintf("Advanced Flight Server v%s. Licensed under the AGPL v3.0.\n Github: https://github.com/AFcPPe/AdvancedFlightServer-Go", Version)
