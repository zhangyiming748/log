package log

import "testing"

func TestSuperLog(t *testing.T) {
	logger := SetLog(LevelWarn)
	logger.Info("Go is best language!", "公众号", "Golang来啦")
	logger.Debug("Go is best language!", "公众号", "Golang来啦")
	logger.Warn("Go is best language!", "公众号", "Golang来啦")
	logger.Error("Go is best language!", "公众号", "Golang来啦")
}
