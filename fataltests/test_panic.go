package main

import log "../../loglevel"

func main() {
	log.SetPriority(log.Pall)
	log.SetFlags(0)
	log.SetPrefix("")
	log.Panic("die")
}
