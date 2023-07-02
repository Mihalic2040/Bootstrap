package app

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Mihalic2040/Bootstrap/src/types"
)

func ShowPeers(w http.ResponseWriter, r *http.Request) {
	peers := App.Host.Peerstore().Peers()

	var peersInfo []types.PeerInfo
	for _, peer := range peers {
		addr := App.Host.Peerstore().PeerInfo(peer).Addrs
		peerInfo := types.PeerInfo{
			Peer: peer.Pretty(),
			Addr: addr,
		}
		peersInfo = append(peersInfo, peerInfo)
	}

	// Convert peersInfo to JSON
	peersJSON, err := json.Marshal(peersInfo)
	if err != nil {
		http.Error(w, "Failed to marshal peers to JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(peersJSON)
}

func ShowAddrs(w http.ResponseWriter, r *http.Request) {
	addrs := App.Host.Addrs()
	log.Println(addrs)

	// Convert multiaddr.Multiaddr to string
	addrsStr := make([]string, len(addrs))
	for i, addr := range addrs {
		addrsStr[i] = addr.String()
	}

	// Include the PeerID and converted addrs in the data structure
	data := struct {
		PeerID string   `json:"peerID"`
		Addrs  []string `json:"addrs"`
	}{
		PeerID: App.Host.ID().Pretty(),
		Addrs:  addrsStr,
	}

	w.Header().Set("Content-Type", "application/json")

	// Serialize the data structure to JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Failed to serialize JSON", http.StatusInternalServerError)
		return
	}

	// Write the JSON data as the response body
	w.Write(jsonData)
}
