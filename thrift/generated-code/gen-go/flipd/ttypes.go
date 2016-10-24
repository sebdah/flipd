// Autogenerated by Thrift Compiler (0.9.3)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package flipd

import (
	"bytes"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = bytes.Equal

var GoUnusedProtection__ int

//****************************************************
//*                       Enums
///****************************************************
type FeatureStatus int64

const (
	FeatureStatus_ENABLED  FeatureStatus = 1
	FeatureStatus_DISABLED FeatureStatus = 2
)

func (p FeatureStatus) String() string {
	switch p {
	case FeatureStatus_ENABLED:
		return "ENABLED"
	case FeatureStatus_DISABLED:
		return "DISABLED"
	}
	return "<UNSET>"
}

func FeatureStatusFromString(s string) (FeatureStatus, error) {
	switch s {
	case "ENABLED":
		return FeatureStatus_ENABLED, nil
	case "DISABLED":
		return FeatureStatus_DISABLED, nil
	}
	return FeatureStatus(0), fmt.Errorf("not a valid FeatureStatus string")
}

func FeatureStatusPtr(v FeatureStatus) *FeatureStatus { return &v }

func (p FeatureStatus) MarshalText() ([]byte, error) {
	return []byte(p.String()), nil
}

func (p *FeatureStatus) UnmarshalText(text []byte) error {
	q, err := FeatureStatusFromString(string(text))
	if err != nil {
		return err
	}
	*p = q
	return nil
}

//****************************************************
//*                       Types
///****************************************************
type Key string

func KeyPtr(v Key) *Key { return &v }

// ****************************************************
// *                      Objects
// /****************************************************
//
// Attributes:
//  - Key
//  - Status
type Feature struct {
	Key    *Key          `thrift:"key,1" json:"key,omitempty"`
	Status FeatureStatus `thrift:"status,2" json:"status,omitempty"`
}

func NewFeature() *Feature {
	return &Feature{
		Status: 2,
	}
}

var Feature_Key_DEFAULT Key

func (p *Feature) GetKey() Key {
	if !p.IsSetKey() {
		return Feature_Key_DEFAULT
	}
	return *p.Key
}

var Feature_Status_DEFAULT FeatureStatus = 2

func (p *Feature) GetStatus() FeatureStatus {
	return p.Status
}
func (p *Feature) IsSetKey() bool {
	return p.Key != nil
}

func (p *Feature) IsSetStatus() bool {
	return p.Status != Feature_Status_DEFAULT
}

func (p *Feature) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
		case 2:
			if err := p.readField2(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *Feature) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		temp := Key(v)
		p.Key = &temp
	}
	return nil
}

func (p *Feature) readField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		temp := FeatureStatus(v)
		p.Status = temp
	}
	return nil
}

func (p *Feature) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("Feature"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *Feature) writeField1(oprot thrift.TProtocol) (err error) {
	if p.IsSetKey() {
		if err := oprot.WriteFieldBegin("key", thrift.STRING, 1); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:key: ", p), err)
		}
		if err := oprot.WriteString(string(*p.Key)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.key (1) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 1:key: ", p), err)
		}
	}
	return err
}

func (p *Feature) writeField2(oprot thrift.TProtocol) (err error) {
	if p.IsSetStatus() {
		if err := oprot.WriteFieldBegin("status", thrift.I32, 2); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:status: ", p), err)
		}
		if err := oprot.WriteI32(int32(p.Status)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.status (2) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 2:status: ", p), err)
		}
	}
	return err
}

func (p *Feature) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("Feature(%+v)", *p)
}

// ****************************************************
// *               Requests and responses
// /****************************************************
//
// Attributes:
//  - Key
type DeregisterFeatureRequest struct {
	Key *Key `thrift:"key,1" json:"key,omitempty"`
}

func NewDeregisterFeatureRequest() *DeregisterFeatureRequest {
	return &DeregisterFeatureRequest{}
}

var DeregisterFeatureRequest_Key_DEFAULT Key

func (p *DeregisterFeatureRequest) GetKey() Key {
	if !p.IsSetKey() {
		return DeregisterFeatureRequest_Key_DEFAULT
	}
	return *p.Key
}
func (p *DeregisterFeatureRequest) IsSetKey() bool {
	return p.Key != nil
}

func (p *DeregisterFeatureRequest) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *DeregisterFeatureRequest) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		temp := Key(v)
		p.Key = &temp
	}
	return nil
}

func (p *DeregisterFeatureRequest) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("DeregisterFeatureRequest"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *DeregisterFeatureRequest) writeField1(oprot thrift.TProtocol) (err error) {
	if p.IsSetKey() {
		if err := oprot.WriteFieldBegin("key", thrift.STRING, 1); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:key: ", p), err)
		}
		if err := oprot.WriteString(string(*p.Key)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.key (1) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 1:key: ", p), err)
		}
	}
	return err
}

func (p *DeregisterFeatureRequest) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("DeregisterFeatureRequest(%+v)", *p)
}

// Attributes:
//  - Features
type GetFeaturesResponse struct {
	Features []*Feature `thrift:"features,1" json:"features,omitempty"`
}

