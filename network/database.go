package network

import (
	"strconv"
	"strings"
	"sync"

	"github.com/safing/portbase/database"
	"github.com/safing/portbase/database/iterator"
	"github.com/safing/portbase/database/query"
	"github.com/safing/portbase/database/record"
	"github.com/safing/portbase/database/storage"
	"github.com/safing/portmaster/process"
)

var (
	dnsConns     = make(map[string]*Connection) // key: <PID>/Scope
	dnsConnsLock sync.RWMutex
	conns        = make(map[string]*Connection) // key: Connection ID
	connsLock    sync.RWMutex

	dbController *database.Controller
)

// StorageInterface provices a storage.Interface to the configuration manager.
type StorageInterface struct {
	storage.InjectBase
}

// Get returns a database record.
func (s *StorageInterface) Get(key string) (record.Record, error) {

	splitted := strings.Split(key, "/")
	switch splitted[0] { //nolint:gocritic // TODO: implement full key space
	case "tree":
		switch len(splitted) {
		case 2:
			pid, err := strconv.Atoi(splitted[1])
			if err == nil {
				proc, ok := process.GetProcessFromStorage(pid)
				if ok {
					return proc, nil
				}
			}
		case 3:
			dnsConnsLock.RLock()
			defer dnsConnsLock.RUnlock()
			conn, ok := dnsConns[splitted[1]+"/"+splitted[2]]
			if ok {
				return conn, nil
			}
		case 4:
			connsLock.RLock()
			defer connsLock.RUnlock()
			conn, ok := conns[splitted[3]]
			if ok {
				return conn, nil
			}
		}
	}

	return nil, storage.ErrNotFound
}

// Query returns a an iterator for the supplied query.
func (s *StorageInterface) Query(q *query.Query, local, internal bool) (*iterator.Iterator, error) {
	it := iterator.New()
	go s.processQuery(q, it)
	// TODO: check local and internal

	return it, nil
}

func (s *StorageInterface) processQuery(q *query.Query, it *iterator.Iterator) {
	slashes := strings.Count(q.DatabaseKeyPrefix(), "/")

	if slashes <= 1 {
		// processes
		for _, proc := range process.All() {
			if q.Matches(proc) {
				it.Next <- proc
			}
		}
	}

	if slashes <= 2 {
		// dns scopes only
		dnsConnsLock.RLock()
		for _, dnsConn := range dnsConns {
			if q.Matches(dnsConn) {
				it.Next <- dnsConn
			}
		}
		dnsConnsLock.RUnlock()
	}

	if slashes <= 3 {
		// connections
		connsLock.RLock()
		for _, conn := range conns {
			if q.Matches(conn) {
				it.Next <- conn
			}
		}
		connsLock.RUnlock()
	}

	it.Finish(nil)
}

func registerAsDatabase() error {
	_, err := database.Register(&database.Database{
		Name:        "network",
		Description: "Network and Firewall Data",
		StorageType: "injected",
		PrimaryAPI:  "",
	})
	if err != nil {
		return err
	}

	controller, err := database.InjectDatabase("network", &StorageInterface{})
	if err != nil {
		return err
	}

	dbController = controller
	process.SetDBController(dbController)
	return nil
}
