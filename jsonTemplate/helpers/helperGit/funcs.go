package helperGit


// Non-helper function
//func (me *TypeGit) NonHelperExec(cmd string, args ...string) *helperTypes.TypeExecCommand {
//	var ret helperTypes.TypeExecCommand
//
//	for range only.Once {
//		if cmd == "" {
//			break
//		}
//		ret.Exe = cmd
//		ret.Args = append(ret.Args, me.GitOptions...)
//		ret.Args = append(ret.Args, args...)
//
//		cwd := helperSystem.HelperGetwd()
//		cd := helperSystem.HelperChdir(me.Base.Path)
//		if cd.IsError() {
//			ux.PrintfError("Cannot change directory to '%s'", me.Base.Path)
//			break
//		}
//
//		ret.Output, ret.Error = me.client.Exec(ret.Exe, ret.Args...)
//		if ret.Error != nil {
//			ret.Exit = 1	// Fake an exit code.
//		}
//
//		cd = helperSystem.HelperChdir(me.Base.Path)
//		if cd.IsError() {
//			ux.PrintfError("Cannot change back to directory '%s'", cwd.Path)
//			break
//		}
//	}
//
//	return &ret
//}
