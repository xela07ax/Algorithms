package sort_massive

import "sort"

func (s *subdomainNameDict) groupingIpDomain() []*listSubDomainIp {
	// нужно найти ip адреса которые полностью совпадают для каждого домена, собрать их в список
	// если совпадает частично, то отдельный список ip будет выводиться отдельно для списка дригих доменов
	for ipAddressVal, domainList := range s.groupIpNew {
		for domainNameVal, id := range domainList {
			ipGroup, ok := s.groupDomainIp[domainNameVal]
			if !ok {
				ipGroup = make(map[ipAddress]iD)
				s.groupDomainIp[domainNameVal] = ipGroup
			}
			ipGroup[ipAddressVal] = id
		}
	}
	ffs := make(map[ipAddress]map[subdomainName]map[ipAddress]int)
	for ip, dList := range s.groupIpNew {
		ffs[ip] = make(map[subdomainName]map[ipAddress]int)
		for domain, _ := range dList {
			ipList := s.groupDomainIp[domain]
			for outIp := range ipList {
				outDomainList := s.groupIpNew[outIp]
				for outDomain2, _ := range outDomainList {
					_, ok := dList[outDomain2]
					if ok {
						_, ok = ffs[ip][outDomain2]
						if !ok {
							ffs[ip][outDomain2] = make(map[ipAddress]int)
						}
						ffs[ip][outDomain2][outIp] = 1
					}
				}
			}
		}
	}
	type cont struct {
		ln   int
		dom  subdomainName
		data map[ipAddress]int
	}

	var l []*listSubDomainIp
	// ищем связанные данные по каждому ip и домену
	for ip, domList := range ffs {
		var bigList []cont
		for domain, ipList := range domList {
			ls := cont{
				ln:   len(ipList),
				dom:  domain,
				data: ipList,
			}
			bigList = append(bigList, ls)
		}
		sort.Slice(bigList, func(i, j int) (less bool) {
			return bigList[i].ln < bigList[j].ln
		})

		subs := make(map[subdomainName]iD)
		ips := make(map[ipAddress]iD)
		for _, con := range bigList {
			for ipadrress := range con.data {
				for i, _ := range bigList {
					if _, ok := bigList[i].data[ipadrress]; ok {
						subs[con.dom] = iD{}
						ips[ipadrress] = iD{}
					}
					for i2, _ := range bigList {
						if _, ok := bigList[i2].data[ipadrress]; ok {
							if ipadrress != ip {
								delete(bigList[i2].data, ipadrress)
							}
						}
					}
				}
			}
		}
		l = append(l, &listSubDomainIp{
			d: subs,
			i: ips,
		})
	}

	lDom := make([]*listSubDomainIp, len(l))
	copy(lDom, l)
	// сортируем ip и домены по количеству элементов
	sort.Slice(l, func(i, j int) (less bool) {
		return len(l[i].i) < len(l[j].i)
	})
	sort.Slice(lDom, func(i, j int) (less bool) {
		return len(lDom[i].d) < len(lDom[j].d)
	})
	srtAll := make([]*listSubDomainIp, 0, len(lDom))
	findVal := func(arr []*listSubDomainIp, itm *listSubDomainIp) bool {
		for _, v := range arr {
			if v == itm {
				return true
			}
		}
		return false
	}
	// идем лесенкой по одному, потом по другому списку и добавляем
	for i := 0; i < len(lDom); i++ {
		if !findVal(srtAll, lDom[i]) {
			srtAll = append(srtAll, lDom[i])
		}
		if !findVal(srtAll, l[i]) {
			srtAll = append(srtAll, l[i])
		}

	}
	// инвертируем порядок
	for i, j := 0, len(srtAll)-1; i < j; i, j = i+1, j-1 {
		srtAll[i], srtAll[j] = srtAll[j], srtAll[i]
	}
	// теперь вначале имеем средние по величине показатели для обоих списков
	// нужно убрать пару ip домен которая уже существует
	findPair := func(arr []*listSubDomainIp, dom subdomainName, ip ipAddress) bool {
		for _, v := range arr {
			if _, ok := v.d[dom]; ok {
				if _, ok := v.i[ip]; ok {
					return true
				}
			}
		}
		return false
	}
	findDuble := func(arr []*listSubDomainIp) []*listSubDomainIp {
		out := make([]*listSubDomainIp, len(lDom))
		for i, _ := range out {
			out[i] = &listSubDomainIp{
				d: make(map[subdomainName]iD),
				i: make(map[ipAddress]iD),
			}
		}
		for i, v := range arr {
			for d, _ := range v.d {
				for ip, _ := range v.i {
					if !findPair(out, d, ip) {
						out[i].d[d] = iD{}
						out[i].i[ip] = iD{}
					}
				}
			}
		}
		return out
	}
	cleanedDubles := make([]*listSubDomainIp, 0, len(lDom))
	// удалим списки без элементов
	for _, v := range findDuble(srtAll) {
		if len(v.d) != 0 {
			cleanedDubles = append(cleanedDubles, v)
		}
	}
	return cleanedDubles
}
