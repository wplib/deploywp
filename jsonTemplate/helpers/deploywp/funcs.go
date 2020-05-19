package deploywp

//func LoadDeployWp(jstr string, args ...string) *TypeDeployWp {
//	j := NewJsonFile()
//
//	for range only.Once {
//		var err error
//
//		err = json.Unmarshal([]byte(jstr), &j)
//		j.State.SetError(err)
//		if j.State.IsError() {
//			break
//		}
//
//		err = j.Source.Process()
//		j.State.SetError(err)
//		if j.State.IsError() {
//			break
//		}
//
//		err = j.Target.Process()
//		j.State.SetError(err)
//		if j.State.IsError() {
//			break
//		}
//
//		err = j.Hosts.Process()
//		j.State.SetError(err)
//		if j.State.IsError() {
//			break
//		}
//
//		j.Valid = true
//	}
//
//	return j
//}
