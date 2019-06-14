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
	Informational(format string, v ...interface{})
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
func NewElog(p Provider) *Elog {
	l := &Elog{haslog: 0}
	if p != nil {
		l.logger = p
	} else {
		l.logger = new(emptyProvider)
	}
	return l
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

// Informational Log INFORMATIONAL level message.
func (this *Elog) Informational(format string, v ...interface{}) {
	if atomic.LoadUint32(&this.haslog) == 1 {
		this.logger.Informational(format, v...)
	}
}

// Debug Log DEBUG level message.
func (this *Elog) Debug(format string, v ...interface{}) {
	if atomic.LoadUint32(&this.haslog) == 1 {
		this.logger.Debug(format, v...)
	}
}

/**************************  empty log  ************************************/
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

// Informational Log INFORMATIONAL level message.
func (this *emptyProvider) Informational(string, ...interface{}) {}

// Debug Log DEBUG level message.
func (this *emptyProvider) Debug(string, ...interface{}) {}

/******************  internal default log  ************************************/
var log = NewElog(nil)

// GetElog get internal log instance
func GetElog() *Elog {
	return log
}

// Mode set enable or diable log output when you has set logger
func Mode(enable bool) {
	log.Mode(enable)
}

// SetProvider set logger provider
func SetProvider(p Provider) {
	log.SetProvider(p)
}

// GetProvider Get logger provider
func GetProvider() Provider {
	return log.GetProvider()
}

// Emergency Log EMERGENCY level message.
func Emergency(format string, v ...interface{}) {
	log.Emergency(format, v...)
}

// Alert Log ALERT level message.
func Alert(format string, v ...interface{}) {
	log.Alert(format, v...)
}

// Critical Log CRITICAL level message.
func Critical(format string, v ...interface{}) {
	log.Critical(format, v...)
}

// Error Log ERROR level message.
func Error(format string, v ...interface{}) {
	log.Error(format, v...)
}

// Warning Log WARNING level message.
func Warning(format string, v ...interface{}) {
	log.Warning(format, v...)
}

// Notice Log NOTICE level message.
func Notice(format string, v ...interface{}) {
	log.Notice(format, v...)

}

// Informational Log INFORMATIONAL level message.
func Informational(format string, v ...interface{}) {
	log.Informational(format, v...)
}

// Debug Log DEBUG level message.
func Debug(format string, v ...interface{}) {
	log.Debug(format, v...)
}
