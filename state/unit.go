// launchpad.net/juju/state
//
// Copyright (c) 2011-2012 Canonical Ltd.

package state

import (
	"errors"
	"fmt"
	"launchpad.net/goyaml"
	"launchpad.net/gozk/zookeeper"
	"launchpad.net/juju/go/charm"
	"launchpad.net/juju/go/state/presence"
	"launchpad.net/juju/go/state/watcher"
	"launchpad.net/tomb"
	"strconv"
	"strings"
	"time"
)

// ResolvedMode describes the way state transition errors 
// are resolved. 
type ResolvedMode int

const (
	ResolvedNone       ResolvedMode = 0
	ResolvedRetryHooks ResolvedMode = 1000
	ResolvedNoHooks    ResolvedMode = 1001
)

// agentPingerPeriod defines the period of pinging the
// ZooKeeper to signal that a unit agent is alive. It's
// also used by machine.
const (
	agentPingerPeriod = 1 * time.Second
)

// Port identifies a network port number for a particular protocol.
type Port struct {
	Protocol string `yaml:"proto"`
	Number   int    `yaml:"port"`
}

// Unit represents the state of a service unit.
type Unit struct {
	st          *State
	key         string
	serviceKey  string
	serviceName string
	sequenceNo  int
}

// ServiceName returns the service name.
func (u *Unit) ServiceName() string {
	return u.serviceName
}

// Name returns the unit name.
func (u *Unit) Name() string {
	return fmt.Sprintf("%s/%d", u.serviceName, u.sequenceNo)
}

// PublicAddress returns the public address of the unit.
func (u *Unit) PublicAddress() (string, error) {
	cn, err := readConfigNode(u.st.zk, u.zkPath())
	if err != nil {
		return "", err
	}
	if address, ok := cn.Get("public-address"); ok {
		return address.(string), nil
	}
	return "", errors.New("unit has no public address")
}

// SetPublicAddress sets the public address of the unit.
func (u *Unit) SetPublicAddress(address string) error {
	cn, err := readConfigNode(u.st.zk, u.zkPath())
	if err != nil {
		return err
	}
	cn.Set("public-address", address)
	_, err = cn.Write()
	if err != nil {
		return err
	}
	return nil
}

// PrivateAddress returns the private address of the unit.
func (u *Unit) PrivateAddress() (string, error) {
	cn, err := readConfigNode(u.st.zk, u.zkPath())
	if err != nil {
		return "", err
	}
	if address, ok := cn.Get("private-address"); ok {
		return address.(string), nil
	}
	return "", errors.New("unit has no private address")
}

// SetPrivateAddress sets the private address of the unit.
func (u *Unit) SetPrivateAddress(address string) error {
	cn, err := readConfigNode(u.st.zk, u.zkPath())
	if err != nil {
		return err
	}
	cn.Set("private-address", address)
	_, err = cn.Write()
	if err != nil {
		return err
	}
	return nil
}

// CharmURL returns the charm URL this unit is supposed
// to use.
func (u *Unit) CharmURL() (url *charm.URL, err error) {
	cn, err := readConfigNode(u.st.zk, u.zkPath())
	if err != nil {
		return nil, err
	}
	if id, ok := cn.Get("charm"); ok {
		url, err = charm.ParseURL(id.(string))
		if err != nil {
			return nil, err
		}
		return url, nil
	}
	return nil, errors.New("unit has no charm URL")
}

// SetCharmURL changes the charm URL for the unit.
func (u *Unit) SetCharmURL(url *charm.URL) error {
	cn, err := readConfigNode(u.st.zk, u.zkPath())
	if err != nil {
		return err
	}
	cn.Set("charm", url.String())
	_, err = cn.Write()
	if err != nil {
		return err
	}
	return nil
}

