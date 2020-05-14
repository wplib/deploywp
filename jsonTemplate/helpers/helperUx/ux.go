package helperUx

import (
	"github.com/wplib/deploywp/ux"
)


// Was going to implement as a Getter interface, but no need.

//var _ ux.UxGetter = (*Ux)(nil)
//type Ux ux.Ux
//
//func (u *Ux) Open() error {
//	panic("implement me")
//}
//
//func (u *Ux) Close() {
//	panic("implement me")
//}
//
//func (u *Ux) PrintfWhite(format string, args ...interface{}) {
//	panic("implement me")
//}
//
//func (u *Ux) PrintfCyan(format string, args ...interface{}) {
//	panic("implement me")
//}
//
//func (u *Ux) PrintfYellow(format string, args ...interface{}) {
//	panic("implement me")
//}
//
//func (u *Ux) PrintfRed(format string, args ...interface{}) {
//	panic("implement me")
//}
//
//func (u *Ux) PrintfGreen(format string, args ...interface{}) {
//	panic("implement me")
//}
//
//func (u *Ux) PrintfBlue(format string, args ...interface{}) {
//	panic("implement me")
//}
//
//func (u *Ux) PrintfMagenta(format string, args ...interface{}) {
//	panic("implement me")
//}
//
//func (u *Ux) SprintfWhite(format string, args ...interface{}) string {
//	panic("implement me")
//}
//
//func (u *Ux) SprintfCyan(format string, args ...interface{}) string {
//	panic("implement me")
//}
//
//func (u *Ux) SprintfYellow(format string, args ...interface{}) string {
//	panic("implement me")
//}
//
//func (u *Ux) SprintfRed(format string, args ...interface{}) string {
//	panic("implement me")
//}
//
//func (u *Ux) SprintfGreen(format string, args ...interface{}) string {
//	panic("implement me")
//}
//
//func (u *Ux) SprintfBlue(format string, args ...interface{}) string {
//	panic("implement me")
//}
//
//func (u *Ux) SprintfMagenta(format string, args ...interface{}) string {
//	panic("implement me")
//}
//
//func (u *Ux) Sprintf(format string, args ...interface{}) string {
//	panic("implement me")
//}
//
//func (u *Ux) Printf(format string, args ...interface{}) {
//	panic("implement me")
//}
//
//func (u *Ux) SprintfOk(format string, args ...interface{}) string {
//	panic("implement me")
//}
//
//func (u *Ux) PrintfOk(format string, args ...interface{}) {
//	panic("implement me")
//}
//
//func (u *Ux) SprintfWarning(format string, args ...interface{}) string {
//	panic("implement me")
//}
//
//func (u *Ux) PrintfWarning(format string, args ...interface{}) {
//	panic("implement me")
//}
//
//func (u *Ux) SprintfError(format string, args ...interface{}) string {
//	panic("implement me")
//}
//
//func (u *Ux) PrintfError(format string, args ...interface{}) {
//	panic("implement me")
//}
//
//func (u *Ux) SprintError(err error) string {
//	panic("implement me")
//}
//
//func (u *Ux) PrintError(err error) {
//	panic("implement me")
//}
//
//func (u *Ux) GetTerminalSize() (int, int, error) {
//	panic("implement me")
//}


/////////////////////////////////////

func HelperPrintfWhite(format string, args ...interface{}) string {
	return ux.SprintfWhite(format, args...)
}

func HelperPrintfCyan(format string, args ...interface{}) string {
	return ux.SprintfCyan(format, args...)
}

func HelperPrintfYellow(format string, args ...interface{}) string {
	return ux.SprintfYellow(format, args...)
}

func HelperPrintfRed(format string, args ...interface{}) string {
	return ux.SprintfRed(format, args...)
}

func HelperPrintfGreen(format string, args ...interface{}) string {
	return ux.SprintfGreen(format, args...)
}

func HelperPrintfBlue(format string, args ...interface{}) string {
	return ux.SprintfBlue(format, args...)
}

func HelperPrintfMagenta(format string, args ...interface{}) string {
	return ux.SprintfMagenta(format, args...)
}

func HelperPrintf(format string, args ...interface{}) string {
	return ux.Sprintf(format, args...)
}

func HelperPrintfOk(format string, args ...interface{}) string {
	return ux.SprintfOk(format, args...)
}

func HelperPrintfWarning(format string, args ...interface{}) string {
	return ux.SprintfWarning(format, args...)
}

func HelperPrintfError(format string, args ...interface{}) string {
	return ux.SprintfError(format, args...)
}

func HelperPrintError(err error) string {
	return ux.SprintError(err)
}
