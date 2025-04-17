package aggregate_massive

import (
	"log"
	"sort"
)

func (e *Events) GetResult() ([]*Event, error) {
	e.mu.RLock()
	defer e.mu.RUnlock()
	e.build()
	// инвертируем порядок событий, для удобного чтения в интерфейсе
	if len(e.list) != 0 {
		for i, j := 0, len(e.list)-1; i < j; i, j = i+1, j-1 {
			e.list[i], e.list[j] = e.list[j], e.list[i]
		}
		for _, event := range e.list {
			event.PersonId = e.personId

			if len(event.Comment) > limitMessage {
				log.Printf("error insert event limit comment length detected, current \"%d\": %q", len(event.Comment), event.Comment)
				event.Comment = event.Comment[:limitMessage]
			}
		}
	}

	return e.list, nil
}

func (e *Events) build() {
	msgsDomainNew := compileMessagesDomain(e.groupDomainNew, textDomainNew, variantDomainNew, variantDomainNewList, "{{.variantDomainNew}}", IdDomainNew)
	msgsIpNew := compileMessagesIp(e.groupIpNew, textIpNew, variantIpNew, variantIpNewList, "{{.variantIpNew}}", variantIpNewDomain, variantIpNewDomainList, "{{.variantIpNewDomain}}", IdIpNew)

	// объединяем результаты
	all := make(map[int]*Event)
	e.addMap(all, msgsDomainNew, "domain new")
	e.addMap(all, msgsIpNew, "ip new")

	// сортируем по порядку добавления
	var keys []int
	for k, _ := range all {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, k := range keys {
		e.list = append(e.list, all[k])
	}
}