// AssignedMachineId returns the id of the assigned machine.
func (u *Unit) AssignedMachineId() (int, error) {
	topology, err := readTopology(u.st.zk)
	if err != nil {
		return 0, err
	}
	if !topology.HasService(u.serviceKey) || !topology.HasUnit(u.serviceKey, u.key) {
		return 0, stateChanged
	}
	machineKey, err := topology.UnitMachineKey(u.serviceKey, u.key)
	if err != nil {
		return 0, err
	}
	return machineId(machineKey), nil
}

// AssignToMachine assigns this unit to a given machine.
func (u *Unit) AssignToMachine(machine *Machine) error {
	assignUnit := func(t *topology) error {
		if !t.HasService(u.serviceKey) || !t.HasUnit(u.serviceKey, u.key) {
			return stateChanged
		}
		machineKey, err := t.UnitMachineKey(u.serviceKey, u.key)
		if err == unitNotAssigned {
			return t.AssignUnitToMachine(u.serviceKey, u.key, machine.key)
		} else if err != nil {
			return err
		} else if machineKey == machine.key {
			// Everything is fine, it's already assigned.
			return nil
		}
		return fmt.Errorf("unit %q already assigned to machine %d", u.Name(), machineId(machineKey))
	}
	return retryTopologyChange(u.st.zk, assignUnit)
}

// AssignToUnusedMachine assigns u to a machine without other units.
// If there are no unused machines besides machine 0, an error is returned.
func (u *Unit) AssignToUnusedMachine() (*Machine, error) {
	machineKey := ""
	assignUnusedUnit := func(t *topology) error {
		if !t.HasService(u.serviceKey) || !t.HasUnit(u.serviceKey, u.key) {
			return stateChanged
		}
		for _, machineKey = range t.MachineKeys() {
			if machineId(machineKey) != 0 {
				hasUnits, err := t.MachineHasUnits(machineKey)
				if err != nil {
					return err
				}
				if !hasUnits {
					break
				}
			}
			// Reset machine key.
			machineKey = ""
		}
		if machineKey == "" {
			return errors.New("no unused machine found")
		}
		if err := t.AssignUnitToMachine(u.serviceKey, u.key, machineKey); err != nil {
			return err
		}
		return nil
	}
	if err := retryTopologyChange(u.st.zk, assignUnusedUnit); err != nil {
		return nil, err
	}
	return &Machine{u.st, machineKey}, nil
}

// UnassignFromMachine removes the assignment between this unit and
// the machine it's assigned to.
func (u *Unit) UnassignFromMachine() error {
	unassignUnit := func(t *topology) error {
		if !t.HasService(u.serviceKey) || !t.HasUnit(u.serviceKey, u.key) {
			return stateChanged
		}
		// If for whatever reason it's already not assigned to a
		// machine, ignore it and move forward so that we don't
		// have to deal with conflicts.
		key, err := t.UnitMachineKey(u.serviceKey, u.key)
		if err == nil && key != "" {
			t.UnassignUnitFromMachine(u.serviceKey, u.key)
		}
		return nil
	}
	return retryTopologyChange(u.st.zk, unassignUnit)
}

// NeedsUpgrade returns whether the unit needs an upgrade.
func (u *Unit) NeedsUpgrade() (bool, error) {
	stat, err := u.st.zk.Exists(u.zkNeedsUpgradePath())
	if err != nil {
		return false, err
	}
	return stat != nil, nil
}

// SetNeedsUpgrade informs the unit that it should perform an upgrade.
func (u *Unit) SetNeedsUpgrade() error {
	_, err := u.st.zk.Create(u.zkNeedsUpgradePath(), "", 0, zkPermAll)
	if zookeeper.IsError(err, zookeeper.ZNODEEXISTS) {
		// Node already exists, so same state.
		return nil
	}
	return err
}

// ClearNeedsUpgrade resets the upgrade notification. It is typically
// done by the unit agent before beginning the upgrade.
func (u *Unit) ClearNeedsUpgrade() error {
	err := u.st.zk.Delete(u.zkNeedsUpgradePath(), -1)
	if zookeeper.IsError(err, zookeeper.ZNONODE) {
		// Node doesn't exist, so same state.
		return nil
	}
	return err
}

