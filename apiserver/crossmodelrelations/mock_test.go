// Copyright 2017 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package crossmodelrelations_test

import (
	"fmt"

	"github.com/juju/errors"
	"github.com/juju/testing"
	"gopkg.in/juju/names.v2"

	common "github.com/juju/juju/apiserver/common/crossmodel"
	"github.com/juju/juju/apiserver/crossmodelrelations"
	"github.com/juju/juju/core/crossmodel"
	"github.com/juju/juju/state"
	coretesting "github.com/juju/juju/testing"
)

type mockState struct {
	testing.Stub
	crossmodelrelations.CrossModelRelationsState
	relations          map[string]*mockRelation
	remoteApplications map[string]*mockRemoteApplication
	applications       map[string]*mockApplication
	offers             []crossmodel.ApplicationOffer
	remoteEntities     map[names.Tag]string
}

func newMockState() *mockState {
	return &mockState{
		relations:          make(map[string]*mockRelation),
		remoteApplications: make(map[string]*mockRemoteApplication),
		applications:       make(map[string]*mockApplication),
		remoteEntities:     make(map[names.Tag]string),
	}
}

func (st *mockState) ListOffers(filter ...crossmodel.ApplicationOfferFilter) ([]crossmodel.ApplicationOffer, error) {
	return st.offers, nil
}

func (st *mockState) ModelUUID() string {
	return coretesting.ModelTag.Id()
}

func (st *mockState) AddRelation(eps ...state.Endpoint) (common.Relation, error) {
	rel := &mockRelation{
		key: fmt.Sprintf("%v:%v %v:%v", eps[0].ApplicationName, eps[0].Name, eps[1].ApplicationName, eps[1].Name)}
	st.relations[rel.key] = rel
	return rel, nil
}

func (st *mockState) EndpointsRelation(eps ...state.Endpoint) (common.Relation, error) {
	rel := &mockRelation{
		key: fmt.Sprintf("%v:%v %v:%v", eps[0].ApplicationName, eps[0].Name, eps[1].ApplicationName, eps[1].Name)}
	st.relations[rel.key] = rel
	return rel, nil
}

func (st *mockState) AddRemoteApplication(params state.AddRemoteApplicationParams) (common.RemoteApplication, error) {
	app := &mockRemoteApplication{consumerproxy: params.IsConsumerProxy}
	st.remoteApplications[params.Name] = app
	return app, nil
}

func (st *mockState) ImportRemoteEntity(sourceModel names.ModelTag, entity names.Tag, token string) error {
	st.MethodCall(st, "ImportRemoteEntity", sourceModel, entity, token)
	if err := st.NextErr(); err != nil {
		return err
	}
	if _, ok := st.remoteEntities[entity]; ok {
		return errors.AlreadyExistsf(entity.Id())
	}
	st.remoteEntities[entity] = token
	return nil
}

func (st *mockState) ExportLocalEntity(entity names.Tag) (string, error) {
	st.MethodCall(st, "ExportLocalEntity", entity)
	if err := st.NextErr(); err != nil {
		return "", err
	}
	if token, ok := st.remoteEntities[entity]; ok {
		return token, errors.AlreadyExistsf(entity.Id())
	}
	token := "token-" + entity.Id()
	st.remoteEntities[entity] = token
	return token, nil
}

func (st *mockState) GetRemoteEntity(sourceModel names.ModelTag, token string) (names.Tag, error) {
	st.MethodCall(st, "GetRemoteEntity", sourceModel, token)
	if err := st.NextErr(); err != nil {
		return nil, err
	}
	for e, t := range st.remoteEntities {
		if t == token {
			return e, nil
		}
	}
	return nil, errors.NotFoundf("token %v", token)
}

func (st *mockState) KeyRelation(key string) (common.Relation, error) {
	st.MethodCall(st, "KeyRelation", key)
	if err := st.NextErr(); err != nil {
		return nil, err
	}
	r, ok := st.relations[key]
	if !ok {
		return nil, errors.NotFoundf("relation %q", key)
	}
	return r, nil
}

