package helperUx

import (
	"github.com/wplib/deploywp/ux"
)


//func HelperOpen() error {
//	var err error
//
//	for range only.Once {
//		Color = aurora.NewAurora(true)
//
//		//err = termui.Init();
//		//if err != nil {
//		//	fmt.Printf("failed to initialize termui: %v", err)
//		//	break
//		//}
//
//		_defined = true
//	}
//
//	return err
//}


//func HelperClose() {
//	if _defined {
//		//termui.Close()
//	}
//}


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
