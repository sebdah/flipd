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

type Flipd interface { //****************************************************
	//*                    Services
	///****************************************************

	Ping() (r string, err error)
	// Parameters:
	//  - Request
	GetFeatures(request *GetFeaturesRequest) (r *GetFeaturesResponse, err error)
	// Parameters:
	//  - Request
	RegisterFeature(request *RegisterFeatureRequest) (err error)
}

//****************************************************
//*                    Services
///****************************************************
type FlipdClient struct {
	Transport       thrift.TTransport
	ProtocolFactory thrift.TProtocolFactory
	InputProtocol   thrift.TProtocol
	OutputProtocol  thrift.TProtocol
	SeqId           int32
}

func NewFlipdClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *FlipdClient {
	return &FlipdClient{Transport: t,
		ProtocolFactory: f,
		InputProtocol:   f.GetProtocol(t),
		OutputProtocol:  f.GetProtocol(t),
		SeqId:           0,
	}
}

func NewFlipdClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *FlipdClient {
	return &FlipdClient{Transport: t,
		ProtocolFactory: nil,
		InputProtocol:   iprot,
		OutputProtocol:  oprot,
		SeqId:           0,
	}
}

func (p *FlipdClient) Ping() (r string, err error) {
	if err = p.sendPing(); err != nil {
		return
	}
	return p.recvPing()
}

func (p *FlipdClient) sendPing() (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	if err = oprot.WriteMessageBegin("ping", thrift.CALL, p.SeqId); err != nil {
		return
	}
	args := FlipdPingArgs{}
	if err = args.Write(oprot); err != nil {
		return
	}
	if err = oprot.WriteMessageEnd(); err != nil {
		return
	}
	return oprot.Flush()
}

func (p *FlipdClient) recvPing() (value string, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	method, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if method != "ping" {
		err = thrift.NewTApplicationException(thrift.WRONG_METHOD_NAME, "ping failed: wrong method name")
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "ping failed: out of sequence response")
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error2 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error3 error
		error3, err = error2.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error3
		return
	}
	if mTypeId != thrift.REPLY {
		err = thrift.NewTApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "ping failed: invalid message type")
		return
	}
	result := FlipdPingResult{}
	if err = result.Read(iprot); err != nil {
		return
	}
	if err = iprot.ReadMessageEnd(); err != nil {
		return
	}
	value = result.GetSuccess()
	return
}

// Parameters:
//  - Request
func (p *FlipdClient) GetFeatures(request *GetFeaturesRequest) (r *GetFeaturesResponse, err error) {
	if err = p.sendGetFeatures(request); err != nil {
		return
	}
	return p.recvGetFeatures()
}

func (p *FlipdClient) sendGetFeatures(request *GetFeaturesRequest) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	if err = oprot.WriteMessageBegin("getFeatures", thrift.CALL, p.SeqId); err != nil {
		return
	}
	args := FlipdGetFeaturesArgs{
		Request: request,
	}
	if err = args.Write(oprot); err != nil {
		return
	}
	if err = oprot.WriteMessageEnd(); err != nil {
		return
	}
	return oprot.Flush()
}

func (p *FlipdClient) recvGetFeatures() (value *GetFeaturesResponse, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	method, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if method != "getFeatures" {
		err = thrift.NewTApplicationException(thrift.WRONG_METHOD_NAME, "getFeatures failed: wrong method name")
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "getFeatures failed: out of sequence response")
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error4 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error5 error
		error5, err = error4.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error5
		return
	}
	if mTypeId != thrift.REPLY {
		err = thrift.NewTApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "getFeatures failed: invalid message type")
		return
	}
	result := FlipdGetFeaturesResult{}
	if err = result.Read(iprot); err != nil {
		return
	}
	if err = iprot.ReadMessageEnd(); err != nil {
		return
	}
	if result.InternalError != nil {
		err = result.InternalError
		return
	}
	value = result.GetSuccess()
	return
}

// Parameters:
//  - Request
func (p *FlipdClient) RegisterFeature(request *RegisterFeatureRequest) (err error) {
	if err = p.sendRegisterFeature(request); err != nil {
		return
	}
	return p.recvRegisterFeature()
}