// WatchNeedsUpgrade creates a watcher for the upgrade notification 
// of the unit.
func (u *Unit) WatchNeedsUpgrade() *NeedsUpgradeWatcher {
	return newNeedsUpgradeWatcher(u.st, u.zkNeedsUpgradePath())
}

// Resolved returns the resolved mode for the unit.
func (u *Unit) Resolved() (ResolvedMode, error) {
	yaml, _, err := u.st.zk.Get(u.zkResolvedPath())
	if zookeeper.IsError(err, zookeeper.ZNONODE) {
		// Default value.
		return ResolvedNone, nil
	}
	if err != nil {
		return ResolvedNone, err
	}
	setting := &struct{ Retry ResolvedMode }{}
	if err = goyaml.Unmarshal([]byte(yaml), setting); err != nil {
		return ResolvedNone, err
	}
	mode := setting.Retry
	if err := validResolvedMode(mode, false); err != nil {
		return ResolvedNone, err
	}
	return mode, nil
}

// SetResolved marks the unit as having had any previous state
// transition problems resolved, and informs the unit that it may
// attempt to reestablish normal workflow.
// The resolved mode parameter informs whether to attempt to 
// reexecute previous failed hooks or to continue as if they had 
// succeeded before.
func (u *Unit) SetResolved(mode ResolvedMode) error {
	if err := validResolvedMode(mode, false); err != nil {
		return err
	}
	setting := &struct{ Retry ResolvedMode }{mode}
	yaml, err := goyaml.Marshal(setting)
	if err != nil {
		return err
	}
	_, err = u.st.zk.Create(u.zkResolvedPath(), string(yaml), 0, zkPermAll)
	if zookeeper.IsError(err, zookeeper.ZNODEEXISTS) {
		return fmt.Errorf("unit %q resolved flag already set", u.Name())
	}
	return err
}

// ClearResolved removes any resolved setting on the unit.
func (u *Unit) ClearResolved() error {
	err := u.st.zk.Delete(u.zkResolvedPath(), -1)
	if zookeeper.IsError(err, zookeeper.ZNONODE) {
		// Node doesn't exist, so same state.
		return nil
	}
	return err
}

// WatchResolved returns a watcher that fires when the unit 
// is marked as having had its problems resolved. See 
// SetResolved for details.
func (u *Unit) WatchResolved() *ResolvedWatcher {
	return newResolvedWatcher(u.st, u.zkResolvedPath())
}

// OpenPort sets the policy of the port with protocol and number to be opened.
func (u *Unit) OpenPort(protocol string, number int) error {
	openPort := func(oldYaml string, stat *zookeeper.Stat) (string, error) {
		ports := &struct{ Open []Port }{}
		if oldYaml != "" {
			if err := goyaml.Unmarshal([]byte(oldYaml), ports); err != nil {
				return "", err
			}
		}
		portToOpen := Port{protocol, number}
		found := false
		for _, openPort := range ports.Open {
			if openPort == portToOpen {
				found = true
				break
			}
		}
		if !found {
			ports.Open = append(ports.Open, portToOpen)
		}
		newYaml, err := goyaml.Marshal(ports)
		if err != nil {
			return "", err
		}
		return string(newYaml), nil
	}
	return u.st.zk.RetryChange(u.zkPortsPath(), 0, zkPermAll, openPort)
}

