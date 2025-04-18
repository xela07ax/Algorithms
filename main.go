package main

import (
	"fmt"
	"log"
	"sorting1/aggregate_massive"
	"sorting1/sort_numbers"
)

func main() {
	// Аггрегируем домены с ip адресами
	eventObj := aggregate_massive.NewEvents(13)
	eventObj.AddNewDomain([]string{"calc.domain.ru", "factor.domain.ru"}, "domain.ru")
	eventObj.AddNewDomain([]string{"fox.mozilla.ru", "factor.mozilla.ru"}, "mozilla.ru")
	eventObj.AddNewDomain([]string{"fox.dom.ru", "factor.dom.ru", "doc.dom.ru"}, "dom.ru")
	eventObj.NewIp("14.168.3.2", []string{"calc.domain.ru", "fox.mozilla.ru"})
	eventObj.NewIp("192.168.3.2", []string{"calc.domain.ru", "fox.mozilla.ru"})
	eventObj.NewIp("192.168.3.3", []string{"factor.mozilla.ru", "fox.dom.ru"})
	eventObj.NewIp("192.168.3.4", []string{"factor.mozilla.ru", "fox.dom.ru"})
	eventObj.NewIp("17.0.3.2", []string{"fox.dom.ru"})
	eventObj.NewIp("17.0.3.5", []string{"fox.dom.ru"})
	eventObj.NewIp("176.0.1.3", []string{"fox.dom.ru", "factor.dom.ru", "doc.dom.ru"})

	events, err := eventObj.GetResult()
	if err != nil {
		log.Fatal(err)
	}
	for i, e := range events {
		fmt.Println(i, e)
	}
	fmt.Println("good by!")
	fmt.Println("Erdyakov Aleksey (2025) m")
}
