package peercommands

import (
	"net/http"

	"github.com/gluster/glusterd2/client"
	"github.com/gluster/glusterd2/context"

	"github.com/gorilla/mux"
)

func deletePeer(w http.ResponseWriter, r *http.Request) {
	p := mux.Vars(r)

	id := p["peerid"]
	if id == "" {
		client.SendResponse(w, -1, http.StatusBadRequest, "peerid not present in request", http.StatusBadRequest, nil)
		return
	}

	if !context.Store.PeerExists(id) {
		client.SendResponse(w, -1, http.StatusNotFound, "", http.StatusNotFound, nil)
		return
	}

	if err := context.Store.DeletePeer(id); err != nil {
		client.SendResponse(w, -1, http.StatusInternalServerError, e.Error(), http.StatusInternalServerError, nil)
	} else {
		client.SendResponse(w, 0, 0, "", http.StatusOK, nil)
	}
}