// ClosePort sets the policy of the port with protocol and number to be closed.
func (u *Unit) ClosePort(protocol string, number int) error {
	closePort := func(oldYaml string, stat *zookeeper.Stat) (string, error) {
		ports := &struct{ Open []Port }{}
		if oldYaml != "" {
			if err := goyaml.Unmarshal([]byte(oldYaml), ports); err != nil {
				return "", err
			}
		}
		portToClose := Port{protocol, number}
		newOpenPorts := []Port{}
		for _, oldOpenPort := range ports.Open {
			if oldOpenPort != portToClose {
				newOpenPorts = append(newOpenPorts, oldOpenPort)
			}
		}
		ports.Open = newOpenPorts
		newYaml, err := goyaml.Marshal(ports)
		if err != nil {
			return "", err
		}
		return string(newYaml), nil
	}
	return u.st.zk.RetryChange(u.zkPortsPath(), 0, zkPermAll, closePort)
}

// OpenPorts returns a slice containing the open ports of the unit.
func (u *Unit) OpenPorts() ([]Port, error) {
	yaml, _, err := u.st.zk.Get(u.zkPortsPath())
	if zookeeper.IsError(err, zookeeper.ZNONODE) {
		// Default value.
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	ports := &struct{ Open []Port }{}
	if err = goyaml.Unmarshal([]byte(yaml), &ports); err != nil {
		return nil, err
	}
	return ports.Open, nil
}

// AgentAlive returns whether the respective remote agent is alive.
func (u *Unit) AgentAlive() (bool, error) {
	return presence.Alive(u.st.zk, u.zkAgentPath())
}

// WaitAgentAlive blocks until the respective agent is alive.
func (u *Unit) WaitAgentAlive(timeout time.Duration) error {
	err := presence.WaitAlive(u.st.zk, u.zkAgentPath(), timeout)
	if err != nil {
		return fmt.Errorf("state: waiting for agent of unit %q: %v", u.Name(), err)
	}
	return nil
}

// SetAgentAlive signals that the agent for unit u is alive
// by starting a pinger on its presence node. It returns the
// started pinger.
func (u *Unit) SetAgentAlive() (*presence.Pinger, error) {
	return presence.StartPinger(u.st.zk, u.zkAgentPath(), agentPingerPeriod)
}

// zkKey returns the ZooKeeper key of the unit.
func (u *Unit) zkKey() string {
	return u.key
}

// zkPath returns the ZooKeeper base path for the unit.
func (u *Unit) zkPath() string {
	return fmt.Sprintf("/units/%s", u.key)
}

// zkPortsPath returns the ZooKeeper path for the open ports.
func (u *Unit) zkPortsPath() string {
	return fmt.Sprintf("/units/%s/ports", u.key)
}

// zkAgentPath returns the ZooKeeper path for the unit agent.
func (u *Unit) zkAgentPath() string {
	return fmt.Sprintf("/units/%s/agent", u.key)
}

// zkNeedsUpgradePath returns the ZooKeeper path for the upgrade flag.
func (u *Unit) zkNeedsUpgradePath() string {
	return fmt.Sprintf("/units/%s/upgrade", u.key)
}

// zkResolvedPath returns the ZooKeeper path for the mark to resolve a unit.
func (u *Unit) zkResolvedPath() string {
	return fmt.Sprintf("/units/%s/resolved", u.key)
}

// parseUnitName parses a unit name like "wordpress/0" into
// its service name and sequence number parts.
func parseUnitName(name string) (serviceName string, seqNo int, err error) {
	parts := strings.Split(name, "/")
	if len(parts) != 2 {
		return "", 0, fmt.Errorf("%q is not a valid unit name", name)
	}
	sequenceNo, err := strconv.ParseInt(parts[1], 10, 0)
	if err != nil {
		return "", 0, err
	}
	return parts[0], int(sequenceNo), nil
}

// parseResolvedMode parses a given YAML for the resolve mode
// and checks if it's valid.
func parseResolvedMode(yaml string) (ResolvedMode, error) {
	var setting struct {
		Retry ResolvedMode
	}
	if err := goyaml.Unmarshal([]byte(yaml), &setting); err != nil {
		return ResolvedNone, err
	}
	mode := setting.Retry
	if err := validResolvedMode(mode, true); err != nil {
		return ResolvedNone, err
	}
	return mode, nil
}

// validResolvedMode ensures that only valid values for the
// resolved mode are used.
func validResolvedMode(mode ResolvedMode, acceptNone bool) error {
	if acceptNone && mode == ResolvedNone {
		return nil
	}
	if mode != ResolvedRetryHooks && mode != ResolvedNoHooks {
		return fmt.Errorf("invalid error resolution mode: %d", mode)
	}
	return nil
}

// NeedsUpgradeWatcher observes changes to a unit's upgrade flag.
type NeedsUpgradeWatcher struct {
	st         *State
	path       string
	tomb       tomb.Tomb
	watcher    *watcher.ContentWatcher
	changeChan chan bool
}

// newNeedsUpgradeWatcher creates and starts a new resolved flag node 
// watcher for the given path.
func newNeedsUpgradeWatcher(st *State, path string) *NeedsUpgradeWatcher {
	w := &NeedsUpgradeWatcher{
		st:         st,
		path:       path,
		changeChan: make(chan bool),
		watcher:    watcher.NewContentWatcher(st.zk, path),
	}
	go w.loop()
	return w
}

// Changes returns a channel that will receive the new
// resolved mode when a change is detected. Note that multiple
// changes may be observed as a single event in the channel.
func (w *NeedsUpgradeWatcher) Changes() <-chan bool {
	return w.changeChan
}

// Stop stops the watch and returns any error encountered
// while watching. This method should always be called
// before discarding the watcher.
func (w *NeedsUpgradeWatcher) Stop() error {
	w.tomb.Kill(nil)
	if err := w.watcher.Stop(); err != nil {
		w.tomb.Wait()
		return err
	}
	return w.tomb.Wait()
}

// loop is the backend for watching the resolved flag node.
func (w *NeedsUpgradeWatcher) loop() {
	defer w.tomb.Done()
	defer close(w.changeChan)

	for {
		select {
		case <-w.tomb.Dying():
			return
		case change := <-w.watcher.Changes():
			select {
			case <-w.watcher.Dying():
				return
			case <-w.tomb.Dying():
				return
			case w.changeChan <- change.Exists:
			}
		}
	}
}

// ResolvedWatcher observes changes to a resolved flag node.
type ResolvedWatcher struct {
	st         *State
	path       string
	tomb       tomb.Tomb
	watcher    *watcher.ContentWatcher
	changeChan chan ResolvedMode
}

// newResolvedWatcher creates and starts a new resolved flag node 
// watcher for the given path.
func newResolvedWatcher(st *State, path string) *ResolvedWatcher {
	w := &ResolvedWatcher{
		st:         st,
		path:       path,
		changeChan: make(chan ResolvedMode),
		watcher:    watcher.NewContentWatcher(st.zk, path),
	}
	go w.loop()
	return w
}

// Changes returns a channel that will receive the new
// resolved mode when a change is detected. Note that multiple
// changes may be observed as a single event in the channel.
func (w *ResolvedWatcher) Changes() <-chan ResolvedMode {
	return w.changeChan
}

// Stop stops the watch and returns any error encountered
// while watching. This method should always be called
// before discarding the watcher.
func (w *ResolvedWatcher) Stop() error {
	w.tomb.Kill(nil)
	if err := w.watcher.Stop(); err != nil {
		w.tomb.Wait()
		return err
	}
	return w.tomb.Wait()
}

// loop is the backend for watching the resolved flag node.
func (w *ResolvedWatcher) loop() {
	defer w.tomb.Done()
	defer close(w.changeChan)

	for {
		select {
		case <-w.tomb.Dying():
			return
		case change := <-w.watcher.Changes():
			mode, err := parseResolvedMode(change.Content)
			if err != nil {
				w.tomb.Kill(err)
				return
			}
			select {
			case <-w.watcher.Dying():
				return
			case <-w.tomb.Dying():
				return
			case w.changeChan <- mode:
			}
		}
	}
}