func (st *mockState) RemoteApplication(id string) (common.RemoteApplication, error) {
	st.MethodCall(st, "RemoteApplication", id)
	if err := st.NextErr(); err != nil {
		return nil, err
	}
	a, ok := st.remoteApplications[id]
	if !ok {
		return nil, errors.NotFoundf("remote application %q", id)
	}
	return a, nil
}

func (st *mockState) Application(id string) (common.Application, error) {
	st.MethodCall(st, "Application", id)
	if err := st.NextErr(); err != nil {
		return nil, err
	}
	a, ok := st.applications[id]
	if !ok {
		return nil, errors.NotFoundf("application %q", id)
	}
	return a, nil
}

type mockRelation struct {
	common.Relation
	testing.Stub
	id    int
	key   string
	units map[string]common.RelationUnit
}

func newMockRelation(id int) *mockRelation {
	return &mockRelation{
		id:    id,
		units: make(map[string]common.RelationUnit),
	}
}

func (r *mockRelation) Tag() names.Tag {
	r.MethodCall(r, "Tag")
	return names.NewRelationTag(r.key)
}

func (r *mockRelation) Destroy() error {
	r.MethodCall(r, "Destroy")
	return r.NextErr()
}

func (r *mockRelation) RemoteUnit(unitId string) (common.RelationUnit, error) {
	r.MethodCall(r, "RemoteUnit", unitId)
	if err := r.NextErr(); err != nil {
		return nil, err
	}
	u, ok := r.units[unitId]
	if !ok {
		return nil, errors.NotFoundf("unit %q", unitId)
	}
	return u, nil
}

func (r *mockRelation) Unit(unitId string) (common.RelationUnit, error) {
	r.MethodCall(r, "Unit", unitId)
	if err := r.NextErr(); err != nil {
		return nil, err
	}
	u, ok := r.units[unitId]
	if !ok {
		return nil, errors.NotFoundf("unit %q", unitId)
	}
	return u, nil
}

func (u *mockRelationUnit) Settings() (map[string]interface{}, error) {
	u.MethodCall(u, "Settings")
	return u.settings, u.NextErr()
}

type mockRemoteApplication struct {
	common.RemoteApplication
	testing.Stub
	consumerproxy bool
}

func (r *mockRemoteApplication) IsConsumerProxy() bool {
	r.MethodCall(r, "IsConsumerProxy")
	return r.consumerproxy
}

func (r *mockRemoteApplication) Destroy() error {
	r.MethodCall(r, "Destroy")
	return r.NextErr()
}

type mockApplication struct {
	common.Application
	testing.Stub
	life state.Life
	eps  []state.Endpoint
}

func (a *mockApplication) Endpoints() ([]state.Endpoint, error) {
	a.MethodCall(a, "Endpoints")
	return a.eps, nil
}

func (a *mockApplication) Life() state.Life {
	a.MethodCall(a, "Life")
	return a.life
}

type mockRelationUnit struct {
	common.RelationUnit
	testing.Stub
	inScope  bool
	settings map[string]interface{}
}

func newMockRelationUnit() *mockRelationUnit {
	return &mockRelationUnit{
		settings: make(map[string]interface{}),
	}
}

func (u *mockRelationUnit) InScope() (bool, error) {
	u.MethodCall(u, "InScope")
	return u.inScope, u.NextErr()
}

func (u *mockRelationUnit) LeaveScope() error {
	u.MethodCall(u, "LeaveScope")
	if err := u.NextErr(); err != nil {
		return err
	}
	u.inScope = false
	return nil
}

func (u *mockRelationUnit) EnterScope(settings map[string]interface{}) error {
	u.MethodCall(u, "EnterScope", settings)
	if err := u.NextErr(); err != nil {
		return err
	}
	u.inScope = true
	u.settings = make(map[string]interface{})
	for k, v := range settings {
		u.settings[k] = v
	}
	return nil
}

func (u *mockRelationUnit) ReplaceSettings(settings map[string]interface{}) error {
	u.MethodCall(u, "ReplaceSettings", settings)
	if err := u.NextErr(); err != nil {
		return err
	}
	u.settings = make(map[string]interface{})
	for k, v := range settings {
		u.settings[k] = v
	}
	return nil
}