func NewGetFeaturesResponse() *GetFeaturesResponse {
	return &GetFeaturesResponse{}
}

var GetFeaturesResponse_Features_DEFAULT []*Feature

func (p *GetFeaturesResponse) GetFeatures() []*Feature {
	return p.Features
}
func (p *GetFeaturesResponse) IsSetFeatures() bool {
	return p.Features != nil
}

func (p *GetFeaturesResponse) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *GetFeaturesResponse) readField1(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return thrift.PrependError("error reading list begin: ", err)
	}
	tSlice := make([]*Feature, 0, size)
	p.Features = tSlice
	for i := 0; i < size; i++ {
		_elem0 := &Feature{
			Status: 2,
		}
		if err := _elem0.Read(iprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", _elem0), err)
		}
		p.Features = append(p.Features, _elem0)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return thrift.PrependError("error reading list end: ", err)
	}
	return nil
}

func (p *GetFeaturesResponse) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("GetFeaturesResponse"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *GetFeaturesResponse) writeField1(oprot thrift.TProtocol) (err error) {
	if p.IsSetFeatures() {
		if err := oprot.WriteFieldBegin("features", thrift.LIST, 1); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:features: ", p), err)
		}
		if err := oprot.WriteListBegin(thrift.STRUCT, len(p.Features)); err != nil {
			return thrift.PrependError("error writing list begin: ", err)
		}
		for _, v := range p.Features {
			if err := v.Write(oprot); err != nil {
				return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", v), err)
			}
		}
		if err := oprot.WriteListEnd(); err != nil {
			return thrift.PrependError("error writing list end: ", err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 1:features: ", p), err)
		}
	}
	return err
}

func (p *GetFeaturesResponse) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("GetFeaturesResponse(%+v)", *p)
}

// Attributes:
//  - Keys
type GetFeaturesRequest struct {
	Keys []string `thrift:"keys,1" json:"keys,omitempty"`
}

func NewGetFeaturesRequest() *GetFeaturesRequest {
	return &GetFeaturesRequest{}
}

var GetFeaturesRequest_Keys_DEFAULT []string

func (p *GetFeaturesRequest) GetKeys() []string {
	return p.Keys
}
func (p *GetFeaturesRequest) IsSetKeys() bool {
	return p.Keys != nil
}

func (p *GetFeaturesRequest) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *GetFeaturesRequest) readField1(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return thrift.PrependError("error reading list begin: ", err)
	}
	tSlice := make([]string, 0, size)
	p.Keys = tSlice
	for i := 0; i < size; i++ {
		var _elem1 string
		if v, err := iprot.ReadString(); err != nil {
			return thrift.PrependError("error reading field 0: ", err)
		} else {
			_elem1 = v
		}
		p.Keys = append(p.Keys, _elem1)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return thrift.PrependError("error reading list end: ", err)
	}
	return nil
}

func (p *GetFeaturesRequest) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("GetFeaturesRequest"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *GetFeaturesRequest) writeField1(oprot thrift.TProtocol) (err error) {
	if p.IsSetKeys() {
		if err := oprot.WriteFieldBegin("keys", thrift.LIST, 1); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:keys: ", p), err)
		}
		if err := oprot.WriteListBegin(thrift.STRING, len(p.Keys)); err != nil {
			return thrift.PrependError("error writing list begin: ", err)
		}
		for _, v := range p.Keys {
			if err := oprot.WriteString(string(v)); err != nil {
				return thrift.PrependError(fmt.Sprintf("%T. (0) field write error: ", p), err)
			}
		}
		if err := oprot.WriteListEnd(); err != nil {
			return thrift.PrependError("error writing list end: ", err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 1:keys: ", p), err)
		}
	}
	return err
}

func (p *GetFeaturesRequest) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("GetFeaturesRequest(%+v)", *p)
}

// Attributes:
//  - Feature
type RegisterFeatureRequest struct {
	Feature *Feature `thrift:"feature,1" json:"feature,omitempty"`
}

func NewRegisterFeatureRequest() *RegisterFeatureRequest {
	return &RegisterFeatureRequest{}
}

var RegisterFeatureRequest_Feature_DEFAULT *Feature

func (p *RegisterFeatureRequest) GetFeature() *Feature {
	if !p.IsSetFeature() {
		return RegisterFeatureRequest_Feature_DEFAULT
	}
	return p.Feature
}
func (p *RegisterFeatureRequest) IsSetFeature() bool {
	return p.Feature != nil
}

func (p *RegisterFeatureRequest) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *RegisterFeatureRequest) readField1(iprot thrift.TProtocol) error {
	p.Feature = &Feature{
		Status: 2,
	}
	if err := p.Feature.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Feature), err)
	}
	return nil
}

func (p *RegisterFeatureRequest) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("RegisterFeatureRequest"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *RegisterFeatureRequest) writeField1(oprot thrift.TProtocol) (err error) {
	if p.IsSetFeature() {
		if err := oprot.WriteFieldBegin("feature", thrift.STRUCT, 1); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:feature: ", p), err)
		}
		if err := p.Feature.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Feature), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 1:feature: ", p), err)
		}
	}
	return err
}

