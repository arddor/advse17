// ase_api

package main

func main() {
	s := Server{}
	s.Initialize("ase_timeseries:28015")
	s.Run(":8000")
}

// Needs endpoint to register terms
// terms need to be stored into db and set to active (a simple flag will do)

// Needs endpoint to DE-register term
// term needs to be deactivated (a simple flag will do)