func (p *FlipdClient) sendRegisterFeature(request *RegisterFeatureRequest) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	if err = oprot.WriteMessageBegin("registerFeature", thrift.CALL, p.SeqId); err != nil {
		return
	}
	args := FlipdRegisterFeatureArgs{
		Request: request,
	}
	if err = args.Write(oprot); err != nil {
		return
	}
	if err = oprot.WriteMessageEnd(); err != nil {
		return
	}
	return oprot.Flush()
}

func (p *FlipdClient) recvRegisterFeature() (err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	method, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if method != "registerFeature" {
		err = thrift.NewTApplicationException(thrift.WRONG_METHOD_NAME, "registerFeature failed: wrong method name")
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "registerFeature failed: out of sequence response")
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error6 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error7 error
		error7, err = error6.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error7
		return
	}
	if mTypeId != thrift.REPLY {
		err = thrift.NewTApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "registerFeature failed: invalid message type")
		return
	}
	result := FlipdRegisterFeatureResult{}
	if err = result.Read(iprot); err != nil {
		return
	}
	if err = iprot.ReadMessageEnd(); err != nil {
		return
	}
	if result.InternalError != nil {
		err = result.InternalError
		return
	} else if result.InvalidInput != nil {
		err = result.InvalidInput
		return
	} else if result.Duplicate != nil {
		err = result.Duplicate
		return
	}
	return
}

type FlipdProcessor struct {
	processorMap map[string]thrift.TProcessorFunction
	handler      Flipd
}

func (p *FlipdProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
	p.processorMap[key] = processor
}

func (p *FlipdProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
	processor, ok = p.processorMap[key]
	return processor, ok
}

func (p *FlipdProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
	return p.processorMap
}

func NewFlipdProcessor(handler Flipd) *FlipdProcessor {

	self8 := &FlipdProcessor{handler: handler, processorMap: make(map[string]thrift.TProcessorFunction)}
	self8.processorMap["ping"] = &flipdProcessorPing{handler: handler}
	self8.processorMap["getFeatures"] = &flipdProcessorGetFeatures{handler: handler}
	self8.processorMap["registerFeature"] = &flipdProcessorRegisterFeature{handler: handler}
	return self8
}

func (p *FlipdProcessor) Process(iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	name, _, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return false, err
	}
	if processor, ok := p.GetProcessorFunction(name); ok {
		return processor.Process(seqId, iprot, oprot)
	}
	iprot.Skip(thrift.STRUCT)
	iprot.ReadMessageEnd()
	x9 := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function "+name)
	oprot.WriteMessageBegin(name, thrift.EXCEPTION, seqId)
	x9.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Flush()
	return false, x9

}

type flipdProcessorPing struct {
	handler Flipd
}