func (p *RegisterFeatureRequest) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("RegisterFeatureRequest(%+v)", *p)
}

// ****************************************************
// *                     Exceptions
// /****************************************************
//
// Attributes:
//  - Message
type DuplicateException struct {
	Message *string `thrift:"message,1" json:"message,omitempty"`
}

func NewDuplicateException() *DuplicateException {
	return &DuplicateException{}
}

var DuplicateException_Message_DEFAULT string

func (p *DuplicateException) GetMessage() string {
	if !p.IsSetMessage() {
		return DuplicateException_Message_DEFAULT
	}
	return *p.Message
}
func (p *DuplicateException) IsSetMessage() bool {
	return p.Message != nil
}

func (p *DuplicateException) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *DuplicateException) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Message = &v
	}
	return nil
}

func (p *DuplicateException) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("DuplicateException"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *DuplicateException) writeField1(oprot thrift.TProtocol) (err error) {
	if p.IsSetMessage() {
		if err := oprot.WriteFieldBegin("message", thrift.STRING, 1); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:message: ", p), err)
		}
		if err := oprot.WriteString(string(*p.Message)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.message (1) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 1:message: ", p), err)
		}
	}
	return err
}

func (p *DuplicateException) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("DuplicateException(%+v)", *p)
}

func (p *DuplicateException) Error() string {
	return p.String()
}

// Attributes:
//  - Message
type InternalErrorException struct {
	Message *string `thrift:"message,1" json:"message,omitempty"`
}

func NewInternalErrorException() *InternalErrorException {
	return &InternalErrorException{}
}

var InternalErrorException_Message_DEFAULT string

func (p *InternalErrorException) GetMessage() string {
	if !p.IsSetMessage() {
		return InternalErrorException_Message_DEFAULT
	}
	return *p.Message
}
func (p *InternalErrorException) IsSetMessage() bool {
	return p.Message != nil
}

func (p *InternalErrorException) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *InternalErrorException) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Message = &v
	}
	return nil
}

func (p *InternalErrorException) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("InternalErrorException"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *InternalErrorException) writeField1(oprot thrift.TProtocol) (err error) {
	if p.IsSetMessage() {
		if err := oprot.WriteFieldBegin("message", thrift.STRING, 1); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:message: ", p), err)
		}
		if err := oprot.WriteString(string(*p.Message)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.message (1) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 1:message: ", p), err)
		}
	}
	return err
}

func (p *InternalErrorException) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("InternalErrorException(%+v)", *p)
}

func (p *InternalErrorException) Error() string {
	return p.String()
}

// Attributes:
//  - Message
type InvalidInputException struct {
	Message *string `thrift:"message,1" json:"message,omitempty"`
}

func NewInvalidInputException() *InvalidInputException {
	return &InvalidInputException{}
}

var InvalidInputException_Message_DEFAULT string

func (p *InvalidInputException) GetMessage() string {
	if !p.IsSetMessage() {
		return InvalidInputException_Message_DEFAULT
	}
	return *p.Message
}
func (p *InvalidInputException) IsSetMessage() bool {
	return p.Message != nil
}

func (p *InvalidInputException) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *InvalidInputException) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Message = &v
	}
	return nil
}

func (p *InvalidInputException) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("InvalidInputException"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *InvalidInputException) writeField1(oprot thrift.TProtocol) (err error) {
	if p.IsSetMessage() {
		if err := oprot.WriteFieldBegin("message", thrift.STRING, 1); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:message: ", p), err)
		}
		if err := oprot.WriteString(string(*p.Message)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.message (1) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 1:message: ", p), err)
		}
	}
	return err
}

func (p *InvalidInputException) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("InvalidInputException(%+v)", *p)
}

func (p *InvalidInputException) Error() string {
	return p.String()
}

// Attributes:
//  - Message
type NotFoundException struct {
	Message *string `thrift:"message,1" json:"message,omitempty"`
}

func NewNotFoundException() *NotFoundException {
	return &NotFoundException{}
}

var NotFoundException_Message_DEFAULT string

func (p *NotFoundException) GetMessage() string {
	if !p.IsSetMessage() {
		return NotFoundException_Message_DEFAULT
	}
	return *p.Message
}
func (p *NotFoundException) IsSetMessage() bool {
	return p.Message != nil
}

func (p *NotFoundException) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *NotFoundException) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Message = &v
	}
	return nil
}

func (p *NotFoundException) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("NotFoundException"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *NotFoundException) writeField1(oprot thrift.TProtocol) (err error) {
	if p.IsSetMessage() {
		if err := oprot.WriteFieldBegin("message", thrift.STRING, 1); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:message: ", p), err)
		}
		if err := oprot.WriteString(string(*p.Message)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.message (1) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 1:message: ", p), err)
		}
	}
	return err
}

func (p *NotFoundException) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("NotFoundException(%+v)", *p)
}

func (p *NotFoundException) Error() string {
	return p.String()
}
