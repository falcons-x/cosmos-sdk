// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package types

import (
	context "context"
	types "github.com/cosmos/cosmos-sdk/types"
	grpc "google.golang.org/grpc"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MsgClient interface {
	// SubmitProposal defines a method to create new proposal given a content.
	SubmitProposal(ctx context.Context, in *MsgSubmitProposal, opts ...grpc.CallOption) (*MsgSubmitProposalResponse, error)
	// Vote defines a method to add a vote on a specific proposal.
	Vote(ctx context.Context, in *MsgVote, opts ...grpc.CallOption) (*MsgVoteResponse, error)
	// WeightedVote defines a method to add a weighted vote on a specific proposal.
	VoteWeighted(ctx context.Context, in *MsgVoteWeighted, opts ...grpc.CallOption) (*MsgVoteWeightedResponse, error)
	// Deposit defines a method to add deposit on a specific proposal.
	Deposit(ctx context.Context, in *MsgDeposit, opts ...grpc.CallOption) (*MsgDepositResponse, error)
}

type msgClient struct {
	cc              grpc.ClientConnInterface
	_SubmitProposal types.Invoker
	_Vote           types.Invoker
	_VoteWeighted   types.Invoker
	_Deposit        types.Invoker
}

func NewMsgClient(cc grpc.ClientConnInterface) MsgClient {
	return &msgClient{cc: cc}
}

func (c *msgClient) SubmitProposal(ctx context.Context, in *MsgSubmitProposal, opts ...grpc.CallOption) (*MsgSubmitProposalResponse, error) {
	if invoker := c._SubmitProposal; invoker != nil {
		var out MsgSubmitProposalResponse
		err := invoker(ctx, in, &out)
		return &out, err
	}
	if invokerConn, ok := c.cc.(types.InvokerConn); ok {
		var err error
		c._SubmitProposal, err = invokerConn.Invoker("/cosmos.gov.v1beta1.Msg/SubmitProposal")
		if err != nil {
			var out MsgSubmitProposalResponse
			err = c._SubmitProposal(ctx, in, &out)
			return &out, err
		}
	}
	out := new(MsgSubmitProposalResponse)
	err := c.cc.Invoke(ctx, "/cosmos.gov.v1beta1.Msg/SubmitProposal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) Vote(ctx context.Context, in *MsgVote, opts ...grpc.CallOption) (*MsgVoteResponse, error) {
	if invoker := c._Vote; invoker != nil {
		var out MsgVoteResponse
		err := invoker(ctx, in, &out)
		return &out, err
	}
	if invokerConn, ok := c.cc.(types.InvokerConn); ok {
		var err error
		c._Vote, err = invokerConn.Invoker("/cosmos.gov.v1beta1.Msg/Vote")
		if err != nil {
			var out MsgVoteResponse
			err = c._Vote(ctx, in, &out)
			return &out, err
		}
	}
	out := new(MsgVoteResponse)
	err := c.cc.Invoke(ctx, "/cosmos.gov.v1beta1.Msg/Vote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) VoteWeighted(ctx context.Context, in *MsgVoteWeighted, opts ...grpc.CallOption) (*MsgVoteWeightedResponse, error) {
	if invoker := c._VoteWeighted; invoker != nil {
		var out MsgVoteWeightedResponse
		err := invoker(ctx, in, &out)
		return &out, err
	}
	if invokerConn, ok := c.cc.(types.InvokerConn); ok {
		var err error
		c._VoteWeighted, err = invokerConn.Invoker("/cosmos.gov.v1beta1.Msg/VoteWeighted")
		if err != nil {
			var out MsgVoteWeightedResponse
			err = c._VoteWeighted(ctx, in, &out)
			return &out, err
		}
	}
	out := new(MsgVoteWeightedResponse)
	err := c.cc.Invoke(ctx, "/cosmos.gov.v1beta1.Msg/VoteWeighted", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) Deposit(ctx context.Context, in *MsgDeposit, opts ...grpc.CallOption) (*MsgDepositResponse, error) {
	if invoker := c._Deposit; invoker != nil {
		var out MsgDepositResponse
		err := invoker(ctx, in, &out)
		return &out, err
	}
	if invokerConn, ok := c.cc.(types.InvokerConn); ok {
		var err error
		c._Deposit, err = invokerConn.Invoker("/cosmos.gov.v1beta1.Msg/Deposit")
		if err != nil {
			var out MsgDepositResponse
			err = c._Deposit(ctx, in, &out)
			return &out, err
		}
	}
	out := new(MsgDepositResponse)
	err := c.cc.Invoke(ctx, "/cosmos.gov.v1beta1.Msg/Deposit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// SubmitProposal defines a method to create new proposal given a content.
	SubmitProposal(types.Context, *MsgSubmitProposal) (*MsgSubmitProposalResponse, error)
	// Vote defines a method to add a vote on a specific proposal.
	Vote(types.Context, *MsgVote) (*MsgVoteResponse, error)
	// WeightedVote defines a method to add a weighted vote on a specific proposal.
	VoteWeighted(types.Context, *MsgVoteWeighted) (*MsgVoteWeightedResponse, error)
	// Deposit defines a method to add deposit on a specific proposal.
	Deposit(types.Context, *MsgDeposit) (*MsgDepositResponse, error)
}

func RegisterMsgServer(s grpc.ServiceRegistrar, srv MsgServer) {
	s.RegisterService(&Msg_ServiceDesc, srv)
}

func _Msg_SubmitProposal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSubmitProposal)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SubmitProposal(types.UnwrapSDKContext(ctx), in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.gov.v1beta1.Msg/SubmitProposal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SubmitProposal(types.UnwrapSDKContext(ctx), req.(*MsgSubmitProposal))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_Vote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgVote)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Vote(types.UnwrapSDKContext(ctx), in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.gov.v1beta1.Msg/Vote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Vote(types.UnwrapSDKContext(ctx), req.(*MsgVote))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_VoteWeighted_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgVoteWeighted)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).VoteWeighted(types.UnwrapSDKContext(ctx), in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.gov.v1beta1.Msg/VoteWeighted",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).VoteWeighted(types.UnwrapSDKContext(ctx), req.(*MsgVoteWeighted))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_Deposit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgDeposit)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Deposit(types.UnwrapSDKContext(ctx), in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.gov.v1beta1.Msg/Deposit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Deposit(types.UnwrapSDKContext(ctx), req.(*MsgDeposit))
	}
	return interceptor(ctx, in, info, handler)
}

// Msg_ServiceDesc is the grpc.ServiceDesc for Msg service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Msg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cosmos.gov.v1beta1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SubmitProposal",
			Handler:    _Msg_SubmitProposal_Handler,
		},
		{
			MethodName: "Vote",
			Handler:    _Msg_Vote_Handler,
		},
		{
			MethodName: "VoteWeighted",
			Handler:    _Msg_VoteWeighted_Handler,
		},
		{
			MethodName: "Deposit",
			Handler:    _Msg_Deposit_Handler,
		},
	},
	Metadata: "cosmos/gov/v1beta1/tx.proto",
}

const (
	MsgSubmitProposalMethod = "/cosmos.gov.v1beta1.Msg/SubmitProposal"
	MsgVoteMethod           = "/cosmos.gov.v1beta1.Msg/Vote"
	MsgVoteWeightedMethod   = "/cosmos.gov.v1beta1.Msg/VoteWeighted"
	MsgDepositMethod        = "/cosmos.gov.v1beta1.Msg/Deposit"
)
