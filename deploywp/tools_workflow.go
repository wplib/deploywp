package deploywp

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/newclarity/scribeHelpers/ux"
	"os"
	"time"
)


// This is an alternative to running templates.
// In theory, the code here, should be able to be replicated in a template file without modification.
func (dwp *TypeDeployWp) Build() *ux.State {
	if state := dwp.IsNil(); state.IsError() {
		return state
	}

	for range onlyOnce {
		dwp.Print.Notify(0, "%s v%s", dwp.Runtime.CmdFile, dwp.Runtime.CmdVersion)
		dwp.Print.Notify(0, "args: %s", dwp.Runtime.GetArgs())


		{
			srcGitRef := dwp.OpenSourceRepo()
			if srcGitRef.State.IsError() {
				dwp.State = srcGitRef.State
				break
			}
		}


		{
			destinationGitRef := dwp.OpenDestinationRepo()
			if destinationGitRef.State.IsError() {
				dwp.State = destinationGitRef.State
				break
			}
			dwp.State = dwp.CleanRepo(destinationGitRef, true)
			if dwp.State.IsError() {
				break
			}
		}


		{
			srcPath := dwp.GetSourceAbsPaths()
			dstPath := dwp.GetDestinationAbsPaths()
			dwp.State = dwp.UpdateDestination(srcPath, dstPath)
			if dwp.State.IsError() {
				break
			}
		}


		{
			dwp.State = dwp.RunComposer(dwp.GetDestinationAbsPaths().GetBasePath(), "install")
			if dwp.State.IsError() {
				break
			}
		}

		os.Exit(1)

		// dwp.Print.Intent("Increment BUILD within destination repository")
		//dwp.State = dwp.OpenDestinationRepo()
		time.Sleep(time.Second * 2)	// Simulate
		if dwp.State.IsError() {
			break
		}


		// dwp.Print.Intent("Commit destination repository to Pantheon")
		//dwp.State = dwp.OpenDestinationRepo()
		time.Sleep(time.Second * 2)	// Simulate
		if dwp.State.IsError() {
			break
		}


		// dwp.Print.Intent("Commit source repository")
		//dwp.State = dwp.OpenDestinationRepo()
		time.Sleep(time.Second * 2)	// Simulate
		if dwp.State.IsError() {
			break
		}
	}

	return dwp.State
}




//func (dwp *TypeDeployWp) test1() {
//	stateError := ux.NewState("test", false)
//	stateError.SetError("Hey it failed")
//
//	dwp.Print.Intent("test 1 part 1")
//	dwp.Print.IntentResponse(stateError)
//	dwp.Print.Intent("test 1 part 2")
//	dwp.Print.IntentResponse(stateError)
//}
//func (dwp *TypeDeployWp) test2() {
//	stateOk := ux.NewState("test2", false)
//	stateOk.SetOk("All OK")
//
//	dwp.Print.Intent("test 2 part 1")
//	dwp.Print.IntentResponse(stateOk)
//	dwp.Print.Intent("test 2 part 2")
//	dwp.test3()
//	dwp.Print.IntentResponse(stateOk)
//	dwp.Print.Intent("test 2 part 3")
//}
//func (dwp *TypeDeployWp) test3() {
//	stateOk := ux.NewState("test3", false)
//	stateOk.SetOk("All OK")
//
//	dwp.Print.Intent( "test 3 part 1")
//	dwp.Print.IntentResponse(stateOk)
//	dwp.Print.Intent("test 3 part 2")
//	dwp.Print.IntentResponse(stateOk)
//}


func (dwp *TypeDeployWp) TestNewUxState() *ux.State {
	if state := dwp.IsNil(); state.IsError() {
		return state
	}

	//stateOk := ux.NewState("error", false)
	//stateOk.SetOk()
	//stateError := ux.NewState("ok", false)
	//stateError.SetError("Hey it failed")
	//
	//dwp.Print.Intent( "First thingy - item 1")
	//dwp.Print.IntentResponse(stateOk)
	//dwp.Print.Intent( "First thingy - item 2")
	//dwp.Print.IntentResponse(stateOk)
	//dwp.Print.Intent( "First thingy - item 3")
	//dwp.Print.IntentResponse(stateError)
	//
	//dwp.test1()
	//
	//dwp.Print.Intent( "Second thingy - item 1")
	//dwp.Print.IntentResponse(stateOk)
	//dwp.Print.Intent( "Second thingy - item 2")
	//dwp.Print.Intent( "Second thingy - item 3")
	//dwp.Print.IntentResponse(stateError)
	//
	//dwp.test2()
	//
	//dwp.Print.Intent( "Third thingy - item 1")
	//dwp.Print.IntentResponse(stateOk)
	//dwp.Print.IntentResponse(stateOk)
	//dwp.Print.Intent( "Third thingy - item 3")
	//dwp.Print.IntentResponse(stateError)

	// Testing ux.State changes.
	var foo []string
	foo = []string{"one", "two", ""}
	dwp.State.SetResponse(&foo)
	foor := dwp.State.GetResponse()
	spew.Dump(foor)
	fmt.Printf("GetType: %s\n", foor.GetType().String())

	var foo2 *[]string
	foo2 = &foo
	dwp.State.SetResponse(&foo2)
	foo2r := dwp.State.GetResponse()
	spew.Dump(foo2r)
	fmt.Printf("GetType: %s\n", foo2r.GetType().String())

	foo3i := "hello"
	foo3 := &foo3i
	dwp.State.SetResponse(foo3)
	foo3r := dwp.State.GetResponse()
	spew.Dump(foo3r)
	fmt.Printf("GetType: %s\n", foo3r.GetType().String())

	// Testing ux.State changes.
	os.Exit(1)
	return dwp.State
}
