package sort_massive

import (
	"log"
	"sync"
)

func NewEvents(personId uint32) *Events {
	return &Events{
		sq:             9,
		mu:             sync.RWMutex{},
		personId:       personId,
		groupDomainNew: make(map[domainName]map[subdomainName]iD),
		groupIpNew:     make(map[ipAddress]map[subdomainName]iD),
		groupIpRemoved: make(map[ipAddress]map[subdomainName]iD),
	}
}

func (e *Events) NewIp(address string, domains []string) {
	e.mu.Lock()
	defer e.mu.Unlock()
	sq := e.getId()
	name := ipAddress(address)
	group, ok := e.groupIpNew[name]
	if !ok {
		group = make(map[subdomainName]iD)
		e.groupIpNew[name] = group
	}
	for _, v := range domains {
		group[subdomainName(v)] = iD{
			sq: sq,
			id: 0,
		}
	}
}

func (e *Events) AddNewDomain(subdomains []string, headDomain string) {
	e.mu.Lock()
	defer e.mu.Unlock()
	sq := e.getId()
	name := domainName(headDomain)
	group, ok := e.groupDomainNew[name]
	if !ok {
		group = make(map[subdomainName]iD)
		e.groupDomainNew[name] = group
	}
	for _, domain := range subdomains {
		group[subdomainName(domain)] = iD{
			sq: sq,
			id: 0,
		}
	}
}

func (e *Events) getId() uint32 {
	e.sq = e.sq + 1
	return e.sq
}

func (e *Events) addMap(all map[int]*Event, items map[uint32]*Event, nameGroup string) {
	for k, v := range items {
		if _, ok := all[int(k)]; !ok {
			all[int(k)] = v
		} else {
			log.Printf("error dublicate %q sequence \"%d\":%v", nameGroup, k, v)
		}
	}
}
