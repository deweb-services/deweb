package types

func (m *TransferOffer) Equal(offer *TransferOffer) bool {
	if m.Price != offer.Price {
		return false
	}
	if m.ExpectedOwnerAddress != offer.ExpectedOwnerAddress {
		return false
	}
	return true
}

func (m *DNSRecords) Equal(dnsRec *DNSRecords) bool {
	if m.Type != dnsRec.Type {
		return false
	}
	for _, mVal := range m.Values {
		var recFound bool
		for _, tVal := range dnsRec.Values {
			if tVal == mVal {
				recFound = true
				break
			}
		}
		if !recFound {
			return false
		}
	}
	return true
}
