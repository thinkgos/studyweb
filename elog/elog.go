package elog

import (
	"sync/atomic"
)

// Provider is the interface implements log message levels method
// RFC5424 log message levels.
type Provider interface {
	Emergency(format string, v ...interface{})
	Alert(format string, v ...interface{})
	Critical(format string, v ...interface{})
	Error(format string, v ...interface{})
	Warning(format string, v ...interface{})
	Notice(format string, v ...interface{})
	Info(format string, v ...interface{})
	Debug(format string, v ...interface{})
}

var _ Provider = (*Elog)(nil)

// Elog log control
type Elog struct {
	logger Provider
	// has log output
	haslog uint32
}

// NewElog new a elog control
// default disable log, if p nil empty log,other use p log
func NewElog() *Elog {
	return &Elog{new(emptyProvider), 0}
}

// Mode set enable or diable log output when you has set logger
func (sf *Elog) Mode(enable bool) {
	if enable {
		atomic.StoreUint32(&sf.haslog, 1)
	} else {
		atomic.StoreUint32(&sf.haslog, 0)
	}
}

// SetProvider set logger provider
func (sf *Elog) SetProvider(p Provider) {
	if p != nil {
		sf.logger = p
	}
}

// GetProvider Get logger provider
func (sf *Elog) GetProvider() Provider {
	return sf.logger
}

// Emergency Log EMERGENCY level message.
func (sf *Elog) Emergency(format string, v ...interface{}) {
	if atomic.LoadUint32(&sf.haslog) == 1 {
		sf.logger.Emergency(format, v...)
	}
}

// Alert Log ALERT level message.
func (sf *Elog) Alert(format string, v ...interface{}) {
	if atomic.LoadUint32(&sf.haslog) == 1 {
		sf.logger.Alert(format, v...)
	}
}

// Critical Log CRITICAL level message.
func (sf *Elog) Critical(format string, v ...interface{}) {
	if atomic.LoadUint32(&sf.haslog) == 1 {
		sf.logger.Critical(format, v...)
	}
}

// Error Log ERROR level message.
func (sf *Elog) Error(format string, v ...interface{}) {
	if atomic.LoadUint32(&sf.haslog) == 1 {
		sf.logger.Error(format, v...)
	}
}

// Warning Log WARNING level message.
func (sf *Elog) Warning(format string, v ...interface{}) {
	if atomic.LoadUint32(&sf.haslog) == 1 {
		sf.logger.Warning(format, v...)
	}
}

// Notice Log NOTICE level message.
func (sf *Elog) Notice(format string, v ...interface{}) {
	if atomic.LoadUint32(&sf.haslog) == 1 {
		sf.logger.Notice(format, v...)
	}
}

// Info Log INFORMATIONAL level message.
func (sf *Elog) Info(format string, v ...interface{}) {
	if atomic.LoadUint32(&sf.haslog) == 1 {
		sf.logger.Info(format, v...)
	}
}

// Debug Log DEBUG level message.
func (sf *Elog) Debug(format string, v ...interface{}) {
	if atomic.LoadUint32(&sf.haslog) == 1 {
		sf.logger.Debug(format, v...)
	}
}

/**************************  default log ************************************/
var _ Provider = (*emptyProvider)(nil)

// empty log
type emptyProvider struct{}

// NewEmptyProvider new empty provider, not any log out
func NewEmptyProvider() Provider {
	return new(emptyProvider)
}

// Emergency Log EMERGENCY level message.
func (sf *emptyProvider) Emergency(string, ...interface{}) {}

// Alert Log ALERT level message.
func (sf *emptyProvider) Alert(string, ...interface{}) {}

// Critical Log CRITICAL level message.
func (sf *emptyProvider) Critical(string, ...interface{}) {}

// Error Log ERROR level message.
func (sf *emptyProvider) Error(string, ...interface{}) {}

// Warning Log WARNING level message.
func (sf *emptyProvider) Warning(string, ...interface{}) {}

// Notice Log NOTICE level message.
func (sf *emptyProvider) Notice(string, ...interface{}) {}

// Info Log INFORMATIONAL level message.
func (sf *emptyProvider) Info(string, ...interface{}) {}

// Debug Log DEBUG level message.
func (sf *emptyProvider) Debug(string, ...interface{}) {}

/******************  internal default log  ************************************/
var logger = NewElog()

// GetElog get internal log instance
func GetElog() *Elog {
	return logger
}

// Mode set enable or diable log output when you has set logger
func Mode(enable bool) {
	logger.Mode(enable)
}

// SetProvider set logger provider
func SetProvider(p Provider) {
	logger.SetProvider(p)
}

// GetProvider Get logger provider
func GetProvider() Provider {
	return logger.GetProvider()
}

// Emergency Log EMERGENCY level message.
func Emergency(format string, v ...interface{}) {
	logger.Emergency(format, v...)
}

// Alert Log ALERT level message.
func Alert(format string, v ...interface{}) {
	logger.Alert(format, v...)
}

// Critical Log CRITICAL level message.
func Critical(format string, v ...interface{}) {
	logger.Critical(format, v...)
}

// Error Log ERROR level message.
func Error(format string, v ...interface{}) {
	logger.Error(format, v...)
}

// Warning Log WARNING level message.
func Warning(format string, v ...interface{}) {
	logger.Warning(format, v...)
}

// Notice Log NOTICE level message.
func Notice(format string, v ...interface{}) {
	logger.Notice(format, v...)

}

// Info Log INFORMATIONAL level message.
func Info(format string, v ...interface{}) {
	logger.Info(format, v...)
}

// Debug Log DEBUG level message.
func Debug(format string, v ...interface{}) {
	logger.Debug(format, v...)
}