func (p *flipdProcessorPing) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := FlipdPingArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("ping", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := FlipdPingResult{}
	var retval string
	var err2 error
	if retval, err2 = p.handler.Ping(); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing ping: "+err2.Error())
		oprot.WriteMessageBegin("ping", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return true, err2
	} else {
		result.Success = &retval
	}
	if err2 = oprot.WriteMessageBegin("ping", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type flipdProcessorGetFeatures struct {
	handler Flipd
}

func (p *flipdProcessorGetFeatures) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := FlipdGetFeaturesArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("getFeatures", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := FlipdGetFeaturesResult{}
	var retval *GetFeaturesResponse
	var err2 error
	if retval, err2 = p.handler.GetFeatures(args.Request); err2 != nil {
		switch v := err2.(type) {
		case *InternalErrorException:
			result.InternalError = v
		default:
			x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing getFeatures: "+err2.Error())
			oprot.WriteMessageBegin("getFeatures", thrift.EXCEPTION, seqId)
			x.Write(oprot)
			oprot.WriteMessageEnd()
			oprot.Flush()
			return true, err2
		}
	} else {
		result.Success = retval
	}
	if err2 = oprot.WriteMessageBegin("getFeatures", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type flipdProcessorRegisterFeature struct {
	handler Flipd
}

func (p *flipdProcessorRegisterFeature) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := FlipdRegisterFeatureArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("registerFeature", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := FlipdRegisterFeatureResult{}
	var err2 error
	if err2 = p.handler.RegisterFeature(args.Request); err2 != nil {
		switch v := err2.(type) {
		case *InternalErrorException:
			result.InternalError = v
		case *InvalidInputException:
			result.InvalidInput = v
		case *DuplicateException:
			result.Duplicate = v
		default:
			x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing registerFeature: "+err2.Error())
			oprot.WriteMessageBegin("registerFeature", thrift.EXCEPTION, seqId)
			x.Write(oprot)
			oprot.WriteMessageEnd()
			oprot.Flush()
			return true, err2
		}
	}
	if err2 = oprot.WriteMessageBegin("registerFeature", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

// HELPER FUNCTIONS AND STRUCTURES

type FlipdPingArgs struct {
}

func NewFlipdPingArgs() *FlipdPingArgs {
	return &FlipdPingArgs{}
}

func (p *FlipdPingArgs) Read(iprot thrift.TProtocol) error {
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
		if err := iprot.Skip(fieldTypeId); err != nil {
			return err
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

func (p *FlipdPingArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("ping_args"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *FlipdPingArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("FlipdPingArgs(%+v)", *p)
}

// Attributes:
//  - Success
type FlipdPingResult struct {
	Success *string `thrift:"success,0" json:"success,omitempty"`
}

func NewFlipdPingResult() *FlipdPingResult {
	return &FlipdPingResult{}
}

var FlipdPingResult_Success_DEFAULT string

func (p *FlipdPingResult) GetSuccess() string {
	if !p.IsSetSuccess() {
		return FlipdPingResult_Success_DEFAULT
	}
	return *p.Success
}
func (p *FlipdPingResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *FlipdPingResult) Read(iprot thrift.TProtocol) error {
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
		case 0:
			if err := p.readField0(iprot); err != nil {
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

func (p *FlipdPingResult) readField0(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 0: ", err)
	} else {
		p.Success = &v
	}
	return nil
}

func (p *FlipdPingResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("ping_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField0(oprot); err != nil {
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

func (p *FlipdPingResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin("success", thrift.STRING, 0); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err)
		}
		if err := oprot.WriteString(string(*p.Success)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.success (0) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err)
		}
	}
	return err
}

func (p *FlipdPingResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("FlipdPingResult(%+v)", *p)
}

// Attributes:
//  - Request
type FlipdGetFeaturesArgs struct {
	Request *GetFeaturesRequest `thrift:"request,1" json:"request"`
}

func NewFlipdGetFeaturesArgs() *FlipdGetFeaturesArgs {
	return &FlipdGetFeaturesArgs{}
}

var FlipdGetFeaturesArgs_Request_DEFAULT *GetFeaturesRequest

func (p *FlipdGetFeaturesArgs) GetRequest() *GetFeaturesRequest {
	if !p.IsSetRequest() {
		return FlipdGetFeaturesArgs_Request_DEFAULT
	}
	return p.Request
}
func (p *FlipdGetFeaturesArgs) IsSetRequest() bool {
	return p.Request != nil
}

func (p *FlipdGetFeaturesArgs) Read(iprot thrift.TProtocol) error {
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

func (p *FlipdGetFeaturesArgs) readField1(iprot thrift.TProtocol) error {
	p.Request = &GetFeaturesRequest{}
	if err := p.Request.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Request), err)
	}
	return nil
}

func (p *FlipdGetFeaturesArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("getFeatures_args"); err != nil {
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

func (p *FlipdGetFeaturesArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("request", thrift.STRUCT, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:request: ", p), err)
	}
	if err := p.Request.Write(oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Request), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:request: ", p), err)
	}
	return err
}

func (p *FlipdGetFeaturesArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("FlipdGetFeaturesArgs(%+v)", *p)
}

// Attributes:
//  - Success
//  - InternalError
type FlipdGetFeaturesResult struct {
	Success       *GetFeaturesResponse    `thrift:"success,0" json:"success,omitempty"`
	InternalError *InternalErrorException `thrift:"internalError,1" json:"internalError,omitempty"`
}

func NewFlipdGetFeaturesResult() *FlipdGetFeaturesResult {
	return &FlipdGetFeaturesResult{}
}

var FlipdGetFeaturesResult_Success_DEFAULT *GetFeaturesResponse

func (p *FlipdGetFeaturesResult) GetSuccess() *GetFeaturesResponse {
	if !p.IsSetSuccess() {
		return FlipdGetFeaturesResult_Success_DEFAULT
	}
	return p.Success
}

var FlipdGetFeaturesResult_InternalError_DEFAULT *InternalErrorException

func (p *FlipdGetFeaturesResult) GetInternalError() *InternalErrorException {
	if !p.IsSetInternalError() {
		return FlipdGetFeaturesResult_InternalError_DEFAULT
	}
	return p.InternalError
}
func (p *FlipdGetFeaturesResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *FlipdGetFeaturesResult) IsSetInternalError() bool {
	return p.InternalError != nil
}

func (p *FlipdGetFeaturesResult) Read(iprot thrift.TProtocol) error {
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
		case 0:
			if err := p.readField0(iprot); err != nil {
				return err
			}
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

func (p *FlipdGetFeaturesResult) readField0(iprot thrift.TProtocol) error {
	p.Success = &GetFeaturesResponse{}
	if err := p.Success.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Success), err)
	}
	return nil
}

func (p *FlipdGetFeaturesResult) readField1(iprot thrift.TProtocol) error {
	p.InternalError = &InternalErrorException{}
	if err := p.InternalError.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.InternalError), err)
	}
	return nil
}

func (p *FlipdGetFeaturesResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("getFeatures_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField0(oprot); err != nil {
		return err
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

func (p *FlipdGetFeaturesResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin("success", thrift.STRUCT, 0); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err)
		}
		if err := p.Success.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Success), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err)
		}
	}
	return err
}

