// Copyright (c) 2018-2020, The Decred developers
// Copyright (c) 2017, The dcrdata developers
// See LICENSE for details.

package explorer

import (
	"context"
	"net/http"
	"net/url"

	"github.com/go-chi/chi"
)

type contextKey int

const (
	ctxBlockIndex contextKey = iota
	ctxBlockHash
	ctxTxHash
	ctxTxInOut
	ctxTxInOutId
	ctxAddress
	ctxAgendaId
	ctxProposalRefID
)

const (
	darkModeCoookie   = "pdapiDarkBG"
	darkModeFormKey   = "darkmode"
	requestURIFormKey = "requestURI"
)

// SyncStatusAPIIntercept returns a json response back instead of a web page
// when display sync status is active for the api endpoints supported.
func (exp *explorerUI) SyncStatusAPIIntercept(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if exp.ShowingSyncStatusPage() {
			exp.HandleApiRequestsOnSync(w, r)
			return
		}
		// Otherwise, proceed to the next http handler.
		next.ServeHTTP(w, r)
	})
}

func getBlockHashCtx(r *http.Request) string {
	hash, ok := r.Context().Value(ctxBlockHash).(string)
	if !ok {
		log.Trace("Block Hash not set")
		return ""
	}
	return hash
}

func getAgendaIDCtx(r *http.Request) string {
	hash, ok := r.Context().Value(ctxAgendaId).(string)
	if !ok {
		log.Trace("Agendaid not set")
		return ""
	}
	return hash
}

func getProposalTokenCtx(r *http.Request) string {
	hash, ok := r.Context().Value(ctxProposalRefID).(string)
	if !ok {
		log.Trace("Proposal ref ID not set")
		return ""
	}
	return hash
}

// TransactionHashCtx embeds "txid" into the request context
func TransactionHashCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		txid := chi.URLParam(r, "txid")
		ctx := context.WithValue(r.Context(), ctxTxHash, txid)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// TransactionIoIndexCtx embeds "inout" and "inoutid" into the request context
func TransactionIoIndexCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		inout := chi.URLParam(r, "inout")
		inoutid := chi.URLParam(r, "inoutid")
		ctx := context.WithValue(r.Context(), ctxTxInOut, inout)
		ctx = context.WithValue(ctx, ctxTxInOutId, inoutid)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// AddressPathCtx embeds "address" into the request context
func AddressPathCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		address := chi.URLParam(r, "address")
		ctx := context.WithValue(r.Context(), ctxAddress, address)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// AgendaPathCtx embeds "agendaid" into the request context
func AgendaPathCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		agendaid := chi.URLParam(r, "agendaid")
		ctx := context.WithValue(r.Context(), ctxAgendaId, agendaid)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// ProposalPathCtx embeds "proposalrefID" into the request context
func ProposalPathCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		proposalRefID := chi.URLParam(r, "proposalrefid")
		ctx := context.WithValue(r.Context(), ctxProposalRefID, proposalRefID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// MenuFormParser parses a form submission from the navigation menu.
func MenuFormParser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.FormValue(darkModeFormKey) != "" {
			cookie, err := r.Cookie(darkModeCoookie)
			if err != nil && err != http.ErrNoCookie {
				log.Errorf("Cookie pdapiDarkBG retrieval error: %v", err)
			} else {
				if err == http.ErrNoCookie {
					cookie = &http.Cookie{
						Name:   darkModeCoookie,
						Value:  "1",
						MaxAge: 0,
					}
				} else {
					cookie.Value = "0"
					cookie.MaxAge = -1
				}

				// Redirect to the specified relative path.
				requestURI := r.FormValue(requestURIFormKey)
				if requestURI == "" {
					requestURI = "/"
				}
				URL, err := url.Parse(requestURI)
				if err != nil {
					http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
					return
				}
				http.SetCookie(w, cookie)
				http.Redirect(w, r, URL.EscapedPath(), http.StatusFound)
				return
			}
		}
		next.ServeHTTP(w, r)
	})
}
