package urlservice

import (
	"backend/pkg/ent"
	"backend/pkg/ent/url"
	"backend/pkg/proto"
	"context"

	gonanoid "github.com/matoous/go-nanoid/v2"
)

type Service struct {
	proto.UnimplementedUrlServiceServer
	db *ent.Client
}

func (u *Service) Get(ctx context.Context, in *proto.GetUrlRequest) (*proto.Url, error) {
	url, err := u.db.Url.Query().Where(url.ShortID(in.ShortId)).Only(ctx)
	if err != nil {
		return nil, err
	}

	return &proto.Url{
		Id:        url.ID.String(),
		ShortId:   url.ShortID,
		Url:       url.URL,
		CreatedAt: url.CreatedAt.String(),
	}, nil
}

func (u *Service) Create(ctx context.Context, in *proto.CreateUrlRequest) (*proto.Url, error) {
	sid, err := gonanoid.New(5)
	if err != nil {
		return nil, err
	}

	url, err := u.db.Url.Create().SetURL(in.GetUrl()).SetShortID(sid).Save(ctx)
	if err != nil {
		return nil, err
	}

	return &proto.Url{
		Id:        url.ID.String(),
		ShortId:   url.ShortID,
		Url:       url.URL,
		CreatedAt: url.CreatedAt.String(),
	}, nil
}