func (p *FlipdGetFeaturesResult) writeField1(oprot thrift.TProtocol) (err error) {
	if p.IsSetInternalError() {
		if err := oprot.WriteFieldBegin("internalError", thrift.STRUCT, 1); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:internalError: ", p), err)
		}
		if err := p.InternalError.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.InternalError), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 1:internalError: ", p), err)
		}
	}
	return err
}

func (p *FlipdGetFeaturesResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("FlipdGetFeaturesResult(%+v)", *p)
}

// Attributes:
//  - Request
type FlipdRegisterFeatureArgs struct {
	Request *RegisterFeatureRequest `thrift:"request,1" json:"request"`
}

func NewFlipdRegisterFeatureArgs() *FlipdRegisterFeatureArgs {
	return &FlipdRegisterFeatureArgs{}
}

var FlipdRegisterFeatureArgs_Request_DEFAULT *RegisterFeatureRequest

func (p *FlipdRegisterFeatureArgs) GetRequest() *RegisterFeatureRequest {
	if !p.IsSetRequest() {
		return FlipdRegisterFeatureArgs_Request_DEFAULT
	}
	return p.Request
}
func (p *FlipdRegisterFeatureArgs) IsSetRequest() bool {
	return p.Request != nil
}

func (p *FlipdRegisterFeatureArgs) Read(iprot thrift.TProtocol) error {
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

func (p *FlipdRegisterFeatureArgs) readField1(iprot thrift.TProtocol) error {
	p.Request = &RegisterFeatureRequest{}
	if err := p.Request.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Request), err)
	}
	return nil
}

func (p *FlipdRegisterFeatureArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("registerFeature_args"); err != nil {
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

func (p *FlipdRegisterFeatureArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("request", thrift.STRUCT, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:request: ", p), err)
	}
	if err := p.Request.Write(oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Request), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:request: ", p), err)
	}
	return err
}

func (p *FlipdRegisterFeatureArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("FlipdRegisterFeatureArgs(%+v)", *p)
}

// Attributes:
//  - InternalError
//  - InvalidInput
//  - Duplicate
type FlipdRegisterFeatureResult struct {
	InternalError *InternalErrorException `thrift:"internalError,1" json:"internalError,omitempty"`
	InvalidInput  *InvalidInputException  `thrift:"InvalidInput,2" json:"InvalidInput,omitempty"`
	Duplicate     *DuplicateException     `thrift:"duplicate,3" json:"duplicate,omitempty"`
}

func NewFlipdRegisterFeatureResult() *FlipdRegisterFeatureResult {
	return &FlipdRegisterFeatureResult{}
}

var FlipdRegisterFeatureResult_InternalError_DEFAULT *InternalErrorException

func (p *FlipdRegisterFeatureResult) GetInternalError() *InternalErrorException {
	if !p.IsSetInternalError() {
		return FlipdRegisterFeatureResult_InternalError_DEFAULT
	}
	return p.InternalError
}

var FlipdRegisterFeatureResult_InvalidInput_DEFAULT *InvalidInputException

func (p *FlipdRegisterFeatureResult) GetInvalidInput() *InvalidInputException {
	if !p.IsSetInvalidInput() {
		return FlipdRegisterFeatureResult_InvalidInput_DEFAULT
	}
	return p.InvalidInput
}

var FlipdRegisterFeatureResult_Duplicate_DEFAULT *DuplicateException

func (p *FlipdRegisterFeatureResult) GetDuplicate() *DuplicateException {
	if !p.IsSetDuplicate() {
		return FlipdRegisterFeatureResult_Duplicate_DEFAULT
	}
	return p.Duplicate
}
func (p *FlipdRegisterFeatureResult) IsSetInternalError() bool {
	return p.InternalError != nil
}

func (p *FlipdRegisterFeatureResult) IsSetInvalidInput() bool {
	return p.InvalidInput != nil
}

func (p *FlipdRegisterFeatureResult) IsSetDuplicate() bool {
	return p.Duplicate != nil
}

func (p *FlipdRegisterFeatureResult) Read(iprot thrift.TProtocol) error {
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
		case 3:
			if err := p.readField3(iprot); err != nil {
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

func (p *FlipdRegisterFeatureResult) readField1(iprot thrift.TProtocol) error {
	p.InternalError = &InternalErrorException{}
	if err := p.InternalError.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.InternalError), err)
	}
	return nil
}

func (p *FlipdRegisterFeatureResult) readField2(iprot thrift.TProtocol) error {
	p.InvalidInput = &InvalidInputException{}
	if err := p.InvalidInput.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.InvalidInput), err)
	}
	return nil
}

func (p *FlipdRegisterFeatureResult) readField3(iprot thrift.TProtocol) error {
	p.Duplicate = &DuplicateException{}
	if err := p.Duplicate.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Duplicate), err)
	}
	return nil
}

func (p *FlipdRegisterFeatureResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("registerFeature_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := p.writeField3(oprot); err != nil {
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

func (p *FlipdRegisterFeatureResult) writeField1(oprot thrift.TProtocol) (err error) {
	if p.IsSetInternalError() {
		if err := oprot.WriteFieldBegin("internalError", thrift.STRUCT, 1); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:internalError: ", p), err)
		}
		if err := p.InternalError.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.InternalError), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 1:internalError: ", p), err)
		}
	}
	return err
}

func (p *FlipdRegisterFeatureResult) writeField2(oprot thrift.TProtocol) (err error) {
	if p.IsSetInvalidInput() {
		if err := oprot.WriteFieldBegin("InvalidInput", thrift.STRUCT, 2); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:InvalidInput: ", p), err)
		}
		if err := p.InvalidInput.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.InvalidInput), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 2:InvalidInput: ", p), err)
		}
	}
	return err
}

func (p *FlipdRegisterFeatureResult) writeField3(oprot thrift.TProtocol) (err error) {
	if p.IsSetDuplicate() {
		if err := oprot.WriteFieldBegin("duplicate", thrift.STRUCT, 3); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:duplicate: ", p), err)
		}
		if err := p.Duplicate.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Duplicate), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 3:duplicate: ", p), err)
		}
	}
	return err
}

func (p *FlipdRegisterFeatureResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("FlipdRegisterFeatureResult(%+v)", *p)
}