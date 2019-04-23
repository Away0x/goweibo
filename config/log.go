package config

import (
	"github.com/lexkong/log"
)

func initLog() {
	logConfig := log.PassLagerCfg{
		// 输出位置，有两个可选项 —— file 和 stdout。选择 file 会将日志记录到 logger_file 指定的日志文件中，选择 stdout 会将日志输出到标准输出，当然也可以两者同时选择
		Writers: "file,stdout",
		// 日志级别，DEBUG、INFO、WARN、ERROR、FATAL
		LoggerLevel: "DEBUG",
		// 日志文件
		LoggerFile: logFilePath,
		// 日志的输出格式，JSON 或者 plaintext，true 会输出成 JSON 格式，false 会输出成非 JSON 格式
		LogFormatText: false,
		// rotate 依据，可选的有 daily 和 size。如果选 daily 则根据天进行转存，如果是 size 则根据大小进行转存
		RollingPolicy: "size",
		// rotate 转存时间，配 合rollingPolicy: daily 使用
		LogRotateDate: 1,
		// rotate 转存大小，配合 rollingPolicy: size 使用 (大于 1mb 会压缩为 zip)
		LogRotateSize: 1,
		// 当日志文件达到转存标准时，log 系统会将该日志文件进行压缩备份，这里指定了备份文件的最大个数
		LogBackupCount: 7,
	}

	log.InitWithConfig(&logConfig)
}
