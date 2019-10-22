package hgrpc

import (
	"context"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
	"time"
)

type GrpcInterceptor struct {
	infoLog   *logrus.Logger
	warnLog   *logrus.Logger
	accessLog *logrus.Logger
}

func NewGrpcInterceptor(info, warn, access *logrus.Logger) *GrpcInterceptor {
	return &GrpcInterceptor{
		infoLog:   info,
		warnLog:   warn,
		accessLog: access,
	}
}

func (s *GrpcInterceptor) Interceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	ts := time.Now()
	p, ok := peer.FromContext(ctx)
	clientIP := ""
	if ok && p != nil {
		clientIP = p.Addr.String()
	}

	res, err := handler(ctx, req)

	s.accessLog.WithFields(logrus.Fields{
		"client":    clientIP,
		"url":       info.FullMethod,
		"req":       req,
		"res":       res,
		"err":       err,
		"resTimeNs": time.Now().Sub(ts).Nanoseconds(),
	}).Info()

	return res, err
}
