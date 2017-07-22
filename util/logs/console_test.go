package logs

import "testing"

func TestConsoleLogger(t *testing.T) {
	logger := NewConsoleLogger(DEBUG)
	//all print
	logger.D("%v", "Debug")
	logger.I("%v", "Info")
	logger.W("%v", "Warning")
	logger.E("%v", "Error")
	logger.F("%v", "Fatal")

	logger.SetTraceLevel(ERROR)
	//only print error and fatal
	logger.D("%v", "Debug")
	logger.I("%v", "Info")
	logger.W("%v", "Warning")
	logger.E("%v", "Error")
	logger.F("%v", "Fatal")

}
