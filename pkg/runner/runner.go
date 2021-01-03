package runner

import (
	"fmt"
	"go_psum/pkg/log"
	"os"
	"strings"

	"github.com/cheynewallace/tabby"
	"github.com/shirou/gopsutil/process"
)

type ProcessPair struct {
	Process *process.Process
	Name    string
}

func Run(nameRaw string, exclude string, showDetail int, verbose int) {
	var log = func(text string) {
		if verbose > 0 {
			log.DefaultLogger.Println(text)
		}
	}
	var currentPid = os.Getpid()

	log(fmt.Sprintf("current pid:%d", currentPid))

	log(fmt.Sprintf("nameRaw:[%s], exclude:[%s], showDetail:[%d]", nameRaw, exclude, showDetail))
	processList, err := process.Processes()
	if err != nil {
		panic(err)
	}

	var filteredList []*process.Process
	for _, proc := range processList {
		if proc.Pid != int32(currentPid) {
			filteredList = append(filteredList, proc)
		}
	}
	convertList := filterProcessList(filteredList, nameRaw, exclude, log)
	log("after filter")
	if showDetail == 1 {
		printDetail(convertList, log)
	} else {
		print(convertList, log)
	}

}

func filterProcessList(list []*process.Process, nameRaw string, exclude string, log func(string)) []*ProcessPair {
	var outList []*ProcessPair
	nameRaw = strings.Trim(nameRaw, " ")
	if nameRaw == "" {
		return outList
	}
	var nameList = strings.Split(nameRaw, ",")
	if len(nameList) == 0 {
		return outList
	}
	var isFirst int = 1
	var fnIsInExclude = func(processName string) bool {

		lExclude := strings.Trim(exclude, " ")
		if lExclude == "" {
			log("process filter is empty")
			return false
		}
		excludeList := strings.Split(exclude, ",")
		if len(excludeList) == 0 {
			log("process filter list is empty")
			return false
		}

		if isFirst == 1 {
			isFirst = 0
			for i := 0; i < len(excludeList); i++ {
				piece := excludeList[i]
				piece = strings.Trim(piece, " ")
				log(fmt.Sprintf("exclude name piece %s", piece))
			}
		}

		for i := 0; i < len(excludeList); i++ {
			piece := excludeList[i]
			piece = strings.Trim(piece, " ")
			if piece == "" {
				continue
			}
			if strings.Index(processName, piece) > -1 {
				return true
			}
		}
		log("process filter end")
		return false
	}

	for _, namePiece := range nameList {
		namePiece = strings.Trim(namePiece, " ")
		log(fmt.Sprintf("namePiece:[%s]", namePiece))
		if namePiece == "" {
			continue
		}
	}
	// for _, excludePiece := range
	for i := 0; i < len(list); i++ {
		currentProcess := list[i]
		processName, err := currentProcess.Cmdline()
		if err != nil {
			// panic: open /proc/5155/status: no such file or directory
			if strings.Index(err.Error(), "no such file or directory") > -1 {
				continue
			}
			panic(err)
		}
		log(fmt.Sprintf("processName:%s", processName))
		for _, namePiece := range nameList {
			namePiece = strings.Trim(namePiece, " ")
			// log(fmt.Sprintf("namePiece:[%s]", namePiece))
			if namePiece == "" {
				continue
			}
			interPieceList := strings.Split(namePiece, "__")

			isMatch := 1
			for _, interPiece := range interPieceList {
				interPiece = strings.Trim(interPiece, " ")
				if interPiece == "" {
					continue
				}
				if strings.Index(processName, interPiece) == -1 {
					isMatch = 0
					break
				}
			}

			if isMatch == 1 {
				log("name contains")
				if !fnIsInExclude(processName) {
					log("name filter by exclude")
					outList = append(outList, &ProcessPair{
						Name:    namePiece,
						Process: currentProcess,
					})
				}
			}

		}

	}

	return outList
}

func print(list []*ProcessPair, log func(string)) {
	t := tabby.New()
	var nameList []string
	for _, p := range list {
		isContain := 0
		for _, v := range nameList {
			if v == p.Name {
				isContain = 1
				break
			}
		}
		if isContain == 0 {
			nameList = append(nameList, p.Name)
		}
	}
	// strings.Join()
	log(fmt.Sprintf("available name :[%s]", strings.Join(nameList, ",")))

	t.AddHeader("name", "count", "mem", "open_files", "net_connections")
	for _, name := range nameList {
		var memSize uint64 = 0
		var open_file_count = 0
		var net_connection_count = 0
		var count = 0
		for _, proc := range list {
			processName := proc.Name
			rawProcess := proc.Process

			if processName == name {
				count++
				memInfo, err := rawProcess.MemoryInfo()
				// var memSize uint64 = 0
				if err == nil {
					memSize = memInfo.RSS + memSize
				}
				// pid := rawProcess.Pid
				// openFileCount := 0
				openFileStat, err := rawProcess.OpenFiles()
				if err == nil {
					open_file_count = open_file_count + len(openFileStat)
				} else {
					// errStr := err.Error()
					// if strings.In
					panic(err)

				}

				// net_connection_count := 0
				netcountList, err := rawProcess.Connections()
				if err == nil {
					net_connection_count = net_connection_count + len(netcountList)
				} else {
					panic(err)
				}
			} // processName == name

		} // for _, process

		t.AddLine(
			name,                     // name
			fmt.Sprintf("%d", count), // count
			// fmt.Sprintf("%d", pid),                           // pid
			fmt.Sprintf("%.2fm", float64(memSize)/1024/1024), // mem
			fmt.Sprintf("%d", open_file_count),               // open_files,
			fmt.Sprintf("%d", net_connection_count),          // net_connections
			// cmd,                                              // cmdline
		)
	} // for _, name := range nameList

	t.Print()
}

func printDetail(list []*ProcessPair, log func(string)) {
	t := tabby.New()
	t.AddHeader("name", "pid", "mem", "open_files", "net_connections", "cmdline")
	for _, proc := range list {
		rawProcess := proc.Process
		// 如果 获取失败就是0
		memInfo, err := rawProcess.MemoryInfo()
		var memSize uint64 = 0
		if err == nil {
			memSize = memInfo.RSS + memSize
		}
		pid := rawProcess.Pid
		openFileCount := 0
		openFileStat, err := rawProcess.OpenFiles()
		if err == nil {
			openFileCount = len(openFileStat)
		} else {
			// errStr := err.Error()
			// if strings.In
			panic(err)

		}

		net_connection_count := 0
		netcountList, err := rawProcess.Connections()
		if err == nil {
			net_connection_count = len(netcountList)
		} else {
			panic(err)
		}

		cmd := ""
		cmd, err = rawProcess.Cmdline()

		t.AddLine(proc.Name,
			fmt.Sprintf("%d", pid),                           // pid
			fmt.Sprintf("%.2fm", float64(memSize)/1024/1024), // mem
			fmt.Sprintf("%d", openFileCount),                 // open_files,
			fmt.Sprintf("%d", net_connection_count),          // net_connections
			cmd,                                              // cmdline
		)
	}
	t.Print()
}
