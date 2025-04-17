package sort_massive

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func compileMessagesDomain(src map[domainName]map[subdomainName]iD, text, variant, variantList, formatPattern string, tip uint32) map[uint32]*Event {
	arrRes := make(map[uint32]*Event, len(src))
	for domainNameVal, subdomainNameList := range src {
		subdomains := make([]string, 0, len(subdomainNameList))
		var sq uint32
		var i int
		for subdomainNameVal, idObj := range subdomainNameList {
			i++
			sq = idObj.sq
			subdomains = append(subdomains, string(subdomainNameVal))
		}
		textSubdomains := subdomains[0]
		variantResult := variant
		if i > 1 {
			variantResult = variantList
			textSubdomains = strings.Join(subdomains, joinPattern)
		}
		msg := strings.ReplaceAll(text, formatPattern, variantResult)
		arrRes[sq] = &Event{
			Comment: fmt.Sprintf(msg, textSubdomains, domainNameVal),
			TypeId:  tip,
		}
	}
	return arrRes
}

func getGroupingIpDomain(groupIpNew map[ipAddress]map[subdomainName]iD) []*listSubDomainIp {
	r := &subdomainNameDict{
		groupIpNew:    groupIpNew,
		groupDomainIp: make(map[subdomainName]map[ipAddress]iD),
	}
	return r.groupingIpDomain()
}

func compileMessagesIp(src map[ipAddress]map[subdomainName]iD, text, variantIp, variantIpList, formatPatternIp, variantDom, variantDomList, formatPatternDom string, tip uint32) map[uint32]*Event {
	arrRes := make(map[uint32]*Event, len(src))
	for _, ls := range getGroupingIpDomain(src) {
		ipList := make([]string, 0, len(ls.i))
		dmList := make([]string, 0, len(ls.d))
		var sq uint32
		for ipAddressVal, _ := range ls.i {
			ipList = append(ipList, string(ipAddressVal))
		}

		for dOMAINVal, _ := range ls.d {
			// после группировки затерялись id и порядок, проставим рандомный из сгруппированного списка
			for _, srcIp := range src {
				if v, ok := srcIp[dOMAINVal]; ok {
					if _, ok := arrRes[sq]; !ok {
						sq = v.sq
						break
					}
				}
			}
			dmList = append(dmList, string(dOMAINVal))
		}
		if sq == 0 { // при неудаче найти свободный порядковый номер, проставим рандомный
			sq = uint32(rand.Intn(1000000))
			if _, ok := arrRes[sq]; ok {
				sq = uint32(rand.Intn(1000000000))
			}
		}
		variantDomResult := variantDom
		textDMaddress := dmList[0]
		if len(dmList) > 1 {
			variantDomResult = variantDomList
			textDMaddress = strings.Join(dmList, joinPattern)
		}
		textipAddress := ipList[0]
		variantIpResult := variantIp
		if len(ipList) > 1 {
			variantIpResult = variantIpList
			textipAddress = strings.Join(ipList, joinPattern)
		}
		variantAvailableResult := variantAvailable
		if len(ipList) > 1 {
			variantAvailableResult = variantAvailableList
		}

		textResult := strings.ReplaceAll(text, formatPatternIp, variantIpResult)
		textResult = strings.ReplaceAll(textResult, formatPatternDom, variantDomResult)
		textResult = strings.ReplaceAll(textResult, "{{.variantAvailable}}", variantAvailableResult) // если паттерна нет, то ничего и не заменит
		arrRes[sq] = &Event{
			Comment:   fmt.Sprintf(textResult, textipAddress, textDMaddress),
			TypeId:    tip,
			CreatedAt: time.Now(),
		}
	}
	return arrRes
}
