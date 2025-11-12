package mv

import "fmt"

var packageIDs = map[ServerRegion]string{
	ServerRegionJP: "com.sega.ColorfulStage",
	ServerRegionEN: "com.sega.ColorfulStage.en",
	ServerRegionTW: "com.sega.ColorfulStage.tw",
	ServerRegionKR: "com.sega.ColorfulStage.kr",
	ServerRegionCN: "com.tencent.pjsk",
}

// MVPath returns the file path for a 2DMV video based on the song ID, kind, and region.
func MVPath(songID int, kind MVKind, region ServerRegion) string {
	return fmt.Sprintf("/storage/emulated/0/Android/data/%s/cache/movie/live/2dmode/%s_mv/%04d/%04d.usm.bytes", packageIDs[region], kind, songID, songID)
}
