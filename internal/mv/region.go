package mv

import "fmt"

type ServerRegion uint8

const (
	ServerRegionJP ServerRegion = iota
	ServerRegionEN
	ServerRegionTW
	ServerRegionKR
	ServerRegionCN
)

func (r ServerRegion) String() string {
	switch r {
	case ServerRegionJP:
		return "jp"
	case ServerRegionEN:
		return "en"
	case ServerRegionTW:
		return "tw"
	case ServerRegionKR:
		return "kr"
	case ServerRegionCN:
		return "cn"
	default:
		panic(fmt.Sprintf("invalid ServerRegion: %d", r))
	}
}

func (r *ServerRegion) Set(s string) error {
	switch s {
	case "jp":
		*r = ServerRegionJP
	case "en":
		*r = ServerRegionEN
	case "tw":
		*r = ServerRegionTW
	case "kr":
		*r = ServerRegionKR
	case "cn":
		*r = ServerRegionCN
	default:
		return fmt.Errorf("invalid ServerRegion: %s", s)
	}
	return nil
}

func (r ServerRegion) Type() string {
	return "ServerRegion"
}
