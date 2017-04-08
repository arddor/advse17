package main

func main() {
	s := Server{}
	s.Initialize("ase_timeseries:27017")
	s.Run(":8080")
}
