package aggregate_massive

import (
	"sync"
	"time"
)

type Event struct {
	Id        uint32    `db:"id"`
	PersonId  uint32    `db:"person_id"`
	TypeId    uint32    `db:"type_id"`
	Comment   string    `db:"comment"`
	CreatedAt time.Time `db:"created_at"`
}

type Events struct {
	sq             uint32
	mu             sync.RWMutex
	personId       uint32
	list           []*Event
	groupDomainNew map[domainName]map[subdomainName]iD
	groupIpNew     map[ipAddress]map[subdomainName]iD
	groupIpRemoved map[ipAddress]map[subdomainName]iD
}

type iD struct {
	sq, id, criticalId uint32 // нужно сохранить порядок добавления - sq
	val                string
}
type ipAddress string
type domainName string
type subdomainName string
type listSubDomainIp struct {
	d map[subdomainName]iD
	i map[ipAddress]iD
}
type subdomainNameDict struct {
	groupIpNew    map[ipAddress]map[subdomainName]iD
	groupDomainIp map[subdomainName]map[ipAddress]iD
}

const (
	IdIpNew     = iota + 1 // IP адрес - новый
	IdDomainNew            // Домен - новый
)

const joinPattern = ", "
const limitMessage = 32768

const (
	variantDomainNew       = " новый связанный поддомен %s"
	variantDomainNewList   = "ы новые связанные поддомены %s"
	variantIpNew           = " связанный ip адрес %s"
	variantIpNewList       = "ы связанные ip адреса %s"
	variantIpNewDomain     = "а %s"
	variantIpNewDomainList = "ов %s"
	variantAvailable       = "ен"
	variantAvailableList   = "ны"
)

const (
	textDomainNew = "Обнаружен{{.variantDomainNew}} для домена %s"
	textIpNew     = "Обнаружен{{.variantIpNew}} для домен{{.variantIpNewDomain}}"
)
