package media

import (
	"os"
	"time"

	"github.com/livekit/protocol/auth"
)

type MediaService struct {
	apiKey    string
	apiSecret string
}

func NewMediaService() *MediaService {
	return &MediaService{
		apiKey:    os.Getenv("LIVEKIT_API_KEY"),
		apiSecret: os.Getenv("LIVEKIT_API_SECRET"),
	}
}

func (s *MediaService) GenerateJoinToken(roomName, identity, alias string) (string, error) {
	at := auth.NewAccessToken(s.apiKey, s.apiSecret)
	
	grant := &auth.VideoGrant{
		RoomJoin: true,
		Room:     roomName,
	}

	at.SetVideoGrant(grant).
		SetIdentity(identity).
		SetName(alias).
		SetValidFor(time.Hour)

	return at.ToJWT()
}
