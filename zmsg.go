// *** DO NOT MODIFY ***
// AUTOGENERATED BY go generate

package dns

import (
	"encoding/base64"
	"net"
)

func (rr *A) pack() int {
	l := rr.Hdr.len()
	l += net.IPv4len // A
	return l
}
func (rr *A) unpack() int {
	l := rr.Hdr.len()
	l += net.IPv4len // A
	return l
}
func (rr *AAAA) pack() int {
	l := rr.Hdr.len()
	l += net.IPv6len // AAAA
	return l
}
func (rr *AAAA) unpack() int {
	l := rr.Hdr.len()
	l += net.IPv6len // AAAA
	return l
}
func (rr *AFSDB) pack() int {
	l := rr.Hdr.len()
	l += 2 // Subtype
	l += len(rr.Hostname) + 1
	return l
}
func (rr *AFSDB) unpack() int {
	l := rr.Hdr.len()
	l += 2 // Subtype
	l += len(rr.Hostname) + 1
	return l
}
func (rr *ANY) pack() int {
	l := rr.Hdr.len()
	return l
}
func (rr *ANY) unpack() int {
	l := rr.Hdr.len()
	return l
}
func (rr *CAA) pack() int {
	l := rr.Hdr.len()
	l += 1 // Flag
	l += len(rr.Tag) + 1
	l += len(rr.Value)
	return l
}
func (rr *CAA) unpack() int {
	l := rr.Hdr.len()
	l += 1 // Flag
	l += len(rr.Tag) + 1
	l += len(rr.Value)
	return l
}
func (rr *CERT) pack() int {
	l := rr.Hdr.len()
	l += 2 // Type
	l += 2 // KeyTag
	l += 1 // Algorithm
	l += base64.StdEncoding.DecodedLen(len(rr.Certificate))
	return l
}
func (rr *CERT) unpack() int {
	l := rr.Hdr.len()
	l += 2 // Type
	l += 2 // KeyTag
	l += 1 // Algorithm
	l += base64.StdEncoding.DecodedLen(len(rr.Certificate))
	return l
}
func (rr *CNAME) pack() int {
	l := rr.Hdr.len()
	l += len(rr.Target) + 1
	return l
}
func (rr *CNAME) unpack() int {
	l := rr.Hdr.len()
	l += len(rr.Target) + 1
	return l
}
func (rr *DHCID) pack() int {
	l := rr.Hdr.len()
	l += base64.StdEncoding.DecodedLen(len(rr.Digest))
	return l
}
func (rr *DHCID) unpack() int {
	l := rr.Hdr.len()
	l += base64.StdEncoding.DecodedLen(len(rr.Digest))
	return l
}
func (rr *DNAME) pack() int {
	l := rr.Hdr.len()
	l += len(rr.Target) + 1
	return l
}
func (rr *DNAME) unpack() int {
	l := rr.Hdr.len()
	l += len(rr.Target) + 1
	return l
}
func (rr *DNSKEY) pack() int {
	l := rr.Hdr.len()
	l += 2 // Flags
	l += 1 // Protocol
	l += 1 // Algorithm
	l += base64.StdEncoding.DecodedLen(len(rr.PublicKey))
	return l
}
func (rr *DNSKEY) unpack() int {
	l := rr.Hdr.len()
	l += 2 // Flags
	l += 1 // Protocol
	l += 1 // Algorithm
	l += base64.StdEncoding.DecodedLen(len(rr.PublicKey))
	return l
}
func (rr *DS) pack() int {
	l := rr.Hdr.len()
	l += 2 // KeyTag
	l += 1 // Algorithm
	l += 1 // DigestType
	l += len(rr.Digest)/2 + 1
	return l
}
func (rr *DS) unpack() int {
	l := rr.Hdr.len()
	l += 2 // KeyTag
	l += 1 // Algorithm
	l += 1 // DigestType
	l += len(rr.Digest)/2 + 1
	return l
}
func (rr *EID) pack() int {
	l := rr.Hdr.len()
	l += len(rr.Endpoint)/2 + 1
	return l
}
func (rr *EID) unpack() int {
	l := rr.Hdr.len()
	l += len(rr.Endpoint)/2 + 1
	return l
}
func (rr *EUI48) pack() int {
	l := rr.Hdr.len()
	l += 6 // Address
	return l
}
func (rr *EUI48) unpack() int {
	l := rr.Hdr.len()
	l += 6 // Address
	return l
}
func (rr *EUI64) pack() int {
	l := rr.Hdr.len()
	l += 8 // Address
	return l
}
func (rr *EUI64) unpack() int {
	l := rr.Hdr.len()
	l += 8 // Address
	return l
}
func (rr *GID) pack() int {
	l := rr.Hdr.len()
	l += 4 // Gid
	return l
}
func (rr *GID) unpack() int {
	l := rr.Hdr.len()
	l += 4 // Gid
	return l
}
func (rr *GPOS) pack() int {
	l := rr.Hdr.len()
	l += len(rr.Longitude) + 1
	l += len(rr.Latitude) + 1
	l += len(rr.Altitude) + 1
	return l
}
func (rr *GPOS) unpack() int {
	l := rr.Hdr.len()
	l += len(rr.Longitude) + 1
	l += len(rr.Latitude) + 1
	l += len(rr.Altitude) + 1
	return l
}
func (rr *HINFO) pack() int {
	l := rr.Hdr.len()
	l += len(rr.Cpu) + 1
	l += len(rr.Os) + 1
	return l
}
func (rr *HINFO) unpack() int {
	l := rr.Hdr.len()
	l += len(rr.Cpu) + 1
	l += len(rr.Os) + 1
	return l
}
func (rr *HIP) pack() int {
	l := rr.Hdr.len()
	l += 1 // HitLength
	l += 1 // PublicKeyAlgorithm
	l += 2 // PublicKeyLength
	l += len(rr.Hit)/2 + 1
	l += base64.StdEncoding.DecodedLen(len(rr.PublicKey))
	for _, x := range rr.RendezvousServers {
		l += len(x) + 1
	}
	return l
}
func (rr *HIP) unpack() int {
	l := rr.Hdr.len()
	l += 1 // HitLength
	l += 1 // PublicKeyAlgorithm
	l += 2 // PublicKeyLength
	l += len(rr.Hit)/2 + 1
	l += base64.StdEncoding.DecodedLen(len(rr.PublicKey))
	for _, x := range rr.RendezvousServers {
		l += len(x) + 1
	}
	return l
}
func (rr *KX) pack() int {
	l := rr.Hdr.len()
	l += 2 // Preference
	l += len(rr.Exchanger) + 1
	return l
}
func (rr *KX) unpack() int {
	l := rr.Hdr.len()
	l += 2 // Preference
	l += len(rr.Exchanger) + 1
	return l
}
func (rr *L32) pack() int {
	l := rr.Hdr.len()
	l += 2           // Preference
	l += net.IPv4len // Locator32
	return l
}
func (rr *L32) unpack() int {
	l := rr.Hdr.len()
	l += 2           // Preference
	l += net.IPv4len // Locator32
	return l
}
func (rr *L64) pack() int {
	l := rr.Hdr.len()
	l += 2 // Preference
	l += 8 // Locator64
	return l
}
func (rr *L64) unpack() int {
	l := rr.Hdr.len()
	l += 2 // Preference
	l += 8 // Locator64
	return l
}
func (rr *LOC) pack() int {
	l := rr.Hdr.len()
	l += 1 // Version
	l += 1 // Size
	l += 1 // HorizPre
	l += 1 // VertPre
	l += 4 // Latitude
	l += 4 // Longitude
	l += 4 // Altitude
	return l
}
func (rr *LOC) unpack() int {
	l := rr.Hdr.len()
	l += 1 // Version
	l += 1 // Size
	l += 1 // HorizPre
	l += 1 // VertPre
	l += 4 // Latitude
	l += 4 // Longitude
	l += 4 // Altitude
	return l
}
func (rr *LP) pack() int {
	l := rr.Hdr.len()
	l += 2 // Preference
	l += len(rr.Fqdn) + 1
	return l
}
func (rr *LP) unpack() int {
	l := rr.Hdr.len()
	l += 2 // Preference
	l += len(rr.Fqdn) + 1
	return l
}
func (rr *MB) pack() int {
	l := rr.Hdr.len()
	l += len(rr.Mb) + 1
	return l
}
func (rr *MB) unpack() int {
	l := rr.Hdr.len()
	l += len(rr.Mb) + 1
	return l
}
func (rr *MD) pack() int {
	l := rr.Hdr.len()
	l += len(rr.Md) + 1
	return l
}
func (rr *MD) unpack() int {
	l := rr.Hdr.len()
	l += len(rr.Md) + 1
	return l
}
func (rr *MF) pack() int {
	l := rr.Hdr.len()
	l += len(rr.Mf) + 1
	return l
}
func (rr *MF) unpack() int {
	l := rr.Hdr.len()
	l += len(rr.Mf) + 1
	return l
}
func (rr *MG) pack() int {
	l := rr.Hdr.len()
	l += len(rr.Mg) + 1
	return l
}
func (rr *MG) unpack() int {
	l := rr.Hdr.len()
	l += len(rr.Mg) + 1
	return l
}
func (rr *MINFO) pack() int {
	l := rr.Hdr.len()
	l += len(rr.Rmail) + 1
	l += len(rr.Email) + 1
	return l
}
func (rr *MINFO) unpack() int {
	l := rr.Hdr.len()
	l += len(rr.Rmail) + 1
	l += len(rr.Email) + 1
	return l
}
func (rr *MR) pack() int {
	l := rr.Hdr.len()
	l += len(rr.Mr) + 1
	return l
}
func (rr *MR) unpack() int {
	l := rr.Hdr.len()
	l += len(rr.Mr) + 1
	return l
}
func (rr *MX) pack() int {
	l := rr.Hdr.len()
	l += 2 // Preference
	l += len(rr.Mx) + 1
	return l
}
func (rr *MX) unpack() int {
	l := rr.Hdr.len()
	l += 2 // Preference
	l += len(rr.Mx) + 1
	return l
}
func (rr *NAPTR) pack() int {
	l := rr.Hdr.len()
	l += 2 // Order
	l += 2 // Preference
	l += len(rr.Flags) + 1
	l += len(rr.Service) + 1
	l += len(rr.Regexp) + 1
	l += len(rr.Replacement) + 1
	return l
}
func (rr *NAPTR) unpack() int {
	l := rr.Hdr.len()
	l += 2 // Order
	l += 2 // Preference
	l += len(rr.Flags) + 1
	l += len(rr.Service) + 1
	l += len(rr.Regexp) + 1
	l += len(rr.Replacement) + 1
	return l
}
func (rr *NID) pack() int {
	l := rr.Hdr.len()
	l += 2 // Preference
	l += 8 // NodeID
	return l
}
func (rr *NID) unpack() int {
	l := rr.Hdr.len()
	l += 2 // Preference
	l += 8 // NodeID
	return l
}
func (rr *NIMLOC) pack() int {
	l := rr.Hdr.len()
	l += len(rr.Locator)/2 + 1
	return l
}
func (rr *NIMLOC) unpack() int {
	l := rr.Hdr.len()
	l += len(rr.Locator)/2 + 1
	return l
}
func (rr *NINFO) pack() int {
	l := rr.Hdr.len()
	for _, x := range rr.ZSData {
		l += len(x) + 1
	}
	return l
}
func (rr *NINFO) unpack() int {
	l := rr.Hdr.len()
	for _, x := range rr.ZSData {
		l += len(x) + 1
	}
	return l
}
func (rr *NS) pack() int {
	l := rr.Hdr.len()
	l += len(rr.Ns) + 1
	return l
}
func (rr *NS) unpack() int {
	l := rr.Hdr.len()
	l += len(rr.Ns) + 1
	return l
}
func (rr *NSAPPTR) pack() int {
	l := rr.Hdr.len()
	l += len(rr.Ptr) + 1
	return l
}
func (rr *NSAPPTR) unpack() int {
	l := rr.Hdr.len()
	l += len(rr.Ptr) + 1
	return l
}
func (rr *NSEC3PARAM) pack() int {
	l := rr.Hdr.len()
	l += 1 // Hash
	l += 1 // Flags
	l += 2 // Iterations
	l += 1 // SaltLength
	l += len(rr.Salt)/2 + 1
	return l
}
func (rr *NSEC3PARAM) unpack() int {
	l := rr.Hdr.len()
	l += 1 // Hash
	l += 1 // Flags
	l += 2 // Iterations
	l += 1 // SaltLength
	l += len(rr.Salt)/2 + 1
	return l
}
func (rr *OPENPGPKEY) pack() int {
	l := rr.Hdr.len()
	l += base64.StdEncoding.DecodedLen(len(rr.PublicKey))
	return l
}
func (rr *OPENPGPKEY) unpack() int {
	l := rr.Hdr.len()
	l += base64.StdEncoding.DecodedLen(len(rr.PublicKey))
	return l
}
func (rr *PTR) pack() int {
	l := rr.Hdr.len()
	l += len(rr.Ptr) + 1
	return l
}
func (rr *PTR) unpack() int {
	l := rr.Hdr.len()
	l += len(rr.Ptr) + 1
	return l
}
func (rr *PX) pack() int {
	l := rr.Hdr.len()
	l += 2 // Preference
	l += len(rr.Map822) + 1
	l += len(rr.Mapx400) + 1
	return l
}
func (rr *PX) unpack() int {
	l := rr.Hdr.len()
	l += 2 // Preference
	l += len(rr.Map822) + 1
	l += len(rr.Mapx400) + 1
	return l
}
func (rr *RFC3597) pack() int {
	l := rr.Hdr.len()
	l += len(rr.Rdata)/2 + 1
	return l
}
func (rr *RFC3597) unpack() int {
	l := rr.Hdr.len()
	l += len(rr.Rdata)/2 + 1
	return l
}
func (rr *RKEY) pack() int {
	l := rr.Hdr.len()
	l += 2 // Flags
	l += 1 // Protocol
	l += 1 // Algorithm
	l += base64.StdEncoding.DecodedLen(len(rr.PublicKey))
	return l
}
func (rr *RKEY) unpack() int {
	l := rr.Hdr.len()
	l += 2 // Flags
	l += 1 // Protocol
	l += 1 // Algorithm
	l += base64.StdEncoding.DecodedLen(len(rr.PublicKey))
	return l
}
func (rr *RP) pack() int {
	l := rr.Hdr.len()
	l += len(rr.Mbox) + 1
	l += len(rr.Txt) + 1
	return l
}
func (rr *RP) unpack() int {
	l := rr.Hdr.len()
	l += len(rr.Mbox) + 1
	l += len(rr.Txt) + 1
	return l
}
func (rr *RRSIG) pack() int {
	l := rr.Hdr.len()
	l += 2 // TypeCovered
	l += 1 // Algorithm
	l += 1 // Labels
	l += 4 // OrigTtl
	l += 4 // Expiration
	l += 4 // Inception
	l += 2 // KeyTag
	l += len(rr.SignerName) + 1
	l += base64.StdEncoding.DecodedLen(len(rr.Signature))
	return l
}
func (rr *RRSIG) unpack() int {
	l := rr.Hdr.len()
	l += 2 // TypeCovered
	l += 1 // Algorithm
	l += 1 // Labels
	l += 4 // OrigTtl
	l += 4 // Expiration
	l += 4 // Inception
	l += 2 // KeyTag
	l += len(rr.SignerName) + 1
	l += base64.StdEncoding.DecodedLen(len(rr.Signature))
	return l
}
func (rr *RT) pack() int {
	l := rr.Hdr.len()
	l += 2 // Preference
	l += len(rr.Host) + 1
	return l
}
func (rr *RT) unpack() int {
	l := rr.Hdr.len()
	l += 2 // Preference
	l += len(rr.Host) + 1
	return l
}
func (rr *SOA) pack() int {
	l := rr.Hdr.len()
	l += len(rr.Ns) + 1
	l += len(rr.Mbox) + 1
	l += 4 // Serial
	l += 4 // Refresh
	l += 4 // Retry
	l += 4 // Expire
	l += 4 // Minttl
	return l
}
func (rr *SOA) unpack() int {
	l := rr.Hdr.len()
	l += len(rr.Ns) + 1
	l += len(rr.Mbox) + 1
	l += 4 // Serial
	l += 4 // Refresh
	l += 4 // Retry
	l += 4 // Expire
	l += 4 // Minttl
	return l
}
func (rr *SPF) pack() int {
	l := rr.Hdr.len()
	for _, x := range rr.Txt {
		l += len(x) + 1
	}
	return l
}
func (rr *SPF) unpack() int {
	l := rr.Hdr.len()
	for _, x := range rr.Txt {
		l += len(x) + 1
	}
	return l
}
func (rr *SRV) pack() int {
	l := rr.Hdr.len()
	l += 2 // Priority
	l += 2 // Weight
	l += 2 // Port
	l += len(rr.Target) + 1
	return l
}
func (rr *SRV) unpack() int {
	l := rr.Hdr.len()
	l += 2 // Priority
	l += 2 // Weight
	l += 2 // Port
	l += len(rr.Target) + 1
	return l
}
func (rr *SSHFP) pack() int {
	l := rr.Hdr.len()
	l += 1 // Algorithm
	l += 1 // Type
	l += len(rr.FingerPrint)/2 + 1
	return l
}
func (rr *SSHFP) unpack() int {
	l := rr.Hdr.len()
	l += 1 // Algorithm
	l += 1 // Type
	l += len(rr.FingerPrint)/2 + 1
	return l
}
func (rr *TA) pack() int {
	l := rr.Hdr.len()
	l += 2 // KeyTag
	l += 1 // Algorithm
	l += 1 // DigestType
	l += len(rr.Digest)/2 + 1
	return l
}
func (rr *TA) unpack() int {
	l := rr.Hdr.len()
	l += 2 // KeyTag
	l += 1 // Algorithm
	l += 1 // DigestType
	l += len(rr.Digest)/2 + 1
	return l
}
func (rr *TALINK) pack() int {
	l := rr.Hdr.len()
	l += len(rr.PreviousName) + 1
	l += len(rr.NextName) + 1
	return l
}
func (rr *TALINK) unpack() int {
	l := rr.Hdr.len()
	l += len(rr.PreviousName) + 1
	l += len(rr.NextName) + 1
	return l
}
func (rr *TKEY) pack() int {
	l := rr.Hdr.len()
	l += len(rr.Algorithm) + 1
	l += 4 // Inception
	l += 4 // Expiration
	l += 2 // Mode
	l += 2 // Error
	l += 2 // KeySize
	l += len(rr.Key) + 1
	l += 2 // OtherLen
	l += len(rr.OtherData) + 1
	return l
}
func (rr *TKEY) unpack() int {
	l := rr.Hdr.len()
	l += len(rr.Algorithm) + 1
	l += 4 // Inception
	l += 4 // Expiration
	l += 2 // Mode
	l += 2 // Error
	l += 2 // KeySize
	l += len(rr.Key) + 1
	l += 2 // OtherLen
	l += len(rr.OtherData) + 1
	return l
}
func (rr *TLSA) pack() int {
	l := rr.Hdr.len()
	l += 1 // Usage
	l += 1 // Selector
	l += 1 // MatchingType
	l += len(rr.Certificate)/2 + 1
	return l
}
func (rr *TLSA) unpack() int {
	l := rr.Hdr.len()
	l += 1 // Usage
	l += 1 // Selector
	l += 1 // MatchingType
	l += len(rr.Certificate)/2 + 1
	return l
}
func (rr *TSIG) pack() int {
	l := rr.Hdr.len()
	l += len(rr.Algorithm) + 1
	l += 6 // TimeSigned
	l += 2 // Fudge
	l += 2 // MACSize
	l += len(rr.MAC)/2 + 1
	l += 2 // OrigId
	l += 2 // Error
	l += 2 // OtherLen
	l += len(rr.OtherData)/2 + 1
	return l
}
func (rr *TSIG) unpack() int {
	l := rr.Hdr.len()
	l += len(rr.Algorithm) + 1
	l += 6 // TimeSigned
	l += 2 // Fudge
	l += 2 // MACSize
	l += len(rr.MAC)/2 + 1
	l += 2 // OrigId
	l += 2 // Error
	l += 2 // OtherLen
	l += len(rr.OtherData)/2 + 1
	return l
}
func (rr *TXT) pack() int {
	l := rr.Hdr.len()
	for _, x := range rr.Txt {
		l += len(x) + 1
	}
	return l
}
func (rr *TXT) unpack() int {
	l := rr.Hdr.len()
	for _, x := range rr.Txt {
		l += len(x) + 1
	}
	return l
}
func (rr *UID) pack() int {
	l := rr.Hdr.len()
	l += 4 // Uid
	return l
}
func (rr *UID) unpack() int {
	l := rr.Hdr.len()
	l += 4 // Uid
	return l
}
func (rr *UINFO) pack() int {
	l := rr.Hdr.len()
	l += len(rr.Uinfo) + 1
	return l
}
func (rr *UINFO) unpack() int {
	l := rr.Hdr.len()
	l += len(rr.Uinfo) + 1
	return l
}
func (rr *URI) pack() int {
	l := rr.Hdr.len()
	l += 2 // Priority
	l += 2 // Weight
	l += len(rr.Target)
	return l
}
func (rr *URI) unpack() int {
	l := rr.Hdr.len()
	l += 2 // Priority
	l += 2 // Weight
	l += len(rr.Target)
	return l
}
func (rr *X25) pack() int {
	l := rr.Hdr.len()
	l += len(rr.PSDNAddress) + 1
	return l
}
func (rr *X25) unpack() int {
	l := rr.Hdr.len()
	l += len(rr.PSDNAddress) + 1
	return l
}
