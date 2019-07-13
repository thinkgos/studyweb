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
func (this *Elog) Mode(enable bool) {
	if enable {
		atomic.StoreUint32(&this.haslog, 1)
	} else {
		atomic.StoreUint32(&this.haslog, 0)
	}
}

// SetProvider set logger provider
func (this *Elog) SetProvider(p Provider) {
	if p != nil {
		this.logger = p
	}
}

// GetProvider Get logger provider
func (this *Elog) GetProvider() Provider {
	return this.logger
}

// Emergency Log EMERGENCY level message.
func (this *Elog) Emergency(format string, v ...interface{}) {
	if atomic.LoadUint32(&this.haslog) == 1 {
		this.logger.Emergency(format, v...)
	}
}

// Alert Log ALERT level message.
func (this *Elog) Alert(format string, v ...interface{}) {
	if atomic.LoadUint32(&this.haslog) == 1 {
		this.logger.Alert(format, v...)
	}
}

// Critical Log CRITICAL level message.
func (this *Elog) Critical(format string, v ...interface{}) {
	if atomic.LoadUint32(&this.haslog) == 1 {
		this.logger.Critical(format, v...)
	}
}

// Error Log ERROR level message.
func (this *Elog) Error(format string, v ...interface{}) {
	if atomic.LoadUint32(&this.haslog) == 1 {
		this.logger.Error(format, v...)
	}
}

// Warning Log WARNING level message.
func (this *Elog) Warning(format string, v ...interface{}) {
	if atomic.LoadUint32(&this.haslog) == 1 {
		this.logger.Warning(format, v...)
	}
}

// Notice Log NOTICE level message.
func (this *Elog) Notice(format string, v ...interface{}) {
	if atomic.LoadUint32(&this.haslog) == 1 {
		this.logger.Notice(format, v...)
	}
}

// Info Log INFORMATIONAL level message.
func (this *Elog) Info(format string, v ...interface{}) {
	if atomic.LoadUint32(&this.haslog) == 1 {
		this.logger.Info(format, v...)
	}
}

// Debug Log DEBUG level message.
func (this *Elog) Debug(format string, v ...interface{}) {
	if atomic.LoadUint32(&this.haslog) == 1 {
		this.logger.Debug(format, v...)
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
func (this *emptyProvider) Emergency(string, ...interface{}) {}

// Alert Log ALERT level message.
func (this *emptyProvider) Alert(string, ...interface{}) {}

// Critical Log CRITICAL level message.
func (this *emptyProvider) Critical(string, ...interface{}) {}

// Error Log ERROR level message.
func (this *emptyProvider) Error(string, ...interface{}) {}

// Warning Log WARNING level message.
func (this *emptyProvider) Warning(string, ...interface{}) {}

// Notice Log NOTICE level message.
func (this *emptyProvider) Notice(string, ...interface{}) {}

// Info Log INFORMATIONAL level message.
func (this *emptyProvider) Info(string, ...interface{}) {}

// Debug Log DEBUG level message.
func (this *emptyProvider) Debug(string, ...interface{}) {}

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
