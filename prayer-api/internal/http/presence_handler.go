
package httpapi

import (
    "encoding/json"
    "net/http"
    "strings"
    "context"
    "github.com/clerk/clerk-sdk-go/v2"

    "prayer-api/internal/domain/identity"
    domainpresence "prayer-api/internal/domain/presence"
    presenceapp "prayer-api/internal/application/presence"
)

type HeartbeatExecutor interface {
	Execute(
		ctx context.Context,
		sessionID identity.PrayerSessionID,
		userID identity.UserID,
	) (*domainpresence.Presence, error)
}

type ParticipantsLister interface {
	Execute(
		ctx context.Context,
		sessionID identity.PrayerSessionID,
	) ([]presenceapp.Participant, error)
}


type PresenceHandler struct {
    heartbeat          HeartbeatExecutor
    listParticipants   ParticipantsLister
}
 
func NewPresenceHandler(
	heartbeat HeartbeatExecutor,
	listParticipants ParticipantsLister,
) *PresenceHandler {
    return &PresenceHandler{
        heartbeat:        heartbeat,
        listParticipants: listParticipants,
    }
}

type heartbeatResponse struct {
    SessionID string `json:"sessionId"`
    Status    string `json:"status"`
    LastSeen  string `json:"lastSeen"`
}

type participantResponse struct {
    UserID   string `json:"userId"`
    Name     string `json:"name"`
    LastSeen string `json:"lastSeen"`
}

type participantsResponse struct {
    SessionID   string                 `json:"sessionId"`
    Participants []participantResponse `json:"participants"`
}

func (h *PresenceHandler) Heartbeat(
    w http.ResponseWriter,
    r *http.Request,
) {
    claims, ok := clerk.SessionClaimsFromContext(r.Context())
    if !ok || claims == nil || strings.TrimSpace(claims.Subject) == "" {
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
    }

    sessionID := strings.TrimSpace(r.PathValue("sessionID"))
    if sessionID == "" {
        http.Error(w, "Session ID is required", http.StatusBadRequest)
        return
    }

    userID := identity.UserID(strings.TrimSpace(claims.Subject))

    result, err := h.heartbeat.Execute(
        r.Context(),
        identity.PrayerSessionID(sessionID),
        userID,
    )
    if err != nil {
        // Map domain/application errors to appropriate HTTP statuses.
        http.Error(w, "Unable to record heartbeat", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)

    _ = json.NewEncoder(w).Encode(heartbeatResponse{
        SessionID: string(result.SessionID),
        Status:    "active",
        LastSeen:  result.LastSeen.Format("2006-01-02T15:04:05Z07:00"),
    })
}

func (h *PresenceHandler) ListParticipants(
    w http.ResponseWriter,
    r *http.Request,
) {
    claims, ok := clerk.SessionClaimsFromContext(r.Context())
    if !ok || claims == nil || strings.TrimSpace(claims.Subject) == "" {
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
    }

    sessionID := strings.TrimSpace(r.PathValue("sessionID"))
    if sessionID == "" {
        http.Error(w, "Session ID is required", http.StatusBadRequest)
        return
    }

    participants, err := h.listParticipants.Execute(
        r.Context(),
        identity.PrayerSessionID(sessionID),
    )
    if err != nil {
        http.Error(w, "Unable to list participants", http.StatusInternalServerError)
        return
    }

    response := participantsResponse{
        SessionID:    sessionID,
        Participants: make([]participantResponse, 0, len(participants)),
    }

    for _, participant := range participants {
        response.Participants = append(
            response.Participants,
            participantResponse{
                UserID:   participant.UserID.String(),
                Name:     participant.Name,
                LastSeen: participant.LastSeen.Format("2006-01-02T15:04:05Z07:00"),
            },
        )
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)

    _ = json.NewEncoder(w).Encode(response)
}