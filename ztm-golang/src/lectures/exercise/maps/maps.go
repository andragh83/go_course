//--Summary:
//  Create a program to display server status. The server statuses are
//  defined as constants, and the servers are represented as strings
//  in the `servers` slice.
//
//--Requirements:
//* Create a function to print server status displaying:
//  - number of servers
//  - number of servers for each status (Online, Offline, Maintenance, Retired)
//* Create a map using the server names as the key and the server status
//  as the value
//* Set all of the server statuses to `Online` when creating the map
//* After creating the map, perform the following actions:
//  - call display server info function
//  - change server status of `darkstar` to `Retired`
//  - change server status of `aiur` to `Offline`
//  - call display server info function
//  - change server status of all servers to `Maintenance`
//  - call display server info function

package main

import "fmt"

const (
	Online      = 0
	Offline     = 1
	Maintenance = 2
	Retired     = 3
)

func printServerstatus(serverStatus map[string]int) {
	
	fmt.Println("number of servers", len(serverStatus))
	// on, off, mnt, ret := 0,0,0,0
	// for _, status := range serverStatus {
	// 	if status == 0 {
	// 		on++
	// 	} else if status == 1 {
	// 		off++
	// 	} else if status == 2 {
	// 		mnt++
	// 	} else {
	// 		ret++
	// 	}
	// }
	// fmt.Println("number of servers online", on)
	// fmt.Println("number of servers offline", off)
	// fmt.Println("number of servers in maintenance", mnt)
	// fmt.Println("number of servers retired", ret)

	stats := make(map[int]int)

	for _, status := range serverStatus {
		switch status {
		case Online: stats[Online] += 1
		case Offline: stats[Offline] +=1
		case Maintenance: stats[Maintenance] +=1
		case Retired: stats[Retired] +=1
		default: panic("unhandled server status")
		}
	}
	fmt.Println("number of servers online", stats[Online])
	fmt.Println("number of servers offline", stats[Offline])
	fmt.Println("number of servers in maintenance", stats[Maintenance])
	fmt.Println("number of servers retired", stats[Retired])
}

func main() {
	servers := []string{"darkstar", "aiur", "omicron", "w359", "baseline"}
	serverStatus := make(map[string]int)
	for _, server :=range servers {
		serverStatus[server] = 0
	}
	printServerstatus(serverStatus)
	serverStatus["darkstar"] = 3
	serverStatus["aiur"] = 1
	printServerstatus(serverStatus)

	for server := range serverStatus {
		serverStatus[server] = 2
	}

	printServerstatus(serverStatus)

}
