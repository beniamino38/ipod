package extremote

import (
	"github.com/oandrew/ipod"
)

type DeviceExtRemote interface {
	PlaybackStatus() (trackLength, trackPos uint32, state PlayerState)
}

func ackSuccess(req ipod.Packet) ACK {
	return ACK{Status: ACKStatusSuccess, CmdID: req.ID.CmdID()}
}

// func ackPending(req ipod.Packet, maxWait uint32) ACKPending {
// 	return ACKPending{Status: ACKStatusPending, CmdID: uint8(req.ID.CmdID()), MaxWait: maxWait}
// }

func HandleExtRemote(req ipod.Packet, tr ipod.PacketWriter, dev DeviceExtRemote) error {
	//log.Printf("Req: %#v", req)
	switch msg := req.Payload.(type) {

	case GetCurrentPlayingTrackChapterInfo:
	// ReturnCurrentPlayingTrackChapterInfo:
	case SetCurrentPlayingTrackChapter:
	case GetCurrentPlayingTrackChapterPlayStatus:
	// ReturnCurrentPlayingTrackChapterPlayStatus:
	case GetCurrentPlayingTrackChapterName:
	// ReturnCurrentPlayingTrackChapterName:
	case GetAudiobookSpeed:
	// ReturnAudiobookSpeed:
	case SetAudiobookSpeed:
	case GetIndexedPlayingTrackInfo:
	// ReturnIndexedPlayingTrackInfo:
	case GetArtworkFormats:
	// RetArtworkFormats:
	case GetTrackArtworkData:
	// RetTrackArtworkData:
	case ResetDBSelection:
		ipod.Respond(req, tr, ackSuccess(req))
	case SelectDBRecord:
		ipod.Respond(req, tr, ackSuccess(req))
	case GetNumberCategorizedDBRecords:
		ipod.Respond(req, tr, ReturnNumberCategorizedDBRecords{
			RecordCount: 0,
		})
	case RetrieveCategorizedDatabaseRecords:
		ipod.Respond(req, tr, ReturnCategorizedDatabaseRecord{})
	case GetPlayStatus:
		ipod.Respond(req, tr, ReturnPlayStatus{
			TrackLength:   300 * 1000,
			TrackPosition: 20 * 1000,
			State:         PlayerStatePaused,
		})
	case GetCurrentPlayingTrackIndex:
		ipod.Respond(req, tr, ReturnCurrentPlayingTrackIndex{
			TrackIndex: 0,
		})
	case GetIndexedPlayingTrackTitle:
		ipod.Respond(req, tr, ReturnIndexedPlayingTrackTitle{
			Title: ipod.StringToBytes("title"),
		})
	case GetIndexedPlayingTrackArtistName:
		ipod.Respond(req, tr, ReturnIndexedPlayingTrackArtistName{
			ArtistName: ipod.StringToBytes("artist"),
		})
	case GetIndexedPlayingTrackAlbumName:
		ipod.Respond(req, tr, ReturnIndexedPlayingTrackAlbumName{
			AlbumName: ipod.StringToBytes("album"),
		})
	case SetPlayStatusChangeNotification:
		ipod.Respond(req, tr, ackSuccess(req))
	//case PlayStatusChangeNotification:
	case PlayCurrentSelection:
	case PlayControl:
		ipod.Respond(req, tr, ackSuccess(req))
	case GetTrackArtworkTimes:
	// RetTrackArtworkTimes:

	case GetShuffle:
		ipod.Respond(req, tr, ReturnShuffle{Mode: ShuffleOff})
	case SetShuffle:
		ipod.Respond(req, tr, ackSuccess(req))

	case GetRepeat:
		ipod.Respond(req, tr, ReturnRepeat{Mode: RepeatOff})
	case SetRepeat:
		ipod.Respond(req, tr, ackSuccess(req))

	case SetDisplayImage:
		ipod.Respond(req, tr, ackSuccess(req))
	case GetMonoDisplayImageLimits:
	// ReturnMonoDisplayImageLimits:
	case GetNumPlayingTracks:
		ipod.Respond(req, tr, ReturnNumPlayingTracks{
			NumTracks: 0,
		})
	case SetCurrentPlayingTrack:
	case SelectSortDBRecord:
	case GetColorDisplayImageLimits:
		ipod.Respond(req, tr, ReturnColorDisplayImageLimits{
			MaxWidth:    320,
			MaxHeight:   240,
			PixelFormat: 0x01,
		})
	case ResetDBSelectionHierarchy:
	case GetDBiTunesInfo:
	// RetDBiTunesInfo:
	case GetUIDTrackInfo:
	// RetUIDTrackInfo:
	case GetDBTrackInfo:
	// RetDBTrackInfo:
	case GetPBTrackInfo:
	// RetPBTrackInfo:

	default:
		_ = msg
	}
	return nil
}
